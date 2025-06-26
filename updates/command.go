package updates

import (
	"errors"
	"os"
	"path"

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
		return
	}

	if latestRelease.TagName == utils.VERSION {
		println("You're already on the latest release")
		os.Exit(0)
	}

	binaryData := latestRelease.GetAsset(binaryAssetName).Download()
	homeDir, err := os.UserHomeDir()
	utils.HandleErr(err)
	err = os.MkdirAll(path.Join(homeDir, ".sysmig"), utils.READWRITE_PERMS)
	utils.HandleErr(err)
	err = os.WriteFile(path.Join(homeDir, ".sysmig", "sysmig"), *binaryData, utils.READWRITE_PERMS)
	utils.HandleErr(err)

	utils.HandleErr(err)
	println("Downloaded the latest version")
}
