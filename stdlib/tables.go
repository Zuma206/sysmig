package stdlib

import "github.com/Shopify/go-lua"

type LuaData interface {
	Open(*lua.State)
}

type LuaTable struct {
	Name   string
	Values []LuaData
}

func (luaTable *LuaTable) Open(state *lua.State) {
	state.CreateTable(0, len(luaTable.Values))
	for _, value := range luaTable.Values {
		value.Open(state)
	}
	state.SetField(-2, luaTable.Name)
}
