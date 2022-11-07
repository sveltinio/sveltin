/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package helpers contains helper functions used across commands.
package helpers

import (
	"bytes"

	"github.com/spf13/cobra"
)

// ExecuteCommandC is a @DEPRECATED function, not used anymore.
func ExecuteCommandC(root *cobra.Command, args ...string) (c *cobra.Command, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, err
}
