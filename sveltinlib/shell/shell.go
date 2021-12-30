/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package shell

import "context"

type Shell interface {
	Execute(string, string, bool) error
	BackgroundExecute(context.Context, string, string, string) ([]byte, error)
}
