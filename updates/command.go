package updates

import (
	"errors"
	"os"
	"os/user"

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
	checkPrivilege()
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
	executablePath := GetExecutablePath()
	install(latestRelease, executablePath)
	run(executablePath)
}

func checkPrivilege() {
	currentUser, err := user.Current()
	utils.HandleErr(err)
	if currentUser.Uid != "0" {
		utils.HandleErr(errors.New("updates must be run as root"))
	}
}
