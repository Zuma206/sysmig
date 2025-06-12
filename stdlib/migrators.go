package stdlib

import (
	"github.com/Shopify/go-lua"
)

// Pushes the migrator function for the migrator `index` onto the stack
func PushMigratorFunc(state *lua.State, index int) {
	state.RawGetInt(index, 2)
}

// Pushes a field of the resolution table `index` onto the stack
var (
	PushResolutionMigration = DefinePusher("migration")
	PushResolutionSync      = DefinePusher("sync")
	PushResolutionNextState = DefinePusher("next_state")
)

// Gets a field of the resolution table `index`
var (
	GetResolutionMigration = DefineGetter(
		PushResolutionMigration, lstring,
		"invalid resolution table: migration must be a string",
	)
	GetResolutionSync = DefineGetter(
		PushResolutionSync, lstring,
		"invalid resolution table: sync must be a string",
	)
)
