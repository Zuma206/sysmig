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
	requireStdModules(state,
		"entries",
		"map",
		"copy",
		"Set",
		"migrator",
		"sequence",
		"dnf",
		"apt",
		"system",
		"nothing",
		"file",
		"files",
		"symlinks",
		"flatpak",
		"components",
	)
}

func getCode(name string) string {
	content, err := stdLuaSources.ReadFile("lua/" + name + ".lua")
	utils.HandleErr(err)
	return string(content)
}

func requireStdModules(state *lua.State, moduleNames ...string) {
	for _, moduleName := range moduleNames {
		require(state, "@std."+moduleName, getCode(moduleName))
	}
	requireRoot(state, &moduleNames)
}

func requireRoot(state *lua.State, moduleNames *[]string) {
	lua.Require(state, "@std", func(state *lua.State) int {
		state.CreateTable(0, len(*moduleNames))
		for _, moduleName := range *moduleNames {
			state.Global("require")
			state.PushString("@std." + moduleName)
			state.Call(1, 1)
			state.SetField(-2, moduleName)
		}
		return 1
	}, false)
}

func require(state *lua.State, name string, code string) {
	lua.Require(state, name, func(state *lua.State) int {
		err := lua.LoadString(state, code)
		utils.HandleErr(err)
		state.Call(0, 1)
		return 1
	}, false)
}
