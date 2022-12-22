package cmd

import (
	"github.com/spf13/cobra"

	"github.com/opensourceways/issue-cli/util"
)

type createOption struct {
	Streams
}

func newCreateOption(s base) *createOption {
	return &createOption{Streams: s.Streams}
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
