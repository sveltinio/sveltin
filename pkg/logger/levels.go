/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package logger ...
package logger

import (
	"strconv"
)

// Level represents the level of severity.
type Level int8

// Log Levels.
const (
	NoLevel Level = iota
	DebugLevel
	FatalLevel
	ErrorLevel
	WarningLevel
	InfoLevel
	SuccessLevel
	ImportantLevel
	DefaultLevel
)

var levelLabelMap = map[Level]string{
	DebugLevel:     "DEBU",
	FatalLevel:     "FATA",
	ErrorLevel:     "ERRO",
	InfoLevel:      "INFO",
	WarningLevel:   "WARN",
	SuccessLevel:   "SUCC",
	ImportantLevel: "NOTI",
	DefaultLevel:   "",
}

// String returns the log level as string.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "LevelDebug"
	case FatalLevel:
		return "LevelFatal"
	case ErrorLevel:
		return "LevelError"
	case InfoLevel:
		return "LevelInfo"
	case WarningLevel:
		return "LevelWarning"
	case SuccessLevel:
		return "LevelSuccess"
	case ImportantLevel:
		return "LevelImportant"
	case DefaultLevel:
		return ""
	default:
		return strconv.Itoa(int(l))
	}

}

func getLevelLabel(level Level) string {
	if _, ok := levelLabelMap[level]; ok {
		return levelLabelMap[level]
	}
	return "undefined log level"
}
