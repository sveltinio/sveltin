/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

// names for the available thems options.
const (
	BlankTheme    string = "blank"
	SveltinTheme  string = "sveltin"
	ExistingTheme string = "existing"
)

// AvailableThemes is an array literal representing the available theme options.
var AvailableThemes = [3]string{BlankTheme, SveltinTheme, ExistingTheme}

// Theme represents the theme folder structure in a Sveltin project.
type Theme struct {
	File       string `mapstructure:"file"`
	Components string `mapStructure:"components"`
	Partials   string `mapstructure:"partials"`
}

// ThemeData contains info about the theme.
type ThemeData struct {
	ID     string
	IsNew  bool
	Name   string
	CSSLib string
	URL    string
}
