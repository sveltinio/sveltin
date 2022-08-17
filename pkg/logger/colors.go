/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package logger ...
package logger

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	white  = lipgloss.AdaptiveColor{Light: "#FFFFFF", Dark: "#FFFFFF"}
	cyan   = lipgloss.AdaptiveColor{Light: "#4f46e5", Dark: "#c7d2fe"}
	red    = lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"}
	orange = lipgloss.AdaptiveColor{Light: "#facc15", Dark: "#fb923c"}
	//yellow  = lipgloss.AdaptiveColor{Light: "#eab308", Dark: "#facc15"}
	green   = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"}
	blue    = lipgloss.AdaptiveColor{Light: "#1d4ed8", Dark: "#3b82f6"}
	magenta = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"}
	muted   = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
)

var levelColorMap = map[Level]lipgloss.AdaptiveColor{
	DefaultLevel:   white,
	DebugLevel:     cyan,
	FatalLevel:     red,
	ErrorLevel:     red,
	WarningLevel:   orange,
	SuccessLevel:   green,
	InfoLevel:      blue,
	ImportantLevel: magenta,
}

func getColorByLevel(level Level) (lipgloss.AdaptiveColor, error) {
	if _, ok := levelColorMap[level]; ok {
		return levelColorMap[level], nil
	}
	return lipgloss.AdaptiveColor{}, fmt.Errorf("%s is not a valid Level", level.String())
}

func getColor(level Level) lipgloss.AdaptiveColor {
	if c, err := getColorByLevel(level); err == nil {
		return c
	}
	return white
}
