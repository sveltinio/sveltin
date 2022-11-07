/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package markup

import "github.com/charmbracelet/lipgloss"

var (
	nocolor = lipgloss.AdaptiveColor{Light: "#0f172a", Dark: "#d1d5db"} // Light: gray-900, Dark: gray-300
	slate   = lipgloss.AdaptiveColor{Light: "#cbd5e1", Dark: "#f8fafc"} // Light: slate-300, Dark: slate-50
	green   = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"} // Light: green-800, Dark: green-500
	gray    = lipgloss.AdaptiveColor{Light: "#6b7280", Dark: "#9ca3af"} // Light: gray-500, Dark: gray-400
	amber   = lipgloss.AdaptiveColor{Light: "#fcd34d", Dark: "#fffbeb"} // Light: amber-300, Dark: amber-50
	purple  = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500
)
