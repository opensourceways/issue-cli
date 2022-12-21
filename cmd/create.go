package cmd

import (
	"github.com/opensourceways/issue-cli/util"
	"github.com/spf13/cobra"
)

type createOption struct {
	Streams
	h util.ReqImpl
}

func newCreateOption(s base) *createOption {
	return &createOption{Streams: s.Streams, h: s.ReqImpl}
}

func newCmdCreate(s base) *cobra.Command {
	o := newCreateOption(s)

	cmd := &cobra.Command{
		Use:     "create [command]",
		Short:   "create resources for openeuler",
		Example: createExample,
		Run: func(cmd *cobra.Command, args []string) {
			checkErr(o.Validate(cmd, args))
			checkErr(cmd.Help())
		},
	}

	cmd.AddCommand(newCmdIssue(s))

	return cmd
}

func (c *createOption) Validate(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return util.UsageErrorf(cmd, "extra arguments: %v", args)
	}

	return nil
}
