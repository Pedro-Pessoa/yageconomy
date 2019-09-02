// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// EconomyPlant is an object representing the database table.
type EconomyPlant struct {
	MessageID int64     `boil:"message_id" json:"message_id" toml:"message_id" yaml:"message_id"`
	ChannelID int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	GuildID   int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	AuthorID  int64     `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	Amount    int64     `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	Password  string    `boil:"password" json:"password" toml:"password" yaml:"password"`

	R *economyPlantR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L economyPlantL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EconomyPlantColumns = struct {
	MessageID string
	ChannelID string
	GuildID   string
	CreatedAt string
	AuthorID  string
	Amount    string
	Password  string
}{
	MessageID: "message_id",
	ChannelID: "channel_id",
	GuildID:   "guild_id",
	CreatedAt: "created_at",
	AuthorID:  "author_id",
	Amount:    "amount",
	Password:  "password",
}

// Generated where

var EconomyPlantWhere = struct {
	MessageID whereHelperint64
	ChannelID whereHelperint64
	GuildID   whereHelperint64
	CreatedAt whereHelpertime_Time
	AuthorID  whereHelperint64
	Amount    whereHelperint64
	Password  whereHelperstring
}{
	MessageID: whereHelperint64{field: "\"economy_plants\".\"message_id\""},
	ChannelID: whereHelperint64{field: "\"economy_plants\".\"channel_id\""},
	GuildID:   whereHelperint64{field: "\"economy_plants\".\"guild_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"economy_plants\".\"created_at\""},
	AuthorID:  whereHelperint64{field: "\"economy_plants\".\"author_id\""},
	Amount:    whereHelperint64{field: "\"economy_plants\".\"amount\""},
	Password:  whereHelperstring{field: "\"economy_plants\".\"password\""},
}

// EconomyPlantRels is where relationship names are stored.
var EconomyPlantRels = struct {
}{}

// economyPlantR is where relationships are stored.
type economyPlantR struct {
}

// NewStruct creates a new relationship struct
func (*economyPlantR) NewStruct() *economyPlantR {
	return &economyPlantR{}
}

// economyPlantL is where Load methods for each relationship are stored.
type economyPlantL struct{}

var (
	economyPlantAllColumns            = []string{"message_id", "channel_id", "guild_id", "created_at", "author_id", "amount", "password"}
	economyPlantColumnsWithoutDefault = []string{"message_id", "channel_id", "guild_id", "created_at", "author_id", "amount", "password"}
	economyPlantColumnsWithDefault    = []string{}
	economyPlantPrimaryKeyColumns     = []string{"message_id"}
)

type (
	// EconomyPlantSlice is an alias for a slice of pointers to EconomyPlant.
	// This should generally be used opposed to []EconomyPlant.
	EconomyPlantSlice []*EconomyPlant

	economyPlantQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	economyPlantType                 = reflect.TypeOf(&EconomyPlant{})
	economyPlantMapping              = queries.MakeStructMapping(economyPlantType)
	economyPlantPrimaryKeyMapping, _ = queries.BindMapping(economyPlantType, economyPlantMapping, economyPlantPrimaryKeyColumns)
	economyPlantInsertCacheMut       sync.RWMutex
	economyPlantInsertCache          = make(map[string]insertCache)
	economyPlantUpdateCacheMut       sync.RWMutex
	economyPlantUpdateCache          = make(map[string]updateCache)
	economyPlantUpsertCacheMut       sync.RWMutex
	economyPlantUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single economyPlant record from the query using the global executor.
func (q economyPlantQuery) OneG(ctx context.Context) (*EconomyPlant, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single economyPlant record from the query.
func (q economyPlantQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EconomyPlant, error) {
	o := &EconomyPlant{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for economy_plants")
	}

	return o, nil
}

// AllG returns all EconomyPlant records from the query using the global executor.
func (q economyPlantQuery) AllG(ctx context.Context) (EconomyPlantSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all EconomyPlant records from the query.
func (q economyPlantQuery) All(ctx context.Context, exec boil.ContextExecutor) (EconomyPlantSlice, error) {
	var o []*EconomyPlant

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EconomyPlant slice")
	}

	return o, nil
}

// CountG returns the count of all EconomyPlant records in the query, and panics on error.
func (q economyPlantQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all EconomyPlant records in the query.
func (q economyPlantQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count economy_plants rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q economyPlantQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q economyPlantQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if economy_plants exists")
	}

	return count > 0, nil
}

// EconomyPlants retrieves all the records using an executor.
func EconomyPlants(mods ...qm.QueryMod) economyPlantQuery {
	mods = append(mods, qm.From("\"economy_plants\""))
	return economyPlantQuery{NewQuery(mods...)}
}

// FindEconomyPlantG retrieves a single record by ID.
func FindEconomyPlantG(ctx context.Context, messageID int64, selectCols ...string) (*EconomyPlant, error) {
	return FindEconomyPlant(ctx, boil.GetContextDB(), messageID, selectCols...)
}

// FindEconomyPlant retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEconomyPlant(ctx context.Context, exec boil.ContextExecutor, messageID int64, selectCols ...string) (*EconomyPlant, error) {
	economyPlantObj := &EconomyPlant{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"economy_plants\" where \"message_id\"=$1", sel,
	)

	q := queries.Raw(query, messageID)

	err := q.Bind(ctx, exec, economyPlantObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from economy_plants")
	}

	return economyPlantObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *EconomyPlant) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EconomyPlant) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no economy_plants provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(economyPlantColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	economyPlantInsertCacheMut.RLock()
	cache, cached := economyPlantInsertCache[key]
	economyPlantInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			economyPlantAllColumns,
			economyPlantColumnsWithDefault,
			economyPlantColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(economyPlantType, economyPlantMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(economyPlantType, economyPlantMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"economy_plants\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"economy_plants\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into economy_plants")
	}

	if !cached {
		economyPlantInsertCacheMut.Lock()
		economyPlantInsertCache[key] = cache
		economyPlantInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single EconomyPlant record using the global executor.
// See Update for more documentation.
func (o *EconomyPlant) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the EconomyPlant.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EconomyPlant) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	economyPlantUpdateCacheMut.RLock()
	cache, cached := economyPlantUpdateCache[key]
	economyPlantUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			economyPlantAllColumns,
			economyPlantPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update economy_plants, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"economy_plants\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, economyPlantPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(economyPlantType, economyPlantMapping, append(wl, economyPlantPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update economy_plants row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for economy_plants")
	}

	if !cached {
		economyPlantUpdateCacheMut.Lock()
		economyPlantUpdateCache[key] = cache
		economyPlantUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q economyPlantQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q economyPlantQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for economy_plants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for economy_plants")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EconomyPlantSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EconomyPlantSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyPlantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"economy_plants\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, economyPlantPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in economyPlant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all economyPlant")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *EconomyPlant) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EconomyPlant) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no economy_plants provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(economyPlantColumnsWithDefault, o)

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

	economyPlantUpsertCacheMut.RLock()
	cache, cached := economyPlantUpsertCache[key]
	economyPlantUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			economyPlantAllColumns,
			economyPlantColumnsWithDefault,
			economyPlantColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			economyPlantAllColumns,
			economyPlantPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert economy_plants, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(economyPlantPrimaryKeyColumns))
			copy(conflict, economyPlantPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"economy_plants\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(economyPlantType, economyPlantMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(economyPlantType, economyPlantMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
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
		return errors.Wrap(err, "models: unable to upsert economy_plants")
	}

	if !cached {
		economyPlantUpsertCacheMut.Lock()
		economyPlantUpsertCache[key] = cache
		economyPlantUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single EconomyPlant record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *EconomyPlant) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single EconomyPlant record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EconomyPlant) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EconomyPlant provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), economyPlantPrimaryKeyMapping)
	sql := "DELETE FROM \"economy_plants\" WHERE \"message_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from economy_plants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for economy_plants")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q economyPlantQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no economyPlantQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economy_plants")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_plants")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o EconomyPlantSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EconomyPlantSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyPlantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"economy_plants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyPlantPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economyPlant slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_plants")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *EconomyPlant) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no EconomyPlant provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EconomyPlant) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEconomyPlant(ctx, exec, o.MessageID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyPlantSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty EconomyPlantSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyPlantSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EconomyPlantSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyPlantPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"economy_plants\".* FROM \"economy_plants\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyPlantPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EconomyPlantSlice")
	}

	*o = slice

	return nil
}

// EconomyPlantExistsG checks if the EconomyPlant row exists.
func EconomyPlantExistsG(ctx context.Context, messageID int64) (bool, error) {
	return EconomyPlantExists(ctx, boil.GetContextDB(), messageID)
}

// EconomyPlantExists checks if the EconomyPlant row exists.
func EconomyPlantExists(ctx context.Context, exec boil.ContextExecutor, messageID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"economy_plants\" where \"message_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, messageID)
	}

	row := exec.QueryRowContext(ctx, sql, messageID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if economy_plants exists")
	}

	return exists, nil
}
