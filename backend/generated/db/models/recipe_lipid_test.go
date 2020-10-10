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

func testRecipeLipids(t *testing.T) {
	t.Parallel()

	query := RecipeLipids()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRecipeLipidsSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
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

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLipidsQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RecipeLipids().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLipidsSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecipeLipidSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLipidsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
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

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLipidsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RecipeLipids().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLipidsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecipeLipidSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLipidsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RecipeLipidExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RecipeLipid exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RecipeLipidExists to return true, but got false.")
	}
}

func testRecipeLipidsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	recipeLipidFound, err := FindRecipeLipid(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if recipeLipidFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRecipeLipidsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RecipeLipids().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRecipeLipidsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RecipeLipids().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRecipeLipidsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	recipeLipidOne := &RecipeLipid{}
	recipeLipidTwo := &RecipeLipid{}
	if err = randomize.Struct(seed, recipeLipidOne, recipeLipidDBTypes, false, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}
	if err = randomize.Struct(seed, recipeLipidTwo, recipeLipidDBTypes, false, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = recipeLipidOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = recipeLipidTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RecipeLipids().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRecipeLipidsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recipeLipidOne := &RecipeLipid{}
	recipeLipidTwo := &RecipeLipid{}
	if err = randomize.Struct(seed, recipeLipidOne, recipeLipidDBTypes, false, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}
	if err = randomize.Struct(seed, recipeLipidTwo, recipeLipidDBTypes, false, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = recipeLipidOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = recipeLipidTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func recipeLipidBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLipid) error {
	*o = RecipeLipid{}
	return nil
}

func recipeLipidAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLipid) error {
	*o = RecipeLipid{}
	return nil
}

func recipeLipidAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLipid) error {
	*o = RecipeLipid{}
	return nil
}

func recipeLipidBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLipid) error {
	*o = RecipeLipid{}
	return nil
}

func recipeLipidAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLipid) error {
	*o = RecipeLipid{}
	return nil
}

func recipeLipidBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLipid) error {
	*o = RecipeLipid{}
	return nil
}

func recipeLipidAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLipid) error {
	*o = RecipeLipid{}
	return nil
}

func recipeLipidBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLipid) error {
	*o = RecipeLipid{}
	return nil
}

func recipeLipidAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLipid) error {
	*o = RecipeLipid{}
	return nil
}

func testRecipeLipidsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RecipeLipid{}
	o := &RecipeLipid{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RecipeLipid object: %s", err)
	}

	AddRecipeLipidHook(boil.BeforeInsertHook, recipeLipidBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	recipeLipidBeforeInsertHooks = []RecipeLipidHook{}

	AddRecipeLipidHook(boil.AfterInsertHook, recipeLipidAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	recipeLipidAfterInsertHooks = []RecipeLipidHook{}

	AddRecipeLipidHook(boil.AfterSelectHook, recipeLipidAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	recipeLipidAfterSelectHooks = []RecipeLipidHook{}

	AddRecipeLipidHook(boil.BeforeUpdateHook, recipeLipidBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	recipeLipidBeforeUpdateHooks = []RecipeLipidHook{}

	AddRecipeLipidHook(boil.AfterUpdateHook, recipeLipidAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	recipeLipidAfterUpdateHooks = []RecipeLipidHook{}

	AddRecipeLipidHook(boil.BeforeDeleteHook, recipeLipidBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	recipeLipidBeforeDeleteHooks = []RecipeLipidHook{}

	AddRecipeLipidHook(boil.AfterDeleteHook, recipeLipidAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	recipeLipidAfterDeleteHooks = []RecipeLipidHook{}

	AddRecipeLipidHook(boil.BeforeUpsertHook, recipeLipidBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	recipeLipidBeforeUpsertHooks = []RecipeLipidHook{}

	AddRecipeLipidHook(boil.AfterUpsertHook, recipeLipidAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	recipeLipidAfterUpsertHooks = []RecipeLipidHook{}
}

func testRecipeLipidsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecipeLipidsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(recipeLipidColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecipeLipidToOneLipidUsingLipid(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RecipeLipid
	var foreign Lipid

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, recipeLipidDBTypes, false, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, lipidDBTypes, false, lipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lipid struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.LipidID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Lipid().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RecipeLipidSlice{&local}
	if err = local.L.LoadLipid(ctx, tx, false, (*[]*RecipeLipid)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Lipid == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Lipid = nil
	if err = local.L.LoadLipid(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Lipid == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRecipeLipidToOneRecipeUsingRecipe(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RecipeLipid
	var foreign Recipe

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, recipeLipidDBTypes, false, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, recipeDBTypes, false, recipeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Recipe struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RecipeID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Recipe().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RecipeLipidSlice{&local}
	if err = local.L.LoadRecipe(ctx, tx, false, (*[]*RecipeLipid)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Recipe == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Recipe = nil
	if err = local.L.LoadRecipe(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Recipe == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRecipeLipidToOneSetOpLipidUsingLipid(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RecipeLipid
	var b, c Lipid

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, recipeLipidDBTypes, false, strmangle.SetComplement(recipeLipidPrimaryKeyColumns, recipeLipidColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, lipidDBTypes, false, strmangle.SetComplement(lipidPrimaryKeyColumns, lipidColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, lipidDBTypes, false, strmangle.SetComplement(lipidPrimaryKeyColumns, lipidColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Lipid{&b, &c} {
		err = a.SetLipid(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Lipid != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RecipeLipids[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.LipidID != x.ID {
			t.Error("foreign key was wrong value", a.LipidID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.LipidID))
		reflect.Indirect(reflect.ValueOf(&a.LipidID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.LipidID != x.ID {
			t.Error("foreign key was wrong value", a.LipidID, x.ID)
		}
	}
}
func testRecipeLipidToOneSetOpRecipeUsingRecipe(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RecipeLipid
	var b, c Recipe

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, recipeLipidDBTypes, false, strmangle.SetComplement(recipeLipidPrimaryKeyColumns, recipeLipidColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, recipeDBTypes, false, strmangle.SetComplement(recipePrimaryKeyColumns, recipeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, recipeDBTypes, false, strmangle.SetComplement(recipePrimaryKeyColumns, recipeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Recipe{&b, &c} {
		err = a.SetRecipe(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Recipe != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RecipeLipids[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RecipeID != x.ID {
			t.Error("foreign key was wrong value", a.RecipeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RecipeID))
		reflect.Indirect(reflect.ValueOf(&a.RecipeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RecipeID != x.ID {
			t.Error("foreign key was wrong value", a.RecipeID, x.ID)
		}
	}
}

func testRecipeLipidsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
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

func testRecipeLipidsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecipeLipidSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRecipeLipidsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RecipeLipids().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	recipeLipidDBTypes = map[string]string{`ID`: `integer`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `Percentage`: `double precision`, `LipidID`: `integer`, `RecipeID`: `integer`}
	_                  = bytes.MinRead
)

func testRecipeLipidsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(recipeLipidPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(recipeLipidAllColumns) == len(recipeLipidPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRecipeLipidsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(recipeLipidAllColumns) == len(recipeLipidPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLipid{}
	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, recipeLipidDBTypes, true, recipeLipidPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(recipeLipidAllColumns, recipeLipidPrimaryKeyColumns) {
		fields = recipeLipidAllColumns
	} else {
		fields = strmangle.SetComplement(
			recipeLipidAllColumns,
			recipeLipidPrimaryKeyColumns,
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

	slice := RecipeLipidSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRecipeLipidsUpsert(t *testing.T) {
	t.Parallel()

	if len(recipeLipidAllColumns) == len(recipeLipidPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RecipeLipid{}
	if err = randomize.Struct(seed, &o, recipeLipidDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RecipeLipid: %s", err)
	}

	count, err := RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, recipeLipidDBTypes, false, recipeLipidPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RecipeLipid struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RecipeLipid: %s", err)
	}

	count, err = RecipeLipids().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
