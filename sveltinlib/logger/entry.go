/**
 * Copyright © 2022 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package logger ...
package logger

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// LogEntry represents a single log entry.
type LogEntry struct {
	Logger     *Logger
	Level      LogLevel
	Message    string
	Entries    []LogEntry
	prefix     string
	indentChar string
	indentSize int
}

// NewLogEntry returns a new entry.
func NewLogEntry(logger *Logger) *LogEntry {
	return &LogEntry{
		Logger: logger,
	}
}

// WithList returns a new entry with `entries` set.
func (item *LogEntry) WithList() *LogEntry {
	return &LogEntry{
		Logger:     item.Logger,
		Entries:    []LogEntry{},
		indentChar: "\u0020", //space
		indentSize: 2,
	}
}

func (item *LogEntry) isWithList() bool {
	return len(item.Entries) > 0
}

// String returns the entry as string.
func (item *LogEntry) String() string {
	var buffer bytes.Buffer
	if item.isWithList() {
		fmt.Fprintf(&buffer, "%s", item.Logger.Printer.FormatList(item))
	} else {
		fmt.Fprintf(&buffer, "%s", item.Logger.Printer.Format(item, true))
	}

	return buffer.String()
}

// SetIndentChar sets the char used as prefix when indent list entries. Whitespace as default.
func (item *LogEntry) SetIndentChar(c string) {
	item.indentChar = c
}

// SetIndentSize sets how many time the indent char has to be repeated when indent list entries.
func (item *LogEntry) SetIndentSize(size int) {
	item.indentSize = size
}

// Append add an entry to the `entries`.
func (item *LogEntry) Append(level LogLevel, msg string) {
	elem := LogEntry{
		Logger:  item.Logger,
		Level:   level,
		Message: msg,
		prefix:  item.prefix,
	}
	item.Entries = append(item.Entries, elem)
}

// Indent add a string as prefix when indent list entries.
func (item *LogEntry) Indent() {
	item.prefix = strings.Repeat(item.indentChar, item.indentSize)
}

// Unindent resets the indent size.
func (item *LogEntry) Unindent() {
	item.indentSize = 0
	item.Indent()
}

// Plain level message.
func (item *LogEntry) Plain(msg string) {
	item.Level = LevelDefault
	item.Message = msg
	item.Logger.log(item)
}

// Debug level message.
func (item *LogEntry) Debug(msg string) {
	item.Level = LevelDebug
	item.Message = msg
	item.Logger.log(item)
}

// Info level message.
func (item *LogEntry) Info(msg string) {
	item.Level = LevelInfo
	item.Message = msg
	item.Logger.log(item)
}

// Error level message.
func (item *LogEntry) Error(msg string) {
	item.Level = LevelError
	item.Message = msg
	item.Logger.log(item)
}

// Warning level message.
func (item *LogEntry) Warning(msg string) {
	item.Level = LevelWarning
	item.Message = msg
	item.Logger.log(item)
}

// Success level message.
func (item *LogEntry) Success(msg string) {
	item.Level = LevelSuccess
	item.Message = msg
	item.Logger.log(item)
}

// Important level message.
func (item *LogEntry) Important(msg string) {
	item.Level = LevelImportant
	item.Message = msg
	item.Logger.log(item)
}

// Fatal level message.
func (item *LogEntry) Fatal(msg string) {
	item.Level = LevelFatal
	item.Message = msg
	item.Logger.log(item)
	os.Exit(1)
}

// Debugf level formatted message.
func (item *LogEntry) Debugf(msg string, v ...interface{}) {
	item.Debug(fmt.Sprintf(msg, v...))
}

// Infof level formatted message.
func (item *LogEntry) Infof(msg string, v ...interface{}) {

	item.Info(fmt.Sprintf(msg, v...))
}

// Errorf level formatted message.
func (item *LogEntry) Errorf(msg string, v ...interface{}) {
	item.Error(fmt.Sprintf(msg, v...))
}

// Successf level formatted message.
func (item *LogEntry) Successf(msg string, v ...interface{}) {
	item.Success(fmt.Sprintf(msg, v...))
}

// Warningf level formatted message.
func (item *LogEntry) Warningf(msg string, v ...interface{}) {
	item.Warning(fmt.Sprintf(msg, v...))
}

// Importantf level formatted message.
func (item *LogEntry) Importantf(msg string, v ...interface{}) {
	item.Important(fmt.Sprintf(msg, v...))
}

// Fatalf level formatted message.
func (item *LogEntry) Fatalf(msg string, v ...interface{}) {
	item.Fatal(fmt.Sprintf(msg, v...))
}
