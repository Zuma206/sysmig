package updates

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/utils"
)

var Command = cobra.Command{
	Use:   "update",
	Short: "Update sysmig to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		performUpdate()
	},
}

const binaryAssetName = "sysmig"

func performUpdate() {
	println("Checking for updates...")
	latestRelease := GetReleases().GetLatestRelease()
	if latestRelease == nil {
		utils.HandleErr(errors.New("cannot find the latest release"))
		return
	}
	if latestRelease.TagName == utils.VERSION {
		println("You're already on the latest release")
		os.Exit(0)
	} else {
		println("Update found, downloading...")
	}
	install(latestRelease)
}
