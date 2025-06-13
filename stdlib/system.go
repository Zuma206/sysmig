package stdlib

import "github.com/Shopify/go-lua"

// Represents the global 'system' function
// Is used to merge migrators
var System = LuaFunc{
	Name: "system",
	Args: []LuaArg{
		// A sequence of migrators to merge
		{"migrators", ltable},
	},
	Body: func(state *lua.State) int {
		// Create a migrator, with name/closure, and return it
		state.CreateTable(0, 2)
		state.PushString("system")
		MigratorName.Set(state, -2)
		state.PushValue(-2)
		state.PushGoClosure(systemMigratorFunc.GoFunction, 1)
		MigratorFunc.Set(state, -2)
		return 1
	},
}

var systemMigratorFunc = LuaFunc{
	Name: "system_migrator.func",
	Args: []LuaArg{
		{"current_state", LuaUnion(ltable, lnil)},
	},
	Body: func(state *lua.State) int {
		// Establish some positions/data in the stack
		migratorsIndex := lua.UpValueIndex(1)
		currentStateIndex := 1
		nMigrators := state.RawLength(migratorsIndex)
		pushSystemResolution(state, nMigrators)
		resolutionIndex := 2
		ResolutionNextState.Push(state, resolutionIndex)
		nextStateIndex := 3

		runMigrations(state, nMigrators, migratorsIndex, currentStateIndex, nextStateIndex)
		collapseMigrations(state, nMigrators, resolutionIndex)

		state.Pop(1)
		return 1
	},
}

// Create a blank resolution table
func pushSystemResolution(state *lua.State, nMigrators int) {
	state.CreateTable(0, 3)
	state.CreateTable(0, nMigrators)
	ResolutionNextState.Set(state, -2)
	state.PushString("")
	ResolutionMigration.Set(state, -2)
	state.PushString("")
	ResolutionSync.Set(state, -2)
}

// Run all system migrations in order, getting their current state
// and saving the new state
func runMigrations(
	state *lua.State,
	nMigrators int,
	migratorsIndex int,
	currentStateIndex int,
	nextStateIndex int,
) {
	for i := range nMigrators {
		state.RawGetInt(migratorsIndex, i+1)
		MigratorFunc.Push(state, -1)
		if state.IsNil(currentStateIndex) {
			state.PushNil()
		} else {
			MigratorName.Push(state, -2)
			state.Table(currentStateIndex)
		}
		state.Call(1, 1)
		MigratorName.Push(state, -2)
		ResolutionNextState.Push(state, -2)
		state.SetTable(nextStateIndex)
	}
}

// Go back through the migration results and merge the migration and sync scripts
func collapseMigrations(state *lua.State, nMigrators int, resolutionIndex int) {
	for range nMigrators {
		ResolutionMigration.Push(state, -1)
		ResolutionMigration.Push(state, resolutionIndex)
		state.Concat(2)
		ResolutionMigration.Set(state, resolutionIndex)
		ResolutionSync.Push(state, -1)
		ResolutionSync.Push(state, resolutionIndex)
		state.Concat(2)
		ResolutionSync.Set(state, resolutionIndex)
		state.Pop(2)
	}
}
