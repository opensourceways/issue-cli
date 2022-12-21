package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/opensourceways/issue-cli/util"
	"github.com/spf13/cobra"
)

var issue = &cobra.Command{
	Use:   "issue-cli",
	Long:  "issue-cli command can create openeuler issue",
	Short: "issue-cli command can create openeuler issue",
	Run: func(cmd *cobra.Command, args []string) {
		checkErr(cmd.Help())
	},
}

func checkErr(err error) {
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error()+"\n")
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
	util.ReqImpl
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
	s := base{
		Streams: Streams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
		ReqImpl: util.NewRequest(nil),
	}

	group := CommandGroups{
		{
			Commands: []*cobra.Command{
				newCmdGet(s),
			},
		},
	}

	group.Add(issue)

	//not create a default 'completion' command
	issue.CompletionOptions.DisableDefaultCmd = true

	return issue
}
