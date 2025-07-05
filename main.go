package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/misc"
	"github.com/zuma206/sysmig/resolve"
	"github.com/zuma206/sysmig/updates"
	"github.com/zuma206/sysmig/utils"
)

var command = cobra.Command{
	Use:   "sysmig",
	Short: "Declarative system configuration for POSIX operating systems",
}

func init() {
	command.AddCommand(&resolve.Command)
	command.AddCommand(&misc.Version)
	command.AddCommand(&misc.Clean)
	command.AddCommand(&misc.Init)
	command.AddCommand(&updates.Command)
}

func main() {
	executable, err := os.Executable()
	utils.HandleErr(err)
	updateExecutablePath := updates.GetExecutablePath()
	if executable == updateExecutablePath {
		updates.ContinueInstall(executable)
	} else {
		command.Execute()
	}
}
