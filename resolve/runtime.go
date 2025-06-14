package resolve

import (
	"encoding/json"
	"errors"
	"path"

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
	openLibraries(state)
	utils.HandleErr(lua.DoFile(state, flags.configPath))
	stdlib.MigratorFunc.Push(state, -1)
	deserialize(oldStateJson, state)
	state.Call(1, 1)
	return getResolution(state, -1)
}

// Takes a resolution table `index` and converts it into a resolution struct
// May panic!
func getResolution(state *lua.State, index int) *Resolution {
	if !state.IsTable(index) {
		err := errors.New("migrator did not return a resolution table")
		utils.HandleErr(err)
	}
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

// Opens the sysmig native stdlib and any required lua libraries
func openLibraries(state *lua.State) {
	stdlib.OpenLibraries(state)
	lua.Require(state, "package", lua.PackageOpen, true)
	patchPackagePath(state)
}

// Patches the package.path variable to look for packages in the directory
// of the root config file, rather than the current directory
func patchPackagePath(state *lua.State) {
	path := path.Join(path.Dir(flags.configPath), "?.lua;")
	state.Global("package")
	state.PushString(path)
	state.SetField(-2, "path")
	state.Pop(1)
}
