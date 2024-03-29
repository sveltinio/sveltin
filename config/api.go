/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package config

// API represents the API folder structure in a Sveltin project.
type API struct {
	Version  string `mapstructure:"version"`
	Filename string `mapstructure:"filename"`
}
