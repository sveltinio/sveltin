/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package markup

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// ListTitle is used to style a list header text.
	listTitle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(slate).
			Foreground(amber).
			MarginRight(2).
			Render

	// CheckMark is used to print a checkmark char.
	CheckMark = lipgloss.NewStyle().SetString("✓").
			Foreground(green).
			PaddingRight(1).
			String()

	// Divider prints a dot char.
	Divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Faint(true).
		String()

	// Inline is used to style strings horizontally joined along their center.
	Inline = func(strs ...string) string {
		return lipgloss.NewStyle().Render(lipgloss.JoinHorizontal(lipgloss.Center, strings.Join(strs, " ")))
	}
	// LIWithIcon is used to style a list item with a check mark as prefix.
	LIWithIcon = func(s, v, icon string) string {
		return icon + lipgloss.NewStyle().
			Foreground(gray).
			Render(s) + Bold(v)
	}

	// NewUL is an utility function to help creating a new unordered list.
	NewUL = func(entries []string) string {
		return UL.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				strings.Join(entries, "\n")))
	}

	// NewULWithIconPrefix is an utility function to help creating a new unordered list with an icon as prefix to the list items.
	NewULWithIconPrefix = func(title string, entries map[string]string, icon string) string {
		items := []string{}
		for k, v := range entries {
			items = append(items, LIWithIcon(k, v, icon))
		}

		return UL.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listTitle(title), strings.Join(items, "\n")))
	}

	// NewOL is an utility function to help creating a new ordered list.
	NewOL = func(entries []string) string {
		items := []string{}
		for k, v := range entries {
			items = append(items, LI(fmt.Sprintf("%d. %s", k+1, v)))
		}

		return OL.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				strings.Join(items, "\n")))
	}

	// NewOLWithTitle is an utility function to help creating a new ordered list with a title text.
	NewOLWithTitle = func(title string, entries []string) string {
		return OL.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				listTitle(title), NewOL(entries)))
	}
)
