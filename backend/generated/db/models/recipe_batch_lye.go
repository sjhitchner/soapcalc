// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// RecipeBatchLye is an object representing the database table.
type RecipeBatchLye struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Weight    float64   `boil:"weight" json:"weight" toml:"weight" yaml:"weight"`
	Discount  float64   `boil:"discount" json:"discount" toml:"discount" yaml:"discount"`
	Cost      float64   `boil:"cost" json:"cost" toml:"cost" yaml:"cost"`
	LyeID     int       `boil:"lye_id" json:"lye_id" toml:"lye_id" yaml:"lye_id"`
	BatchID   int       `boil:"batch_id" json:"batch_id" toml:"batch_id" yaml:"batch_id"`

	R *recipeBatchLyeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L recipeBatchLyeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RecipeBatchLyeColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Weight    string
	Discount  string
	Cost      string
	LyeID     string
	BatchID   string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Weight:    "weight",
	Discount:  "discount",
	Cost:      "cost",
	LyeID:     "lye_id",
	BatchID:   "batch_id",
}

// Generated where

var RecipeBatchLyeWhere = struct {
	ID        whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
	Weight    whereHelperfloat64
	Discount  whereHelperfloat64
	Cost      whereHelperfloat64
	LyeID     whereHelperint
	BatchID   whereHelperint
}{
	ID:        whereHelperint{field: "\"recipe_batch_lye\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"recipe_batch_lye\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"recipe_batch_lye\".\"updated_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"recipe_batch_lye\".\"deleted_at\""},
	Weight:    whereHelperfloat64{field: "\"recipe_batch_lye\".\"weight\""},
	Discount:  whereHelperfloat64{field: "\"recipe_batch_lye\".\"discount\""},
	Cost:      whereHelperfloat64{field: "\"recipe_batch_lye\".\"cost\""},
	LyeID:     whereHelperint{field: "\"recipe_batch_lye\".\"lye_id\""},
	BatchID:   whereHelperint{field: "\"recipe_batch_lye\".\"batch_id\""},
}

// RecipeBatchLyeRels is where relationship names are stored.
var RecipeBatchLyeRels = struct {
	Batch string
	Lye   string
}{
	Batch: "Batch",
	Lye:   "Lye",
}

// recipeBatchLyeR is where relationships are stored.
type recipeBatchLyeR struct {
	Batch *RecipeBatch `boil:"Batch" json:"Batch" toml:"Batch" yaml:"Batch"`
	Lye   *Lye         `boil:"Lye" json:"Lye" toml:"Lye" yaml:"Lye"`
}

// NewStruct creates a new relationship struct
func (*recipeBatchLyeR) NewStruct() *recipeBatchLyeR {
	return &recipeBatchLyeR{}
}

// recipeBatchLyeL is where Load methods for each relationship are stored.
type recipeBatchLyeL struct{}

var (
	recipeBatchLyeAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "weight", "discount", "cost", "lye_id", "batch_id"}
	recipeBatchLyeColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "weight", "discount", "cost", "lye_id", "batch_id"}
	recipeBatchLyeColumnsWithDefault    = []string{"id"}
	recipeBatchLyePrimaryKeyColumns     = []string{"id"}
)

type (
	// RecipeBatchLyeSlice is an alias for a slice of pointers to RecipeBatchLye.
	// This should generally be used opposed to []RecipeBatchLye.
	RecipeBatchLyeSlice []*RecipeBatchLye
	// RecipeBatchLyeHook is the signature for custom RecipeBatchLye hook methods
	RecipeBatchLyeHook func(context.Context, boil.ContextExecutor, *RecipeBatchLye) error

	recipeBatchLyeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	recipeBatchLyeType                 = reflect.TypeOf(&RecipeBatchLye{})
	recipeBatchLyeMapping              = queries.MakeStructMapping(recipeBatchLyeType)
	recipeBatchLyePrimaryKeyMapping, _ = queries.BindMapping(recipeBatchLyeType, recipeBatchLyeMapping, recipeBatchLyePrimaryKeyColumns)
	recipeBatchLyeInsertCacheMut       sync.RWMutex
	recipeBatchLyeInsertCache          = make(map[string]insertCache)
	recipeBatchLyeUpdateCacheMut       sync.RWMutex
	recipeBatchLyeUpdateCache          = make(map[string]updateCache)
	recipeBatchLyeUpsertCacheMut       sync.RWMutex
	recipeBatchLyeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var recipeBatchLyeBeforeInsertHooks []RecipeBatchLyeHook
var recipeBatchLyeBeforeUpdateHooks []RecipeBatchLyeHook
var recipeBatchLyeBeforeDeleteHooks []RecipeBatchLyeHook
var recipeBatchLyeBeforeUpsertHooks []RecipeBatchLyeHook

var recipeBatchLyeAfterInsertHooks []RecipeBatchLyeHook
var recipeBatchLyeAfterSelectHooks []RecipeBatchLyeHook
var recipeBatchLyeAfterUpdateHooks []RecipeBatchLyeHook
var recipeBatchLyeAfterDeleteHooks []RecipeBatchLyeHook
var recipeBatchLyeAfterUpsertHooks []RecipeBatchLyeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RecipeBatchLye) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLyeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RecipeBatchLye) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLyeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RecipeBatchLye) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLyeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RecipeBatchLye) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLyeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RecipeBatchLye) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLyeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RecipeBatchLye) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLyeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RecipeBatchLye) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLyeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RecipeBatchLye) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLyeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RecipeBatchLye) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeBatchLyeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRecipeBatchLyeHook registers your hook function for all future operations.
func AddRecipeBatchLyeHook(hookPoint boil.HookPoint, recipeBatchLyeHook RecipeBatchLyeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		recipeBatchLyeBeforeInsertHooks = append(recipeBatchLyeBeforeInsertHooks, recipeBatchLyeHook)
	case boil.BeforeUpdateHook:
		recipeBatchLyeBeforeUpdateHooks = append(recipeBatchLyeBeforeUpdateHooks, recipeBatchLyeHook)
	case boil.BeforeDeleteHook:
		recipeBatchLyeBeforeDeleteHooks = append(recipeBatchLyeBeforeDeleteHooks, recipeBatchLyeHook)
	case boil.BeforeUpsertHook:
		recipeBatchLyeBeforeUpsertHooks = append(recipeBatchLyeBeforeUpsertHooks, recipeBatchLyeHook)
	case boil.AfterInsertHook:
		recipeBatchLyeAfterInsertHooks = append(recipeBatchLyeAfterInsertHooks, recipeBatchLyeHook)
	case boil.AfterSelectHook:
		recipeBatchLyeAfterSelectHooks = append(recipeBatchLyeAfterSelectHooks, recipeBatchLyeHook)
	case boil.AfterUpdateHook:
		recipeBatchLyeAfterUpdateHooks = append(recipeBatchLyeAfterUpdateHooks, recipeBatchLyeHook)
	case boil.AfterDeleteHook:
		recipeBatchLyeAfterDeleteHooks = append(recipeBatchLyeAfterDeleteHooks, recipeBatchLyeHook)
	case boil.AfterUpsertHook:
		recipeBatchLyeAfterUpsertHooks = append(recipeBatchLyeAfterUpsertHooks, recipeBatchLyeHook)
	}
}

// One returns a single recipeBatchLye record from the query.
func (q recipeBatchLyeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RecipeBatchLye, error) {
	o := &RecipeBatchLye{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for recipe_batch_lye")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RecipeBatchLye records from the query.
func (q recipeBatchLyeQuery) All(ctx context.Context, exec boil.ContextExecutor) (RecipeBatchLyeSlice, error) {
	var o []*RecipeBatchLye

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RecipeBatchLye slice")
	}

	if len(recipeBatchLyeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RecipeBatchLye records in the query.
func (q recipeBatchLyeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count recipe_batch_lye rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q recipeBatchLyeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if recipe_batch_lye exists")
	}

	return count > 0, nil
}

// Batch pointed to by the foreign key.
func (o *RecipeBatchLye) Batch(mods ...qm.QueryMod) recipeBatchQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.BatchID),
		qmhelper.WhereIsNull("deleted_at"),
	}

	queryMods = append(queryMods, mods...)

	query := RecipeBatches(queryMods...)
	queries.SetFrom(query.Query, "\"recipe_batch\"")

	return query
}

// Lye pointed to by the foreign key.
func (o *RecipeBatchLye) Lye(mods ...qm.QueryMod) lyeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.LyeID),
		qmhelper.WhereIsNull("deleted_at"),
	}

	queryMods = append(queryMods, mods...)

	query := Lyes(queryMods...)
	queries.SetFrom(query.Query, "\"lye\"")

	return query
}

// LoadBatch allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeBatchLyeL) LoadBatch(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeBatchLye interface{}, mods queries.Applicator) error {
	var slice []*RecipeBatchLye
	var object *RecipeBatchLye

	if singular {
		object = maybeRecipeBatchLye.(*RecipeBatchLye)
	} else {
		slice = *maybeRecipeBatchLye.(*[]*RecipeBatchLye)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeBatchLyeR{}
		}
		args = append(args, object.BatchID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeBatchLyeR{}
			}

			for _, a := range args {
				if a == obj.BatchID {
					continue Outer
				}
			}

			args = append(args, obj.BatchID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`recipe_batch`),
		qm.WhereIn(`recipe_batch.id in ?`, args...),
		qmhelper.WhereIsNull(`recipe_batch.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load RecipeBatch")
	}

	var resultSlice []*RecipeBatch
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice RecipeBatch")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for recipe_batch")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for recipe_batch")
	}

	if len(recipeBatchLyeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Batch = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BatchID == foreign.ID {
				local.R.Batch = foreign
				break
			}
		}
	}

	return nil
}

// LoadLye allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (recipeBatchLyeL) LoadLye(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRecipeBatchLye interface{}, mods queries.Applicator) error {
	var slice []*RecipeBatchLye
	var object *RecipeBatchLye

	if singular {
		object = maybeRecipeBatchLye.(*RecipeBatchLye)
	} else {
		slice = *maybeRecipeBatchLye.(*[]*RecipeBatchLye)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &recipeBatchLyeR{}
		}
		args = append(args, object.LyeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &recipeBatchLyeR{}
			}

			for _, a := range args {
				if a == obj.LyeID {
					continue Outer
				}
			}

			args = append(args, obj.LyeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`lye`),
		qm.WhereIn(`lye.id in ?`, args...),
		qmhelper.WhereIsNull(`lye.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Lye")
	}

	var resultSlice []*Lye
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Lye")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for lye")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for lye")
	}

	if len(recipeBatchLyeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Lye = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.LyeID == foreign.ID {
				local.R.Lye = foreign
				break
			}
		}
	}

	return nil
}

// SetBatch of the recipeBatchLye to the related item.
// Sets o.R.Batch to related.
// Adds o to related.R.BatchRecipeBatchLyes.
func (o *RecipeBatchLye) SetBatch(ctx context.Context, exec boil.ContextExecutor, insert bool, related *RecipeBatch) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_batch_lye\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"batch_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeBatchLyePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BatchID = related.ID
	if o.R == nil {
		o.R = &recipeBatchLyeR{
			Batch: related,
		}
	} else {
		o.R.Batch = related
	}

	if related.R == nil {
		related.R = &recipeBatchR{
			BatchRecipeBatchLyes: RecipeBatchLyeSlice{o},
		}
	} else {
		related.R.BatchRecipeBatchLyes = append(related.R.BatchRecipeBatchLyes, o)
	}

	return nil
}

// SetLye of the recipeBatchLye to the related item.
// Sets o.R.Lye to related.
// Adds o to related.R.RecipeBatchLye.
func (o *RecipeBatchLye) SetLye(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Lye) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"recipe_batch_lye\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"lye_id"}),
		strmangle.WhereClause("\"", "\"", 2, recipeBatchLyePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.LyeID = related.ID
	if o.R == nil {
		o.R = &recipeBatchLyeR{
			Lye: related,
		}
	} else {
		o.R.Lye = related
	}

	if related.R == nil {
		related.R = &lyeR{
			RecipeBatchLye: o,
		}
	} else {
		related.R.RecipeBatchLye = o
	}

	return nil
}

// RecipeBatchLyes retrieves all the records using an executor.
func RecipeBatchLyes(mods ...qm.QueryMod) recipeBatchLyeQuery {
	mods = append(mods, qm.From("\"recipe_batch_lye\""), qmhelper.WhereIsNull("\"recipe_batch_lye\".\"deleted_at\""))
	return recipeBatchLyeQuery{NewQuery(mods...)}
}

// FindRecipeBatchLye retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRecipeBatchLye(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RecipeBatchLye, error) {
	recipeBatchLyeObj := &RecipeBatchLye{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"recipe_batch_lye\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, recipeBatchLyeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from recipe_batch_lye")
	}

	return recipeBatchLyeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RecipeBatchLye) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_batch_lye provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(recipeBatchLyeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	recipeBatchLyeInsertCacheMut.RLock()
	cache, cached := recipeBatchLyeInsertCache[key]
	recipeBatchLyeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			recipeBatchLyeAllColumns,
			recipeBatchLyeColumnsWithDefault,
			recipeBatchLyeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(recipeBatchLyeType, recipeBatchLyeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(recipeBatchLyeType, recipeBatchLyeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"recipe_batch_lye\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"recipe_batch_lye\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into recipe_batch_lye")
	}

	if !cached {
		recipeBatchLyeInsertCacheMut.Lock()
		recipeBatchLyeInsertCache[key] = cache
		recipeBatchLyeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RecipeBatchLye.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RecipeBatchLye) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	recipeBatchLyeUpdateCacheMut.RLock()
	cache, cached := recipeBatchLyeUpdateCache[key]
	recipeBatchLyeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			recipeBatchLyeAllColumns,
			recipeBatchLyePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update recipe_batch_lye, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"recipe_batch_lye\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, recipeBatchLyePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(recipeBatchLyeType, recipeBatchLyeMapping, append(wl, recipeBatchLyePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update recipe_batch_lye row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for recipe_batch_lye")
	}

	if !cached {
		recipeBatchLyeUpdateCacheMut.Lock()
		recipeBatchLyeUpdateCache[key] = cache
		recipeBatchLyeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q recipeBatchLyeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for recipe_batch_lye")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for recipe_batch_lye")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RecipeBatchLyeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchLyePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"recipe_batch_lye\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, recipeBatchLyePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in recipeBatchLye slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all recipeBatchLye")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RecipeBatchLye) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_batch_lye provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(recipeBatchLyeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	recipeBatchLyeUpsertCacheMut.RLock()
	cache, cached := recipeBatchLyeUpsertCache[key]
	recipeBatchLyeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			recipeBatchLyeAllColumns,
			recipeBatchLyeColumnsWithDefault,
			recipeBatchLyeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			recipeBatchLyeAllColumns,
			recipeBatchLyePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert recipe_batch_lye, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(recipeBatchLyePrimaryKeyColumns))
			copy(conflict, recipeBatchLyePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"recipe_batch_lye\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(recipeBatchLyeType, recipeBatchLyeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(recipeBatchLyeType, recipeBatchLyeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert recipe_batch_lye")
	}

	if !cached {
		recipeBatchLyeUpsertCacheMut.Lock()
		recipeBatchLyeUpsertCache[key] = cache
		recipeBatchLyeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RecipeBatchLye record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RecipeBatchLye) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RecipeBatchLye provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), recipeBatchLyePrimaryKeyMapping)
		sql = "DELETE FROM \"recipe_batch_lye\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"recipe_batch_lye\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(recipeBatchLyeType, recipeBatchLyeMapping, append(wl, recipeBatchLyePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from recipe_batch_lye")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for recipe_batch_lye")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q recipeBatchLyeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no recipeBatchLyeQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipe_batch_lye")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_batch_lye")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RecipeBatchLyeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(recipeBatchLyeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchLyePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"recipe_batch_lye\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeBatchLyePrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchLyePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"recipe_batch_lye\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, recipeBatchLyePrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipeBatchLye slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_batch_lye")
	}

	if len(recipeBatchLyeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RecipeBatchLye) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRecipeBatchLye(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RecipeBatchLyeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RecipeBatchLyeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeBatchLyePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"recipe_batch_lye\".* FROM \"recipe_batch_lye\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, recipeBatchLyePrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RecipeBatchLyeSlice")
	}

	*o = slice

	return nil
}

// RecipeBatchLyeExists checks if the RecipeBatchLye row exists.
func RecipeBatchLyeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"recipe_batch_lye\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if recipe_batch_lye exists")
	}

	return exists, nil
}
