package main

import (
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
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

		color.Cyan("\n* " + path)
		return fn(path, info, err)
	}
}
