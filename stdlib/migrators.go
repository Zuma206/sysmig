package stdlib

import (
	"github.com/Shopify/go-lua"
)

var (
	MigratorName = LuaAttribute[string]{"name", lstring, "migrator"}
	MigratorFunc = LuaAttribute[lua.Function]{"func", lfunction, "migrator"}
)

var (
	ResolutionMigration = LuaAttribute[string]{"migration", lstring, "resolution"}
	ResolutionSync      = LuaAttribute[string]{"sync", lstring, "resolution"}
	ResolutionNextState = LuaAttribute[any]{"next_state", ltable, "resolution"}
)
