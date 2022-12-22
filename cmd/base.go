package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var issue = &cobra.Command{
	Use:   "issue-cli",
	Long:  "issue-cli command can create openeuler resources, example: issue ...",
	Short: "issue-cli command can create openeuler resources",
	Run: func(cmd *cobra.Command, args []string) {
		checkErr(cmd.Help())
	},
}

func checkErr(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// CommandGroup command group
type CommandGroup struct {
	Message  string
	Commands []*cobra.Command
}

type CommandGroups []CommandGroup

func (g CommandGroups) Add(c *cobra.Command) {
	for _, group := range g {
		c.AddCommand(group.Commands...)
	}
}

type base struct {
	Streams
}

type baseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type basePageResp struct {
	Total   int `json:"total"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type Streams struct {
	// In default os.Stdin
	In io.Reader

	// Out default os.Stdout
	Out io.Writer

	// ErrOut default os.Stderr
	ErrOut io.Writer
}

func Cmd() *cobra.Command {
	b := base{Streams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}}

	group := CommandGroups{
		{
			Commands: []*cobra.Command{
				newCmdGet(b),
				newCmdCreate(b),
			},
		},
	}

	group.Add(issue)

	//not create a default 'completion' command
	issue.CompletionOptions.DisableDefaultCmd = true

	issue.SetIn(b.In)
	issue.SetOut(b.Out)
	issue.SetErr(b.ErrOut)

	return issue
}
