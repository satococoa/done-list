package main

import (
	"fmt"
	"io"
	"os"

	"github.com/satococoa/done-list/lib"
)

const (
	ExitCodeOK = iota
	ExitCodeSettingError
)

type CLI struct{ outStream, errStream io.Writer }

func (c *CLI) Run(args []string) int {
	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if githubAccessToken == "" {
		fmt.Fprintln(c.errStream, "Please set GITHUB_ACCESS_TOKEN")
		return ExitCodeSettingError
	}
	client := donelist.CreateClient(githubAccessToken)
	issues, err := donelist.FetchIssues(client)
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		os.Exit(ExitCodeSettingError)
	}
	donelist.PrintDoneList(c.outStream, issues)
	return ExitCodeOK
}

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stderr,
	}
	os.Exit(cli.Run(os.Args))
}
