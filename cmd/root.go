
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


//RootCmd is the base command
var RootCmd = &cobra.Command{
	Use:   "Sleipnir",
	Short: "Sleipnir is the preferred compiler by the norse gods",
	Long: `Sleipnir compiles and runs Ygg-Drasill code. Below are the listed subcommands:
	-compile ()
	-ruuuuuuuuuuun ()`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Sleipnir.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


