package updates

import (
	"os"
	"path"

	"github.com/zuma206/sysmig/utils"
)

func install(release *GithubRelease) {
	binaryData := release.GetAsset(binaryAssetName).Download()
	executablePath := getExecutablePath()
	utils.HandleErr(os.MkdirAll(path.Dir(executablePath), utils.READWRITE_PERMS))
	utils.HandleErr(os.WriteFile(executablePath, *binaryData, utils.EXECUTABLE_PERMS))
	println("Download complete")
	println("Installing...")
}

func getExecutablePath() string {
	userHomeDir, err := os.UserHomeDir()
	utils.HandleErr(err)
	return path.Join(userHomeDir, ".sysmig", "sysmig")
}
