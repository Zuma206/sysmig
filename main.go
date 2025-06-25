package main

import (
	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/misc"
	"github.com/zuma206/sysmig/resolve"
	"github.com/zuma206/sysmig/updates"
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
	command.Execute()
}
