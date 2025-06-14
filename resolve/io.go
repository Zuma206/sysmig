package resolve

import (
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
		[]byte(resolution.sync),
		utils.EXECUTABLE_PERMS,
	))
}

// Read the state file
func readState() string {
	data, err := os.ReadFile(flags.statePath)
	utils.HandleErr(err)
	return string(data)
}
