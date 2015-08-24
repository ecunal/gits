package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
)

// TODO verbose option
// TODO fetch

func main() {
	app := cli.NewApp()
	app.Name = "gits"
	app.Usage = "check git status"
	app.Action = func(c *cli.Context) {
		filepath.Walk(".", check)
	}

	app.Run(os.Args)
}

func check(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if !info.IsDir() {
		return nil
	}

	if !isGit(path) {
		return nil
	}

	fmt.Println("* " + path)

	// branch
	// git rev-parse --abbrev-ref HEAD
	cmdb := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmdb.Dir = path
	resb, _ := cmdb.CombinedOutput()
	fmt.Println("On branch " + strings.TrimSpace(string(resb)))

	// status
	// git status --porcelain
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = path
	res, _ := cmd.CombinedOutput()
	clean := strings.TrimSpace(string(res))
	if clean == "" {
		fmt.Println("Working tree clean")
	} else {
		fmt.Println(clean)
	}

	return filepath.SkipDir
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
