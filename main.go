package main

import (
	"basketManager/cmd/migrate"
	"basketManager/cmd/server"
	"basketManager/config"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const exitFailure = 1

func main() {
	cfg := config.InitConfig()

	var cmd = &cobra.Command{
		Use:   "go-basket-manager",
		Short: "basket",
	}

	logrus.Debugf("config loaded: %+v", cfg)

	server.Register(cmd, cfg)
	migrate.Register(cmd, cfg)

	if err := cmd.Execute(); err != nil {
		logrus.Error(err.Error())
		os.Exit(exitFailure)
	}
}
