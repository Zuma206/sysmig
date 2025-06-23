package std

import (
	"embed"

	"github.com/Shopify/go-lua"
	"github.com/zuma206/sysmig/utils"
)

//go:embed lua/*
var stdLuaSources embed.FS

func OpenStd(state *lua.State, configDir string) {
	lua.Require(state, "@std.path", pathLua(configDir), false)
	lua.Require(state, "@std.serialize", serializeLua, false)
	lua.Require(state, "@std.dir", dirLua, false)
	requireModules(state,
		"entries",
		"map",
		"copy",
		"Set",
		"migrator",
		"sequence",
		"rhel",
		"deb",
		"system",
		"nothing",
		"files",
	)
	require(state, "@std", getCode("std"))
}

func getCode(name string) string {
	content, err := stdLuaSources.ReadFile("lua/" + name + ".lua")
	utils.HandleErr(err)
	return string(content)
}

func requireModules(state *lua.State, moduleNames ...string) {
	for _, moduleName := range moduleNames {
		require(state, "@std."+moduleName, getCode(moduleName))
	}
}

func require(state *lua.State, name string, code string) {
	lua.Require(state, name, func(state *lua.State) int {
		err := lua.LoadString(state, code)
		utils.HandleErr(err)
		state.Call(0, 1)
		return 1
	}, false)
}
