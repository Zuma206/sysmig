package resolve

import (
	"os"

	"github.com/zuma206/sysmig/utils"
)

func writeResolution(resolution *Resolution) {
	utils.HandleErr(os.WriteFile(
		flags.migrationPath,
		[]byte(resolution.migrationScript),
		utils.EXECUTABLE_PERMS,
	))
	utils.HandleErr(os.WriteFile(
		flags.statePath,
		[]byte(resolution.newStateJson),
		utils.READWRITE_PERMS,
	))
	utils.HandleErr(os.WriteFile(
		flags.syncPath,
		[]byte(resolution.syncScript),
		utils.EXECUTABLE_PERMS,
	))
}
