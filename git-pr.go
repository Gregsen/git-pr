package main

import (
	"fmt"
	"os/exec"
)

func main() {
	branch, err := currentBranch()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(branch)
}

type GitRepo struct {
	branch string
}

func NewGitRepo() *GitRepo {
	var branch string
	var err error
	if IsGitRepo() {
		if branch, err = currentBranch(); err != nil {
			return nil
		}
	}
	return &GitRepo{branch: branch}

}

// IsGitRepo checks if the current directory is a Git repository.
// It returns a bool.
func IsGitRepo() bool {
	repo := exec.Command("git", "rev-parse", "-q", "--git-dir")
	if err := repo.Run(); err != nil {
		return false
	}
	return true
}

// IsClean checks if the repository has no uncommitted changes
// or untracked files. Returns a bool.
func IsClean() bool {
	comits := exec.Command("git", "diff-index", "--quiet", "HEAD", "--ignore-submodules", "--")
	if err := comits.Run(); err != nil {
		return false
	}
	files := exec.Command("git", "diff-files", "--quiet", "HEAD", "--ignore-submodules", "--")
	if err := files.Run(); err != nil {
		return false
	}
	return true
}

func currentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branch, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(branch), err
}
