package main

import (
	"log"

	"github.com/opensourceways/issue-cli/cmd"
)

func main() {
	issueCmd := cmd.Cmd()

	err := issueCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
