/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package npmc ...
package npmc

import (
	"encoding/json"
	"fmt"

	jww "github.com/spf13/jwalterweatherman"
)

// Parse parses the JSON-encoded data and stores the result in the value pointed to by v.
func Parse(content []byte) *PackageJSON {
	var pkgParsed PackageJSON
	err := json.Unmarshal(content, &pkgParsed)
	if err != nil {
		jww.FATAL.Fatalf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	}
	return &pkgParsed
}
