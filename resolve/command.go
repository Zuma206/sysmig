package resolve

import (
	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/utils"
)

type Flags struct {
	configPath    string
	migrationPath string
	statePath     string
	syncPath      string
}

// Global instance of the flags needed to resolve
// Kinda cursed but it's how cobra works for some reason?
var flags = new(Flags)

var Command = cobra.Command{
	Use:   "resolve",
	Short: "Resolves a lua system configuration into a migration script",
	Run: func(cmd *cobra.Command, args []string) {
		writeResolution(resolve(readState()))
		println("Done")
	},
}

func init() {
	Command.Flags().StringVarP(
		&flags.configPath, "config", "c",
		utils.GetConfigPath(),
		"the path of the system configuration to resolve",
	)
	Command.Flags().StringVarP(
		&flags.migrationPath, "output", "o",
		utils.GetMigrationPath(),
		"the path to write the output migration script to",
	)
	Command.Flags().StringVarP(
		&flags.statePath, "state", "s",
		utils.GetStatePath(),
		"the path to read/write the system state to",
	)
	Command.Flags().StringVarP(
		&flags.syncPath, "sync", "n",
		utils.GetSyncPath(),
		"the path to write the output sync script to",
	)
}
