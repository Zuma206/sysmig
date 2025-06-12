package stdlib

import (
	"fmt"

	"github.com/Shopify/go-lua"
)

type LuaArgType struct {
	Name      string
	Validator func(state *lua.State, index int) bool
}

type LuaArg struct {
	Name string
	Type LuaArgType
}

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
		if !arg.Type.Validator(state, top+index) {
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
