package utils

import (
	"errors"
	"os"
	"path"
)

// $HOME/.sysmig
// or /etc/.sysmig
func GetDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		HandleErr(errors.New("cannot find your home directory"))
	}
	return path.Join(dir, ".sysmig")
}

func GetConfigDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		HandleErr(errors.New("cannot find your config directory"))
	}
	return dir
}

func GetConfigPath() string {
	return path.Join(GetConfigDir(), "sysmig", "system.lua")
}

func GetMigrationPath() string {
	return path.Join(GetDir(), "migrate.sh")
}

func GetStatePath() string {
	return path.Join(GetDir(), "current-system-state.json")
}

func GetSyncPath() string {
	return path.Join(GetDir(), "sync.sh")
}
