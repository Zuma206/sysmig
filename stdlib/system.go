package stdlib

import "github.com/Shopify/go-lua"

var System = LuaFunc{
	Name: "system",
	Args: []LuaArg{
		{"migrators", ltable},
	},
	Body: func(state *lua.State) int {
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
		migratorsIndex := lua.UpValueIndex(1)
		currentStateIndex := 1
		nMigrators := state.RawLength(migratorsIndex)
		pushSystemResolution(state, nMigrators)
		resolutionIndex := 2
		ResolutionNextState.Push(state, resolutionIndex)
		nextStateIndex := 3

		for i := range nMigrators {
			state.RawGetInt(migratorsIndex, i+1)
			MigratorFunc.Push(state, -1)
			MigratorName.Push(state, -2)
			state.Table(currentStateIndex)
			state.Call(1, 1)
			MigratorName.Push(state, -2)
			ResolutionNextState.Push(state, -1)
			state.SetTable(nextStateIndex)
		}

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

		state.Pop(1)
		return 1
	},
}

func pushSystemResolution(state *lua.State, nMigrators int) {
	state.CreateTable(0, 3)
	state.CreateTable(0, nMigrators)
	ResolutionNextState.Set(state, -2)
	state.PushString("")
	ResolutionMigration.Set(state, -2)
	state.PushString("")
	ResolutionSync.Set(state, -2)
}
