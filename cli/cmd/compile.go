
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "compiles Ygg-Drasill code",
	Long: `Compile YggDrasill code to a webassembly binary.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("compile called")
	},
}

func init() {
	RootCmd.AddCommand(compileCmd)

	//Todo: add flags for run.

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
