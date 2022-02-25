/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package utils ...
package utils

import (
	"fmt"

	jww "github.com/spf13/jwalterweatherman"
)

// CheckIfError panics on os.Exit(1) if error
func CheckIfError(err error) {
	if err == nil {
		return
	}
	jww.FATAL.Fatalf("\x1b[31;1m✘ %s\x1b[0m\n", fmt.Sprintf("error: %s", err))
}
