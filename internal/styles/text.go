/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package styles ...
package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Plain sets an as it is formatting rule for the text.
	Plain = lipgloss.NewStyle().Foreground(nocolor).Render
	// Italic sets an italic formatting rule for the text.
	Italic = lipgloss.NewStyle().Italic(true).Render
	// Bold sets a bold formatting rule for the text.
	Bold = lipgloss.NewStyle().Bold(true).Render
	// Gray sets foreground color to gray for the text.
	Gray = lipgloss.NewStyle().Foreground(gray).Render
	// Faint sets a rule for rendering the foreground color in a dimmer shade.
	Faint = lipgloss.NewStyle().Faint(true).Render
	// Bordered prints a bordered text.
	Bordered = lipgloss.NewStyle().
			Margin(1).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder(), true, true, true, true).
			BorderForeground(lipgloss.Color("#874BFD")).
			Render
)
