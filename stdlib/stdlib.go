package stdlib

import "github.com/Shopify/go-lua"

func OpenLibraries(state *lua.State) {
	state.PushGlobalTable()
	System.Open(state)
	state.Pop(1)
}
