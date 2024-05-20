package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:        "list",
		Description: "List existing branches and their repos",
		Aliases:     []string{"ls"},
		Flags: []cli.Flag{
			GetDestPathFlag(),
			GetBranchNameFlag(false),
		},
		Action: func(ctx *cli.Context) error {
			destPath := ctx.String("dest_path")
			selectedBranchName := ctx.String("branch_name")

			// List all the repos and their content
			branches, err := os.ReadDir(destPath)
			if err != nil {
				return cli.Exit(err, 1)
			}

			for _, branchName := range branches {
				if selectedBranchName != "" && selectedBranchName != branchName.Name() {
					continue
				}
				fmt.Println("Branch:", branchName.Name())

				branchFodler, _ := filepath.Abs(filepath.Join(destPath, branchName.Name()))
				repoNames, err := os.ReadDir(branchFodler)

				for _, repoName := range repoNames {
					fmt.Println(repoName.Name())
				}

				if err != nil {
					return cli.Exit(err, 1)
				}
			}

			return nil
		},
	}
}
