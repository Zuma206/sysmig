package stdlib

import "github.com/Shopify/go-lua"

var (
	ltable = LuaType[any]{"table", func(state *lua.State, index int) (any, bool) {
		value := state.ToValue(index)
		ok := state.IsTable(index)
		return value, ok
	}}
	lstring = LuaType[string]{"string", func(state *lua.State, index int) (string, bool) {
		return state.ToString(index)
	}}
	lfunction = LuaType[lua.Function]{"function", func(state *lua.State, index int) (lua.Function, bool) {
		if state.IsGoFunction(index) {
			return state.ToGoFunction(index), true
		}
		return nil, state.IsFunction(index)
	}}
)
