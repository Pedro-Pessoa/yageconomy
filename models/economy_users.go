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
	"github.com/volatiletech/sqlboiler/types"
)

// EconomyUser is an object representing the database table.
type EconomyUser struct {
	GuildID                 int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	UserID                  int64            `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	MoneyBank               int64            `boil:"money_bank" json:"money_bank" toml:"money_bank" yaml:"money_bank"`
	MoneyWallet             int64            `boil:"money_wallet" json:"money_wallet" toml:"money_wallet" yaml:"money_wallet"`
	LastDailyClaim          time.Time        `boil:"last_daily_claim" json:"last_daily_claim" toml:"last_daily_claim" yaml:"last_daily_claim"`
	LastChatmoneyClaim      time.Time        `boil:"last_chatmoney_claim" json:"last_chatmoney_claim" toml:"last_chatmoney_claim" yaml:"last_chatmoney_claim"`
	LastFishing             time.Time        `boil:"last_fishing" json:"last_fishing" toml:"last_fishing" yaml:"last_fishing"`
	WaifudBy                int64            `boil:"waifud_by" json:"waifud_by" toml:"waifud_by" yaml:"waifud_by"`
	Waifus                  types.Int64Array `boil:"waifus" json:"waifus,omitempty" toml:"waifus" yaml:"waifus,omitempty"`
	WaifuItemWorth          int64            `boil:"waifu_item_worth" json:"waifu_item_worth" toml:"waifu_item_worth" yaml:"waifu_item_worth"`
	WaifuLastClaimAmount    int64            `boil:"waifu_last_claim_amount" json:"waifu_last_claim_amount" toml:"waifu_last_claim_amount" yaml:"waifu_last_claim_amount"`
	WaifuExtraWorth         int64            `boil:"waifu_extra_worth" json:"waifu_extra_worth" toml:"waifu_extra_worth" yaml:"waifu_extra_worth"`
	WaifuAffinityTowards    int64            `boil:"waifu_affinity_towards" json:"waifu_affinity_towards" toml:"waifu_affinity_towards" yaml:"waifu_affinity_towards"`
	WaifuDivorces           int              `boil:"waifu_divorces" json:"waifu_divorces" toml:"waifu_divorces" yaml:"waifu_divorces"`
	WaifuAffinityChanges    int              `boil:"waifu_affinity_changes" json:"waifu_affinity_changes" toml:"waifu_affinity_changes" yaml:"waifu_affinity_changes"`
	FishCaugth              int64            `boil:"fish_caugth" json:"fish_caugth" toml:"fish_caugth" yaml:"fish_caugth"`
	GamblingBoostPercentage int              `boil:"gambling_boost_percentage" json:"gambling_boost_percentage" toml:"gambling_boost_percentage" yaml:"gambling_boost_percentage"`
	LastInterestUpdate      time.Time        `boil:"last_interest_update" json:"last_interest_update" toml:"last_interest_update" yaml:"last_interest_update"`
	LastRobAttempt          time.Time        `boil:"last_rob_attempt" json:"last_rob_attempt" toml:"last_rob_attempt" yaml:"last_rob_attempt"`
	LastFailedHeist         time.Time        `boil:"last_failed_heist" json:"last_failed_heist" toml:"last_failed_heist" yaml:"last_failed_heist"`

	R *economyUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L economyUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EconomyUserColumns = struct {
	GuildID                 string
	UserID                  string
	MoneyBank               string
	MoneyWallet             string
	LastDailyClaim          string
	LastChatmoneyClaim      string
	LastFishing             string
	WaifudBy                string
	Waifus                  string
	WaifuItemWorth          string
	WaifuLastClaimAmount    string
	WaifuExtraWorth         string
	WaifuAffinityTowards    string
	WaifuDivorces           string
	WaifuAffinityChanges    string
	FishCaugth              string
	GamblingBoostPercentage string
	LastInterestUpdate      string
	LastRobAttempt          string
	LastFailedHeist         string
}{
	GuildID:                 "guild_id",
	UserID:                  "user_id",
	MoneyBank:               "money_bank",
	MoneyWallet:             "money_wallet",
	LastDailyClaim:          "last_daily_claim",
	LastChatmoneyClaim:      "last_chatmoney_claim",
	LastFishing:             "last_fishing",
	WaifudBy:                "waifud_by",
	Waifus:                  "waifus",
	WaifuItemWorth:          "waifu_item_worth",
	WaifuLastClaimAmount:    "waifu_last_claim_amount",
	WaifuExtraWorth:         "waifu_extra_worth",
	WaifuAffinityTowards:    "waifu_affinity_towards",
	WaifuDivorces:           "waifu_divorces",
	WaifuAffinityChanges:    "waifu_affinity_changes",
	FishCaugth:              "fish_caugth",
	GamblingBoostPercentage: "gambling_boost_percentage",
	LastInterestUpdate:      "last_interest_update",
	LastRobAttempt:          "last_rob_attempt",
	LastFailedHeist:         "last_failed_heist",
}

// Generated where

var EconomyUserWhere = struct {
	GuildID                 whereHelperint64
	UserID                  whereHelperint64
	MoneyBank               whereHelperint64
	MoneyWallet             whereHelperint64
	LastDailyClaim          whereHelpertime_Time
	LastChatmoneyClaim      whereHelpertime_Time
	LastFishing             whereHelpertime_Time
	WaifudBy                whereHelperint64
	Waifus                  whereHelpertypes_Int64Array
	WaifuItemWorth          whereHelperint64
	WaifuLastClaimAmount    whereHelperint64
	WaifuExtraWorth         whereHelperint64
	WaifuAffinityTowards    whereHelperint64
	WaifuDivorces           whereHelperint
	WaifuAffinityChanges    whereHelperint
	FishCaugth              whereHelperint64
	GamblingBoostPercentage whereHelperint
	LastInterestUpdate      whereHelpertime_Time
	LastRobAttempt          whereHelpertime_Time
	LastFailedHeist         whereHelpertime_Time
}{
	GuildID:                 whereHelperint64{field: "\"economy_users\".\"guild_id\""},
	UserID:                  whereHelperint64{field: "\"economy_users\".\"user_id\""},
	MoneyBank:               whereHelperint64{field: "\"economy_users\".\"money_bank\""},
	MoneyWallet:             whereHelperint64{field: "\"economy_users\".\"money_wallet\""},
	LastDailyClaim:          whereHelpertime_Time{field: "\"economy_users\".\"last_daily_claim\""},
	LastChatmoneyClaim:      whereHelpertime_Time{field: "\"economy_users\".\"last_chatmoney_claim\""},
	LastFishing:             whereHelpertime_Time{field: "\"economy_users\".\"last_fishing\""},
	WaifudBy:                whereHelperint64{field: "\"economy_users\".\"waifud_by\""},
	Waifus:                  whereHelpertypes_Int64Array{field: "\"economy_users\".\"waifus\""},
	WaifuItemWorth:          whereHelperint64{field: "\"economy_users\".\"waifu_item_worth\""},
	WaifuLastClaimAmount:    whereHelperint64{field: "\"economy_users\".\"waifu_last_claim_amount\""},
	WaifuExtraWorth:         whereHelperint64{field: "\"economy_users\".\"waifu_extra_worth\""},
	WaifuAffinityTowards:    whereHelperint64{field: "\"economy_users\".\"waifu_affinity_towards\""},
	WaifuDivorces:           whereHelperint{field: "\"economy_users\".\"waifu_divorces\""},
	WaifuAffinityChanges:    whereHelperint{field: "\"economy_users\".\"waifu_affinity_changes\""},
	FishCaugth:              whereHelperint64{field: "\"economy_users\".\"fish_caugth\""},
	GamblingBoostPercentage: whereHelperint{field: "\"economy_users\".\"gambling_boost_percentage\""},
	LastInterestUpdate:      whereHelpertime_Time{field: "\"economy_users\".\"last_interest_update\""},
	LastRobAttempt:          whereHelpertime_Time{field: "\"economy_users\".\"last_rob_attempt\""},
	LastFailedHeist:         whereHelpertime_Time{field: "\"economy_users\".\"last_failed_heist\""},
}

// EconomyUserRels is where relationship names are stored.
var EconomyUserRels = struct {
}{}

// economyUserR is where relationships are stored.
type economyUserR struct {
}

// NewStruct creates a new relationship struct
func (*economyUserR) NewStruct() *economyUserR {
	return &economyUserR{}
}

// economyUserL is where Load methods for each relationship are stored.
type economyUserL struct{}

var (
	economyUserAllColumns            = []string{"guild_id", "user_id", "money_bank", "money_wallet", "last_daily_claim", "last_chatmoney_claim", "last_fishing", "waifud_by", "waifus", "waifu_item_worth", "waifu_last_claim_amount", "waifu_extra_worth", "waifu_affinity_towards", "waifu_divorces", "waifu_affinity_changes", "fish_caugth", "gambling_boost_percentage", "last_interest_update", "last_rob_attempt", "last_failed_heist"}
	economyUserColumnsWithoutDefault = []string{"guild_id", "user_id", "money_bank", "money_wallet", "last_daily_claim", "last_chatmoney_claim", "last_fishing", "waifud_by", "waifus", "waifu_item_worth", "waifu_last_claim_amount", "waifu_extra_worth", "waifu_affinity_towards", "waifu_divorces", "waifu_affinity_changes", "fish_caugth", "gambling_boost_percentage", "last_interest_update", "last_rob_attempt"}
	economyUserColumnsWithDefault    = []string{"last_failed_heist"}
	economyUserPrimaryKeyColumns     = []string{"guild_id", "user_id"}
)

type (
	// EconomyUserSlice is an alias for a slice of pointers to EconomyUser.
	// This should generally be used opposed to []EconomyUser.
	EconomyUserSlice []*EconomyUser

	economyUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	economyUserType                 = reflect.TypeOf(&EconomyUser{})
	economyUserMapping              = queries.MakeStructMapping(economyUserType)
	economyUserPrimaryKeyMapping, _ = queries.BindMapping(economyUserType, economyUserMapping, economyUserPrimaryKeyColumns)
	economyUserInsertCacheMut       sync.RWMutex
	economyUserInsertCache          = make(map[string]insertCache)
	economyUserUpdateCacheMut       sync.RWMutex
	economyUserUpdateCache          = make(map[string]updateCache)
	economyUserUpsertCacheMut       sync.RWMutex
	economyUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single economyUser record from the query using the global executor.
func (q economyUserQuery) OneG(ctx context.Context) (*EconomyUser, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single economyUser record from the query.
func (q economyUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EconomyUser, error) {
	o := &EconomyUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for economy_users")
	}

	return o, nil
}

// AllG returns all EconomyUser records from the query using the global executor.
func (q economyUserQuery) AllG(ctx context.Context) (EconomyUserSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all EconomyUser records from the query.
func (q economyUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (EconomyUserSlice, error) {
	var o []*EconomyUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EconomyUser slice")
	}

	return o, nil
}

// CountG returns the count of all EconomyUser records in the query, and panics on error.
func (q economyUserQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all EconomyUser records in the query.
func (q economyUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count economy_users rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q economyUserQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q economyUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if economy_users exists")
	}

	return count > 0, nil
}

// EconomyUsers retrieves all the records using an executor.
func EconomyUsers(mods ...qm.QueryMod) economyUserQuery {
	mods = append(mods, qm.From("\"economy_users\""))
	return economyUserQuery{NewQuery(mods...)}
}

// FindEconomyUserG retrieves a single record by ID.
func FindEconomyUserG(ctx context.Context, guildID int64, userID int64, selectCols ...string) (*EconomyUser, error) {
	return FindEconomyUser(ctx, boil.GetContextDB(), guildID, userID, selectCols...)
}

// FindEconomyUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEconomyUser(ctx context.Context, exec boil.ContextExecutor, guildID int64, userID int64, selectCols ...string) (*EconomyUser, error) {
	economyUserObj := &EconomyUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"economy_users\" where \"guild_id\"=$1 AND \"user_id\"=$2", sel,
	)

	q := queries.Raw(query, guildID, userID)

	err := q.Bind(ctx, exec, economyUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from economy_users")
	}

	return economyUserObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *EconomyUser) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EconomyUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no economy_users provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(economyUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	economyUserInsertCacheMut.RLock()
	cache, cached := economyUserInsertCache[key]
	economyUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			economyUserAllColumns,
			economyUserColumnsWithDefault,
			economyUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(economyUserType, economyUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(economyUserType, economyUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"economy_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"economy_users\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into economy_users")
	}

	if !cached {
		economyUserInsertCacheMut.Lock()
		economyUserInsertCache[key] = cache
		economyUserInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single EconomyUser record using the global executor.
// See Update for more documentation.
func (o *EconomyUser) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the EconomyUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EconomyUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	economyUserUpdateCacheMut.RLock()
	cache, cached := economyUserUpdateCache[key]
	economyUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			economyUserAllColumns,
			economyUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update economy_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"economy_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, economyUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(economyUserType, economyUserMapping, append(wl, economyUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update economy_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for economy_users")
	}

	if !cached {
		economyUserUpdateCacheMut.Lock()
		economyUserUpdateCache[key] = cache
		economyUserUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q economyUserQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q economyUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for economy_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for economy_users")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EconomyUserSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EconomyUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"economy_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, economyUserPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in economyUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all economyUser")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *EconomyUser) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EconomyUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no economy_users provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(economyUserColumnsWithDefault, o)

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

	economyUserUpsertCacheMut.RLock()
	cache, cached := economyUserUpsertCache[key]
	economyUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			economyUserAllColumns,
			economyUserColumnsWithDefault,
			economyUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			economyUserAllColumns,
			economyUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert economy_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(economyUserPrimaryKeyColumns))
			copy(conflict, economyUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"economy_users\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(economyUserType, economyUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(economyUserType, economyUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert economy_users")
	}

	if !cached {
		economyUserUpsertCacheMut.Lock()
		economyUserUpsertCache[key] = cache
		economyUserUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single EconomyUser record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *EconomyUser) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single EconomyUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EconomyUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EconomyUser provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), economyUserPrimaryKeyMapping)
	sql := "DELETE FROM \"economy_users\" WHERE \"guild_id\"=$1 AND \"user_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from economy_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for economy_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q economyUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no economyUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economy_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_users")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o EconomyUserSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EconomyUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"economy_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyUserPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economyUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_users")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *EconomyUser) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no EconomyUser provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EconomyUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEconomyUser(ctx, exec, o.GuildID, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyUserSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty EconomyUserSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EconomyUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"economy_users\".* FROM \"economy_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EconomyUserSlice")
	}

	*o = slice

	return nil
}

// EconomyUserExistsG checks if the EconomyUser row exists.
func EconomyUserExistsG(ctx context.Context, guildID int64, userID int64) (bool, error) {
	return EconomyUserExists(ctx, boil.GetContextDB(), guildID, userID)
}

// EconomyUserExists checks if the EconomyUser row exists.
func EconomyUserExists(ctx context.Context, exec boil.ContextExecutor, guildID int64, userID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"economy_users\" where \"guild_id\"=$1 AND \"user_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID, userID)
	}

	row := exec.QueryRowContext(ctx, sql, guildID, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if economy_users exists")
	}

	return exists, nil
}
