package utils

import (
	"fmt"
	"os"
)

// If an error is not nil, it writes it to stderr and exits with status code 1
func HandleErr(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: ")
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
