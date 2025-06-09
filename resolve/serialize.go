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

func isArray(state *lua.State) bool {
	state.PushNil()
	if !state.Next(-2) {
		return true
	}
	isArray := state.IsNumber(-2)
	state.Pop(2)
	return isArray
}

func serializeTable(state *lua.State) any {
	if isArray(state) {
		return serializeArrayTable(state)
	} else {
		return serializeHashTable(state)
	}
}

func serializeArrayTable(state *lua.State) []any {
	length := getLength(state)
	array := make([]any, length)
	for i := range length {
		state.RawGetInt(-1, i+1)
		value := serialize(state)
		array[i] = value
		state.Pop(1)
	}
	return array
}

func getLength(state *lua.State) int {
	state.Length(-1)
	length, ok := state.ToInteger(-1)
	if !ok {
		err := errors.New("cannot get length of non-table")
		utils.HandleErr(err)
	}
	state.Pop(1)
	return length
}

func serializeHashTable(state *lua.State) map[string]any {
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
	state.Pop(1)
	return table
}
