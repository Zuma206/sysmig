package resolve

import (
	"encoding/json"
	"errors"

	"github.com/Shopify/go-lua"
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
	deserialize(oldStateJson, state)
	state.Call(1, 1)
	return toResolution(state)
}

// Keys for the 'resolution' table returned by a system script
const (
	RESOLUTION_MIGRATION = "migration"
	RESOLUTION_SYNC      = "sync"
	RESOLUTION_STATE     = "next_state"
)

// Takes a resolution table on the lua stack and converts it into a resolution struct
// May panic!
func toResolution(state *lua.State) *Resolution {
	state.Field(-1, RESOLUTION_MIGRATION)
	migration, ok := state.ToString(-1)
	if !ok {
		err := errors.New("resolution table must include a string migration field")
		utils.HandleErr(err)
	}

	state.Pop(1)
	state.Field(-1, RESOLUTION_SYNC)
	sync, ok := state.ToString(-1)
	if !ok {
		err := errors.New("resolution table must include a string sync field")
		utils.HandleErr(err)
	}

	state.Pop(1)
	state.Field(-1, RESOLUTION_STATE)
	if state.IsNoneOrNil(-1) {
		err := errors.New("resolution table must include a next_state field")
		utils.HandleErr(err)
	}

	newState := serialize(state)
	newStateJson, err := json.Marshal(newState)
	if err != nil {
		utils.HandleErr(err)
	}

	return &Resolution{
		migrationScript: migration,
		syncScript:      sync,
		nextStateJson:   string(newStateJson),
	}
}
