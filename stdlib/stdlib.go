package stdlib

import "github.com/Shopify/go-lua"

// Opens all standard library functions/tables into the global scope
func OpenLibraries(state *lua.State) {
	state.PushGlobalTable()
	System.Open(state)
	state.Pop(1)
}
