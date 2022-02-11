/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package packagejson

import (
	"encoding/json"

	"github.com/sveltinio/sveltin/utils"
)

func Parse(content []byte) *PackageJson {
	var pkgParsed PackageJson

	err := json.Unmarshal(content, &pkgParsed)
	utils.CheckIfError(err)

	return &pkgParsed
}
