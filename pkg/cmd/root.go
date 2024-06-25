package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"

	"github.com/Ygg-Drasill/Sleipnir/pkg/compiler"
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

func getCurrentVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	return info.Main.Version
}

func getLatestVersion() (string, error) {
	const repoURL = "https://api.github.com/repos/Ygg-Drasill/Sleipnir/releases/latest"
	resp, err := http.Get(repoURL)
	if err != nil {
		return "", fmt.Errorf("error fetching latest version: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error fetching latest version: received status code %d", resp.StatusCode)
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return "", fmt.Errorf("error parsing latest version response: %w", err)
	}

	return release.TagName, nil
}
