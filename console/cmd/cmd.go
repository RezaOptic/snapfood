package cmd

import (
	"github.com/spf13/cobra"
	"snapfood/app"
)

//const prefixLog string = "________________"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "datacmd",
	Short: "starting point of all other commands.",
	Long:  `starting point of all other commands.`,
}

var Application *app.App

// Execute execute cmd
func init() {
	Application = app.NewApp(nil)
}

// Execute execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		return
	}
}
