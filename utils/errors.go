/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package utils

import (
	"fmt"

	jww "github.com/spf13/jwalterweatherman"
)

// Panics if an error is not nil
func CheckIfError(err error) {
	if err == nil {
		return
	}
	jww.FATAL.Fatalf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
}
