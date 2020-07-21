// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Userinfo is an object representing the database table.
type Userinfo struct {
	ID           int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Username     string `boil:"username" json:"username" toml:"username" yaml:"username"`
	PasswordHash string `boil:"password_hash" json:"password_hash" toml:"password_hash" yaml:"password_hash"`
	UserType     string `boil:"user_type" json:"user_type" toml:"user_type" yaml:"user_type"`

	R *userinfoR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userinfoL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserinfoColumns = struct {
	ID           string
	Username     string
	PasswordHash string
	UserType     string
}{
	ID:           "id",
	Username:     "username",
	PasswordHash: "password_hash",
	UserType:     "user_type",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var UserinfoWhere = struct {
	ID           whereHelperint
	Username     whereHelperstring
	PasswordHash whereHelperstring
	UserType     whereHelperstring
}{
	ID:           whereHelperint{field: "\"userinfo\".\"id\""},
	Username:     whereHelperstring{field: "\"userinfo\".\"username\""},
	PasswordHash: whereHelperstring{field: "\"userinfo\".\"password_hash\""},
	UserType:     whereHelperstring{field: "\"userinfo\".\"user_type\""},
}

// UserinfoRels is where relationship names are stored.
var UserinfoRels = struct {
}{}

// userinfoR is where relationships are stored.
type userinfoR struct {
}

// NewStruct creates a new relationship struct
func (*userinfoR) NewStruct() *userinfoR {
	return &userinfoR{}
}

// userinfoL is where Load methods for each relationship are stored.
type userinfoL struct{}

var (
	userinfoAllColumns            = []string{"id", "username", "password_hash", "user_type"}
	userinfoColumnsWithoutDefault = []string{"username", "password_hash", "user_type"}
	userinfoColumnsWithDefault    = []string{"id"}
	userinfoPrimaryKeyColumns     = []string{"id"}
)

type (
	// UserinfoSlice is an alias for a slice of pointers to Userinfo.
	// This should generally be used opposed to []Userinfo.
	UserinfoSlice []*Userinfo
	// UserinfoHook is the signature for custom Userinfo hook methods
	UserinfoHook func(context.Context, boil.ContextExecutor, *Userinfo) error

	userinfoQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userinfoType                 = reflect.TypeOf(&Userinfo{})
	userinfoMapping              = queries.MakeStructMapping(userinfoType)
	userinfoPrimaryKeyMapping, _ = queries.BindMapping(userinfoType, userinfoMapping, userinfoPrimaryKeyColumns)
	userinfoInsertCacheMut       sync.RWMutex
	userinfoInsertCache          = make(map[string]insertCache)
	userinfoUpdateCacheMut       sync.RWMutex
	userinfoUpdateCache          = make(map[string]updateCache)
	userinfoUpsertCacheMut       sync.RWMutex
	userinfoUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userinfoBeforeInsertHooks []UserinfoHook
var userinfoBeforeUpdateHooks []UserinfoHook
var userinfoBeforeDeleteHooks []UserinfoHook
var userinfoBeforeUpsertHooks []UserinfoHook

var userinfoAfterInsertHooks []UserinfoHook
var userinfoAfterSelectHooks []UserinfoHook
var userinfoAfterUpdateHooks []UserinfoHook
var userinfoAfterDeleteHooks []UserinfoHook
var userinfoAfterUpsertHooks []UserinfoHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Userinfo) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userinfoBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Userinfo) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userinfoBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Userinfo) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userinfoBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Userinfo) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userinfoBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Userinfo) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userinfoAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Userinfo) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userinfoAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Userinfo) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userinfoAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Userinfo) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userinfoAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Userinfo) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userinfoAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserinfoHook registers your hook function for all future operations.
func AddUserinfoHook(hookPoint boil.HookPoint, userinfoHook UserinfoHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		userinfoBeforeInsertHooks = append(userinfoBeforeInsertHooks, userinfoHook)
	case boil.BeforeUpdateHook:
		userinfoBeforeUpdateHooks = append(userinfoBeforeUpdateHooks, userinfoHook)
	case boil.BeforeDeleteHook:
		userinfoBeforeDeleteHooks = append(userinfoBeforeDeleteHooks, userinfoHook)
	case boil.BeforeUpsertHook:
		userinfoBeforeUpsertHooks = append(userinfoBeforeUpsertHooks, userinfoHook)
	case boil.AfterInsertHook:
		userinfoAfterInsertHooks = append(userinfoAfterInsertHooks, userinfoHook)
	case boil.AfterSelectHook:
		userinfoAfterSelectHooks = append(userinfoAfterSelectHooks, userinfoHook)
	case boil.AfterUpdateHook:
		userinfoAfterUpdateHooks = append(userinfoAfterUpdateHooks, userinfoHook)
	case boil.AfterDeleteHook:
		userinfoAfterDeleteHooks = append(userinfoAfterDeleteHooks, userinfoHook)
	case boil.AfterUpsertHook:
		userinfoAfterUpsertHooks = append(userinfoAfterUpsertHooks, userinfoHook)
	}
}

// OneG returns a single userinfo record from the query using the global executor.
func (q userinfoQuery) OneG(ctx context.Context) (*Userinfo, error) {
	return q.One(ctx, boil.GetContextDB())
}

// OneGP returns a single userinfo record from the query using the global executor, and panics on error.
func (q userinfoQuery) OneGP(ctx context.Context) *Userinfo {
	o, err := q.One(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// OneP returns a single userinfo record from the query, and panics on error.
func (q userinfoQuery) OneP(ctx context.Context, exec boil.ContextExecutor) *Userinfo {
	o, err := q.One(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single userinfo record from the query.
func (q userinfoQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Userinfo, error) {
	o := &Userinfo{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for userinfo")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Userinfo records from the query using the global executor.
func (q userinfoQuery) AllG(ctx context.Context) (UserinfoSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// AllGP returns all Userinfo records from the query using the global executor, and panics on error.
func (q userinfoQuery) AllGP(ctx context.Context) UserinfoSlice {
	o, err := q.All(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// AllP returns all Userinfo records from the query, and panics on error.
func (q userinfoQuery) AllP(ctx context.Context, exec boil.ContextExecutor) UserinfoSlice {
	o, err := q.All(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Userinfo records from the query.
func (q userinfoQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserinfoSlice, error) {
	var o []*Userinfo

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to Userinfo slice")
	}

	if len(userinfoAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Userinfo records in the query, and panics on error.
func (q userinfoQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// CountGP returns the count of all Userinfo records in the query using the global executor, and panics on error.
func (q userinfoQuery) CountGP(ctx context.Context) int64 {
	c, err := q.Count(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// CountP returns the count of all Userinfo records in the query, and panics on error.
func (q userinfoQuery) CountP(ctx context.Context, exec boil.ContextExecutor) int64 {
	c, err := q.Count(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Userinfo records in the query.
func (q userinfoQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count userinfo rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q userinfoQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// ExistsGP checks if the row exists in the table using the global executor, and panics on error.
func (q userinfoQuery) ExistsGP(ctx context.Context) bool {
	e, err := q.Exists(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ExistsP checks if the row exists in the table, and panics on error.
func (q userinfoQuery) ExistsP(ctx context.Context, exec boil.ContextExecutor) bool {
	e, err := q.Exists(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q userinfoQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if userinfo exists")
	}

	return count > 0, nil
}

// Userinfos retrieves all the records using an executor.
func Userinfos(mods ...qm.QueryMod) userinfoQuery {
	mods = append(mods, qm.From("\"userinfo\""))
	return userinfoQuery{NewQuery(mods...)}
}

// FindUserinfoG retrieves a single record by ID.
func FindUserinfoG(ctx context.Context, iD int, selectCols ...string) (*Userinfo, error) {
	return FindUserinfo(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindUserinfoP retrieves a single record by ID with an executor, and panics on error.
func FindUserinfoP(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) *Userinfo {
	retobj, err := FindUserinfo(ctx, exec, iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindUserinfoGP retrieves a single record by ID, and panics on error.
func FindUserinfoGP(ctx context.Context, iD int, selectCols ...string) *Userinfo {
	retobj, err := FindUserinfo(ctx, boil.GetContextDB(), iD, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindUserinfo retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserinfo(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Userinfo, error) {
	userinfoObj := &Userinfo{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"userinfo\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userinfoObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from userinfo")
	}

	return userinfoObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Userinfo) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Userinfo) InsertP(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) {
	if err := o.Insert(ctx, exec, columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Userinfo) InsertGP(ctx context.Context, columns boil.Columns) {
	if err := o.Insert(ctx, boil.GetContextDB(), columns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Userinfo) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no userinfo provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userinfoColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userinfoInsertCacheMut.RLock()
	cache, cached := userinfoInsertCache[key]
	userinfoInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userinfoAllColumns,
			userinfoColumnsWithDefault,
			userinfoColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userinfoType, userinfoMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userinfoType, userinfoMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"userinfo\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"userinfo\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "db: unable to insert into userinfo")
	}

	if !cached {
		userinfoInsertCacheMut.Lock()
		userinfoInsertCache[key] = cache
		userinfoInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Userinfo record using the global executor.
// See Update for more documentation.
func (o *Userinfo) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// UpdateP uses an executor to update the Userinfo, and panics on error.
// See Update for more documentation.
func (o *Userinfo) UpdateP(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) int64 {
	rowsAff, err := o.Update(ctx, exec, columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateGP a single Userinfo record using the global executor. Panics on error.
// See Update for more documentation.
func (o *Userinfo) UpdateGP(ctx context.Context, columns boil.Columns) int64 {
	rowsAff, err := o.Update(ctx, boil.GetContextDB(), columns)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Update uses an executor to update the Userinfo.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Userinfo) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userinfoUpdateCacheMut.RLock()
	cache, cached := userinfoUpdateCache[key]
	userinfoUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userinfoAllColumns,
			userinfoPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update userinfo, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"userinfo\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userinfoPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userinfoType, userinfoMapping, append(wl, userinfoPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "db: unable to update userinfo row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for userinfo")
	}

	if !cached {
		userinfoUpdateCacheMut.Lock()
		userinfoUpdateCache[key] = cache
		userinfoUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q userinfoQuery) UpdateAllP(ctx context.Context, exec boil.ContextExecutor, cols M) int64 {
	rowsAff, err := q.UpdateAll(ctx, exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAllG updates all rows with the specified column values.
func (q userinfoQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q userinfoQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for userinfo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for userinfo")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o UserinfoSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o UserinfoSlice) UpdateAllGP(ctx context.Context, cols M) int64 {
	rowsAff, err := o.UpdateAll(ctx, boil.GetContextDB(), cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o UserinfoSlice) UpdateAllP(ctx context.Context, exec boil.ContextExecutor, cols M) int64 {
	rowsAff, err := o.UpdateAll(ctx, exec, cols)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserinfoSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("db: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userinfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"userinfo\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userinfoPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in userinfo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all userinfo")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Userinfo) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Userinfo) UpsertGP(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Userinfo) UpsertP(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) {
	if err := o.Upsert(ctx, exec, updateOnConflict, conflictColumns, updateColumns, insertColumns); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Userinfo) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no userinfo provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userinfoColumnsWithDefault, o)

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

	userinfoUpsertCacheMut.RLock()
	cache, cached := userinfoUpsertCache[key]
	userinfoUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userinfoAllColumns,
			userinfoColumnsWithDefault,
			userinfoColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			userinfoAllColumns,
			userinfoPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("db: unable to upsert userinfo, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userinfoPrimaryKeyColumns))
			copy(conflict, userinfoPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"userinfo\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userinfoType, userinfoMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userinfoType, userinfoMapping, ret)
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
		return errors.Wrap(err, "db: unable to upsert userinfo")
	}

	if !cached {
		userinfoUpsertCacheMut.Lock()
		userinfoUpsertCache[key] = cache
		userinfoUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Userinfo record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Userinfo) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// DeleteP deletes a single Userinfo record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Userinfo) DeleteP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := o.Delete(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteGP deletes a single Userinfo record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Userinfo) DeleteGP(ctx context.Context) int64 {
	rowsAff, err := o.Delete(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// Delete deletes a single Userinfo record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Userinfo) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no Userinfo provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userinfoPrimaryKeyMapping)
	sql := "DELETE FROM \"userinfo\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from userinfo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for userinfo")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q userinfoQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAllP deletes all rows, and panics on error.
func (q userinfoQuery) DeleteAllP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := q.DeleteAll(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all matching rows.
func (q userinfoQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no userinfoQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from userinfo")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for userinfo")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o UserinfoSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o UserinfoSlice) DeleteAllP(ctx context.Context, exec boil.ContextExecutor) int64 {
	rowsAff, err := o.DeleteAll(ctx, exec)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o UserinfoSlice) DeleteAllGP(ctx context.Context) int64 {
	rowsAff, err := o.DeleteAll(ctx, boil.GetContextDB())
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return rowsAff
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserinfoSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userinfoBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userinfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"userinfo\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userinfoPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from userinfo slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for userinfo")
	}

	if len(userinfoAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Userinfo) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("db: no Userinfo provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Userinfo) ReloadP(ctx context.Context, exec boil.ContextExecutor) {
	if err := o.Reload(ctx, exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Userinfo) ReloadGP(ctx context.Context) {
	if err := o.Reload(ctx, boil.GetContextDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Userinfo) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserinfo(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserinfoSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("db: empty UserinfoSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *UserinfoSlice) ReloadAllP(ctx context.Context, exec boil.ContextExecutor) {
	if err := o.ReloadAll(ctx, exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *UserinfoSlice) ReloadAllGP(ctx context.Context) {
	if err := o.ReloadAll(ctx, boil.GetContextDB()); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserinfoSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserinfoSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userinfoPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"userinfo\".* FROM \"userinfo\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userinfoPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in UserinfoSlice")
	}

	*o = slice

	return nil
}

// UserinfoExistsG checks if the Userinfo row exists.
func UserinfoExistsG(ctx context.Context, iD int) (bool, error) {
	return UserinfoExists(ctx, boil.GetContextDB(), iD)
}

// UserinfoExistsP checks if the Userinfo row exists. Panics on error.
func UserinfoExistsP(ctx context.Context, exec boil.ContextExecutor, iD int) bool {
	e, err := UserinfoExists(ctx, exec, iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// UserinfoExistsGP checks if the Userinfo row exists. Panics on error.
func UserinfoExistsGP(ctx context.Context, iD int) bool {
	e, err := UserinfoExists(ctx, boil.GetContextDB(), iD)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// UserinfoExists checks if the Userinfo row exists.
func UserinfoExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"userinfo\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if userinfo exists")
	}

	return exists, nil
}
