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

var branch string

func main() {
	app := cli.NewApp()
	app.Name = "gits"
	app.Usage = "check git status"
	app.Action = func(c *cli.Context) {
		filepath.Walk(".", walker(status))
	}

	app.Commands = []cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Action: func(c *cli.Context) {
				filepath.Walk(".", walker(fetch))
			},
		},
		{
			Name:    "pull",
			Aliases: []string{"p"},
			Action: func(c *cli.Context) {
				branch = c.Args().First()
				filepath.Walk(".", walker(pull))
			},
		},
		{
			Name:    "checkout",
			Aliases: []string{"co"},
			Action: func(c *cli.Context) {
				branch = c.Args().First()
				filepath.Walk(".", walker(checkout))
			},
		},
	}

	app.Run(os.Args)
}

func walker(fn filepath.WalkFunc) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
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
		return fn(path, info, err)
	}
}

func status(path string, info os.FileInfo, err error) error {
	fmt.Println("On branch " + currentBranch(path))

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

func fetch(path string, info os.FileInfo, err error) error {
	cmd := exec.Command("git", "fetch")
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
	args := []string{"pull"}
	if branch != "" {
		args = append(args, "origin", branch)
	}
	cmd := exec.Command("git", args...)
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

func isGit(dir string) bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmd.Dir = dir
	res, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	return strings.TrimSpace(string(res)) == "true"
}

func currentBranch(path string) string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = path
	res, _ := cmd.CombinedOutput()
	return strings.TrimSpace(string(res))
}
