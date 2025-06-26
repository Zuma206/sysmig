package updates

import (
	"os"
	"os/exec"
	"path"

	"github.com/zuma206/sysmig/utils"
)

func install(release *GithubRelease) {
	binaryData := release.GetAsset(binaryAssetName).Download()
	executablePath := getExecutablePath()
	utils.HandleErr(os.MkdirAll(path.Dir(executablePath), utils.READWRITE_PERMS))
	utils.HandleErr(os.WriteFile(executablePath, *binaryData, utils.EXECUTABLE_PERMS))
	println("Download complete")
	run(executablePath)
}

func run(executablePath string) {
	cmd := exec.Command(executablePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	utils.HandleErr(cmd.Start())
}

func getExecutablePath() string {
	userHomeDir, err := os.UserHomeDir()
	utils.HandleErr(err)
	return path.Join(userHomeDir, ".sysmig", "sysmig")
}
