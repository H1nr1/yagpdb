package rsvp

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jonas747/dcmd"
	"github.com/jonas747/discordgo"
	"github.com/jonas747/yagpdb/bot"
	"github.com/jonas747/yagpdb/bot/eventsystem"
	"github.com/jonas747/yagpdb/commands"
	"github.com/jonas747/yagpdb/common"
	"github.com/jonas747/yagpdb/common/scheduledevents2"
	eventModels "github.com/jonas747/yagpdb/common/scheduledevents2/models"
	"github.com/jonas747/yagpdb/rsvp/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"strconv"
	"strings"
	"sync"
	"time"
)

var _ bot.BotInitHandler = (*Plugin)(nil)

func (p *Plugin) BotInit() {
	eventsystem.AddHandlerAsyncLast(p.handleMessageCreate, eventsystem.EventMessageCreate)
	eventsystem.AddHandlerAsyncLast(p.handleMessageReactionAdd, eventsystem.EventMessageReactionAdd)
	scheduledevents2.RegisterHandler("rsvp_update_session", int64(0), p.handleScheduledUpdate)
}

var _ commands.CommandProvider = (*Plugin)(nil)

func (p *Plugin) AddCommands() {
	catEvents := &dcmd.Category{
		Name:        "Events",
		Description: "Event commands",
		HelpEmoji:   "🎫",
		EmbedColor:  0x42b9f4,
	}
	container := commands.CommandSystem.Root.Sub("events", "event")
	container.NotFound = commands.CommonContainerNotFoundHandler(container, "")

	cmdCreateEvent := &commands.YAGCommand{
		CmdCategory: catEvents,
		Name:        "Create",
		Aliases:     []string{"new", "make"},
		Description: "Creates a event, You will be led through an interactive setup",
		RunFunc: func(parsed *dcmd.Data) (interface{}, error) {

			count, err := models.RSVPSessions(models.RSVPSessionWhere.GuildID.EQ(parsed.GS.ID)).CountG(parsed.Context())
			if err != nil {
				return nil, err
			}

			if count > 25 {
				return "Max 25 active events at a time", nil
			}

			p.setupSessionsMU.Lock()
			for _, v := range p.setupSessions {
				if v.SetupChannel == parsed.CS.ID {
					p.setupSessionsMU.Unlock()
					return "Already a setup process going on in this channel, if you want to exit it type `exit`, admins can force cancel setups with `events stopsetup`", nil
				}
			}

			setupSession := &SetupSession{
				CreatedOnMessageID: parsed.Msg.ID,
				GuildID:            parsed.GS.ID,
				SetupChannel:       parsed.CS.ID,
				AuthorID:           parsed.Msg.Author.ID,
				LastAction:         time.Now(),
				plugin:             p,

				stopCH: make(chan bool),
			}
			go setupSession.loopCheckActive()

			p.setupSessions = append(p.setupSessions, setupSession)
			p.setupSessionsMU.Unlock()

			return "Started interactive setup:\nWhat channel should i put the event embed in? (type `this` or `here` for the current one)", nil
		},
	}

	cmdList := &commands.YAGCommand{
		CmdCategory:         catEvents,
		Name:                "List",
		Aliases:             []string{"ls"},
		Description:         "Lists all events in this server",
		RequireDiscordPerms: []int64{discordgo.PermissionManageServer},
		RunFunc: func(parsed *dcmd.Data) (interface{}, error) {
			events, err := models.RSVPSessions(models.RSVPSessionWhere.GuildID.EQ(parsed.GS.ID), qm.OrderBy("starts_at asc")).AllG(parsed.Context())
			if err != nil {
				return nil, err
			}

			if len(events) < 1 {
				return "No active events on this server.", nil
			}

			var output strings.Builder
			for _, v := range events {
				timeUntil := v.StartsAt.Sub(time.Now())
				humanized := common.HumanizeDuration(common.DurationPrecisionMinutes, timeUntil)

				output.WriteString(fmt.Sprintf("#%2d: **%s** in `%s` https://ptb.discordapp.com/channels/%d/%d/%d\n",
					v.LocalID, v.Title, humanized, parsed.GS.ID, v.ChannelID, v.MessageID))
			}

			return output.String(), nil
		},
	}

	cmdDel := &commands.YAGCommand{
		CmdCategory:         catEvents,
		Name:                "Delete",
		Aliases:             []string{"rm", "del"},
		Description:         "Deletes a event, specify the event ID of the event you wanna delete",
		RequireDiscordPerms: []int64{discordgo.PermissionManageServer},
		RequiredArgs:        1,
		Arguments: []*dcmd.ArgDef{
			&dcmd.ArgDef{Name: "ID", Type: dcmd.Int},
		},
		RunFunc: func(parsed *dcmd.Data) (interface{}, error) {

			m, err := models.RSVPSessions(
				models.RSVPSessionWhere.GuildID.EQ(parsed.GS.ID),
				models.RSVPSessionWhere.LocalID.EQ(parsed.Args[0].Int64()),
			).OneG(parsed.Context())

			if err != nil {
				if err == sql.ErrNoRows {
					return "Unknown event", nil
				}

				return nil, err
			}

			_, err = m.DeleteG(parsed.Context())
			if err != nil {
				return nil, err
			}

			return "Deleted `" + m.Title + "`", nil
		},
	}

	cmdStopSetup := &commands.YAGCommand{
		CmdCategory:         catEvents,
		Name:                "StopSetup",
		Aliases:             []string{"cancelsetup"},
		Description:         "Force cancels the current setup session in this channel",
		RequireDiscordPerms: []int64{discordgo.PermissionManageServer},
		RunFunc: func(parsed *dcmd.Data) (interface{}, error) {

			p.setupSessionsMU.Lock()
			for _, v := range p.setupSessions {
				if v.SetupChannel == parsed.CS.ID {
					p.setupSessionsMU.Unlock()
					go v.remove()
					return "Canceled the current setup in this channel", nil
				}
			}
			p.setupSessionsMU.Unlock()

			return "No ongoing setup in the current channel.", nil
		},
	}

	container.AddCommand(cmdCreateEvent, cmdCreateEvent.GetTrigger())
	container.AddCommand(cmdList, cmdList.GetTrigger())
	container.AddCommand(cmdDel, cmdDel.GetTrigger())
	container.AddCommand(cmdStopSetup, cmdStopSetup.GetTrigger())
}

func (p *Plugin) handleMessageCreate(evt *eventsystem.EventData) {
	m := evt.MessageCreate()
	if m.Author == nil {
		return
	}

	p.setupSessionsMU.Lock()
	defer p.setupSessionsMU.Unlock()

	for _, v := range p.setupSessions {
		if v.SetupChannel == m.ChannelID && m.Author.ID == v.AuthorID {
			go v.handleMessage(m.Message)
			break
		}
	}
}

func UpdateEventEmbed(m *models.RSVPSession) error {

	usersToFetch := []int64{
		m.AuthorID,
	}

	var participants []*models.RSVPParticipant
	if m.R != nil {
		for _, v := range m.R.RSVPSessionsMessageRSVPParticipants {
			usersToFetch = append(usersToFetch, v.UserID)
		}

		participants = m.R.RSVPSessionsMessageRSVPParticipants
	}

	fetchedUsers := bot.GetUsers(m.GuildID, usersToFetch...)

	author := fetchedUsers[0]
	participantUsers := fetchedUsers[1:]

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    author.Username,
			IconURL: author.AvatarURL("64"),
		},
		Title:     fmt.Sprintf("#%d: %s", m.LocalID, m.Title),
		Timestamp: m.StartsAt.Format(time.RFC3339),
		Color:     0x518eef,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Event starts ",
		},
	}

	timeUntil := m.StartsAt.Sub(time.Now())
	timeUntilStr := common.HumanizeDuration(common.DurationPrecisionMinutes, timeUntil)
	if timeUntil > 0 {
		timeUntilStr = "Starts in `" + timeUntilStr + "`"
	} else {
		timeUntilStr = "Started `" + timeUntilStr + "` ago"
	}

	UTCTime := m.StartsAt.UTC()

	const timeFormat = "02 Jan 2006 15:04"

	embed.Description = timeUntilStr

	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name: "Times",
		Value: fmt.Sprintf("UTC: `%s`\nLook at the bottom of this message for a time when it starts in your local time.",
			UTCTime.Format(timeFormat)),
	}, &discordgo.MessageEmbedField{
		Name:  "Reactions usage",
		Value: "React to mark you as a participant, undecided, or not joining",
	})

	participantsEmbed := &discordgo.MessageEmbedField{
		Name:   "Participants",
		Inline: true,
		Value:  "```\n",
	}

	waitingListField := &discordgo.MessageEmbedField{
		Name:   "🕐 Waiting list",
		Inline: true,
		Value:  "```\n",
	}

	addedParticipants := 0
	numWaitingList := 0
	for i, v := range participants {
		if v.JoinState != int16(ParticipantStateJoining) && v.JoinState != int16(ParticipantStateWaitlist) {
			continue
		}

		user := participantUsers[i]
		if (addedParticipants >= m.MaxParticipants && m.MaxParticipants > 0) || v.JoinState == int16(ParticipantStateWaitlist) {
			// we hit the max limit so add them to the waiting list instead
			waitingListField.Value += user.Username + "#" + user.Discriminator + "\n"
			numWaitingList++
			continue
		}

		participantsEmbed.Value += user.Username + "#" + user.Discriminator + "\n"
		addedParticipants++
	}

	if participantsEmbed.Value == "```\n" {
		participantsEmbed.Value += "None"
	}
	participantsEmbed.Value += "```"

	waitingListField.Name += " (" + strconv.Itoa(numWaitingList) + ")"

	if waitingListField.Value == "```\n" {
		waitingListField.Value += "None"
	}
	waitingListField.Value += "```"

	if m.MaxParticipants > 0 {
		participantsEmbed.Name += fmt.Sprintf(" (%d / %d)", addedParticipants, m.MaxParticipants)
	} else {
		participantsEmbed.Name += fmt.Sprintf("(%d)", addedParticipants)
	}

	// The undecided and maybe people
	undecidedField := ParticipantField(ParticipantStateMaybe, participants, participantUsers, "❔ Undecided")
	// notJoiningField := ParticipantField(ParticipantStateNotJoining, participants, participantUsers, "Not joining")

	embed.Fields = append(embed.Fields, participantsEmbed, waitingListField, undecidedField)

	_, err := common.BotSession.ChannelMessageEditEmbed(m.ChannelID, m.MessageID, embed)
	return err
}

func ParticipantField(state ParticipantState, participants []*models.RSVPParticipant, users []*discordgo.User, name string) *discordgo.MessageEmbedField {
	field := &discordgo.MessageEmbedField{
		Name:   name,
		Inline: true,
		Value:  "```\n",
	}

	count := 0
	for i, v := range participants {
		user := users[i]

		if v.JoinState == int16(state) {
			field.Value += user.Username + "#" + user.Discriminator + "\n"
			count++

			if count >= 100 {
				break
			}
		}
	}

	if count == 0 {
		field.Value += "No-one\n"
	} else {
		field.Name += " (" + strconv.Itoa(count) + ")"
	}

	field.Value += "```"

	return field
}

func NextUpdateTime(m *models.RSVPSession) time.Time {
	timeUntil := m.StartsAt.Sub(time.Now())

	if timeUntil < time.Second*15 {
		return time.Now().Add(time.Second * 1)
	} else if timeUntil < time.Minute*2 {
		return time.Now().Add(time.Second * 10)
	} else if timeUntil < time.Minute*15 {
		return time.Now().Add(time.Minute)
	} else {
		return time.Now().Add(time.Minute * 10)
	}
}

func (p *Plugin) handleScheduledUpdate(evt *eventModels.ScheduledEvent, data interface{}) (retry bool, err error) {
	mID := *(data.(*int64))

	m, err := models.RSVPSessions(models.RSVPSessionWhere.MessageID.EQ(mID), qm.Load("RSVPSessionsMessageRSVPParticipants", qm.OrderBy("marked_as_participating_at asc"))).OneG(context.Background())
	if err != nil {
		return false, err
	}

	err = UpdateEventEmbed(m)
	if err != nil {
		code, _ := common.DiscordError(err)
		if code == discordgo.ErrCodeUnknownMessage || code == discordgo.ErrCodeUnknownChannel {
			m.DeleteG(context.Background())
			return false, nil
		}

		return scheduledevents2.CheckDiscordErrRetry(err), err
	}

	if m.StartsAt.Sub(time.Now()) < 1 {
		p.startEvent(m)
		return false, nil
	} else if m.StartsAt.Sub(time.Now()) < time.Minute*30 && !m.SentReminders && m.SendReminders {
		m.SentReminders = true
		_, err := m.UpdateG(context.Background(), boil.Whitelist("sent_reminders"))
		if err != nil {
			return true, err
		}

		p.sendReminders(m, "Event is starting in less than 30 minutes!", "The event you signed up for: **"+m.Title+"** is starting soon!")
	}

	err = scheduledevents2.ScheduleEvent("rsvp_update_session", evt.GuildID, NextUpdateTime(m), m.MessageID)
	return false, err
}

type ParticipantState int16

const (
	ParticipantStateJoining    ParticipantState = 1
	ParticipantStateMaybe      ParticipantState = 2
	ParticipantStateNotJoining ParticipantState = 3
	ParticipantStateWaitlist   ParticipantState = 4
)

func (p *Plugin) startEvent(m *models.RSVPSession) error {

	p.sendReminders(m, "Event starting now!", "The event you signed up for: **"+m.Title+"** is starting now!")

	common.BotSession.MessageReactionsRemoveAll(m.ChannelID, m.MessageID)
	_, err := m.DeleteG(context.Background())
	return err
}

func (p *Plugin) sendReminders(m *models.RSVPSession, title, desc string) {

	serverName := strconv.FormatInt(m.GuildID, 10)
	gs := bot.State.Guild(true, m.GuildID)
	if gs != nil {
		gs.RLock()
		serverName = gs.Guild.Name
		gs.RUnlock()
	}

	for _, v := range m.R.RSVPSessionsMessageRSVPParticipants {

		if v.JoinState != int16(ParticipantStateJoining) && v.JoinState != int16(ParticipantStateMaybe) {
			continue
		}

		err := bot.SendDMEmbed(v.UserID, &discordgo.MessageEmbed{
			Title:       title,
			Description: desc,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "From the server: " + serverName,
			},
		})

		if err != nil {
			logger.WithError(err).WithField("guild", m.GuildID).Error("failed sending reminder")
		}
	}

}

func (p *Plugin) handleMessageReactionAdd(evt *eventsystem.EventData) {
	ra := evt.MessageReactionAdd()
	if ra.UserID == common.BotUser.ID {
		return
	}

	joining := ra.Emoji.Name == EmojiJoining
	notJoining := ra.Emoji.Name == EmojiNotJoining
	maybe := ra.Emoji.Name == EmojiMaybe
	waitlist := ra.Emoji.Name == EmojiWaitlist
	if !joining && !notJoining && !maybe && !waitlist {
		return
	}

	m, err := models.RSVPSessions(models.RSVPSessionWhere.MessageID.EQ(ra.MessageID), qm.Load("RSVPSessionsMessageRSVPParticipants", qm.OrderBy("marked_as_participating_at asc"))).OneG(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		logger.WithError(err).WithField("guild", ra.GuildID).Error("failed retrieving RSVP session")
	}

	foundExisting := false
	var participant *models.RSVPParticipant
	for _, v := range m.R.RSVPSessionsMessageRSVPParticipants {
		if v.UserID == ra.UserID {
			participant = v
			foundExisting = true
			break
		}
	}

	if !foundExisting {
		participant = &models.RSVPParticipant{
			RSVPSessionsMessageID: m.MessageID,
			UserID:                ra.UserID,
			GuildID:               ra.GuildID,
		}
	}

	common.BotSession.MessageReactionRemove(ra.ChannelID, ra.MessageID, ra.Emoji.APIName(), ra.UserID)

	if joining {
		if participant.JoinState == int16(ParticipantStateJoining) {
			// already at this state
			return
		}

		participant.JoinState = int16(ParticipantStateJoining)
		participant.MarkedAsParticipatingAt = time.Now()
	} else if maybe {
		if participant.JoinState == int16(ParticipantStateMaybe) {
			// already at this state
			return
		}

		participant.JoinState = int16(ParticipantStateMaybe)
		participant.MarkedAsParticipatingAt = time.Now()
	} else if waitlist {
		if participant.JoinState == int16(ParticipantStateWaitlist) {
			// already at this state
			return
		}

		participant.JoinState = int16(ParticipantStateWaitlist)
		participant.MarkedAsParticipatingAt = time.Now()
	} else if notJoining {
		participant.JoinState = int16(ParticipantStateNotJoining)
	}

	if foundExisting {
		_, err = participant.UpdateG(context.Background(), boil.Infer())
	} else {
		err = m.AddRSVPSessionsMessageRSVPParticipantsG(context.Background(), true, participant)
	}

	if err != nil {
		logger.WithError(err).WithField("guild", ra.GuildID).Error("failed updating rsvp participant")
	}

	updatingSessiosMU.Lock()
	for _, v := range updatingSessionEmbeds {
		if v.ID == m.MessageID {
			v.lastModelUpdate = time.Now()
			updatingSessiosMU.Unlock()
			return
		}
	}

	s := &UpdatingSession{
		ID:              m.MessageID,
		GuildID:         m.GuildID,
		lastModelUpdate: time.Now(),
	}
	updatingSessionEmbeds = append(updatingSessionEmbeds, s)
	go s.run()
	updatingSessiosMU.Unlock()
}

var (
	updatingSessionEmbeds []*UpdatingSession
	updatingSessiosMU     sync.Mutex
)

// Spam update protection, forces 5 seconds between each update
type UpdatingSession struct {
	ID      int64
	GuildID int64

	lastModelUpdate time.Time
	lastEmbedUpdate time.Time
}

func (u *UpdatingSession) run() {
	for {
		u.update()
		time.Sleep(time.Second * 5)

		updatingSessiosMU.Lock()
		if u.lastEmbedUpdate.After(u.lastModelUpdate) || u.lastEmbedUpdate.Equal(u.lastModelUpdate) {
			// remove, no need for further updates

			logger.Println("removing")
			for i, v := range updatingSessionEmbeds {
				if v == u {
					updatingSessionEmbeds = append(updatingSessionEmbeds[:i], updatingSessionEmbeds[i+1:]...)
					break
				}
			}

			updatingSessiosMU.Unlock()
			return
		}

		updatingSessiosMU.Unlock()
	}
}

func (u *UpdatingSession) update() {
	logger.Println("Updating")

	updatingSessiosMU.Lock()
	u.lastEmbedUpdate = time.Now()
	updatingSessiosMU.Unlock()

	m, err := models.RSVPSessions(models.RSVPSessionWhere.MessageID.EQ(u.ID), qm.Load("RSVPSessionsMessageRSVPParticipants", qm.OrderBy("marked_as_participating_at asc"))).OneG(context.Background())
	if err != nil {
		logger.WithError(err).WithField("guild", u.GuildID).Error("failed retreiving rsvp")
		return
	}

	err = UpdateEventEmbed(m)
	if err != nil {
		logger.WithError(err).WithField("guild", u.GuildID).Error("failed retreiving rsvp")
	}
}
