package cli

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/zh0glikk/mongo-app/internal/config"
	"github.com/zh0glikk/mongo-app/internal/service"
)

const defaultConfigPath = "./config.json"

func Run(args []string) bool {
	var cfg config.Config

	app := cli.NewApp()

	err := config.SetupConfig(defaultConfigPath, &cfg)
	if err != nil {
		logrus.WithError(err).Error("failed to parse config")
		return false
	}

	app.Commands = cli.Commands{
		{
			Name: "web",
			Action: func(_ *cli.Context) error {
				return service.NewService(&cfg).Run(context.Background())
			},
		},
	}

	if err := app.Run(args); err != nil {
		logrus.WithError(err).Error("app finished")
		return false
	}
	return true
}
