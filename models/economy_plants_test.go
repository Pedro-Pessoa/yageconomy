// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testEconomyPlants(t *testing.T) {
	t.Parallel()

	query := EconomyPlants()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEconomyPlantsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
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

	count, err := EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEconomyPlantsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EconomyPlants().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEconomyPlantsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EconomyPlantSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEconomyPlantsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EconomyPlantExists(ctx, tx, o.ChannelID)
	if err != nil {
		t.Errorf("Unable to check if EconomyPlant exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EconomyPlantExists to return true, but got false.")
	}
}

func testEconomyPlantsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	economyPlantFound, err := FindEconomyPlant(ctx, tx, o.ChannelID)
	if err != nil {
		t.Error(err)
	}

	if economyPlantFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEconomyPlantsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EconomyPlants().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEconomyPlantsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EconomyPlants().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEconomyPlantsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	economyPlantOne := &EconomyPlant{}
	economyPlantTwo := &EconomyPlant{}
	if err = randomize.Struct(seed, economyPlantOne, economyPlantDBTypes, false, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}
	if err = randomize.Struct(seed, economyPlantTwo, economyPlantDBTypes, false, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = economyPlantOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = economyPlantTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EconomyPlants().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEconomyPlantsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	economyPlantOne := &EconomyPlant{}
	economyPlantTwo := &EconomyPlant{}
	if err = randomize.Struct(seed, economyPlantOne, economyPlantDBTypes, false, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}
	if err = randomize.Struct(seed, economyPlantTwo, economyPlantDBTypes, false, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = economyPlantOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = economyPlantTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testEconomyPlantsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEconomyPlantsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(economyPlantColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEconomyPlantsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
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

func testEconomyPlantsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EconomyPlantSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEconomyPlantsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EconomyPlants().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	economyPlantDBTypes = map[string]string{`ChannelID`: `bigint`, `GuildID`: `bigint`, `MessageID`: `bigint`, `AuthorID`: `bigint`, `Amount`: `bigint`, `Password`: `text`}
	_                   = bytes.MinRead
)

func testEconomyPlantsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(economyPlantPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(economyPlantColumns) == len(economyPlantPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEconomyPlantsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(economyPlantColumns) == len(economyPlantPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EconomyPlant{}
	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, economyPlantDBTypes, true, economyPlantPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(economyPlantColumns, economyPlantPrimaryKeyColumns) {
		fields = economyPlantColumns
	} else {
		fields = strmangle.SetComplement(
			economyPlantColumns,
			economyPlantPrimaryKeyColumns,
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

	slice := EconomyPlantSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEconomyPlantsUpsert(t *testing.T) {
	t.Parallel()

	if len(economyPlantColumns) == len(economyPlantPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EconomyPlant{}
	if err = randomize.Struct(seed, &o, economyPlantDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EconomyPlant: %s", err)
	}

	count, err := EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, economyPlantDBTypes, false, economyPlantPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EconomyPlant struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EconomyPlant: %s", err)
	}

	count, err = EconomyPlants().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}