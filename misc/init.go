package misc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/utils"
)

var Init = cobra.Command{
	Use:   "init",
	Short: "Creates an initial state file",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runInit(); err != nil {
			utils.HandleErr(err)
		}
	},
}

// Initialises the sysmig directory and files inside, may panic
func runInit() error {
	err := checkDir()
	if err != nil {
		return err
	}
	return checkStateFile()
}

// Checks if the sysmig directory exists, else it creates it
func checkDir() error {
	path := utils.GetDir()
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return fmt.Errorf("%q exists but is not a directory", path)
	} else if err != nil && os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	}
	return err
}

// Checks if the state file exists, and creates it otherwise
func checkStateFile() error {
	path := utils.GetStatePath()
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return createStateFile(path)
	}
	println("Cannot create an initial state file, one already exists")
	return nil
}

// Writes a JSON nil value to the given path
func createStateFile(path string) error {
	err := os.WriteFile(path, []byte("null\n"), utils.READWRITE_PERMS)
	if err != nil {
		return err
	}
	fmt.Printf("Created an initial state file at %s\n", path)
	return nil
}
