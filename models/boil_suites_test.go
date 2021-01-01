// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Users", testUsers)
	t.Run("UserLinks", testUserLinks)
}

func TestDelete(t *testing.T) {
	t.Run("Users", testUsersDelete)
	t.Run("UserLinks", testUserLinksDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("UserLinks", testUserLinksQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("UserLinks", testUserLinksSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Users", testUsersExists)
	t.Run("UserLinks", testUserLinksExists)
}

func TestFind(t *testing.T) {
	t.Run("Users", testUsersFind)
	t.Run("UserLinks", testUserLinksFind)
}

func TestBind(t *testing.T) {
	t.Run("Users", testUsersBind)
	t.Run("UserLinks", testUserLinksBind)
}

func TestOne(t *testing.T) {
	t.Run("Users", testUsersOne)
	t.Run("UserLinks", testUserLinksOne)
}

func TestAll(t *testing.T) {
	t.Run("Users", testUsersAll)
	t.Run("UserLinks", testUserLinksAll)
}

func TestCount(t *testing.T) {
	t.Run("Users", testUsersCount)
	t.Run("UserLinks", testUserLinksCount)
}

func TestHooks(t *testing.T) {
	t.Run("Users", testUsersHooks)
	t.Run("UserLinks", testUserLinksHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("UserLinks", testUserLinksInsert)
	t.Run("UserLinks", testUserLinksInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("UserLinkToUserUsingUser", testUserLinkToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("UserToUserLinks", testUserToManyUserLinks)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("UserLinkToUserUsingUser", testUserLinkToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("UserLinkToUserUsingUser", testUserLinkToOneRemoveOpUserUsingUser)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("UserToUserLinks", testUserToManyAddOpUserLinks)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("UserToUserLinks", testUserToManySetOpUserLinks)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("UserToUserLinks", testUserToManyRemoveOpUserLinks)
}

func TestReload(t *testing.T) {
	t.Run("Users", testUsersReload)
	t.Run("UserLinks", testUserLinksReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Users", testUsersReloadAll)
	t.Run("UserLinks", testUserLinksReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Users", testUsersSelect)
	t.Run("UserLinks", testUserLinksSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Users", testUsersUpdate)
	t.Run("UserLinks", testUserLinksUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("UserLinks", testUserLinksSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Users", testUsersUpsert)
	t.Run("UserLinks", testUserLinksUpsert)
}
