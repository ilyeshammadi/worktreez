package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

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
			GetRepoNamesFlag(),
			&cli.BoolFlag{
				Name:     "force",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			reposPath := ctx.String("repos_path")
			destPath := ctx.String("dest_path")
			branchName := ctx.String("branch_name")

			for _, repoName := range ctx.StringSlice("repo_name") {
				repoPath, _ := filepath.Abs(filepath.Join(reposPath, repoName))
				destRepoPath, _ := filepath.Abs(filepath.Join(destPath, branchName, repoName))

				if err := removeWorktree(ctx, repoPath, destRepoPath); err != nil {
					return err
				}

				if err := removeBranch(ctx, repoPath, branchName); err != nil {
					return err
				}

				if err := removeBranchNameFolder(ctx); err != nil {
					return err
				}

				if err := prune(ctx, repoPath); err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func removeWorktree(ctx *cli.Context, repoPath, destRepoPath string) cli.ExitCoder {
	// Remove work tree
	removeCmd := exec.Command(
		"git",
		"-C",
		repoPath,
		"worktree",
		"remove",
		destRepoPath,
	)
	if ctx.Bool("dry_run") {
		fmt.Println(removeCmd.String())
	} else {
		output, err := removeCmd.CombinedOutput()
		if err != nil {
			errMessage := fmt.Sprintf("Command execution failed: %v\nOutput: %s", err, string(output))
			return cli.Exit(errMessage, 1)
		}
	}
	return nil
}

func prune(ctx *cli.Context, repoPath string) cli.ExitCoder {
	// Prune work tree
	pruneCmd := exec.Command(
		"git",
		"-C",
		repoPath,
		"worktree",
		"prune",
	)
	if ctx.Bool("dry_run") {
		fmt.Println(pruneCmd.String())
	} else {
		output, err := pruneCmd.CombinedOutput()
		if err != nil {
			errMessage := fmt.Sprintf("Command execution failed: %v\nOutput: %s", err, string(output))
			return cli.Exit(errMessage, 1)
		}
	}
	return nil
}

func removeBranch(ctx *cli.Context, repoPath, branchName string) cli.ExitCoder {
	cmdArgs := []string{
		"-C",
		repoPath,
		"branch",
		"-D",
		branchName,
	}

	if ctx.Bool("force") {
		cmdArgs = append(cmdArgs, "--force")
	}

	cmd := exec.Command(
		"git",
		cmdArgs...,
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
	return nil
}

func removeBranchNameFolder(ctx *cli.Context) cli.ExitCoder {
	folderName, err := filepath.Abs(filepath.Join(ctx.String("dest_path"), ctx.String("branch_name")))
	if err != nil {
		errMessage := fmt.Sprintf("Incorrect path to foler %s", folderName)
		return cli.Exit(errMessage, 1)
	}
	cmd := exec.Command(
		"rm",
		"-rf",
		folderName,
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
	return nil
}
