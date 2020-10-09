// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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

func testSoapFragranceinventories(t *testing.T) {
	t.Parallel()

	query := SoapFragranceinventories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSoapFragranceinventoriesSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSoapFragranceinventoriesQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SoapFragranceinventories().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSoapFragranceinventoriesSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SoapFragranceinventorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSoapFragranceinventoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSoapFragranceinventoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SoapFragranceinventories().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSoapFragranceinventoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SoapFragranceinventorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSoapFragranceinventoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SoapFragranceinventoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if SoapFragranceinventory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SoapFragranceinventoryExists to return true, but got false.")
	}
}

func testSoapFragranceinventoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	soapFragranceinventoryFound, err := FindSoapFragranceinventory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if soapFragranceinventoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSoapFragranceinventoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SoapFragranceinventories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSoapFragranceinventoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SoapFragranceinventories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSoapFragranceinventoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	soapFragranceinventoryOne := &SoapFragranceinventory{}
	soapFragranceinventoryTwo := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, soapFragranceinventoryOne, soapFragranceinventoryDBTypes, false, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}
	if err = randomize.Struct(seed, soapFragranceinventoryTwo, soapFragranceinventoryDBTypes, false, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = soapFragranceinventoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = soapFragranceinventoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SoapFragranceinventories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSoapFragranceinventoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	soapFragranceinventoryOne := &SoapFragranceinventory{}
	soapFragranceinventoryTwo := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, soapFragranceinventoryOne, soapFragranceinventoryDBTypes, false, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}
	if err = randomize.Struct(seed, soapFragranceinventoryTwo, soapFragranceinventoryDBTypes, false, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = soapFragranceinventoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = soapFragranceinventoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func soapFragranceinventoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SoapFragranceinventory) error {
	*o = SoapFragranceinventory{}
	return nil
}

func soapFragranceinventoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SoapFragranceinventory) error {
	*o = SoapFragranceinventory{}
	return nil
}

func soapFragranceinventoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SoapFragranceinventory) error {
	*o = SoapFragranceinventory{}
	return nil
}

func soapFragranceinventoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SoapFragranceinventory) error {
	*o = SoapFragranceinventory{}
	return nil
}

func soapFragranceinventoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SoapFragranceinventory) error {
	*o = SoapFragranceinventory{}
	return nil
}

func soapFragranceinventoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SoapFragranceinventory) error {
	*o = SoapFragranceinventory{}
	return nil
}

func soapFragranceinventoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SoapFragranceinventory) error {
	*o = SoapFragranceinventory{}
	return nil
}

func soapFragranceinventoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SoapFragranceinventory) error {
	*o = SoapFragranceinventory{}
	return nil
}

func soapFragranceinventoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SoapFragranceinventory) error {
	*o = SoapFragranceinventory{}
	return nil
}

func testSoapFragranceinventoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SoapFragranceinventory{}
	o := &SoapFragranceinventory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory object: %s", err)
	}

	AddSoapFragranceinventoryHook(boil.BeforeInsertHook, soapFragranceinventoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	soapFragranceinventoryBeforeInsertHooks = []SoapFragranceinventoryHook{}

	AddSoapFragranceinventoryHook(boil.AfterInsertHook, soapFragranceinventoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	soapFragranceinventoryAfterInsertHooks = []SoapFragranceinventoryHook{}

	AddSoapFragranceinventoryHook(boil.AfterSelectHook, soapFragranceinventoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	soapFragranceinventoryAfterSelectHooks = []SoapFragranceinventoryHook{}

	AddSoapFragranceinventoryHook(boil.BeforeUpdateHook, soapFragranceinventoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	soapFragranceinventoryBeforeUpdateHooks = []SoapFragranceinventoryHook{}

	AddSoapFragranceinventoryHook(boil.AfterUpdateHook, soapFragranceinventoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	soapFragranceinventoryAfterUpdateHooks = []SoapFragranceinventoryHook{}

	AddSoapFragranceinventoryHook(boil.BeforeDeleteHook, soapFragranceinventoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	soapFragranceinventoryBeforeDeleteHooks = []SoapFragranceinventoryHook{}

	AddSoapFragranceinventoryHook(boil.AfterDeleteHook, soapFragranceinventoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	soapFragranceinventoryAfterDeleteHooks = []SoapFragranceinventoryHook{}

	AddSoapFragranceinventoryHook(boil.BeforeUpsertHook, soapFragranceinventoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	soapFragranceinventoryBeforeUpsertHooks = []SoapFragranceinventoryHook{}

	AddSoapFragranceinventoryHook(boil.AfterUpsertHook, soapFragranceinventoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	soapFragranceinventoryAfterUpsertHooks = []SoapFragranceinventoryHook{}
}

func testSoapFragranceinventoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSoapFragranceinventoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(soapFragranceinventoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSoapFragranceinventoryToOneSoapFragranceUsingFragrance(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SoapFragranceinventory
	var foreign SoapFragrance

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, soapFragranceinventoryDBTypes, false, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, soapFragranceDBTypes, false, soapFragranceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragrance struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.FragranceID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Fragrance().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SoapFragranceinventorySlice{&local}
	if err = local.L.LoadFragrance(ctx, tx, false, (*[]*SoapFragranceinventory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Fragrance == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Fragrance = nil
	if err = local.L.LoadFragrance(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Fragrance == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSoapFragranceinventoryToOneSoapSupplierUsingSupplier(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SoapFragranceinventory
	var foreign SoapSupplier

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, soapFragranceinventoryDBTypes, false, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, soapSupplierDBTypes, false, soapSupplierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapSupplier struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SupplierID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Supplier().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SoapFragranceinventorySlice{&local}
	if err = local.L.LoadSupplier(ctx, tx, false, (*[]*SoapFragranceinventory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Supplier == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Supplier = nil
	if err = local.L.LoadSupplier(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Supplier == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSoapFragranceinventoryToOneSetOpSoapFragranceUsingFragrance(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SoapFragranceinventory
	var b, c SoapFragrance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, soapFragranceinventoryDBTypes, false, strmangle.SetComplement(soapFragranceinventoryPrimaryKeyColumns, soapFragranceinventoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, soapFragranceDBTypes, false, strmangle.SetComplement(soapFragrancePrimaryKeyColumns, soapFragranceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, soapFragranceDBTypes, false, strmangle.SetComplement(soapFragrancePrimaryKeyColumns, soapFragranceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*SoapFragrance{&b, &c} {
		err = a.SetFragrance(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Fragrance != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FragranceSoapFragranceinventories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.FragranceID != x.ID {
			t.Error("foreign key was wrong value", a.FragranceID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FragranceID))
		reflect.Indirect(reflect.ValueOf(&a.FragranceID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.FragranceID != x.ID {
			t.Error("foreign key was wrong value", a.FragranceID, x.ID)
		}
	}
}
func testSoapFragranceinventoryToOneSetOpSoapSupplierUsingSupplier(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SoapFragranceinventory
	var b, c SoapSupplier

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, soapFragranceinventoryDBTypes, false, strmangle.SetComplement(soapFragranceinventoryPrimaryKeyColumns, soapFragranceinventoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, soapSupplierDBTypes, false, strmangle.SetComplement(soapSupplierPrimaryKeyColumns, soapSupplierColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, soapSupplierDBTypes, false, strmangle.SetComplement(soapSupplierPrimaryKeyColumns, soapSupplierColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*SoapSupplier{&b, &c} {
		err = a.SetSupplier(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Supplier != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SupplierSoapFragranceinventories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SupplierID != x.ID {
			t.Error("foreign key was wrong value", a.SupplierID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SupplierID))
		reflect.Indirect(reflect.ValueOf(&a.SupplierID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SupplierID != x.ID {
			t.Error("foreign key was wrong value", a.SupplierID, x.ID)
		}
	}
}

func testSoapFragranceinventoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
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

func testSoapFragranceinventoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SoapFragranceinventorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSoapFragranceinventoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SoapFragranceinventories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	soapFragranceinventoryDBTypes = map[string]string{`ID`: `integer`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `PurchaseDate`: `timestamp with time zone`, `ExpiryDate`: `timestamp with time zone`, `Cost`: `double precision`, `Weight`: `double precision`, `FragranceID`: `integer`, `SupplierID`: `integer`}
	_                             = bytes.MinRead
)

func testSoapFragranceinventoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(soapFragranceinventoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(soapFragranceinventoryAllColumns) == len(soapFragranceinventoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSoapFragranceinventoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(soapFragranceinventoryAllColumns) == len(soapFragranceinventoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SoapFragranceinventory{}
	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, soapFragranceinventoryDBTypes, true, soapFragranceinventoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(soapFragranceinventoryAllColumns, soapFragranceinventoryPrimaryKeyColumns) {
		fields = soapFragranceinventoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			soapFragranceinventoryAllColumns,
			soapFragranceinventoryPrimaryKeyColumns,
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

	slice := SoapFragranceinventorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSoapFragranceinventoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(soapFragranceinventoryAllColumns) == len(soapFragranceinventoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SoapFragranceinventory{}
	if err = randomize.Struct(seed, &o, soapFragranceinventoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SoapFragranceinventory: %s", err)
	}

	count, err := SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, soapFragranceinventoryDBTypes, false, soapFragranceinventoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SoapFragranceinventory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SoapFragranceinventory: %s", err)
	}

	count, err = SoapFragranceinventories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
