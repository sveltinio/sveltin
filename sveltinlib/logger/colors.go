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

type color int

const (
	noColor = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

func colorCode(color color) string {
	return fmt.Sprintf("\033[0;%dm", int(color))
}

func makeLogLevelColorMap() map[LogLevel]string {
	return map[LogLevel]string{
		LevelDefault:   colorCode(white),
		LevelDebug:     colorCode(cyan),
		LevelFatal:     colorCode(red),
		LevelError:     colorCode(red),
		LevelWarning:   colorCode(yellow),
		LevelSuccess:   colorCode(green),
		LevelInfo:      colorCode(blue),
		LevelImportant: colorCode(magenta),
	}
}

func getColorByLogLevel(level LogLevel) (string, error) {
	colorMap := makeLogLevelColorMap()
	if _, ok := colorMap[level]; ok {
		return colorMap[level], nil
	}
	return "", sveltinerr.NewDefaultError(fmt.Errorf("%s is not a valid LogLevel", level.String()))
}

func getColor(level LogLevel) string {
	if c, err := getColorByLogLevel(level); err == nil {
		return c
	}
	return colorCode(noColor)
}
