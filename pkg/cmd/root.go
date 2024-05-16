package cmd

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/compiler"
	"log"
	"log/slog"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	versionBool   bool
	compileString string
	debugBool     bool
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
			err := c.Compile()
			if err != nil {
				log.Fatal(err)
			}
			c.ConvertWat2Wasm("o.wasm")

			fmt.Println("Compilation done!", compilePath)
			if debugBool {
				debugFolder := "debug/"
				err := os.Mkdir(path.Clean(debugFolder), os.ModePerm)
				if err != nil && !os.IsExist(err) {
					log.Fatal(err)
				}
				c.WriteJsonFile(debugFolder + "ast.json")
				c.WriteWatFile(debugFolder + "o.wat")
			}
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

	rootCmd.Flags().StringVarP(&compileString, "compile", "c", "", "Compile an ygl file to wasm")

	rootCmd.Flags().BoolVarP(&debugBool, "debug", "d", false, "Has to be run with hammer-time\nEnable debug mode.\nWrite JSON and WAT file.")
}
