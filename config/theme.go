/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

// Theme represents the theme folder structure in a Sveltin project.
type Theme struct {
	File       string `mapstructure:"file"`
	Components string `mapStructure:"components"`
	Partials   string `mapstructure:"partials"`
}
