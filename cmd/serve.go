package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/jasonkwh/droneshield-test/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

		if err := Close(clPool); err != nil {
			zl.Error("failed to close the server", zap.Error(err))
		}
		os.Exit(0)
	},
}

func initZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()

	// set the internal logger to INFO because we need all internal logs
	cfg.Level.SetLevel(zapcore.InfoLevel)
	return cfg.Build()
}

func Close(services []io.Closer) error {
	var errs error

	for _, item := range services {
		err := item.Close()
		if err != nil {
			errs = multierr.Append(errs, err)
		}
	}

	return errs
}
