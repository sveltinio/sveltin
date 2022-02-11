/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package npmc

import (
	"encoding/json"
	"fmt"

	jww "github.com/spf13/jwalterweatherman"
)

func Parse(content []byte) *PackageJson {
	var pkgParsed PackageJson
	err := json.Unmarshal(content, &pkgParsed)
	if err != nil {
		jww.FATAL.Fatalf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	}
	return &pkgParsed
}
