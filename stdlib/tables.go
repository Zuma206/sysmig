package stdlib

import "github.com/Shopify/go-lua"

// Represents a lua table or function that can be attached to a table
type LuaData interface {
	Open(*lua.State)
}

type LuaTable struct {
	Name   string
	Values []LuaData
}

// Adds the table to the table at the top of the stack
func (luaTable *LuaTable) Open(state *lua.State) {
	state.CreateTable(0, len(luaTable.Values))
	for _, value := range luaTable.Values {
		value.Open(state)
	}
	state.SetField(-2, luaTable.Name)
}
