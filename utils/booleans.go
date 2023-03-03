/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package utils

// NewTrue set a bool pointer to true and returns it.
func NewTrue() *bool {
	b := true
	return &b
}

// NewFalse set a bool pointer to false and returns it.
func NewFalse() *bool {
	b := false
	return &b
}
