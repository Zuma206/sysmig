package resolve

import (
	"errors"
	"os"
	"path"

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
	dir, err := os.UserHomeDir()
	if err != nil {
		utils.HandleErr(errors.New("cannot find your home directory"))
	}
	return path.Join(dir, ".sysmig")
}

func GetConfigDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		utils.HandleErr(errors.New("cannot find your config directory"))
	}
	return dir
}

func GetConfigPath() string {
	return path.Join(GetConfigDir(), "sysmig", "system.lua")
}

func GetMigrationPath() string {
	return path.Join(getDir(), "migrate.sh")
}

func GetStatePath() string {
	return path.Join(getDir(), "current-system-state.json")
}

func GetSyncPath() string {
	return path.Join(getDir(), "sync.sh")
}
