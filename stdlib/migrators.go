package stdlib

import (
	"github.com/Shopify/go-lua"
)

// Represents attributes on a 'migrator' table
var (
	MigratorName = LuaAttribute[string]{"name", lstring, "migrator"}
	MigratorFunc = LuaAttribute[lua.Function]{"func", lfunction, "migrator"}
)

// Represents attributes on a 'resolution' table
var (
	ResolutionMigration = LuaAttribute[string]{"migration", lstring, "resolution"}
	ResolutionSync      = LuaAttribute[string]{"sync", lstring, "resolution"}
	ResolutionNextState = LuaAttribute[any]{"next_state", ltable, "resolution"}
)
