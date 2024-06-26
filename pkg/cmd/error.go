package cmd

import (
	"fmt"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
