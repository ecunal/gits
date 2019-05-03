package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"
)

func status(path string, info os.FileInfo, err error) error {
	fmt.Printf("On branch %s\n\n", currentBranch(path))

	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = path
	res, _ := cmd.CombinedOutput()
	if str := string(res); !isWhitespace(str) {
		fmt.Println(str)
	}

	return filepath.SkipDir
}

// fetch prunes the remote branches with "-p" option.
func fetch(path string, info os.FileInfo, err error) error {
	clean := executeTrimmed(path, "git", "fetch", "-p")
	if clean == "" {
		fmt.Println("Already up-to-date.")
	} else {
		fmt.Println(clean)
	}

	return filepath.SkipDir
}

func pull(path string, info os.FileInfo, err error) error {
	if branch == "" {
		branch = currentBranch(path)
	}

	fmt.Println(executeTrimmed(path, "git", "pull", "origin", branch))
	return filepath.SkipDir
}

func checkout(path string, info os.FileInfo, err error) error {
	fmt.Println(executeTrimmed(path, "git", "checkout", branch))
	return filepath.SkipDir
}

func currentBranch(path string) string {
	return executeTrimmed(path, "git", "rev-parse", "--abbrev-ref", "HEAD")
}

func isGit(path string) bool {
	res := executeTrimmed(path, "git", "rev-parse", "--is-inside-work-tree")
	return res == "true"
}

func executeTrimmed(path, command string, arg ...string) string {
	cmd := exec.Command(command, arg...)
	cmd.Dir = path

	res, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(res))
}

func isWhitespace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}

func diff(path string, info os.FileInfo, err error) error {
	fmt.Println(executeTrimmed(path, "git", "diff"))
	return filepath.SkipDir
}

func add(path string, info os.FileInfo, err error) error {
	fmt.Println(executeTrimmed(path, "git", "add", "-u"))
	return filepath.SkipDir
}

func commit(path string, info os.FileInfo, err error) error {
	fmt.Println(executeTrimmed(path, "git", "commit", "-m", commitMessage))
	return filepath.SkipDir
}
