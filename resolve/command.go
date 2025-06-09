package resolve

import (
	"os"

	"github.com/spf13/cobra"
)

var flags = new(struct {
	configPath string
})

var Command = cobra.Command{
	Use:   "resolve",
	Short: "Resolves a lua system configuration into a migration script",
	Run: func(cmd *cobra.Command, args []string) {
		resolution := resolve(flags.configPath)
		println(resolution.MigrationScript)
		println(resolution.SyncScript)
		println(resolution.NewStateJson)
	},
}

func init() {
	Command.Flags().StringVarP(
		&flags.configPath, "config", "c",
		getConfigPath(),
		"the path of the system configuration to resolve",
	)
}

func getConfigPath() string {
	dir := os.Getenv("HOME")
	if dir == "" {
		dir = "/etc"
	}
	return dir + "/.sysmig/system.lua"
}
