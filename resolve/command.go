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
	},
}

func init() {
	Command.Flags().StringVarP(
		&flags.configPath, "config", "c",
		getConfigPath(),
		"the path of the system configuration to resolve",
	)
	Command.Flags().StringVarP(
		&flags.migrationPath, "output", "o",
		getMigrationPath(),
		"the path to write the output migration script to",
	)
	Command.Flags().StringVarP(
		&flags.statePath, "state", "s",
		getStatePath(),
		"the path to read/write the system state to",
	)
	Command.Flags().StringVarP(
		&flags.syncPath, "sync", "n",
		getSyncPath(),
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

func getConfigPath() string {
	return getDir() + "/system.lua"
}

func getMigrationPath() string {
	return getDir() + "/migrate.sh"
}

func getStatePath() string {
	return getDir() + "/state.json"
}

func getSyncPath() string {
	return getDir() + "/sync.sh"
}
