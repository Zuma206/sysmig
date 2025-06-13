package stdlib

import "github.com/Shopify/go-lua"

func LuaUnion[A any, B any](typeA LuaType[A], typeB LuaType[B]) LuaType[any] {
	return LuaType[any]{
		Name: typeA.Name + "|" + typeB.Name,
		Getter: func(state *lua.State, index int) (any, bool) {
			a, aOk := typeA.Getter(state, index)
			if aOk {
				return a, aOk
			}
			b, bOk := typeB.Getter(state, index)
			if bOk {
				return b, bOk
			}
			return a, false
		},
	}
}

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
	lnil = LuaType[any]{"nil", func(state *lua.State, index int) (any, bool) {
		return nil, state.IsNil(index)
	}}
)
