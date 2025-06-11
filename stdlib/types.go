package stdlib

import "github.com/Shopify/go-lua"

var (
	ltable = LuaArgType{"table", func(state *lua.State, index int) bool {
		return state.IsTable(index)
	}}
)
