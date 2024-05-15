package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
)

func IsValidPath(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}

	if !info.IsDir() {
		return false
	}
	return true
}

func IsValidBranchName(branchName string) bool {
	// Define a regular expression to match valid branch names
	return regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(branchName)
}

func IsValidGitRepository(basePath, repoName string) bool {
	// Construct the expected path for the repository
	repoPath := filepath.Join(basePath, repoName)

	if !IsValidPath(repoPath) {
		return false
	}

	// Verify if the directory is a valid Git repository by using 'git rev-parse'
	cmd := exec.Command("git", "-C", repoPath, "rev-parse")
	if err := cmd.Run(); err != nil {
		// The command failed, so the directory is not a valid Git repository
		return false
	}

	// The directory is a valid Git repository
	return true
}
