package stdlib

import "github.com/Shopify/go-lua"

var System = LuaFunc{
	Name: "system",
	Args: []LuaArg{
		{"migrators list", ltable},
	},
	Body: func(state *lua.State) int {
		return 0
	},
}
