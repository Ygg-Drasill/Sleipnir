package cmd

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/compiler"
	"log/slog"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	versionBool   bool
	compileString string
)

var rootCmd = &cobra.Command{
	Use:   "Sleipnir",
	Short: "Sleipnir is the preferred compiler by the norse gods",
	Long:  `Sleipnir is the preferred compiler by the norse gods`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionBool {
			fmt.Println("Version 0.0.1")
			return
		}

		if compileString != "" {
			compilePath := path.Clean(compileString)
			c := compiler.NewFromFile(compilePath)
			c.Compile()
			c.WriteOutputToFile("o.wat")
			fmt.Println("Compilation done!", compilePath)
			return
		}

		err := cmd.Help()
		if err != nil {
			slog.Error("Error displaying help:", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&versionBool, "version", "v", false, "shows current version")

	rootCmd.Flags().StringVar(&compileString, "hammer-time", "", "Compile an ygl file to wasm")
}