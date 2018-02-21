package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var branch string

func main() {
	app := cli.NewApp()
	app.Name = "gits"
	app.Usage = "recursive git commands"
	app.Action = func(c *cli.Context) {
		filepath.Walk(".", walker(status))
	}

	app.Commands = []cli.Command{
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   `Recursive "git fetch -p"`,
			Action: func(c *cli.Context) {
				filepath.Walk(".", walker(fetch))
			},
		},
		{
			Name:      "pull",
			Aliases:   []string{"p"},
			Usage:     `Recursive "git pull origin <branch>"`,
			ArgsUsage: "<branch> - If no branch is supplied, current branch is used.",
			Action: func(c *cli.Context) {
				branch = c.Args().First()
				filepath.Walk(".", walker(pull))
			},
		},
		{
			Name:      "checkout",
			Aliases:   []string{"co"},
			Usage:     `Recursive "git checkout <branch>"`,
			ArgsUsage: "<branch> - Branch name is required.",
			Action: func(c *cli.Context) {
				branch = c.Args().First()
				if branch == "" {
					fmt.Println("Branch name to checkout is required.")
					return
				}
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

		color.Cyan("* " + path)
		return fn(path, info, err)
	}
}
