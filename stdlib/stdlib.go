package stdlib

import "github.com/Shopify/go-lua"

// All global objects provided by the stdlib
var globals map[string]lua.Function

// Add the stdlib globals to a given lua runtime
func AddStdlib(state *lua.State) {
	for key, value := range globals {
		state.Register(key, value)
	}
}
