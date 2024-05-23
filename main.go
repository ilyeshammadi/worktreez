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
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			// TODO: Init command to setup the default profile
			// TODO: Profile command to manage non default profile by name
			cmd.Create(),
			cmd.Delete(),
			cmd.List(), // TODO: Fix the CD isch issue
			cmd.Run(),
			// TODO: Add a new command to append new repos to existing worktree branches
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
