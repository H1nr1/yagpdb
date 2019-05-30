// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRSVPSessions(t *testing.T) {
	t.Parallel()

	query := RSVPSessions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRSVPSessionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRSVPSessionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RSVPSessions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRSVPSessionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RSVPSessionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRSVPSessionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RSVPSessionExists(ctx, tx, o.MessageID)
	if err != nil {
		t.Errorf("Unable to check if RSVPSession exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RSVPSessionExists to return true, but got false.")
	}
}

func testRSVPSessionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	rsvpSessionFound, err := FindRSVPSession(ctx, tx, o.MessageID)
	if err != nil {
		t.Error(err)
	}

	if rsvpSessionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRSVPSessionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RSVPSessions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRSVPSessionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RSVPSessions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRSVPSessionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	rsvpSessionOne := &RSVPSession{}
	rsvpSessionTwo := &RSVPSession{}
	if err = randomize.Struct(seed, rsvpSessionOne, rsvpSessionDBTypes, false, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}
	if err = randomize.Struct(seed, rsvpSessionTwo, rsvpSessionDBTypes, false, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = rsvpSessionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = rsvpSessionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RSVPSessions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRSVPSessionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	rsvpSessionOne := &RSVPSession{}
	rsvpSessionTwo := &RSVPSession{}
	if err = randomize.Struct(seed, rsvpSessionOne, rsvpSessionDBTypes, false, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}
	if err = randomize.Struct(seed, rsvpSessionTwo, rsvpSessionDBTypes, false, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = rsvpSessionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = rsvpSessionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testRSVPSessionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRSVPSessionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(rsvpSessionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRSVPSessionToManyRSVPSessionsMessageRSVPParticipants(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RSVPSession
	var b, c RSVPParticipant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, rsvpParticipantDBTypes, false, rsvpParticipantColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, rsvpParticipantDBTypes, false, rsvpParticipantColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.RSVPSessionsMessageID = a.MessageID
	c.RSVPSessionsMessageID = a.MessageID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.RSVPSessionsMessageRSVPParticipants().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.RSVPSessionsMessageID == b.RSVPSessionsMessageID {
			bFound = true
		}
		if v.RSVPSessionsMessageID == c.RSVPSessionsMessageID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RSVPSessionSlice{&a}
	if err = a.L.LoadRSVPSessionsMessageRSVPParticipants(ctx, tx, false, (*[]*RSVPSession)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RSVPSessionsMessageRSVPParticipants); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RSVPSessionsMessageRSVPParticipants = nil
	if err = a.L.LoadRSVPSessionsMessageRSVPParticipants(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RSVPSessionsMessageRSVPParticipants); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testRSVPSessionToManyAddOpRSVPSessionsMessageRSVPParticipants(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RSVPSession
	var b, c, d, e RSVPParticipant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, rsvpSessionDBTypes, false, strmangle.SetComplement(rsvpSessionPrimaryKeyColumns, rsvpSessionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RSVPParticipant{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, rsvpParticipantDBTypes, false, strmangle.SetComplement(rsvpParticipantPrimaryKeyColumns, rsvpParticipantColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*RSVPParticipant{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRSVPSessionsMessageRSVPParticipants(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.MessageID != first.RSVPSessionsMessageID {
			t.Error("foreign key was wrong value", a.MessageID, first.RSVPSessionsMessageID)
		}
		if a.MessageID != second.RSVPSessionsMessageID {
			t.Error("foreign key was wrong value", a.MessageID, second.RSVPSessionsMessageID)
		}

		if first.R.RSVPSessionsMessage != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.RSVPSessionsMessage != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RSVPSessionsMessageRSVPParticipants[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RSVPSessionsMessageRSVPParticipants[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RSVPSessionsMessageRSVPParticipants().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testRSVPSessionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRSVPSessionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RSVPSessionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRSVPSessionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RSVPSessions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	rsvpSessionDBTypes = map[string]string{`MessageID`: `bigint`, `GuildID`: `bigint`, `ChannelID`: `bigint`, `LocalID`: `bigint`, `AuthorID`: `bigint`, `CreatedAt`: `timestamp with time zone`, `StartsAt`: `timestamp with time zone`, `Title`: `text`, `Description`: `text`, `MaxParticipants`: `integer`, `SendReminders`: `boolean`, `SentReminders`: `boolean`}
	_                  = bytes.MinRead
)

func testRSVPSessionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(rsvpSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(rsvpSessionColumns) == len(rsvpSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRSVPSessionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(rsvpSessionColumns) == len(rsvpSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RSVPSession{}
	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, rsvpSessionDBTypes, true, rsvpSessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(rsvpSessionColumns, rsvpSessionPrimaryKeyColumns) {
		fields = rsvpSessionColumns
	} else {
		fields = strmangle.SetComplement(
			rsvpSessionColumns,
			rsvpSessionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RSVPSessionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRSVPSessionsUpsert(t *testing.T) {
	t.Parallel()

	if len(rsvpSessionColumns) == len(rsvpSessionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RSVPSession{}
	if err = randomize.Struct(seed, &o, rsvpSessionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RSVPSession: %s", err)
	}

	count, err := RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, rsvpSessionDBTypes, false, rsvpSessionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RSVPSession struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RSVPSession: %s", err)
	}

	count, err = RSVPSessions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
