package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/docker/docker/cliconfig"
)

const (
	version = "0.1.10-dev"
	usage   = "inspect images on a registry"
)

var inspectCmd = func(c *cli.Context) {
	inspect(c)
}

func main() {
	app := cli.NewApp()
	app.Name = "stackup"
	app.Version = version
	app.Usage = usage
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug output",
		},
		cli.StringFlag{
			Name:  "username",
			Value: "",
			Usage: "registry username",
		},
		cli.StringFlag{
			Name:  "password",
			Value: "",
			Usage: "registry password",
		},
		cli.StringFlag{
			Name:  "docker-cfg",
			Value: cliconfig.ConfigDir(),
			Usage: "Docker's cli config for auth",
		},
	}
	app.Before = func(c *cli.Context) error {
		if c.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}
	app.Action = inspectCmd
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
