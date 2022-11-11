/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package tpltypes

// names for the available thems options.
const (
	BlankTheme    string = "blank"
	SveltinTheme  string = "sveltin"
	ExistingTheme string = "existing"
)

// AvailableThemes is the list of the available themes.
var AvailableThemes = []string{BlankTheme, SveltinTheme, ExistingTheme}

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
	URL    string
	Style  string `mapstructure:"style" json:"style" validate:"required,oneof='blank' 'sveltin'"`
	Name   string `mapstructure:"name" json:"name" validate:"required"`
	CSSLib string `mapstructure:"cssLib" json:"cssLib" validate:"required,oneof='bootstrap' 'bulma' 'scss' 'tailwindcss' 'vanillacss'"`
}
