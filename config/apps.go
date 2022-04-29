/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

// AppTemplate is the stuct presenting the github repository
// used by sveltin to clone the starter project repos.
type AppTemplate struct {
	Name string
	URL  string
}
