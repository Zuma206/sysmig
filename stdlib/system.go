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
		collapseAttribute(state, &ResolutionMigration, resolutionIndex)
		collapseAttribute(state, &ResolutionSync, resolutionIndex)
		// Cleanup artifacts from runMigrations
		state.Pop(2)
	}
}

// Concats (with a newline) a given attribute from the table at -1 and `tableIndex`, onto `tableIndex`
func collapseAttribute(state *lua.State, attribute *LuaAttribute[string], tableIndex int) {
	attribute.Push(state, -1)
	state.PushString("\n")
	attribute.Push(state, tableIndex)
	state.Concat(3)
	attribute.Set(state, tableIndex)
}
