package stdlib

import (
	"errors"

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

type Pusher func(state *lua.State, index int)

// Defines a function that pushes the `field` field of table `index` onto the stack
func DefinePusher(field string) Pusher {
	return func(state *lua.State, index int) {
		state.Field(index, field)
	}
}

type Getter[T any] func(state *lua.State, index int) T

// Converts a pusher into a getter using a to method
func DefineGetter[T any](pusher Pusher, luaType LuaType[T], typeErrMsg string) Getter[T] {
	return func(state *lua.State, index int) T {
		pusher(state, index)
		value, ok := luaType.Getter(state, -1)
		if !ok {
			utils.HandleErr(errors.New(typeErrMsg))
		}
		state.Pop(1)
		return value
	}
}
