package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a WebAssembly binary",
	Long: `Run a WebAssembly binary in a wasmtime instance
(no screeps nodes will run properly)`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You need to specify the binary to run")
			os.Exit(1)
		}
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
