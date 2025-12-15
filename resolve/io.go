package resolve

import (
	"fmt"
	"os"

	"github.com/zuma206/sysmig/scripts"
	"github.com/zuma206/sysmig/utils"
)

// Take a resolution struct and write the output the various appropriate files
func writeResolution(resolution *Resolution) {
	utils.HandleErr(os.WriteFile(
		flags.migrationPath,
		scripts.FmtMigration(resolution.migration, resolution.nextState, flags.statePath),
		utils.EXECUTABLE_PERMS,
	))
	utils.HandleErr(os.WriteFile(
		flags.syncPath,
		scripts.FmtSync(resolution.sync, flags.lastSyncPath),
		utils.EXECUTABLE_PERMS,
	))
}

// Read the state file
func readState() string {
	data, err := os.ReadFile(flags.statePath)
	if os.IsNotExist(err) {
		err = fmt.Errorf("%s (try running `sysmig init` if you haven't yet on this system)", err.Error())
	}
	utils.HandleErr(err)
	return string(data)
}
