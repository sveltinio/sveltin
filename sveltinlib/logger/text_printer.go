/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package logger ...
package logger

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/mattn/go-colorable"
)

// TextPrinter sets stdout as printer.
type TextPrinter struct {
	Options *PrinterOptions
}

// SetPrinterOptions sets the printer options for TextPrinter.
func (tp *TextPrinter) SetPrinterOptions(options *PrinterOptions) {
	tp.Options = options
}

// Print send the message string to the stdout.
func (tp *TextPrinter) Print(msg string) {
	stdOut := bufio.NewWriter(colorable.NewColorableStdout())
	stdOut.WriteString(msg)
	stdOut.WriteString("\n")
	stdOut.Flush()
}

// Format defines how a single log entry will be formatted by the TextLogger printer.
func (tp *TextPrinter) Format(item *LogEntry, isWithList bool) string {
	if isWithList {
		return fmt.Sprintf("%s%s %s", item.prefix, formatMarkup(item.Level, tp.Options), formatMessage(white, item.Message))
	}
	return formatMessage(white, item.Message)
}

// FormatList defines how a log entry with List will be formatted by the TextLogger printer.
func (tp *TextPrinter) FormatList(item *LogEntry) string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "\n%s\n", item.Logger.Printer.Format(item, false))
	buffer.WriteString(strings.Repeat("-", len(item.Message)+1))
	buffer.WriteString("\n")
	for _, line := range item.Entries {
		fmt.Fprintf(&buffer, "%s\n", item.Logger.Printer.Format(&line, true))
	}
	return buffer.String()
}

//=============================================================================

func formatMessage(color color, msg string) string {
	return fmt.Sprintf("%s%s", colorCode(color), msg)
}

func formatMarkup(level LogLevel, options *PrinterOptions) string {
	timestampStr := ""
	if options.Timestamp {
		timestampStr = fmt.Sprintf("%s ", time.Now().Format(options.TimestampFormat))
	}
	colorStr := ""
	if options.Colors {
		colorStr = getColor(level)
	}
	iconStr := ""
	if options.Icons {
		iconStr = getIcon(level)
	}
	labelStr := ""
	if options.Labels {
		labelStr = getLabel(level)
	}
	return fmt.Sprintf("%s%s%s %s", timestampStr, colorStr, iconStr, labelStr)
}
