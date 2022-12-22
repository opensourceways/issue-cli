package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/opensourceways/issue-cli/util"
)

type issueOption struct {
	Streams
	h *util.Request

	filepath string
	title    string
	repoid   int64
	issueid  int64

	email, code string
}

func newIssueOption(s base) *issueOption {
	return &issueOption{Streams: s.Streams, h: util.NewRequest(nil)}
}

var createExample = `
 # Create issue
 issue-cli create issue -f [file path] -t [title] -r [repo] -i [issue id]
`

func newCmdIssue(s base) *cobra.Command {
	o := newIssueOption(s)

	cmd := &cobra.Command{
		Use:     "issue [flags]",
		Short:   "create issue for openeuler",
		Example: createExample,
		Run: func(cmd *cobra.Command, args []string) {
			checkErr(o.Validate(cmd, args))
			checkErr(o.Run())
		},
	}

	cmd.Flags().StringVarP(&o.filepath, "file", "f", o.filepath, "issue body file path")
	cmd.Flags().StringVarP(&o.title, "title", "t", o.title, "issue title")
	cmd.Flags().Int64VarP(&o.repoid, "repo", "r", o.repoid, "create an issue in that repository")
	cmd.Flags().Int64VarP(&o.issueid, "issue", "i", o.issueid, "create an issue template id")

	return cmd
}

func (c *issueOption) Validate(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return util.UsageErrorf(cmd, "unexpected args: %v", args)
	}

	if len(c.filepath) <= 0 {
		return fmt.Errorf("please enter file path")
	}

	if len(c.title) <= 0 {
		return fmt.Errorf("please enter the issue title")
	}

	if c.repoid <= 0 {
		return fmt.Errorf("please specify the repo id,you can use `issue get repo`")
	}

	if c.issueid <= 0 {
		return fmt.Errorf("please specify the issue-type id,you can use `issue get it`")
	}

	return nil
}

func (c *issueOption) Run() error {
	var email string
	fmt.Println("请输入邮箱:")
	_, err := fmt.Fscan(c.In, &email)
	if err != nil {
		return err
	}

	var res = baseResp{}

	_, err = c.h.CustomRequest(VerifyUrl, "post", fmt.Sprintf(`{"email":"%s"}`, email), nil, nil, &res)
	if err != nil {
		return err
	}

	if res.Code != 200 {
		return fmt.Errorf(res.Msg)
	}
	c.email = email

	fmt.Println("验证码已经发送至您的邮箱,请查看后输入:")
	var code string
	_, err = fmt.Fscan(c.In, &code)
	if err != nil {
		return err
	}
	c.code = code
	return c.createIssue()
}

func (c *issueOption) createIssue() error {
	bys, err := os.ReadFile(c.filepath)
	if err != nil {
		return err
	}
	var req = struct {
		Id      int64  `json:"project_id"`
		IssueId int64  `json:"issue_type_id"`
		Title   string `json:"title"`
		Body    string `json:"description"`
		Email   string `json:"email"`
		Code    string `json:"code"`
	}{Id: c.repoid, Title: c.title, Body: string(bys), Email: c.email, Code: c.code, IssueId: c.issueid}

	bys, err = json.Marshal(req)
	if err != nil {
		return err
	}

	var res = struct {
		baseResp
	}{}

	_, err = c.h.CustomRequest(CreateIssueUrl, "POST", bys, nil, nil, &res)
	if err != nil {
		return err
	}

	if res.Code != 201 {
		return fmt.Errorf(res.Msg)
	}

	_, err = fmt.Fprintln(c.Out, "success")
	return err
}
