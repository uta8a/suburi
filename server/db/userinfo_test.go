// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testUserinfos(t *testing.T) {
	t.Parallel()

	query := Userinfos()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testUserinfosDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
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

	count, err := Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserinfosQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Userinfos().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserinfosSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserinfoSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testUserinfosExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := UserinfoExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Userinfo exists: %s", err)
	}
	if !e {
		t.Errorf("Expected UserinfoExists to return true, but got false.")
	}
}

func testUserinfosFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	userinfoFound, err := FindUserinfo(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if userinfoFound == nil {
		t.Error("want a record, got nil")
	}
}

func testUserinfosBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Userinfos().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testUserinfosOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Userinfos().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testUserinfosAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	userinfoOne := &Userinfo{}
	userinfoTwo := &Userinfo{}
	if err = randomize.Struct(seed, userinfoOne, userinfoDBTypes, false, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}
	if err = randomize.Struct(seed, userinfoTwo, userinfoDBTypes, false, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userinfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userinfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Userinfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testUserinfosCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	userinfoOne := &Userinfo{}
	userinfoTwo := &Userinfo{}
	if err = randomize.Struct(seed, userinfoOne, userinfoDBTypes, false, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}
	if err = randomize.Struct(seed, userinfoTwo, userinfoDBTypes, false, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = userinfoOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = userinfoTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func userinfoBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Userinfo) error {
	*o = Userinfo{}
	return nil
}

func userinfoAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Userinfo) error {
	*o = Userinfo{}
	return nil
}

func userinfoAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Userinfo) error {
	*o = Userinfo{}
	return nil
}

func userinfoBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Userinfo) error {
	*o = Userinfo{}
	return nil
}

func userinfoAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Userinfo) error {
	*o = Userinfo{}
	return nil
}

func userinfoBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Userinfo) error {
	*o = Userinfo{}
	return nil
}

func userinfoAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Userinfo) error {
	*o = Userinfo{}
	return nil
}

func userinfoBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Userinfo) error {
	*o = Userinfo{}
	return nil
}

func userinfoAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Userinfo) error {
	*o = Userinfo{}
	return nil
}

func testUserinfosHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Userinfo{}
	o := &Userinfo{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, userinfoDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Userinfo object: %s", err)
	}

	AddUserinfoHook(boil.BeforeInsertHook, userinfoBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	userinfoBeforeInsertHooks = []UserinfoHook{}

	AddUserinfoHook(boil.AfterInsertHook, userinfoAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	userinfoAfterInsertHooks = []UserinfoHook{}

	AddUserinfoHook(boil.AfterSelectHook, userinfoAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	userinfoAfterSelectHooks = []UserinfoHook{}

	AddUserinfoHook(boil.BeforeUpdateHook, userinfoBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	userinfoBeforeUpdateHooks = []UserinfoHook{}

	AddUserinfoHook(boil.AfterUpdateHook, userinfoAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	userinfoAfterUpdateHooks = []UserinfoHook{}

	AddUserinfoHook(boil.BeforeDeleteHook, userinfoBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	userinfoBeforeDeleteHooks = []UserinfoHook{}

	AddUserinfoHook(boil.AfterDeleteHook, userinfoAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	userinfoAfterDeleteHooks = []UserinfoHook{}

	AddUserinfoHook(boil.BeforeUpsertHook, userinfoBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	userinfoBeforeUpsertHooks = []UserinfoHook{}

	AddUserinfoHook(boil.AfterUpsertHook, userinfoAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	userinfoAfterUpsertHooks = []UserinfoHook{}
}

func testUserinfosInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserinfosInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(userinfoColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testUserinfosReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
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

func testUserinfosReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := UserinfoSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testUserinfosSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Userinfos().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	userinfoDBTypes = map[string]string{`ID`: `integer`, `Username`: `character varying`, `PasswordHash`: `text`, `UserType`: `text`}
	_               = bytes.MinRead
)

func testUserinfosUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(userinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(userinfoAllColumns) == len(userinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testUserinfosSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(userinfoAllColumns) == len(userinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Userinfo{}
	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, userinfoDBTypes, true, userinfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(userinfoAllColumns, userinfoPrimaryKeyColumns) {
		fields = userinfoAllColumns
	} else {
		fields = strmangle.SetComplement(
			userinfoAllColumns,
			userinfoPrimaryKeyColumns,
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

	slice := UserinfoSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testUserinfosUpsert(t *testing.T) {
	t.Parallel()

	if len(userinfoAllColumns) == len(userinfoPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Userinfo{}
	if err = randomize.Struct(seed, &o, userinfoDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Userinfo: %s", err)
	}

	count, err := Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, userinfoDBTypes, false, userinfoPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Userinfo struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Userinfo: %s", err)
	}

	count, err = Userinfos().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
