package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"
)

func status(path string, info os.FileInfo, args []string) error {
	fmt.Printf("On branch %s\n\n", currentBranch(path))

	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = path
	res, _ := cmd.CombinedOutput()
	if str := string(res); !isWhitespace(str) {
		fmt.Println(str)
	}

	return filepath.SkipDir
}

func currentBranch(path string) string {
	return executeTrimmed(path, "git", "rev-parse", "--abbrev-ref", "HEAD")
}

func isGit(path string) bool {
	res := executeTrimmed(path, "git", "rev-parse", "--is-inside-work-tree")
	return res == "true"
}

func execute(path string, info os.FileInfo, args []string) error {
	fmt.Println(executeTrimmed(path, "git", args...))
	return filepath.SkipDir
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
