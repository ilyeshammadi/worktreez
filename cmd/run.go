package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"worktreez/utils"

	"github.com/urfave/cli/v2"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:        "run",
		UsageText:   "worktreez run -d dest_folder -b BRANCH -- git status",
		Description: "Run a shell command against the specified branch/repository",
		Aliases:     []string{"r"},
		Flags: []cli.Flag{
			GetDestPathFlag(),
			GetBranchNameFlag(true),
			GetDestRepoNamesFlag(),
			DryRunFlag(),
		},
		Action: func(ctx *cli.Context) error {
			destPath := ctx.String("dest_path")
			branchName := ctx.String("branch_name")

			if !ctx.Args().Present() {
				return cli.Exit("missing command", 1)
			}

			branchFodler, _ := filepath.Abs(filepath.Join(destPath, branchName))

			repoNames, err := os.ReadDir(branchFodler)
			if err != nil {
				return cli.Exit(err, 1)
			}

			for _, repoName := range repoNames {
				if len(ctx.StringSlice("repo_name")) > 0 && !utils.CheckIn(repoName.Name(), ctx.StringSlice("repo_name")) {
					continue
				}

				repoPath, _ := filepath.Abs(filepath.Join(destPath, branchName, repoName.Name()))
				fmt.Println(repoPath)

				cmd := exec.Command(ctx.Args().First(), ctx.Args().Tail()...)
				cmd.Dir = repoPath
				utils.RunCommandWithOutput(cmd, ctx.Bool("dry_run"))
			}

			return nil
		},
	}
}
