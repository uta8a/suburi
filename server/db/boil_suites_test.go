// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Userinfos", testUserinfos)
}

func TestDelete(t *testing.T) {
	t.Run("Userinfos", testUserinfosDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Userinfos", testUserinfosQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Userinfos", testUserinfosSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Userinfos", testUserinfosExists)
}

func TestFind(t *testing.T) {
	t.Run("Userinfos", testUserinfosFind)
}

func TestBind(t *testing.T) {
	t.Run("Userinfos", testUserinfosBind)
}

func TestOne(t *testing.T) {
	t.Run("Userinfos", testUserinfosOne)
}

func TestAll(t *testing.T) {
	t.Run("Userinfos", testUserinfosAll)
}

func TestCount(t *testing.T) {
	t.Run("Userinfos", testUserinfosCount)
}

func TestHooks(t *testing.T) {
	t.Run("Userinfos", testUserinfosHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Userinfos", testUserinfosInsert)
	t.Run("Userinfos", testUserinfosInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Userinfos", testUserinfosReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Userinfos", testUserinfosReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Userinfos", testUserinfosSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Userinfos", testUserinfosUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Userinfos", testUserinfosSliceUpdateAll)
}
