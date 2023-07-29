package cmd

import (
	"io"
	"log"
	"os"
	"os/signal"

	droneCl "github.com/jasonkwh/droneshield-test/internal/client"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "starts the drone client",
	Run:   client,
}

func client(cmd *cobra.Command, args []string) {
	// make grateful close pool
	var clPool []io.Closer

	zl, err := initZapLogger()
	if err != nil {
		log.Fatal("unable to start zap logger")
	}

	cl, err := droneCl.NewClient(cfg.Redis, cfg.WindSimulation, zl)
	if err != nil {
		zl.Fatal("failed to initialize drone client", zap.Error(err))
	}
	clPool = append(clPool, cl)

	zl.Info("client started")

	// handle shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	if err := gratefulClose(clPool); err != nil {
		zl.Error("failed to close the server", zap.Error(err))
	}
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
