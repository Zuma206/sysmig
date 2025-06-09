package resolve

import "github.com/spf13/cobra"

var Command = cobra.Command{
	Use:   "resolve",
	Short: "Resolves a lua system configuration into a migration script",
	Run: func(cmd *cobra.Command, args []string) {
		println("Resolved")
	},
}
