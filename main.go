package main

import (
	"log"
	"os"

	"worktreez/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Authors: []*cli.Author{
			{
				Name:  "Ilyes Hammadi",
				Email: "hammadiilyesahmed@gmail.com",
			},
		},
		Name:                 "worktreez",
		Usage:                "hammadiilyesahmed@gmail.com",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			cmd.Create(),
			cmd.Delete(),
			cmd.List(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
