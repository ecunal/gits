package main

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

var branch string
var commitMessage string

func main() {
	if len(os.Args) == 1 {
		filepath.Walk(".", walker(status))
		return
	}

	filepath.Walk(".", walker(execute))
}

func walker(fn func(path string, info os.FileInfo, args []string) error) filepath.WalkFunc {
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
		return fn(path, info, os.Args[1:])
	}
}
