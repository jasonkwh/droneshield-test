package cmd

import (
	"fmt"
	"io"
	"log"

	"github.com/jasonkwh/droneshield-test/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the websocket server",
}

// CONFIG
var cfgFile string

type Config struct {
	Server config.ServerConfig
	Redis  config.RedisConfig

	PubSubChannel string
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file")
}

// CONFIG
var cfg Config

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Put all the config in a common struct
	viper.Unmarshal(&cfg)
}

func initZapLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	// set the internal logger to INFO because we need all internal logs
	cfg.Level.SetLevel(zapcore.InfoLevel)
	return cfg.Build()
}

func gratefulClose(services []io.Closer) error {
	var errs error

	for _, item := range services {
		err := item.Close()
		if err != nil {
			errs = multierr.Append(errs, err)
		}
	}

	return errs
}
