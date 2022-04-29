/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package common ...
package common

// UnionMap returns a map joining elements from the inputs.
func UnionMap(m1, m2 map[string]string) map[string]string {
	for ia, va := range m1 {
		if it, ok := m2[ia]; ok {
			va += it
		}
		m2[ia] = va
	}
	return m2
}
