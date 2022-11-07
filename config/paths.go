/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package config

// Paths is the struct mapping the folders structure for a sveltin project.
type Paths struct {
	Build   string `mapstructure:"build"`
	Config  string `mapstructure:"config"`
	Content string `mapstructure:"content"`
	Static  string `mapstructure:"static"`
	Themes  string `mapstructure:"themes"`
	Src     string `mapstructure:"src"`
	Params  string `mapstructure:"params"`
	Lib     string `mapstructure:"lib"`
	Routes  string `mapstructure:"routes"`
	API     string `mapstructure:"api"`
}
