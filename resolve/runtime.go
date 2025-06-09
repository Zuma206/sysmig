package resolve

import (
	"encoding/json"
	"errors"

	"github.com/Shopify/go-lua"
	"github.com/zuma206/sysmig/utils"
)

type Resolution struct {
	MigrationScript string
	SyncScript      string
	NewStateJson    string
}

func resolve(configPath string) *Resolution {
	state := lua.NewState()
	utils.HandleErr(lua.DoFile(state, configPath))
	state.PushString("old_state_value")
	state.Call(1, 1)
	return toResolution(state)
}

const (
	RESOLUTION_MIGRATION = "migration"
	RESOLUTION_SYNC      = "sync"
	RESOLUTION_STATE     = "next_state"
)

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
	newState := serialize(state)
	newStateJson, err := json.Marshal(newState)
	if err != nil {
		utils.HandleErr(err)
	}

	return &Resolution{
		MigrationScript: migration,
		SyncScript:      sync,
		NewStateJson:    string(newStateJson),
	}
}
