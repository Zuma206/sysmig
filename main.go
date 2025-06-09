package main

import (
	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/resolve"
	"github.com/zuma206/sysmig/utils"
)

var command = cobra.Command{
	Use:   "sysmig",
	Short: "Declarative system configuration for POSIX operating systems",
}

func init() {
	command.AddCommand(&resolve.Command)
}

func main() {
	utils.HandleErr(command.Execute())
}
