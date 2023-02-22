/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package markup

import "github.com/charmbracelet/lipgloss"

var (
	// Plain sets an as it is formatting rule for the text.
	Plain = lipgloss.NewStyle().Foreground(nocolor).Render
	// Italic sets an italic formatting rule for the text.
	Italic = lipgloss.NewStyle().Italic(true).Render
	// Bold sets a bold formatting rule for the text.
	Bold = lipgloss.NewStyle().Bold(true).Render
	//Underline sets an underlined rule for the text.
	Underline = lipgloss.NewStyle().Underline(true).Render
	// Gray sets foreground color to gray for the text.
	Gray = lipgloss.NewStyle().Foreground(gray).Render
	// Faint sets a rule for rendering the foreground color in a dimmer shade.
	Faint = lipgloss.NewStyle().Faint(true).Render

	// Green renders text in green
	Green = lipgloss.NewStyle().Foreground(green).Render
	// Amber renders text in amber
	Amber = lipgloss.NewStyle().Foreground(amber).Render
	// Yellow renders text in yellow
	Yellow = lipgloss.NewStyle().Foreground(yellow).Render
	// Purple renders text in purple
	Purple = lipgloss.NewStyle().Foreground(purple).Render
	// Blue renders text in blue
	Blue = lipgloss.NewStyle().Foreground(blue).Render

	// Bordered prints a bordered text.
	Bordered = lipgloss.NewStyle().
			Margin(1).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder(), true, true, true, true).
			BorderForeground(lipgloss.Color("#874BFD")).
			Render
	// Centered prints a center aligned text
	Centered = lipgloss.NewStyle().Align(lipgloss.Center).Render
)
