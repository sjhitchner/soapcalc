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

func testRecipeLyes(t *testing.T) {
	t.Parallel()

	query := RecipeLyes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRecipeLyesSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
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

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLyesQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RecipeLyes().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLyesSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecipeLyeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLyesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
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

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLyesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RecipeLyes().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLyesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecipeLyeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecipeLyesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RecipeLyeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RecipeLye exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RecipeLyeExists to return true, but got false.")
	}
}

func testRecipeLyesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	recipeLyeFound, err := FindRecipeLye(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if recipeLyeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRecipeLyesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RecipeLyes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRecipeLyesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RecipeLyes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRecipeLyesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	recipeLyeOne := &RecipeLye{}
	recipeLyeTwo := &RecipeLye{}
	if err = randomize.Struct(seed, recipeLyeOne, recipeLyeDBTypes, false, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}
	if err = randomize.Struct(seed, recipeLyeTwo, recipeLyeDBTypes, false, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = recipeLyeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = recipeLyeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RecipeLyes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRecipeLyesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recipeLyeOne := &RecipeLye{}
	recipeLyeTwo := &RecipeLye{}
	if err = randomize.Struct(seed, recipeLyeOne, recipeLyeDBTypes, false, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}
	if err = randomize.Struct(seed, recipeLyeTwo, recipeLyeDBTypes, false, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = recipeLyeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = recipeLyeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func recipeLyeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLye) error {
	*o = RecipeLye{}
	return nil
}

func recipeLyeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLye) error {
	*o = RecipeLye{}
	return nil
}

func recipeLyeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLye) error {
	*o = RecipeLye{}
	return nil
}

func recipeLyeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLye) error {
	*o = RecipeLye{}
	return nil
}

func recipeLyeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLye) error {
	*o = RecipeLye{}
	return nil
}

func recipeLyeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLye) error {
	*o = RecipeLye{}
	return nil
}

func recipeLyeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLye) error {
	*o = RecipeLye{}
	return nil
}

func recipeLyeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLye) error {
	*o = RecipeLye{}
	return nil
}

func recipeLyeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RecipeLye) error {
	*o = RecipeLye{}
	return nil
}

func testRecipeLyesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RecipeLye{}
	o := &RecipeLye{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RecipeLye object: %s", err)
	}

	AddRecipeLyeHook(boil.BeforeInsertHook, recipeLyeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	recipeLyeBeforeInsertHooks = []RecipeLyeHook{}

	AddRecipeLyeHook(boil.AfterInsertHook, recipeLyeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	recipeLyeAfterInsertHooks = []RecipeLyeHook{}

	AddRecipeLyeHook(boil.AfterSelectHook, recipeLyeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	recipeLyeAfterSelectHooks = []RecipeLyeHook{}

	AddRecipeLyeHook(boil.BeforeUpdateHook, recipeLyeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	recipeLyeBeforeUpdateHooks = []RecipeLyeHook{}

	AddRecipeLyeHook(boil.AfterUpdateHook, recipeLyeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	recipeLyeAfterUpdateHooks = []RecipeLyeHook{}

	AddRecipeLyeHook(boil.BeforeDeleteHook, recipeLyeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	recipeLyeBeforeDeleteHooks = []RecipeLyeHook{}

	AddRecipeLyeHook(boil.AfterDeleteHook, recipeLyeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	recipeLyeAfterDeleteHooks = []RecipeLyeHook{}

	AddRecipeLyeHook(boil.BeforeUpsertHook, recipeLyeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	recipeLyeBeforeUpsertHooks = []RecipeLyeHook{}

	AddRecipeLyeHook(boil.AfterUpsertHook, recipeLyeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	recipeLyeAfterUpsertHooks = []RecipeLyeHook{}
}

func testRecipeLyesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecipeLyesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(recipeLyeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecipeLyeToManyLyeRecipes(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RecipeLye
	var b, c Recipe

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, recipeDBTypes, false, recipeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, recipeDBTypes, false, recipeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.LyeID = a.ID
	c.LyeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.LyeRecipes().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.LyeID == b.LyeID {
			bFound = true
		}
		if v.LyeID == c.LyeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RecipeLyeSlice{&a}
	if err = a.L.LoadLyeRecipes(ctx, tx, false, (*[]*RecipeLye)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LyeRecipes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.LyeRecipes = nil
	if err = a.L.LoadLyeRecipes(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.LyeRecipes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testRecipeLyeToManyAddOpLyeRecipes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RecipeLye
	var b, c, d, e Recipe

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, recipeLyeDBTypes, false, strmangle.SetComplement(recipeLyePrimaryKeyColumns, recipeLyeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Recipe{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, recipeDBTypes, false, strmangle.SetComplement(recipePrimaryKeyColumns, recipeColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Recipe{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddLyeRecipes(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.LyeID {
			t.Error("foreign key was wrong value", a.ID, first.LyeID)
		}
		if a.ID != second.LyeID {
			t.Error("foreign key was wrong value", a.ID, second.LyeID)
		}

		if first.R.Lye != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Lye != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.LyeRecipes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.LyeRecipes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.LyeRecipes().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testRecipeLyeToOneLyeUsingLye(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RecipeLye
	var foreign Lye

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, recipeLyeDBTypes, false, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, lyeDBTypes, false, lyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Lye struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.LyeID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Lye().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RecipeLyeSlice{&local}
	if err = local.L.LoadLye(ctx, tx, false, (*[]*RecipeLye)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Lye == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Lye = nil
	if err = local.L.LoadLye(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Lye == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRecipeLyeToOneSetOpLyeUsingLye(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RecipeLye
	var b, c Lye

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, recipeLyeDBTypes, false, strmangle.SetComplement(recipeLyePrimaryKeyColumns, recipeLyeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, lyeDBTypes, false, strmangle.SetComplement(lyePrimaryKeyColumns, lyeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, lyeDBTypes, false, strmangle.SetComplement(lyePrimaryKeyColumns, lyeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Lye{&b, &c} {
		err = a.SetLye(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Lye != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RecipeLyes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.LyeID != x.ID {
			t.Error("foreign key was wrong value", a.LyeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.LyeID))
		reflect.Indirect(reflect.ValueOf(&a.LyeID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.LyeID != x.ID {
			t.Error("foreign key was wrong value", a.LyeID, x.ID)
		}
	}
}

func testRecipeLyesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
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

func testRecipeLyesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecipeLyeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRecipeLyesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RecipeLyes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	recipeLyeDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `ID`: `character varying`, `Weight`: `double precision`, `Concentration`: `double precision`, `Discount`: `double precision`, `Cost`: `double precision`, `LyeID`: `integer`}
	_                = bytes.MinRead
)

func testRecipeLyesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(recipeLyePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(recipeLyeAllColumns) == len(recipeLyePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRecipeLyesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(recipeLyeAllColumns) == len(recipeLyePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RecipeLye{}
	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, recipeLyeDBTypes, true, recipeLyePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(recipeLyeAllColumns, recipeLyePrimaryKeyColumns) {
		fields = recipeLyeAllColumns
	} else {
		fields = strmangle.SetComplement(
			recipeLyeAllColumns,
			recipeLyePrimaryKeyColumns,
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

	slice := RecipeLyeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRecipeLyesUpsert(t *testing.T) {
	t.Parallel()

	if len(recipeLyeAllColumns) == len(recipeLyePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RecipeLye{}
	if err = randomize.Struct(seed, &o, recipeLyeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RecipeLye: %s", err)
	}

	count, err := RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, recipeLyeDBTypes, false, recipeLyePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RecipeLye struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RecipeLye: %s", err)
	}

	count, err = RecipeLyes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}