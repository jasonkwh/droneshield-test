package cmd

import (
	"fmt"

	"github.com/jasonkwh/droneshield-test/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CONFIG
var cfgFile string

type Config struct {
	Server config.ServerConfig
	Redis  config.RedisConfig
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file")
}

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
