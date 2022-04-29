/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

// Paths is the struct mapping the folders structure for a sveltin project.
type Paths struct {
	Build   string `mapstructure:"build"`
	Config  string `mapstructure:"config"`
	Content string `mapstructure:"content"`
	Static  string `mapstructure:"static"`
	Src     string `mapstructure:"src"`
	Routes  string `mapstructure:"routes"`
	Lib     string `mapstructure:"lib"`
	API     string `mapstructure:"api"`
	Themes  string `mapstructure:"themes"`
}
