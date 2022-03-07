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

// LogLevel represents the level of severity.
type LogLevel int

// Log Levels.
const (
	LevelPanic LogLevel = iota
	LevelFatal
	LevelError
	LevelWarning
	LevelInfo
	LevelSuccess
	LevelDebug
	LevelImportant
	LevelDefault
)

// String returns the log level as string.
func (level LogLevel) String() string {
	switch level {
	case LevelDebug:
		return "LevelDebug"
	case LevelInfo:
		return "LevelInfo"
	case LevelSuccess:
		return "LevelSuccess"
	case LevelError:
		return "LevelError"
	case LevelWarning:
		return "LevelWarning"
	case LevelImportant:
		return "LevelImportant"
	case LevelFatal:
		return "LevelFatal"
	}
	return LevelDefault.String()
}

func makeLogLevelLabelMap() map[LogLevel]string {
	return map[LogLevel]string{
		LevelDefault:   "",
		LevelDebug:     "[DEBU]",
		LevelFatal:     "[FATA]",
		LevelError:     "[ERRO]",
		LevelWarning:   "[WARN]",
		LevelSuccess:   "[SUCC]",
		LevelInfo:      "[INFO]",
		LevelImportant: "[NOTI]",
	}
}

func getLabelByLogLevel(level LogLevel) (string, error) {
	labelMap := makeLogLevelLabelMap()
	if _, ok := labelMap[level]; ok {
		return labelMap[level], nil
	}
	return "", sveltinerr.NewDefaultError(fmt.Errorf("%s is not a valid LogLevel", level.String()))
}

func getLabel(level LogLevel) string {
	if label, err := getLabelByLogLevel(level); err == nil {
		return label
	}
	return "undefined"
}
