package stdlib

import (
	"fmt"

	"github.com/Shopify/go-lua"
	"github.com/zuma206/sysmig/utils"
)

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

// Represents a key `Name` on a table `ParentName` of type `Type`
type LuaAttribute[T any] struct {
	Name       string
	Type       LuaType[T]
	ParentName string
}

// Push an attribute from a table at `index` onto the stack
func (attr *LuaAttribute[T]) Push(state *lua.State, index int) {
	state.Field(index, attr.Name)
}

// Return an attribute from a table at `index` onto the stack
func (attr *LuaAttribute[T]) Get(state *lua.State, index int) T {
	attr.Push(state, index)
	value, ok := attr.Type.Getter(state, -1)
	if !ok {
		typeName := state.TypeOf(-1).String()
		err := fmt.Errorf(
			"invalid %s: %s must be of type %s, not %s",
			attr.ParentName, attr.Name, attr.Type.Name, typeName,
		)
		utils.HandleErr(err)
	}
	state.Pop(1)
	return value
}

// Set the value at the top of the stack to the attribute on the table at `index`
func (attr *LuaAttribute[T]) Set(state *lua.State, index int) {
	state.SetField(index, attr.Name)
}
