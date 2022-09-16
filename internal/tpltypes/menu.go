/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package tpltypes

// MenuData is the struct representing a menu item.
type MenuData struct {
	Items       *MenuItems
	WithContent bool
}

// MenuItems is a struct representing a resource and its content as menu item.
type MenuItems struct {
	Name      string
	Resources []string
	Content   map[string][]string
}
