package misc

import (
	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/manifest"
)

var Version = cobra.Command{
	Use:   "version",
	Short: "get the current version",
	Run: func(cmd *cobra.Command, args []string) {
		println(manifest.Manifest.Version)
	},
}
