package resolve

import (
	"encoding/json"

	"github.com/Shopify/go-lua"
	"github.com/zuma206/sysmig/stdlib"
	"github.com/zuma206/sysmig/utils"
)

// Represents the 'resolution' of a system script:
// migration script + sync script + next state
type Resolution struct {
	migrationScript string
	syncScript      string
	nextStateJson   string
}

// Takes json for the previous state, and passes it through a LUA VM
// to resolve it to a resolution struct
func resolve(oldStateJson string) *Resolution {
	state := lua.NewState()
	utils.HandleErr(lua.DoFile(state, flags.configPath))
	stdlib.MigratorFunc.Push(state, -1)
	deserialize(oldStateJson, state)
	state.Call(1, 1)
	return getResolution(state, -1)
}

// Takes a resolution table `index` and converts it into a resolution struct
// May panic!
func getResolution(state *lua.State, index int) *Resolution {
	stdlib.ResolutionNextState.Push(state, index)
	nextState := serialize(state)
	state.Pop(1)
	nextStateJson, err := json.Marshal(nextState)
	utils.HandleErr(err)

	return &Resolution{
		migrationScript: stdlib.ResolutionMigration.Get(state, index),
		syncScript:      stdlib.ResolutionSync.Get(state, index),
		nextStateJson:   string(nextStateJson),
	}
}
