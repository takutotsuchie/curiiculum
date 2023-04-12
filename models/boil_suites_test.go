// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Labels", testLabels)
	t.Run("TaskLabelRelations", testTaskLabelRelations)
	t.Run("Tasks", testTasks)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Labels", testLabelsDelete)
	t.Run("TaskLabelRelations", testTaskLabelRelationsDelete)
	t.Run("Tasks", testTasksDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Labels", testLabelsQueryDeleteAll)
	t.Run("TaskLabelRelations", testTaskLabelRelationsQueryDeleteAll)
	t.Run("Tasks", testTasksQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Labels", testLabelsSliceDeleteAll)
	t.Run("TaskLabelRelations", testTaskLabelRelationsSliceDeleteAll)
	t.Run("Tasks", testTasksSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Labels", testLabelsExists)
	t.Run("TaskLabelRelations", testTaskLabelRelationsExists)
	t.Run("Tasks", testTasksExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Labels", testLabelsFind)
	t.Run("TaskLabelRelations", testTaskLabelRelationsFind)
	t.Run("Tasks", testTasksFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Labels", testLabelsBind)
	t.Run("TaskLabelRelations", testTaskLabelRelationsBind)
	t.Run("Tasks", testTasksBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Labels", testLabelsOne)
	t.Run("TaskLabelRelations", testTaskLabelRelationsOne)
	t.Run("Tasks", testTasksOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Labels", testLabelsAll)
	t.Run("TaskLabelRelations", testTaskLabelRelationsAll)
	t.Run("Tasks", testTasksAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Labels", testLabelsCount)
	t.Run("TaskLabelRelations", testTaskLabelRelationsCount)
	t.Run("Tasks", testTasksCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Labels", testLabelsHooks)
	t.Run("TaskLabelRelations", testTaskLabelRelationsHooks)
	t.Run("Tasks", testTasksHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Labels", testLabelsInsert)
	t.Run("Labels", testLabelsInsertWhitelist)
	t.Run("TaskLabelRelations", testTaskLabelRelationsInsert)
	t.Run("TaskLabelRelations", testTaskLabelRelationsInsertWhitelist)
	t.Run("Tasks", testTasksInsert)
	t.Run("Tasks", testTasksInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("TaskLabelRelationToLabelUsingLabel", testTaskLabelRelationToOneLabelUsingLabel)
	t.Run("TaskLabelRelationToTaskUsingTask", testTaskLabelRelationToOneTaskUsingTask)
	t.Run("TaskToUserUsingUser", testTaskToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("LabelToTaskLabelRelations", testLabelToManyTaskLabelRelations)
	t.Run("TaskToTaskLabelRelations", testTaskToManyTaskLabelRelations)
	t.Run("UserToTasks", testUserToManyTasks)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("TaskLabelRelationToLabelUsingTaskLabelRelations", testTaskLabelRelationToOneSetOpLabelUsingLabel)
	t.Run("TaskLabelRelationToTaskUsingTaskLabelRelations", testTaskLabelRelationToOneSetOpTaskUsingTask)
	t.Run("TaskToUserUsingTasks", testTaskToOneSetOpUserUsingUser)
}

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
func TestToManyAdd(t *testing.T) {
	t.Run("LabelToTaskLabelRelations", testLabelToManyAddOpTaskLabelRelations)
	t.Run("TaskToTaskLabelRelations", testTaskToManyAddOpTaskLabelRelations)
	t.Run("UserToTasks", testUserToManyAddOpTasks)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Labels", testLabelsReload)
	t.Run("TaskLabelRelations", testTaskLabelRelationsReload)
	t.Run("Tasks", testTasksReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Labels", testLabelsReloadAll)
	t.Run("TaskLabelRelations", testTaskLabelRelationsReloadAll)
	t.Run("Tasks", testTasksReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Labels", testLabelsSelect)
	t.Run("TaskLabelRelations", testTaskLabelRelationsSelect)
	t.Run("Tasks", testTasksSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Labels", testLabelsUpdate)
	t.Run("TaskLabelRelations", testTaskLabelRelationsUpdate)
	t.Run("Tasks", testTasksUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Labels", testLabelsSliceUpdateAll)
	t.Run("TaskLabelRelations", testTaskLabelRelationsSliceUpdateAll)
	t.Run("Tasks", testTasksSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
