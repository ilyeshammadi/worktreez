package cmd

import (
	"fmt"
	"worktreez/utils"

	"github.com/urfave/cli/v2"
)

func DryRunFlag() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:     "dry_run",
		Required: false,
	}
}

func GetReposFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     "repos_path",
		Aliases:  []string{"r"},
		Required: true,
		Action: func(ctx *cli.Context, s string) error {
			if !utils.IsValidPath(s) {
				return cli.Exit(fmt.Sprintf("Invalid repos_path: %s", s), 1)
			}
			return nil
		},
	}
}

func GetDestPathFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     "dest_path",
		Aliases:  []string{"d"},
		Required: true,
		Action: func(ctx *cli.Context, s string) error {
			if !utils.IsValidPath(s) {
				return cli.Exit(fmt.Sprintf("Invalid dest_path: %s", s), 1)
			}
			return nil
		},
	}
}

func GetBranchNameFlag(required bool) *cli.StringFlag {
	return &cli.StringFlag{
		Name:     "branch_name",
		Aliases:  []string{"b"},
		Required: required,
		Action: func(ctx *cli.Context, s string) error {
			if !utils.IsValidBranchName(s) {
				return cli.Exit(fmt.Sprintf("Invalid branch name %s", s), 1)
			}
			return nil
		},
	}
}

func GetRepoNamesFlag() *cli.StringSliceFlag {
	return &cli.StringSliceFlag{
		Name:     "repo_name",
		Aliases:  []string{"n"},
		Required: true,
		Action: func(ctx *cli.Context, s []string) error {
			reposPath := ctx.String("repos_path")
			for _, elem := range s {
				if !utils.IsValidGitRepository(reposPath, elem) {
					return cli.Exit(fmt.Sprintf("Invalid repository %s", elem), 1)
				}
			}
			return nil
		},
	}
}
