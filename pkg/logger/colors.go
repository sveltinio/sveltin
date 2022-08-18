/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package logger ...
package logger

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	nocolor = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"} // Light: black, Dark: white
	gray    = lipgloss.AdaptiveColor{Light: "#0f172a", Dark: "#d1d5db"} // Light: gray-900, Dark: gray-300
	sky     = lipgloss.AdaptiveColor{Light: "#0ea5e9", Dark: "#e0f2fe"} // Light: sky-500, Dark: sky-100
	red     = lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"} // Light: red-500, Dark: red-500
	orange  = lipgloss.AdaptiveColor{Light: "#f97316", Dark: "#fb923c"} // Light: orange-500, Dark: orange-400
	//yellow  = lipgloss.AdaptiveColor{Light: "#eab308", Dark: "#facc15"} // Light: yellow-500, Dark: yellow-400
	green  = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"} // Light: green-800, Dark: green-500
	blue   = lipgloss.AdaptiveColor{Light: "#1d4ed8", Dark: "#3b82f6"} // Light: blue-700, Dark: blue-500
	purple = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500
	muted  = lipgloss.AdaptiveColor{Light: "#6b7280", Dark: "#9ca3af"} // Light: gray-500, Dark: gray-400
)

var levelColorMap = map[Level]lipgloss.AdaptiveColor{
	DefaultLevel:   gray,
	DebugLevel:     sky,
	FatalLevel:     red,
	ErrorLevel:     red,
	WarningLevel:   orange,
	SuccessLevel:   green,
	InfoLevel:      blue,
	ImportantLevel: purple,
}

func getColor(level Level) lipgloss.AdaptiveColor {
	if _, ok := levelColorMap[level]; ok {
		return levelColorMap[level]
	}
	return nocolor
}
