/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package markup defines styles and colors used to print messages on the shell as if they were HTML tags.
package markup

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// H1 is used to style an heading 1, as the HTML <h1> tag.
	H1 = lipgloss.NewStyle().
		Margin(1, 0).
		BorderBottom(true).
		BorderForeground(purple).
		BorderStyle(lipgloss.ThickBorder()).Render

	// H2 is used to style an heading 2, as the HTML <h2> tag.
	H2 = lipgloss.NewStyle().
		MarginTop(1).
		Italic(true).
		BorderBottom(true).
		BorderStyle(lipgloss.NormalBorder()).
		Render

	// P is used to style a paragraph, as the HTML <p> tag.
	P = lipgloss.NewStyle().Render

	// A is used to style a link, as the HTML <a> tag.
	A = lipgloss.NewStyle().Foreground(green).Render

	// BR is used to prints a line breaks, as the HTML <br> tag.
	BR = lipgloss.NewStyle().Render("")

	// HR is used to style a horizontal rule, as the HTML <hr> tag.
	HR = func(width int) string {
		return lipgloss.NewStyle().
			MarginTop(1).
			BorderBottom(true).
			BorderStyle(lipgloss.NormalBorder()).
			Render(strings.Repeat(" ", width))
	}

	// Section is used to style a content section, as the HTML <section> tag.
	Section = func(title string, content []string) string {
		return lipgloss.NewStyle().Margin(1, 0).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				H2(title),
				P(strings.Join(content, "\n")),
			))
	}

	codeBlockContentStyle = lipgloss.NewStyle().Italic(true).Faint(true).Render
	// Code is used to style a code line.
	Code = func(content string) string {
		return lipgloss.NewStyle().Margin(0, 0, 0, 0).Render(
			codeBlockContentStyle(content),
		)
	}

	// CodeBlock is used to style a block of code <block><code>...</code></block>
	CodeBlock = func(strs ...string) string {
		return lipgloss.NewStyle().Margin(0, 0, 1, 0).Render(
			lipgloss.JoinVertical(lipgloss.Left, codeBlockContentStyle(strings.Join(strs, "\n"))),
		)
	}
	// UL is used to style an unordered list, as the HTML <ul> tag.
	UL = lipgloss.NewStyle().Margin(1, 0, 1, 1)

	// OL is used to style an ordered list, as the HTML <ol> tag.
	OL = UL.Copy()

	// LI is used to style a list item, as the HTML <li> tag..
	LI = lipgloss.NewStyle().Render
)
