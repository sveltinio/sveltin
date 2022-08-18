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
	nocolor = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}
	gray    = lipgloss.AdaptiveColor{Light: "#191919", Dark: "#c7c7c7"}
	sky     = lipgloss.AdaptiveColor{Light: "#0ea5e9", Dark: "#e0f2fe"}
	red     = lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"}
	orange  = lipgloss.AdaptiveColor{Light: "#f97316", Dark: "#fb923c"}
	//yellow  = lipgloss.AdaptiveColor{Light: "#eab308", Dark: "#facc15"}
	green   = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"}
	blue    = lipgloss.AdaptiveColor{Light: "#1d4ed8", Dark: "#3b82f6"}
	magenta = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"}
	muted   = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
)

var levelColorMap = map[Level]lipgloss.AdaptiveColor{
	DefaultLevel:   gray,
	DebugLevel:     sky,
	FatalLevel:     red,
	ErrorLevel:     red,
	WarningLevel:   orange,
	SuccessLevel:   green,
	InfoLevel:      blue,
	ImportantLevel: magenta,
}

func getColor(level Level) lipgloss.AdaptiveColor {
	if _, ok := levelColorMap[level]; ok {
		return levelColorMap[level]
	}
	return nocolor
}
