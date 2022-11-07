/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package helpers

import "github.com/sveltinio/sveltin/internal/tpltypes"

// NewMenuItems return a NoPageItems.
func NewMenuItems(resources []string, content map[string][]string) *tpltypes.MenuItems {
	r := new(tpltypes.MenuItems)
	r.Resources = resources
	r.Content = content
	return r
}
