package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func Create() *cli.Command {
	return &cli.Command{
		Name:        "create",
		Description: "Create a new set of worktrees for the selected repos",
		Aliases:     []string{"c"},
		Flags: []cli.Flag{
			DryRunFlag(),
			GetReposFlag(),
			GetDestPathFlag(),
			GetBranchNameFlag(),
			GetRepoNamesFlag(),
		},
		Action: func(ctx *cli.Context) error {
				reposPath := ctx.String("repos_path")
			destPath := ctx.String("dest_path")
			branchName := ctx.String("branch_name")

			for _, repoName := range ctx.StringSlice("repo_name") {
				repoPath, _ := filepath.Abs(filepath.Join(reposPath, repoName))
				repoPathGit, _ := filepath.Abs(filepath.Join(repoPath, ".git"))
				destRepoPath, _ := filepath.Abs(filepath.Join(destPath, branchName, repoName))

				cmd := exec.Command(
					"git",
					"--git-dir",
					repoPathGit,
					"--work-tree",
					repoPath,
					"worktree",
					"add",
					"-B",
					branchName,
					destRepoPath,
				)

				if ctx.Bool("dry_run") {
					fmt.Println(cmd.String())
				} else {
					output, err := cmd.CombinedOutput()
					if err != nil {
						errMessage := fmt.Sprintf("Command execution failed: %v\nOutput: %s", err, string(output))
						return cli.Exit(errMessage, 1)
					}
				}
			}

			return nil
		},
	}
}
