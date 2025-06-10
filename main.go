package main

import (
	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/misc"
	"github.com/zuma206/sysmig/resolve"
)

var command = cobra.Command{
	Use:   "sysmig",
	Short: "Declarative system configuration for POSIX operating systems",
}

func init() {
	command.AddCommand(&resolve.Command)
	command.AddCommand(&misc.Version)
}

func main() {
	command.Execute()
}
