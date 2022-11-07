/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package shell defines OS level ways to interact with node package managers and git command.
package shell

import "context"

// Shell is the interface defining the methods to be implemented by a shell instance.
type Shell interface {
	Execute(string, string, bool) error
	BackgroundExecute(context.Context, string, string, string) ([]byte, error)
}
