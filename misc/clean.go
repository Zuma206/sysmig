package misc

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zuma206/sysmig/utils"
)

type CleanFlags struct {
	clearMigration bool
	clearState     bool
	clearSync      bool
	confirm        bool
}

// The global instance of the flags for the clean command
var cleanFlags = new(CleanFlags)

var Clean = cobra.Command{
	Use:   "clean",
	Short: "Removes state/migration (dangerous!)",
	Run: func(cmd *cobra.Command, args []string) {
		clean()
	},
}

func init() {
	Clean.Flags().BoolVarP(
		&cleanFlags.confirm, "confirm", "y",
		false, "confirm the action, and actually perform it",
	)
	Clean.Flags().BoolVarP(
		&cleanFlags.clearMigration, "migration", "m",
		false, "delete the last migration script",
	)
	Clean.Flags().BoolVarP(
		&cleanFlags.clearState, "state", "s",
		false, "deletes the current state (dangerous!)",
	)
	Clean.Flags().BoolVarP(
		&cleanFlags.clearSync, "sync", "n",
		false, "deletes the last sync script",
	)
}

type CleanInstruction struct {
	name    string
	path    string
	enabled bool
}

func clean() {
	instructions := []*CleanInstruction{
		{"migration script", utils.GetMigrationPath(), cleanFlags.clearMigration},
		{"current state file", utils.GetStatePath(), cleanFlags.clearState},
		{"sync script", utils.GetSyncPath(), cleanFlags.clearSync},
	}
	if !cleanFlags.confirm {
		printCleans(&instructions)
	} else {
		performCleans(&instructions)
	}
}

// Prints out what the clean command would do if it was confirmed
func printCleans(cleanInstructions *[]*CleanInstruction) {
	println("Run the command again with -y to delete:")
	printed := false
	for _, instruction := range *cleanInstructions {
		if instruction.enabled {
			fmt.Printf("\tThe %s (%s)\n", instruction.name, instruction.path)
			printed = true
		}
	}
	if !printed {
		println("\tnothing")
	}
}

// Actually performs the clean, and deletes the files
func performCleans(cleanInstructions *[]*CleanInstruction) {
	for _, instruction := range *cleanInstructions {
		if instruction.enabled {
			utils.HandleErr(os.Remove(instruction.path))
			fmt.Printf("Removed %s (%s)\n", instruction.name, instruction.path)
		}
	}
}
