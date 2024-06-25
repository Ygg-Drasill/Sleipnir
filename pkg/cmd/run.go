package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a WebAssembly binary",
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		binary, err := readBinaryFromFile(path)
		handleError(err)
		result, err := runBinary(binary)
		handleError(err)

		if result != nil {
			fmt.Printf("%v", result)
		}
	},
}
