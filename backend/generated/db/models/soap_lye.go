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

// SoapLye is an object representing the database table.
type SoapLye struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Kind      string    `boil:"kind" json:"kind" toml:"kind" yaml:"kind"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Note      string    `boil:"note" json:"note" toml:"note" yaml:"note"`

	R *soapLyeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L soapLyeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SoapLyeColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Kind      string
	Name      string
	Note      string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Kind:      "kind",
	Name:      "name",
	Note:      "note",
}

// Generated where

var SoapLyeWhere = struct {
	ID        whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
	Kind      whereHelperstring
	Name      whereHelperstring
	Note      whereHelperstring
}{
	ID:        whereHelperint{field: "\"soap_lye\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"soap_lye\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"soap_lye\".\"updated_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"soap_lye\".\"deleted_at\""},
	Kind:      whereHelperstring{field: "\"soap_lye\".\"kind\""},
	Name:      whereHelperstring{field: "\"soap_lye\".\"name\""},
	Note:      whereHelperstring{field: "\"soap_lye\".\"note\""},
}

// SoapLyeRels is where relationship names are stored.
var SoapLyeRels = struct {
	LyeSoapRecipelyes string
}{
	LyeSoapRecipelyes: "LyeSoapRecipelyes",
}

// soapLyeR is where relationships are stored.
type soapLyeR struct {
	LyeSoapRecipelyes SoapRecipelyeSlice `boil:"LyeSoapRecipelyes" json:"LyeSoapRecipelyes" toml:"LyeSoapRecipelyes" yaml:"LyeSoapRecipelyes"`
}

// NewStruct creates a new relationship struct
func (*soapLyeR) NewStruct() *soapLyeR {
	return &soapLyeR{}
}

// soapLyeL is where Load methods for each relationship are stored.
type soapLyeL struct{}

var (
	soapLyeAllColumns            = []string{"id", "created_at", "updated_at", "deleted_at", "kind", "name", "note"}
	soapLyeColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "kind", "name", "note"}
	soapLyeColumnsWithDefault    = []string{"id"}
	soapLyePrimaryKeyColumns     = []string{"id"}
)

type (
	// SoapLyeSlice is an alias for a slice of pointers to SoapLye.
	// This should generally be used opposed to []SoapLye.
	SoapLyeSlice []*SoapLye
	// SoapLyeHook is the signature for custom SoapLye hook methods
	SoapLyeHook func(context.Context, boil.ContextExecutor, *SoapLye) error

	soapLyeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	soapLyeType                 = reflect.TypeOf(&SoapLye{})
	soapLyeMapping              = queries.MakeStructMapping(soapLyeType)
	soapLyePrimaryKeyMapping, _ = queries.BindMapping(soapLyeType, soapLyeMapping, soapLyePrimaryKeyColumns)
	soapLyeInsertCacheMut       sync.RWMutex
	soapLyeInsertCache          = make(map[string]insertCache)
	soapLyeUpdateCacheMut       sync.RWMutex
	soapLyeUpdateCache          = make(map[string]updateCache)
	soapLyeUpsertCacheMut       sync.RWMutex
	soapLyeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var soapLyeBeforeInsertHooks []SoapLyeHook
var soapLyeBeforeUpdateHooks []SoapLyeHook
var soapLyeBeforeDeleteHooks []SoapLyeHook
var soapLyeBeforeUpsertHooks []SoapLyeHook

var soapLyeAfterInsertHooks []SoapLyeHook
var soapLyeAfterSelectHooks []SoapLyeHook
var soapLyeAfterUpdateHooks []SoapLyeHook
var soapLyeAfterDeleteHooks []SoapLyeHook
var soapLyeAfterUpsertHooks []SoapLyeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SoapLye) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapLyeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SoapLye) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapLyeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SoapLye) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapLyeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SoapLye) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapLyeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SoapLye) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapLyeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SoapLye) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapLyeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SoapLye) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapLyeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SoapLye) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapLyeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SoapLye) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapLyeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSoapLyeHook registers your hook function for all future operations.
func AddSoapLyeHook(hookPoint boil.HookPoint, soapLyeHook SoapLyeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		soapLyeBeforeInsertHooks = append(soapLyeBeforeInsertHooks, soapLyeHook)
	case boil.BeforeUpdateHook:
		soapLyeBeforeUpdateHooks = append(soapLyeBeforeUpdateHooks, soapLyeHook)
	case boil.BeforeDeleteHook:
		soapLyeBeforeDeleteHooks = append(soapLyeBeforeDeleteHooks, soapLyeHook)
	case boil.BeforeUpsertHook:
		soapLyeBeforeUpsertHooks = append(soapLyeBeforeUpsertHooks, soapLyeHook)
	case boil.AfterInsertHook:
		soapLyeAfterInsertHooks = append(soapLyeAfterInsertHooks, soapLyeHook)
	case boil.AfterSelectHook:
		soapLyeAfterSelectHooks = append(soapLyeAfterSelectHooks, soapLyeHook)
	case boil.AfterUpdateHook:
		soapLyeAfterUpdateHooks = append(soapLyeAfterUpdateHooks, soapLyeHook)
	case boil.AfterDeleteHook:
		soapLyeAfterDeleteHooks = append(soapLyeAfterDeleteHooks, soapLyeHook)
	case boil.AfterUpsertHook:
		soapLyeAfterUpsertHooks = append(soapLyeAfterUpsertHooks, soapLyeHook)
	}
}

// One returns a single soapLye record from the query.
func (q soapLyeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SoapLye, error) {
	o := &SoapLye{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for soap_lye")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SoapLye records from the query.
func (q soapLyeQuery) All(ctx context.Context, exec boil.ContextExecutor) (SoapLyeSlice, error) {
	var o []*SoapLye

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SoapLye slice")
	}

	if len(soapLyeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SoapLye records in the query.
func (q soapLyeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count soap_lye rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q soapLyeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if soap_lye exists")
	}

	return count > 0, nil
}

// LyeSoapRecipelyes retrieves all the soap_recipelye's SoapRecipelyes with an executor via lye_id column.
func (o *SoapLye) LyeSoapRecipelyes(mods ...qm.QueryMod) soapRecipelyeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"soap_recipelye\".\"lye_id\"=?", o.ID),
		qmhelper.WhereIsNull("\"soap_recipelye\".\"deleted_at\""),
	)

	query := SoapRecipelyes(queryMods...)
	queries.SetFrom(query.Query, "\"soap_recipelye\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"soap_recipelye\".*"})
	}

	return query
}

// LoadLyeSoapRecipelyes allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (soapLyeL) LoadLyeSoapRecipelyes(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSoapLye interface{}, mods queries.Applicator) error {
	var slice []*SoapLye
	var object *SoapLye

	if singular {
		object = maybeSoapLye.(*SoapLye)
	} else {
		slice = *maybeSoapLye.(*[]*SoapLye)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &soapLyeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &soapLyeR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`soap_recipelye`),
		qm.WhereIn(`soap_recipelye.lye_id in ?`, args...),
		qmhelper.WhereIsNull(`soap_recipelye.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load soap_recipelye")
	}

	var resultSlice []*SoapRecipelye
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice soap_recipelye")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on soap_recipelye")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for soap_recipelye")
	}

	if len(soapRecipelyeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.LyeSoapRecipelyes = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.LyeID {
				local.R.LyeSoapRecipelyes = append(local.R.LyeSoapRecipelyes, foreign)
				break
			}
		}
	}

	return nil
}

// AddLyeSoapRecipelyes adds the given related objects to the existing relationships
// of the soap_lye, optionally inserting them as new records.
// Appends related to o.R.LyeSoapRecipelyes.
// Sets related.R.Lye appropriately.
func (o *SoapLye) AddLyeSoapRecipelyes(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SoapRecipelye) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.LyeID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"soap_recipelye\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"lye_id"}),
				strmangle.WhereClause("\"", "\"", 2, soapRecipelyePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.LyeID = o.ID
		}
	}

	if o.R == nil {
		o.R = &soapLyeR{
			LyeSoapRecipelyes: related,
		}
	} else {
		o.R.LyeSoapRecipelyes = append(o.R.LyeSoapRecipelyes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &soapRecipelyeR{
				Lye: o,
			}
		} else {
			rel.R.Lye = o
		}
	}
	return nil
}

// SoapLyes retrieves all the records using an executor.
func SoapLyes(mods ...qm.QueryMod) soapLyeQuery {
	mods = append(mods, qm.From("\"soap_lye\""), qmhelper.WhereIsNull("\"soap_lye\".\"deleted_at\""))
	return soapLyeQuery{NewQuery(mods...)}
}

// FindSoapLye retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSoapLye(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*SoapLye, error) {
	soapLyeObj := &SoapLye{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"soap_lye\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, soapLyeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from soap_lye")
	}

	return soapLyeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SoapLye) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no soap_lye provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(soapLyeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	soapLyeInsertCacheMut.RLock()
	cache, cached := soapLyeInsertCache[key]
	soapLyeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			soapLyeAllColumns,
			soapLyeColumnsWithDefault,
			soapLyeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(soapLyeType, soapLyeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(soapLyeType, soapLyeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"soap_lye\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"soap_lye\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into soap_lye")
	}

	if !cached {
		soapLyeInsertCacheMut.Lock()
		soapLyeInsertCache[key] = cache
		soapLyeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SoapLye.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SoapLye) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	soapLyeUpdateCacheMut.RLock()
	cache, cached := soapLyeUpdateCache[key]
	soapLyeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			soapLyeAllColumns,
			soapLyePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update soap_lye, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"soap_lye\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, soapLyePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(soapLyeType, soapLyeMapping, append(wl, soapLyePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update soap_lye row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for soap_lye")
	}

	if !cached {
		soapLyeUpdateCacheMut.Lock()
		soapLyeUpdateCache[key] = cache
		soapLyeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q soapLyeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for soap_lye")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for soap_lye")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SoapLyeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soapLyePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"soap_lye\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, soapLyePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in soapLye slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all soapLye")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SoapLye) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no soap_lye provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(soapLyeColumnsWithDefault, o)

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

	soapLyeUpsertCacheMut.RLock()
	cache, cached := soapLyeUpsertCache[key]
	soapLyeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			soapLyeAllColumns,
			soapLyeColumnsWithDefault,
			soapLyeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			soapLyeAllColumns,
			soapLyePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert soap_lye, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(soapLyePrimaryKeyColumns))
			copy(conflict, soapLyePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"soap_lye\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(soapLyeType, soapLyeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(soapLyeType, soapLyeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert soap_lye")
	}

	if !cached {
		soapLyeUpsertCacheMut.Lock()
		soapLyeUpsertCache[key] = cache
		soapLyeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SoapLye record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SoapLye) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SoapLye provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), soapLyePrimaryKeyMapping)
		sql = "DELETE FROM \"soap_lye\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"soap_lye\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(soapLyeType, soapLyeMapping, append(wl, soapLyePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to delete from soap_lye")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for soap_lye")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q soapLyeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no soapLyeQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from soap_lye")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for soap_lye")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SoapLyeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(soapLyeBeforeDeleteHooks) != 0 {
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
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soapLyePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"soap_lye\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, soapLyePrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soapLyePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"soap_lye\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, soapLyePrimaryKeyColumns, len(o)),
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
		return 0, errors.Wrap(err, "models: unable to delete all from soapLye slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for soap_lye")
	}

	if len(soapLyeAfterDeleteHooks) != 0 {
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
func (o *SoapLye) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSoapLye(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SoapLyeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SoapLyeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soapLyePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"soap_lye\".* FROM \"soap_lye\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, soapLyePrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SoapLyeSlice")
	}

	*o = slice

	return nil
}

// SoapLyeExists checks if the SoapLye row exists.
func SoapLyeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"soap_lye\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if soap_lye exists")
	}

	return exists, nil
}
