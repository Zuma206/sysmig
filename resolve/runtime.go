package resolve

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/Shopify/go-lua"
	"github.com/zuma206/sysmig/std"
	"github.com/zuma206/sysmig/utils"
)

// Represents the 'resolution' of a system script:
// migration script + sync script + next state
type Resolution struct {
	migration string
	sync      string
	nextState []byte
}

// Migrator keys
const (
	function = "func"
)

// Takes json for the previous state, and passes it through a LUA VM
// to resolve it to a resolution struct
func resolve(oldStateJson string) *Resolution {
	println("Attempting to evaluate", flags.configPath)
	state := lua.NewState()
	openLibraries(state)
	utils.HandleErr(lua.DoFile(state, flags.configPath))
	state.Field(-1, function)
	deserialize(oldStateJson, state)
	state.Call(1, 1)
	return getResolution(state, -1)
}

// Resolution keys
const (
	nextState = "next_state"
	migration = "migration"
	sync      = "sync"
)

// Takes a resolution table `index` and converts it into a resolution struct
// May panic!
func getResolution(state *lua.State, index int) *Resolution {
	if !state.IsTable(index) {
		err := errors.New("migrator did not return a resolution table")
		utils.HandleErr(err)
	}
	state.Field(index, nextState)
	nextStateSerialized := std.Serialize(state)
	state.Pop(1)
	nextStateJson, err := json.Marshal(nextStateSerialized)
	utils.HandleErr(err)

	return &Resolution{
		migration: getStringKey(state, index, migration),
		sync:      getStringKey(state, index, sync),
		nextState: nextStateJson,
	}
}

// Opens the sysmig native stdlib and any required lua libraries
func openLibraries(state *lua.State) {
	configDir := path.Dir(flags.configPath)
	lua.OpenLibraries(state)
	std.OpenStd(state, configDir)
	patchPackagePath(state, configDir)
}

// Patches the package.path variable to look for packages in the directory
// of the root config file, rather than the current directory
func patchPackagePath(state *lua.State, configDir string) {
	localPath := path.Join(configDir, "?.lua;")
	stdPath := path.Join(path.Join(utils.GetDir(), "?.lua;"))
	state.Global("package")
	state.PushString(strings.Join([]string{localPath, stdPath}, ""))
	state.SetField(-2, "path")
	state.Pop(1)
}

// Get a string field from a table, erroring if it's not a string
func getStringKey(state *lua.State, index int, key string) string {
	state.Field(index, key)
	value, ok := state.ToString(-1)
	if !ok {
		utils.HandleErr(fmt.Errorf("cannot get key %s", key))
	}
	state.Pop(1)
	return value
}
