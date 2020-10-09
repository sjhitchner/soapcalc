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

// SoapRecipeadditive is an object representing the database table.
type SoapRecipeadditive struct {
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt  null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	ID         string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Percentage float64   `boil:"percentage" json:"percentage" toml:"percentage" yaml:"percentage"`
	Weight     float64   `boil:"weight" json:"weight" toml:"weight" yaml:"weight"`
	Cost       float64   `boil:"cost" json:"cost" toml:"cost" yaml:"cost"`
	AdditiveID int       `boil:"additive_id" json:"additive_id" toml:"additive_id" yaml:"additive_id"`

	R *soapRecipeadditiveR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L soapRecipeadditiveL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SoapRecipeadditiveColumns = struct {
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
	ID         string
	Percentage string
	Weight     string
	Cost       string
	AdditiveID string
}{
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
	ID:         "id",
	Percentage: "percentage",
	Weight:     "weight",
	Cost:       "cost",
	AdditiveID: "additive_id",
}

// Generated where

var SoapRecipeadditiveWhere = struct {
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
	DeletedAt  whereHelpernull_Time
	ID         whereHelperstring
	Percentage whereHelperfloat64
	Weight     whereHelperfloat64
	Cost       whereHelperfloat64
	AdditiveID whereHelperint
}{
	CreatedAt:  whereHelpertime_Time{field: "\"soap_recipeadditive\".\"created_at\""},
	UpdatedAt:  whereHelpertime_Time{field: "\"soap_recipeadditive\".\"updated_at\""},
	DeletedAt:  whereHelpernull_Time{field: "\"soap_recipeadditive\".\"deleted_at\""},
	ID:         whereHelperstring{field: "\"soap_recipeadditive\".\"id\""},
	Percentage: whereHelperfloat64{field: "\"soap_recipeadditive\".\"percentage\""},
	Weight:     whereHelperfloat64{field: "\"soap_recipeadditive\".\"weight\""},
	Cost:       whereHelperfloat64{field: "\"soap_recipeadditive\".\"cost\""},
	AdditiveID: whereHelperint{field: "\"soap_recipeadditive\".\"additive_id\""},
}

// SoapRecipeadditiveRels is where relationship names are stored.
var SoapRecipeadditiveRels = struct {
	Additive            string
	AdditiveSoapRecipes string
}{
	Additive:            "Additive",
	AdditiveSoapRecipes: "AdditiveSoapRecipes",
}

// soapRecipeadditiveR is where relationships are stored.
type soapRecipeadditiveR struct {
	Additive            *SoapAdditive   `boil:"Additive" json:"Additive" toml:"Additive" yaml:"Additive"`
	AdditiveSoapRecipes SoapRecipeSlice `boil:"AdditiveSoapRecipes" json:"AdditiveSoapRecipes" toml:"AdditiveSoapRecipes" yaml:"AdditiveSoapRecipes"`
}

// NewStruct creates a new relationship struct
func (*soapRecipeadditiveR) NewStruct() *soapRecipeadditiveR {
	return &soapRecipeadditiveR{}
}

// soapRecipeadditiveL is where Load methods for each relationship are stored.
type soapRecipeadditiveL struct{}

var (
	soapRecipeadditiveAllColumns            = []string{"created_at", "updated_at", "deleted_at", "id", "percentage", "weight", "cost", "additive_id"}
	soapRecipeadditiveColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "id", "percentage", "weight", "cost", "additive_id"}
	soapRecipeadditiveColumnsWithDefault    = []string{}
	soapRecipeadditivePrimaryKeyColumns     = []string{"id"}
)

type (
	// SoapRecipeadditiveSlice is an alias for a slice of pointers to SoapRecipeadditive.
	// This should generally be used opposed to []SoapRecipeadditive.
	SoapRecipeadditiveSlice []*SoapRecipeadditive
	// SoapRecipeadditiveHook is the signature for custom SoapRecipeadditive hook methods
	SoapRecipeadditiveHook func(context.Context, boil.ContextExecutor, *SoapRecipeadditive) error

	soapRecipeadditiveQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	soapRecipeadditiveType                 = reflect.TypeOf(&SoapRecipeadditive{})
	soapRecipeadditiveMapping              = queries.MakeStructMapping(soapRecipeadditiveType)
	soapRecipeadditivePrimaryKeyMapping, _ = queries.BindMapping(soapRecipeadditiveType, soapRecipeadditiveMapping, soapRecipeadditivePrimaryKeyColumns)
	soapRecipeadditiveInsertCacheMut       sync.RWMutex
	soapRecipeadditiveInsertCache          = make(map[string]insertCache)
	soapRecipeadditiveUpdateCacheMut       sync.RWMutex
	soapRecipeadditiveUpdateCache          = make(map[string]updateCache)
	soapRecipeadditiveUpsertCacheMut       sync.RWMutex
	soapRecipeadditiveUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var soapRecipeadditiveBeforeInsertHooks []SoapRecipeadditiveHook
var soapRecipeadditiveBeforeUpdateHooks []SoapRecipeadditiveHook
var soapRecipeadditiveBeforeDeleteHooks []SoapRecipeadditiveHook
var soapRecipeadditiveBeforeUpsertHooks []SoapRecipeadditiveHook

var soapRecipeadditiveAfterInsertHooks []SoapRecipeadditiveHook
var soapRecipeadditiveAfterSelectHooks []SoapRecipeadditiveHook
var soapRecipeadditiveAfterUpdateHooks []SoapRecipeadditiveHook
var soapRecipeadditiveAfterDeleteHooks []SoapRecipeadditiveHook
var soapRecipeadditiveAfterUpsertHooks []SoapRecipeadditiveHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SoapRecipeadditive) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapRecipeadditiveBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SoapRecipeadditive) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapRecipeadditiveBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SoapRecipeadditive) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapRecipeadditiveBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SoapRecipeadditive) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapRecipeadditiveBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SoapRecipeadditive) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapRecipeadditiveAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SoapRecipeadditive) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapRecipeadditiveAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SoapRecipeadditive) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapRecipeadditiveAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SoapRecipeadditive) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapRecipeadditiveAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SoapRecipeadditive) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range soapRecipeadditiveAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSoapRecipeadditiveHook registers your hook function for all future operations.
func AddSoapRecipeadditiveHook(hookPoint boil.HookPoint, soapRecipeadditiveHook SoapRecipeadditiveHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		soapRecipeadditiveBeforeInsertHooks = append(soapRecipeadditiveBeforeInsertHooks, soapRecipeadditiveHook)
	case boil.BeforeUpdateHook:
		soapRecipeadditiveBeforeUpdateHooks = append(soapRecipeadditiveBeforeUpdateHooks, soapRecipeadditiveHook)
	case boil.BeforeDeleteHook:
		soapRecipeadditiveBeforeDeleteHooks = append(soapRecipeadditiveBeforeDeleteHooks, soapRecipeadditiveHook)
	case boil.BeforeUpsertHook:
		soapRecipeadditiveBeforeUpsertHooks = append(soapRecipeadditiveBeforeUpsertHooks, soapRecipeadditiveHook)
	case boil.AfterInsertHook:
		soapRecipeadditiveAfterInsertHooks = append(soapRecipeadditiveAfterInsertHooks, soapRecipeadditiveHook)
	case boil.AfterSelectHook:
		soapRecipeadditiveAfterSelectHooks = append(soapRecipeadditiveAfterSelectHooks, soapRecipeadditiveHook)
	case boil.AfterUpdateHook:
		soapRecipeadditiveAfterUpdateHooks = append(soapRecipeadditiveAfterUpdateHooks, soapRecipeadditiveHook)
	case boil.AfterDeleteHook:
		soapRecipeadditiveAfterDeleteHooks = append(soapRecipeadditiveAfterDeleteHooks, soapRecipeadditiveHook)
	case boil.AfterUpsertHook:
		soapRecipeadditiveAfterUpsertHooks = append(soapRecipeadditiveAfterUpsertHooks, soapRecipeadditiveHook)
	}
}

// One returns a single soapRecipeadditive record from the query.
func (q soapRecipeadditiveQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SoapRecipeadditive, error) {
	o := &SoapRecipeadditive{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for soap_recipeadditive")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SoapRecipeadditive records from the query.
func (q soapRecipeadditiveQuery) All(ctx context.Context, exec boil.ContextExecutor) (SoapRecipeadditiveSlice, error) {
	var o []*SoapRecipeadditive

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SoapRecipeadditive slice")
	}

	if len(soapRecipeadditiveAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SoapRecipeadditive records in the query.
func (q soapRecipeadditiveQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count soap_recipeadditive rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q soapRecipeadditiveQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if soap_recipeadditive exists")
	}

	return count > 0, nil
}

// Additive pointed to by the foreign key.
func (o *SoapRecipeadditive) Additive(mods ...qm.QueryMod) soapAdditiveQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.AdditiveID),
		qmhelper.WhereIsNull("deleted_at"),
	}

	queryMods = append(queryMods, mods...)

	query := SoapAdditives(queryMods...)
	queries.SetFrom(query.Query, "\"soap_additive\"")

	return query
}

// AdditiveSoapRecipes retrieves all the soap_recipe's SoapRecipes with an executor via additives_id column.
func (o *SoapRecipeadditive) AdditiveSoapRecipes(mods ...qm.QueryMod) soapRecipeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"soap_recipe\".\"additives_id\"=?", o.ID),
		qmhelper.WhereIsNull("\"soap_recipe\".\"deleted_at\""),
	)

	query := SoapRecipes(queryMods...)
	queries.SetFrom(query.Query, "\"soap_recipe\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"soap_recipe\".*"})
	}

	return query
}

// LoadAdditive allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (soapRecipeadditiveL) LoadAdditive(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSoapRecipeadditive interface{}, mods queries.Applicator) error {
	var slice []*SoapRecipeadditive
	var object *SoapRecipeadditive

	if singular {
		object = maybeSoapRecipeadditive.(*SoapRecipeadditive)
	} else {
		slice = *maybeSoapRecipeadditive.(*[]*SoapRecipeadditive)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &soapRecipeadditiveR{}
		}
		args = append(args, object.AdditiveID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &soapRecipeadditiveR{}
			}

			for _, a := range args {
				if a == obj.AdditiveID {
					continue Outer
				}
			}

			args = append(args, obj.AdditiveID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`soap_additive`),
		qm.WhereIn(`soap_additive.id in ?`, args...),
		qmhelper.WhereIsNull(`soap_additive.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load SoapAdditive")
	}

	var resultSlice []*SoapAdditive
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice SoapAdditive")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for soap_additive")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for soap_additive")
	}

	if len(soapRecipeadditiveAfterSelectHooks) != 0 {
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
		object.R.Additive = foreign
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AdditiveID == foreign.ID {
				local.R.Additive = foreign
				break
			}
		}
	}

	return nil
}

// LoadAdditiveSoapRecipes allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (soapRecipeadditiveL) LoadAdditiveSoapRecipes(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSoapRecipeadditive interface{}, mods queries.Applicator) error {
	var slice []*SoapRecipeadditive
	var object *SoapRecipeadditive

	if singular {
		object = maybeSoapRecipeadditive.(*SoapRecipeadditive)
	} else {
		slice = *maybeSoapRecipeadditive.(*[]*SoapRecipeadditive)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &soapRecipeadditiveR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &soapRecipeadditiveR{}
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
		qm.From(`soap_recipe`),
		qm.WhereIn(`soap_recipe.additives_id in ?`, args...),
		qmhelper.WhereIsNull(`soap_recipe.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load soap_recipe")
	}

	var resultSlice []*SoapRecipe
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice soap_recipe")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on soap_recipe")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for soap_recipe")
	}

	if len(soapRecipeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AdditiveSoapRecipes = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.AdditivesID {
				local.R.AdditiveSoapRecipes = append(local.R.AdditiveSoapRecipes, foreign)
				break
			}
		}
	}

	return nil
}

// SetAdditive of the soapRecipeadditive to the related item.
// Sets o.R.Additive to related.
// Adds o to related.R.AdditiveSoapRecipeadditives.
func (o *SoapRecipeadditive) SetAdditive(ctx context.Context, exec boil.ContextExecutor, insert bool, related *SoapAdditive) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"soap_recipeadditive\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"additive_id"}),
		strmangle.WhereClause("\"", "\"", 2, soapRecipeadditivePrimaryKeyColumns),
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

	o.AdditiveID = related.ID
	if o.R == nil {
		o.R = &soapRecipeadditiveR{
			Additive: related,
		}
	} else {
		o.R.Additive = related
	}

	if related.R == nil {
		related.R = &soapAdditiveR{
			AdditiveSoapRecipeadditives: SoapRecipeadditiveSlice{o},
		}
	} else {
		related.R.AdditiveSoapRecipeadditives = append(related.R.AdditiveSoapRecipeadditives, o)
	}

	return nil
}

// AddAdditiveSoapRecipes adds the given related objects to the existing relationships
// of the soap_recipeadditive, optionally inserting them as new records.
// Appends related to o.R.AdditiveSoapRecipes.
// Sets related.R.Additive appropriately.
func (o *SoapRecipeadditive) AddAdditiveSoapRecipes(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SoapRecipe) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.AdditivesID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"soap_recipe\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"additives_id"}),
				strmangle.WhereClause("\"", "\"", 2, soapRecipePrimaryKeyColumns),
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

			rel.AdditivesID = o.ID
		}
	}

	if o.R == nil {
		o.R = &soapRecipeadditiveR{
			AdditiveSoapRecipes: related,
		}
	} else {
		o.R.AdditiveSoapRecipes = append(o.R.AdditiveSoapRecipes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &soapRecipeR{
				Additive: o,
			}
		} else {
			rel.R.Additive = o
		}
	}
	return nil
}

// SoapRecipeadditives retrieves all the records using an executor.
func SoapRecipeadditives(mods ...qm.QueryMod) soapRecipeadditiveQuery {
	mods = append(mods, qm.From("\"soap_recipeadditive\""), qmhelper.WhereIsNull("\"soap_recipeadditive\".\"deleted_at\""))
	return soapRecipeadditiveQuery{NewQuery(mods...)}
}

// FindSoapRecipeadditive retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSoapRecipeadditive(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*SoapRecipeadditive, error) {
	soapRecipeadditiveObj := &SoapRecipeadditive{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"soap_recipeadditive\" where \"id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, soapRecipeadditiveObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from soap_recipeadditive")
	}

	return soapRecipeadditiveObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SoapRecipeadditive) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no soap_recipeadditive provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(soapRecipeadditiveColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	soapRecipeadditiveInsertCacheMut.RLock()
	cache, cached := soapRecipeadditiveInsertCache[key]
	soapRecipeadditiveInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			soapRecipeadditiveAllColumns,
			soapRecipeadditiveColumnsWithDefault,
			soapRecipeadditiveColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(soapRecipeadditiveType, soapRecipeadditiveMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(soapRecipeadditiveType, soapRecipeadditiveMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"soap_recipeadditive\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"soap_recipeadditive\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into soap_recipeadditive")
	}

	if !cached {
		soapRecipeadditiveInsertCacheMut.Lock()
		soapRecipeadditiveInsertCache[key] = cache
		soapRecipeadditiveInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SoapRecipeadditive.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SoapRecipeadditive) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	soapRecipeadditiveUpdateCacheMut.RLock()
	cache, cached := soapRecipeadditiveUpdateCache[key]
	soapRecipeadditiveUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			soapRecipeadditiveAllColumns,
			soapRecipeadditivePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update soap_recipeadditive, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"soap_recipeadditive\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, soapRecipeadditivePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(soapRecipeadditiveType, soapRecipeadditiveMapping, append(wl, soapRecipeadditivePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update soap_recipeadditive row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for soap_recipeadditive")
	}

	if !cached {
		soapRecipeadditiveUpdateCacheMut.Lock()
		soapRecipeadditiveUpdateCache[key] = cache
		soapRecipeadditiveUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q soapRecipeadditiveQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for soap_recipeadditive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for soap_recipeadditive")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SoapRecipeadditiveSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soapRecipeadditivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"soap_recipeadditive\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, soapRecipeadditivePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in soapRecipeadditive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all soapRecipeadditive")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SoapRecipeadditive) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no soap_recipeadditive provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(soapRecipeadditiveColumnsWithDefault, o)

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

	soapRecipeadditiveUpsertCacheMut.RLock()
	cache, cached := soapRecipeadditiveUpsertCache[key]
	soapRecipeadditiveUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			soapRecipeadditiveAllColumns,
			soapRecipeadditiveColumnsWithDefault,
			soapRecipeadditiveColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			soapRecipeadditiveAllColumns,
			soapRecipeadditivePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert soap_recipeadditive, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(soapRecipeadditivePrimaryKeyColumns))
			copy(conflict, soapRecipeadditivePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"soap_recipeadditive\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(soapRecipeadditiveType, soapRecipeadditiveMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(soapRecipeadditiveType, soapRecipeadditiveMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert soap_recipeadditive")
	}

	if !cached {
		soapRecipeadditiveUpsertCacheMut.Lock()
		soapRecipeadditiveUpsertCache[key] = cache
		soapRecipeadditiveUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SoapRecipeadditive record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SoapRecipeadditive) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SoapRecipeadditive provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), soapRecipeadditivePrimaryKeyMapping)
		sql = "DELETE FROM \"soap_recipeadditive\" WHERE \"id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"soap_recipeadditive\" SET %s WHERE \"id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(soapRecipeadditiveType, soapRecipeadditiveMapping, append(wl, soapRecipeadditivePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to delete from soap_recipeadditive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for soap_recipeadditive")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q soapRecipeadditiveQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no soapRecipeadditiveQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from soap_recipeadditive")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for soap_recipeadditive")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SoapRecipeadditiveSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(soapRecipeadditiveBeforeDeleteHooks) != 0 {
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
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soapRecipeadditivePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"soap_recipeadditive\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, soapRecipeadditivePrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soapRecipeadditivePrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"soap_recipeadditive\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, soapRecipeadditivePrimaryKeyColumns, len(o)),
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
		return 0, errors.Wrap(err, "models: unable to delete all from soapRecipeadditive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for soap_recipeadditive")
	}

	if len(soapRecipeadditiveAfterDeleteHooks) != 0 {
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
func (o *SoapRecipeadditive) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSoapRecipeadditive(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SoapRecipeadditiveSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SoapRecipeadditiveSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soapRecipeadditivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"soap_recipeadditive\".* FROM \"soap_recipeadditive\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, soapRecipeadditivePrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SoapRecipeadditiveSlice")
	}

	*o = slice

	return nil
}

// SoapRecipeadditiveExists checks if the SoapRecipeadditive row exists.
func SoapRecipeadditiveExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"soap_recipeadditive\" where \"id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if soap_recipeadditive exists")
	}

	return exists, nil
}
