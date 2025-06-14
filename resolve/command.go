package resolve

import (
	"os"

	"github.com/spf13/cobra"
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
		GetConfigPath(),
		"the path of the system configuration to resolve",
	)
	Command.Flags().StringVarP(
		&flags.migrationPath, "output", "o",
		GetMigrationPath(),
		"the path to write the output migration script to",
	)
	Command.Flags().StringVarP(
		&flags.statePath, "state", "s",
		GetStatePath(),
		"the path to read/write the system state to",
	)
	Command.Flags().StringVarP(
		&flags.syncPath, "sync", "n",
		GetSyncPath(),
		"the path to write the output sync script to",
	)
}

// $HOME/.sysmig
// or /etc/.sysmig
func getDir() string {
	dir := os.Getenv("HOME")
	if dir == "" {
		dir = "/etc"
	}
	return dir + "/.sysmig"
}

func GetConfigPath() string {
	return getDir() + "/system.lua"
}

func GetMigrationPath() string {
	return getDir() + "/migrate.sh"
}

func GetStatePath() string {
	return getDir() + "/state.json"
}

func GetSyncPath() string {
	return getDir() + "/sync.sh"
}
