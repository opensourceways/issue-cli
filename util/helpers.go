package util

import (
	"fmt"

	"github.com/spf13/cobra"
)

func UsageErrorf(cmd *cobra.Command, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\nsee '%s -h' for help and examples", msg, cmd.CommandPath())
}
