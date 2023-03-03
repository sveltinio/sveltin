/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package utils

import "github.com/tidwall/sjson"

// SetJsonStringValue sets a json string value for the specified path.
func SetJsonStringValue(content []byte, name, value string) ([]byte, error) {
	return sjson.SetBytes(content, name, value)
}
