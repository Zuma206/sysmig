package resolve

import (
	"errors"
	"strconv"

	"github.com/Shopify/go-lua"
	"github.com/zuma206/sysmig/utils"
)

// Takes state on the lua stack and converts it to a go variable
func serialize(state *lua.State) any {
	switch state.TypeOf(-1) {
	case lua.TypeTable:
		return serializeTable(state)
	default:
		return state.ToValue(-1)
	}
}

// Takes a table on the lua stack and converts it to an array or a map
func serializeTable(state *lua.State) any {
	// Start as assuming an array
	array := []any{}
	state.PushNil()
	for state.Next(-2) {
		value := serialize(state)
		state.Pop(1)
		// If the key is not a number, we need to upgrade this to a map!
		if !state.IsNumber(-1) {
			return serializeAsHashtable(state, array, value)
		}
		array = append(array, value)
	}
	return array
}

// Upgrade an array that was being serialized from the stack to a hashmap
func serializeAsHashtable(state *lua.State, array []any, value any) map[string]any {
	// Take all the previous values and put them in a map
	// To be JSON compatible, cast numbers to strings
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

// Convert the value on the stack (the key) and convert it to a string
// Panics with a mixed-key error if the key cannot be converted to a  string
func getKey(state *lua.State) string {
	key, ok := state.ToString(-1)
	if !ok {
		err := errors.New("mixed-key table keys must be able to be serialized as strings")
		utils.HandleErr(err)
	}
	return key
}
