package main

import (
	"log"
	"os"
	"worktreez/cmd"
	"worktreez/config"

	"github.com/urfave/cli/v2"
)

func main() {
	config.SetGlobalConfig()
	app := &cli.App{
		Authors: []*cli.Author{
			{
				Name:  "Ilyes Hammadi",
				Email: "hammadiilyesahmed@gmail.com",
			},
		},
		Name:                 "worktreez",
		EnableBashCompletion: true,
		Version:              "0.1.0",
		Commands: []*cli.Command{
			cmd.Create(),
			cmd.Delete(),
			cmd.List(),
			cmd.Run(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
