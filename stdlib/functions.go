package stdlib

import (
	"fmt"

	"github.com/Shopify/go-lua"
)

type LuaArgType struct {
	Name      string
	Validator func(state *lua.State) bool
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

func (luaFunc *LuaFunc) GoFunction(state *lua.State) int {
	luaFunc.validateNArgs(state)
	luaFunc.validateArgTypes(state)
	return luaFunc.Body(state)
}

func (luaFunc *LuaFunc) validateNArgs(state *lua.State) {
	givenArgs := state.Top()
	requiredArgs := len(luaFunc.Args)
	if givenArgs != requiredArgs {
		msg := fmt.Sprintf("%s takes %d arguments but got %d", luaFunc.Name, requiredArgs, givenArgs)
		state.PushString(msg)
		state.Error()
	}
}

func (luaFunc *LuaFunc) validateArgTypes(state *lua.State) {
	for _, arg := range luaFunc.Args {
		if !arg.Type.Validator(state) {
			msg := fmt.Sprintf("%s must be of type %s", arg.Name, arg.Type.Name)
			state.PushString(msg)
			state.Error()
		}
	}
}

func (luaFunc *LuaFunc) Open(state *lua.State) {
	state.PushGoFunction(luaFunc.GoFunction)
	state.SetField(-2, luaFunc.Name)
}
