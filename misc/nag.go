package misc

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/utils"
)

type NagFlags struct {
	daysUntilNag int
	lastSyncPath string
}

var nagFlags = new(NagFlags)

var Nag = cobra.Command{
	Use:   "nag",
	Short: "print a message when it's been a while since the last sync",
	Run: func(cmd *cobra.Command, args []string) {
		if err := nag(); err != nil {
			utils.HandleErr(err)
		}
	},
}

func init() {
	Nag.Flags().IntVarP(
		&nagFlags.daysUntilNag, "days", "d", 7,
		"the number of days to start nagging after",
	)
	Nag.Flags().StringVarP(
		&nagFlags.lastSyncPath, "last-sync", "l",
		utils.GetLastSyncPath(),
		"the number of days to start nagging after",
	)
}

const secondsInDay = 24 * 60 * 60

func nag() error {
	lastSyncContent, err := os.ReadFile(nagFlags.lastSyncPath)
	if err != nil {
		return err
	}
	lastSyncStr := strings.ReplaceAll(string(lastSyncContent), "\n", "")
	lastSync, err := strconv.ParseInt(lastSyncStr, 10, 64)
	if err != nil {
		return err
	}
	elapsed := time.Now().Unix() - lastSync
	if elapsed > int64(nagFlags.daysUntilNag*secondsInDay) {
		days := elapsed / secondsInDay
		fmt.Printf("It's been %d days since your last sync\n", days)
	}
	return nil
}
