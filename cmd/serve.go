package cmd

import (
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/jasonkwh/droneshield-test/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the websocket server",
	Long:  `starts the websocket server`,

	Run: func(cmd *cobra.Command, args []string) {
		// make grateful close pool
		var clPool []io.Closer

		zl, err := initZapLogger()
		if err != nil {
			log.Fatal("unable to start zap logger")
		}

		// start server
		srv := server.NewServer(cfg.Server, cfg.Redis, zl)
		clPool = append(clPool, srv)
		go func() {
			err := srv.ListenAndServe()
			if err != nil {
				zl.Error("server has stopped", zap.Error(err))
			}
		}()

		zl.Info("server started")

		// handle shutdown
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		if err := gratefulClose(clPool); err != nil {
			zl.Error("failed to close the server", zap.Error(err))
		}
		os.Exit(0)
	},
}
