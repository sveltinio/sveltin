/**
 * Copyright © 2022 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package logger ...
package logger

import (
	"fmt"

	"github.com/sveltinio/sveltin/sveltinlib/sveltinerr"
)

type icon string

const (
	iconDefault   icon = ""
	iconDebug     icon = "\u2734" // ✴
	iconFatal     icon = "\u2718"
	iconError     icon = "\u2718" // ✘
	iconSuccess   icon = "\u2714" // ✔
	iconWarning   icon = "\u26A0" // esclamation mark
	iconInfo      icon = "\u25B6" // ▶
	iconImportant icon = "\u2691" // ⚑
)

func makeLogLevelIconMap() map[LogLevel]icon {
	return map[LogLevel]icon{
		LevelDefault:   iconDefault,
		LevelDebug:     iconDebug,
		LevelFatal:     iconError,
		LevelError:     iconError,
		LevelWarning:   iconWarning,
		LevelSuccess:   iconSuccess,
		LevelInfo:      iconInfo,
		LevelImportant: iconImportant,
	}
}

func getIconByLogLevel(level LogLevel) (icon, error) {
	iconMap := makeLogLevelIconMap()
	if _, ok := iconMap[level]; ok {
		return iconMap[level], nil
	}
	return "", sveltinerr.NewDefaultError(fmt.Errorf("%s is not a valid LogLevel", level.String()))
}

func getIcon(level LogLevel) string {
	if icon, err := getIconByLogLevel(level); err == nil {
		return string(icon)
	}
	return "undefined"
}
