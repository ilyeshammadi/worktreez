package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func runCommand(cmd *exec.Cmd, dryRun, withOutput bool) cli.ExitCoder {
	if dryRun {
		fmt.Println(cmd.String())
	} else {
		output, err := cmd.CombinedOutput()
		if err != nil {
			errMessage := fmt.Sprintf("Command execution failed: %v\nOutput: %s", err, string(output))
			return cli.Exit(errMessage, 1)
		}
		if withOutput {
			fmt.Println(string(output))
		}
	}
	return nil
}

func RunCommand(cmd *exec.Cmd, dryRun bool) cli.ExitCoder {
	return runCommand(cmd, dryRun, false)
}

func RunCommandWithOutput(cmd *exec.Cmd, dryRun bool) cli.ExitCoder {
	return runCommand(cmd, dryRun, true)
}

func CheckIn(value string, values []string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func GetIcon(icon string) string {
	if EnableIcons {
		return icon
	} else {
		return ""
	}
}

func PrintBranch(branchDir string) {
	indent := ""
	fmt.Println(ColorPurple + GetIcon(" ") + filepath.Base(branchDir) + ColorReset)

	repos, err := os.ReadDir(branchDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, repo := range repos {
		last := i == len(repos)-1
		fmt.Print(indent)
		if last {
			fmt.Print(ColorGray + "└── " + ColorReset)
		} else {
			fmt.Print(ColorGray + "├── " + ColorReset)
		}
		// Check if repo has uncommited changes
		cmd := exec.Command("git", "-C", filepath.Join(branchDir, repo.Name()), "status", "--porcelain")
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error executing git command: %v\n", err)
			return
		}
		fmt.Print(ColorBlue + GetIcon(" ") + repo.Name() + ColorReset + " ")
		if len(output) > 0 {
			fmt.Print(ColorRed, UncommitedSymbol, ColorReset)
		}
		fmt.Println()
	}
	fmt.Println()
}
