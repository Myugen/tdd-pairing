package cmd

import (
	"fmt"

	"github.com/myugen/tdd-pairing-go/pkg/httpserver"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Runs the app",
	RunE:  runUpCmd,
}

func runUpCmd(cmd *cobra.Command, args []string) error {
	server := httpserver.Setup()
	return server.Start(fmt.Sprintf(":%s", viper.GetString("port")))
}
