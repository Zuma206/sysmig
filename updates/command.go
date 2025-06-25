package updates

import (
	"errors"

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
	latestRelease := GetReleases().GetLatestRelease()
	if latestRelease == nil {
		utils.HandleErr(errors.New("cannot find the latest release"))
	}

	asset := latestRelease.GetAsset(binaryAssetName)
	if asset == nil {
		utils.HandleErr(errors.New("cannot locate binary to download"))
		return
	}

	println(asset.BrowserDownloadUrl)
}
