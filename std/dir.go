package std

import (
	"path"

	"github.com/Shopify/go-lua"
)

func dirLua(state *lua.State) int {
	state.PushGoFunction(func(state *lua.State) int {
		filePath, ok := state.ToString(1)
		if !ok {
			state.PushString("std.dir must be passed a string")
			state.Error()
		}
		state.PushString(path.Dir(filePath))
		return 1
	})
	return 1
}
