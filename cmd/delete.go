package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"worktreez/utils"

	"github.com/urfave/cli/v2"
)

func Delete() *cli.Command {
	return &cli.Command{
		Name:        "delete",
		Description: "Delete all working trees with branch name",
		Aliases:     []string{"d"},
		Flags: []cli.Flag{
			DryRunFlag(),
			GetReposFlag(),
			GetDestPathFlag(),
			GetBranchNameFlag(),
			&cli.BoolFlag{
				Name:     "force",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			reposPath := ctx.String("repos_path")
			destPath := ctx.String("dest_path")
			branchName := ctx.String("branch_name")

			branchFodler, _ := filepath.Abs(filepath.Join(destPath, branchName))

			// Make a copy of the repo names in the branch folder
			repoNames, err := os.ReadDir(branchFodler)
			if err != nil {
				return cli.Exit(err, 1)
			}

			// Delete the branch folder
			deleteBranchCmd := exec.Command(
				"rm",
				"-rf",
				branchFodler,
			)
			utils.RunCommand(deleteBranchCmd, ctx.Bool("dry_run"))

			// Prune each folder in the repos path
			for _, repoName := range repoNames {
				repoPath, _ := filepath.Abs(filepath.Join(reposPath, repoName.Name()))
				pruneCmd := exec.Command(
					"git",
					"-C",
					repoPath,
					"worktree",
					"prune",
				)
				utils.RunCommand(pruneCmd, ctx.Bool("dry_run"))
			}

			return nil
		},
	}
}
