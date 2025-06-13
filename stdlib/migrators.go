package stdlib

import (
	"github.com/Shopify/go-lua"
)

var (
	MigratorName = LuaAttribute[string]{"name", lstring, "migrator"}
	MigratorFunc = LuaAttribute[lua.Function]{"func", lfunction, "migrator"}
)

func PushMigrator(state *lua.State, name string, function lua.Function, upValues uint8) {
	state.CreateTable(0, 2)
	state.PushString(name)
	MigratorName.Set(state, -2)
	state.PushGoClosure(function, upValues)
	MigratorFunc.Set(state, -2)
}

var (
	ResolutionMigration = LuaAttribute[string]{"migration", lstring, "resolution"}
	ResolutionSync      = LuaAttribute[string]{"sync", lstring, "resolution"}
	ResolutionNextState = LuaAttribute[any]{"next_state", ltable, "resolution"}
)
