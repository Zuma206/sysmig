package misc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/resolve"
	"github.com/zuma206/sysmig/utils"
)

var Init = cobra.Command{
	Use:   "init",
	Short: "Creates an initial state file",
	Run: func(cmd *cobra.Command, args []string) {
		runInit()
	},
}

func runInit() {
	path := resolve.GetStatePath()
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		createStateFile(path)
	} else {
		println("Cannot create an initial state file, one already exists")
	}
}

func createStateFile(path string) {
	utils.HandleErr(os.WriteFile(path, []byte("null\n"), utils.READWRITE_PERMS))
	fmt.Printf("Created an initial state file at %s\n", path)
}
