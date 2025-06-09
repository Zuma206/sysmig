package utils

import (
	"fmt"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "error: ")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
