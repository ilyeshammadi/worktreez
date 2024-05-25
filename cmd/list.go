package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"worktreez/utils"

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

			fmt.Printf("%s%s%s Uncommited changes\n", utils.ColorRed, utils.UncommitedSymbol, utils.ColorReset)
			fmt.Println()
			for _, branchName := range branches {
				if selectedBranchName != "" && selectedBranchName != branchName.Name() {
					continue
				}
				branchFodler, _ := filepath.Abs(filepath.Join(destPath, branchName.Name()))
				utils.PrintBranch(branchFodler)

			}

			return nil
		},
	}
}
