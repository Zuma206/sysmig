package std

import (
	"path"
	"strings"

	"github.com/Shopify/go-lua"
)

func pathLua(configPath string) lua.Function {
	return func(state *lua.State) int {
		state.PushGoFunction(func(state *lua.State) int {
			desiredPath, ok := state.ToString(1)
			if !ok {
				state.PushString("std.path must be passed a string")
				state.Error()
			}
			if !strings.HasPrefix(desiredPath, "/") && !strings.HasPrefix(desiredPath, "~") {
				state.PushString(path.Join(configPath, desiredPath))
			}
			return 1
		})
		return 1
	}
}
