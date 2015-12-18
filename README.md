# gits
A small tool to execute recursive git commands in a directory.

When run without any arguments, executes `git status` on current directory recursively and lists the output(s).

[![Go Report Card](http://goreportcard.com/badge/ecunal/gits)](http://goreportcard.com/report/ecunal/gits)

### Commands

```
* fetch, f
    Recursive "git fetch -p"
* pull, p [branch]
    Recursive "git pull origin [branch]"
    If no branch is supplied, current branch is used.
* checkout, co branch
    Recursive "git checkout branch"
    Branch name is required.
* help, h
    Shows a list of commands or help for one command
```

### TODO
- [ ] Enable custom git commands
