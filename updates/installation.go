package updates

import (
	"errors"
	"os"
	"os/exec"
	"os/user"
	"path"
	"time"

	"github.com/zuma206/sysmig/utils"
)

func install(release *GithubRelease, executablePath string) {
	binaryData := release.GetAsset(binaryAssetName).Download()
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

func GetExecutablePath() string {
	var homeDir string
	currentUser, err := user.Current()
	utils.HandleErr(err)
	if currentUser.Uid == "0" {
		sudoUsername, ok := os.LookupEnv("SUDO_USER")
		if !ok {
			utils.HandleErr(errors.New("ran as root, but SUDO_USER was not set"))
		}
		sudoUser, err := user.Lookup(sudoUsername)
		utils.HandleErr(err)
		homeDir = sudoUser.HomeDir
	} else {
		homeDir, err = os.UserHomeDir()
		utils.HandleErr(err)
	}
	return path.Join(homeDir, ".sysmig", "sysmig")
}

func ContinueInstall(executable string) {
	assertPrivilege()
	println("Installing...")
	time.Sleep(1 * time.Second)
	data, err := os.ReadFile(executable)
	utils.HandleErr(err)
	utils.HandleErr(os.WriteFile(path.Join("/", "usr", "local", "bin", "sysmig"), data, utils.EXECUTABLE_PERMS))
	println("sysmig was successfully updated")
}
