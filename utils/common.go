package utils

import (
	"fmt"
	"os/exec"

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
