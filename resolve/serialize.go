package resolve

import (
	"errors"
	"strconv"

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

func serializeTable(state *lua.State) any {
	array := []any{}
	state.PushNil()
	for state.Next(-2) {
		value := serialize(state)
		state.Pop(1)
		if !state.IsNumber(-1) {
			return serializeAsHashtable(state, array, value)
		}
		array = append(array, value)
	}
	return array
}

func serializeAsHashtable(state *lua.State, array []any, value any) map[string]any {
	hashmap := map[string]any{getKey(state): value}
	for i, value := range array {
		hashmap[strconv.Itoa(i)] = value
	}
	for state.Next(-2) {
		value := serialize(state)
		state.Pop(1)
		hashmap[getKey(state)] = value
	}
	return hashmap
}

func getKey(state *lua.State) string {
	key, ok := state.ToString(-1)
	if !ok {
		err := errors.New("mixed-key table keys must be able to be serialized as strings")
		utils.HandleErr(err)
	}
	return key
}
