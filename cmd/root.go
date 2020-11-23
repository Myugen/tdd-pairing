package cmd

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "App CLI Interface",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.app/config.yaml or current location)")
	viper.SetDefault("author", "John Doe <john@doe>")
	viper.SetDefault("version", "0.0.1")
	viper.SetDefault("server.port", "8080")

	rootCmd.AddCommand(versionCmd, upCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalf("home directory detection error: %s", err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(fmt.Sprintf("%s/.app", home))
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.SetEnvPrefix("app")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
