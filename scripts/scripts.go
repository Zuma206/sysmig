package scripts

import (
	_ "embed"
	"encoding/base64"
	"fmt"
)

//go:embed migrate.sh
var migrate string

//go:embed sync.sh
var sync string

func FmtMigration(script string, nextState []byte, statePath string) []byte {
	nextStateB64 := base64.StdEncoding.EncodeToString(nextState)
	fmtScript := fmt.Sprintf(migrate, script, nextStateB64, statePath)
	return []byte(fmtScript)
}

func FmtSync(script string, lastSyncPath string) []byte {
	fmtScript := fmt.Sprintf(sync, script, lastSyncPath, "%s")
	return []byte(fmtScript)
}
