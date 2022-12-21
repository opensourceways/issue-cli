package main

import (
	"log"
	"os"

	"github.com/opensourceways/issue-cli/cmd"
)

func main() {
	issueCmd := cmd.Cmd()
	issueCmd.SetErr(os.Stderr)
	issueCmd.SetOut(os.Stdout)
	issueCmd.SetIn(os.Stdin)

	err := issueCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
