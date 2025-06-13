package stdlib

import (
	"fmt"

	"github.com/Shopify/go-lua"
)

// Represents a type in the lua vm
// has a name, and a function to get
// and/or validate an index is that type
type LuaType[T any] struct {
	Name   string
	Getter func(state *lua.State, index int) (T, bool)
}

// Represents an stdlib function argument's name and type
type LuaArg struct {
	Name string
	Type LuaType[any]
}

// Represents a stdlib function name, arguments, and go body
type LuaFunc struct {
	Name string
	Args []LuaArg
	Body lua.Function
}

// A pushable go-function for a LuaFunc
func (luaFunc *LuaFunc) GoFunction(state *lua.State) int {
	luaFunc.validateNArgs(state)
	luaFunc.validateArgTypes(state)
	return luaFunc.Body(state)
}

// Validates that there are the correct amount of arguments passed to a LuaFunc
func (luaFunc *LuaFunc) validateNArgs(state *lua.State) {
	givenArgs := state.Top()
	requiredArgs := len(luaFunc.Args)
	if givenArgs != requiredArgs {
		msg := fmt.Sprintf("%s takes %d arguments but got %d", luaFunc.Name, requiredArgs, givenArgs)
		state.PushString(msg)
		state.Error()
	}
}

// Validate that all passed functions are of the correct type
func (luaFunc *LuaFunc) validateArgTypes(state *lua.State) {
	top := state.Top()
	for index, arg := range luaFunc.Args {
		if _, ok := arg.Type.Getter(state, top+index); !ok {
			msg := fmt.Sprintf("%s must be of type %s", arg.Name, arg.Type.Name)
			state.PushString(msg)
			state.Error()
		}
	}
}

// Adds the lua func to the table at the top of the stack
func (luaFunc *LuaFunc) Open(state *lua.State) {
	state.PushGoFunction(luaFunc.GoFunction)
	state.SetField(-2, luaFunc.Name)
}
