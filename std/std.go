package std

import (
	_ "embed"

	"github.com/Shopify/go-lua"
	"github.com/zuma206/sysmig/utils"
)

//go:embed std.lua
var std string

//go:embed migrator.lua
var migrator string

//go:embed system.lua
var system string

func OpenStd(state *lua.State) {
	require(state, "@std.migrator", migrator)
	require(state, "@std.system", system)
	require(state, "@std", std)
}

func require(state *lua.State, name string, code string) {
	lua.Require(state, name, func(state *lua.State) int {
		err := lua.LoadString(state, code)
		utils.HandleErr(err)
		state.Call(0, 1)
		return 1
	}, false)
}
