// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Apartments", testApartments)
	t.Run("Buildings", testBuildings)
}

func TestDelete(t *testing.T) {
	t.Run("Apartments", testApartmentsDelete)
	t.Run("Buildings", testBuildingsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Apartments", testApartmentsQueryDeleteAll)
	t.Run("Buildings", testBuildingsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Apartments", testApartmentsSliceDeleteAll)
	t.Run("Buildings", testBuildingsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Apartments", testApartmentsExists)
	t.Run("Buildings", testBuildingsExists)
}

func TestFind(t *testing.T) {
	t.Run("Apartments", testApartmentsFind)
	t.Run("Buildings", testBuildingsFind)
}

func TestBind(t *testing.T) {
	t.Run("Apartments", testApartmentsBind)
	t.Run("Buildings", testBuildingsBind)
}

func TestOne(t *testing.T) {
	t.Run("Apartments", testApartmentsOne)
	t.Run("Buildings", testBuildingsOne)
}

func TestAll(t *testing.T) {
	t.Run("Apartments", testApartmentsAll)
	t.Run("Buildings", testBuildingsAll)
}

func TestCount(t *testing.T) {
	t.Run("Apartments", testApartmentsCount)
	t.Run("Buildings", testBuildingsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Apartments", testApartmentsHooks)
	t.Run("Buildings", testBuildingsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Apartments", testApartmentsInsert)
	t.Run("Apartments", testApartmentsInsertWhitelist)
	t.Run("Buildings", testBuildingsInsert)
	t.Run("Buildings", testBuildingsInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Apartments", testApartmentsReload)
	t.Run("Buildings", testBuildingsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Apartments", testApartmentsReloadAll)
	t.Run("Buildings", testBuildingsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Apartments", testApartmentsSelect)
	t.Run("Buildings", testBuildingsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Apartments", testApartmentsUpdate)
	t.Run("Buildings", testBuildingsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Apartments", testApartmentsSliceUpdateAll)
	t.Run("Buildings", testBuildingsSliceUpdateAll)
}
