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
	fmt.Println("On branch", currentBranch(path))

	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = path
	res, _ := cmd.CombinedOutput()
	str := string(res)
	if !isWhitespace(str) {
		fmt.Println("\n" + str)
	}

	return filepath.SkipDir
}

// fetch prunes the remote branches with "-p" option.
func fetch(path string, info os.FileInfo, err error) error {
	cmd := exec.Command("git", "fetch", "-p")
	cmd.Dir = path
	res, _ := cmd.CombinedOutput()
	clean := strings.TrimSpace(string(res))
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

	cmd := exec.Command("git", "pull", "origin", branch)
	cmd.Dir = path
	res, _ := cmd.CombinedOutput()
	fmt.Println(strings.TrimSpace(string(res)))

	return filepath.SkipDir
}

func checkout(path string, info os.FileInfo, err error) error {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = path
	res, _ := cmd.CombinedOutput()
	fmt.Println(strings.TrimSpace(string(res)))

	return filepath.SkipDir
}

func currentBranch(path string) string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = path
	res, _ := cmd.CombinedOutput()
	return strings.TrimSpace(string(res))
}

func isGit(dir string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = dir
	res, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(res)) == "true"
}

func isWhitespace(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}

	return true
}
