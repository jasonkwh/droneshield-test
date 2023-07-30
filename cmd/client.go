package cmd

import (
	"io"
	"log"
	"net/http"
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

	cl, err := droneCl.NewClient(cfg.Redis, cfg.Client.MovementServer, cfg.Client.WindSimulation, zl)
	if err != nil {
		zl.Fatal("failed to initialize drone client", zap.Error(err))
	}
	clPool = append(clPool, cl)
	go func() {
		if err := cl.Run(); err != nil && err != http.ErrServerClosed {
			zl.Fatal("drone client movement server failed to serve", zap.Error(err))
		}
	}()

	zl.Info("client started")

	// handle shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	if err := gracefulClose(clPool); err != nil {
		zl.Error("failed to close the server", zap.Error(err))
	}
	os.Exit(0)
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
