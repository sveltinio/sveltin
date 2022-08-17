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
	"runtime"
)

type icon string

const (
	iconDefault   icon = ""
	iconDebug     icon = "\u25B8" // ▸
	iconFatal     icon = "\u2718" // ✘
	iconError     icon = "\u2718" // ✘
	iconSuccess   icon = "\u2714" // ✔
	iconWarning   icon = "\u0021" // esclamation mark
	iconInfo      icon = "\u25B8" // ▸
	iconImportant icon = "\u2691" // ⚑
)

var levelIconMap = map[Level]icon{
	DefaultLevel:   iconDefault,
	DebugLevel:     iconDebug,
	FatalLevel:     iconFatal,
	ErrorLevel:     iconError,
	WarningLevel:   iconWarning,
	SuccessLevel:   iconSuccess,
	InfoLevel:      iconInfo,
	ImportantLevel: iconImportant,
}

func getIconByLevel(level Level) (icon, error) {
	if _, ok := levelIconMap[level]; ok {
		return levelIconMap[level], nil
	}
	return "", fmt.Errorf("%s is not a valid LogLevel", level.String())
}

func getIcon(level Level) string {
	if isWindows() {
		return ">"
	}
	if icon, err := getIconByLevel(level); err == nil {
		return string(icon)
	}
	return "undefined"
}

// isWindows returns true if the OS is MS Windows.
func isWindows() bool {
	return runtime.GOOS == "windows"
}
