package helpers

import (
	"bytes"

	"github.com/spf13/cobra"
)

func ExecuteCommandC(root *cobra.Command, args ...string) (c *cobra.Command, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, err
}
