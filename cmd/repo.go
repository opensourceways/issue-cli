package cmd

import (
	"fmt"
	"io"
	"net/url"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/opensourceways/issue-cli/util"
)

type repoOption struct {
	Streams
	h *util.Request

	page int
	size int

	name string
}

func newRepoOption(s Streams) *repoOption {
	return &repoOption{Streams: s, h: util.NewRequest(nil)}
}

func newRepoCmd(s Streams) *cobra.Command {
	o := newRepoOption(s)

	cmd := &cobra.Command{
		Use:   "repo [options]",
		Short: "obtain information about the repository that openeuler can use to create an issue",
		Long:  "obtain information about the repository that openeuler can use to create an issue",
		Run: func(cmd *cobra.Command, args []string) {
			checkErr(o.Validate(cmd, args))
			checkErr(o.Run())
		},
	}

	cmd.Flags().StringVarP(&o.name, "name", "n", "", "repo name")
	cmd.Flags().IntVarP(&o.page, "page", "p", 1, "get the number of pages for the repo")
	cmd.Flags().IntVarP(&o.size, "size", "s", 20, "get the number of sizes for the repo")

	return cmd
}

func (o *repoOption) Run() error {
	var v = url.Values{}
	v.Add("page", strconv.Itoa(o.page))
	v.Add("per_page", strconv.Itoa(o.size))
	v.Add("keyword", o.name)

	var res = struct {
		basePageResp
		Result []struct {
			Name string `json:"repo"`
			Id   int64  `json:"enterprise_number"`
		} `json:"data"`
	}{}

	_, err := o.h.CustomRequest(RepoUrl, "GET", nil, nil, v, &res)
	if err != nil {
		return err
	}

	if err = o.printContextHeaders(o.Out); err != nil {
		return err
	}
	for _, s := range res.Result {
		_, err = fmt.Fprintf(o.Out, "%-15d\t%s\n", s.Id, s.Name)
	}
	return err
}

func (o *repoOption) Validate(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return util.UsageErrorf(cmd, "unexpected args: %v", args)
	}
	return nil
}

func (o *repoOption) printContextHeaders(out io.Writer) error {
	columnNames := []any{"REPONUMBER", "REPONAME"}
	_, err := fmt.Fprintf(out, "%-15s\t%s\n", columnNames...)
	return err
}
