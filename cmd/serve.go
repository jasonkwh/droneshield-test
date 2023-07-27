package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// CONFIG
var cfg Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the websocket server",
	Long:  `starts the websocket server`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cfg.Server.Port)
	},
}
