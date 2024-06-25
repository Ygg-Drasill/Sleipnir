package cmd

import (
	"fmt"
	"github.com/Ygg-Drasill/Sleipnir/pkg/compiler"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

var (
	versionBool   bool
	compileString string
	debugBool     bool
)

var rootCmd = &cobra.Command{
	Use:   "sleipnir",
	Short: "sleipnir compiles Ygg-Drasill code to webassembly",
	Long:  `sleipnir is the preferred compiler by the norse gods`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionBool {
			currentVersion := getCurrentVersion()
			latestVersion, err := getLatestVersion()
			if err != nil {
				log.Printf("Error fetching latest version: %v", err)
				fmt.Printf("Current version: %s\n", currentVersion)
			} else {
				fmt.Printf("Current version: %s\n", currentVersion)
				fmt.Printf("Latest version: %s\n", latestVersion)
			}
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
			log.Fatal("Error displaying help:", err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(runCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&versionBool, "version", "v", false, "shows current version")
	rootCmd.Flags().StringVarP(&compileString, "compile", "c", "", "Compile an ygl file to wasm")
	rootCmd.Flags().BoolVarP(&debugBool, "debug", "d", false, "Has to be run with compile\nEnables debug mode and outputs a JSON and WAT file.")
}
