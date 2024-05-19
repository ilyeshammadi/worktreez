package utils

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func RunCommand(cmd *exec.Cmd, dryRun bool) cli.ExitCoder {
	if dryRun {
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
