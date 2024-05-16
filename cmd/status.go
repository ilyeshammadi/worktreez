package cmd

import "github.com/urfave/cli/v2"

func Status() *cli.Command {
	return &cli.Command{
		Name:        "status",
		Description: "Get the status for each repo",
		Aliases:     []string{"s"},
		Flags: []cli.Flag{
			GetReposFlag(),
			GetDestPathFlag(),
			GetBranchNameFlag(),
			GetRepoNamesFlag(),
		},
		Action: func(ctx *cli.Context) error {
			return nil
		},
	}
}
