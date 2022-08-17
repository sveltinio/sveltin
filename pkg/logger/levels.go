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
	"strconv"
)

// Level represents the level of severity.
type Level int8

// Log Levels.
const (
	DebugLevel Level = iota
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
	}
	return strconv.Itoa(int(l))
}

func getLabelByLevel(level Level) (string, error) {
	if _, ok := levelLabelMap[level]; ok {
		return levelLabelMap[level], nil
	}
	return "", fmt.Errorf("%s is not a valid Level", level.String())
}

func getLevelLabel(level Level) string {
	if label, err := getLabelByLevel(level); err == nil {
		return label
	}
	return "undefined"
}
