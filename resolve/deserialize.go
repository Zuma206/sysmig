package resolve

import (
	"encoding/json"

	"github.com/Shopify/go-lua"
	"github.com/zuma206/sysmig/utils"
)

func deserialize(jsonData string, state *lua.State) {
	var result any
	if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
		utils.HandleErr(err)
	}
	hydrate(result, state)
}

func hydrate(value any, state *lua.State) {
	switch typedValue := value.(type) {
	case string:
		state.PushString(typedValue)
	case int:
		state.PushInteger(typedValue)
	case float64:
		state.PushNumber(typedValue)
	case map[string]any:
		hydrateMap(typedValue, state)
	case []any:
		hydrateArray(typedValue, state)
	case nil:
		state.PushNil()
	}
}

func hydrateMap(hashmap map[string]any, state *lua.State) {
	state.CreateTable(0, len(hashmap))
	for key, value := range hashmap {
		hydrate(value, state)
		state.SetField(-2, key)
	}
}

func hydrateArray(array []any, state *lua.State) {
	state.CreateTable(len(array), 0)
	for i, value := range array {
		hydrate(value, state)
		state.RawSetInt(-2, i)
	}
}
