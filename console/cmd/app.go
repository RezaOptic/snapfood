// Package cmd contains the console commands to run the HTTP server or other scripts.
package cmd

import (
	"github.com/spf13/cobra"
	"snapfood/utils/logger"
)

var App = &cobra.Command{
	Use:   "http_server",
	Short: "http_server",
	Long:  `this job will run main app on http server`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.ZLogger.Info("Main app started to run ...")
		Application.Initialize()
	},
}

func init() {
	rootCmd.AddCommand(App)
}
