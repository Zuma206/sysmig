package resolve

import (
	"errors"

	"github.com/Shopify/go-lua"
	"github.com/zuma206/sysmig/utils"
)

func serialize(state *lua.State) any {
	switch state.TypeOf(-1) {
	case lua.TypeTable:
		return serializeTable(state)
	default:
		return state.ToValue(-1)
	}
}

func serializeTable(state *lua.State) map[string]any {
	table := make(map[string]any)
	state.PushNil()
	for state.Next(-2) {
		value := serialize(state)
		state.Pop(1)
		key, ok := state.ToString(-1)
		if !ok {
			err := errors.New("cannot serialize state tables with non-string keys")
			utils.HandleErr(err)
		}
		table[key] = value
	}
	return table
}
