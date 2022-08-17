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
	"os"
	"strings"
)

// LogEntry represents a single log entry.
type LogEntry struct {
	Logger     *Logger
	Level      Level
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

// NewListLogger returns a new entry with `entries` set.
func NewListLogger() *LogEntry {
	l := &Logger{
		Printer: iconAndColorOnlyTextPrinter(),
		Level:   DebugLevel,
	}
	return NewLogEntry(l).withList()
}

// Title set the test to be used as title for ListLogger.
func (item *LogEntry) Title(title string) {
	item.Message = title
}

// WithList returns a new entry with `entries` set.
func (item *LogEntry) withList() *LogEntry {
	return &LogEntry{
		Logger:     item.Logger,
		Entries:    []LogEntry{},
		indentChar: "\u0020", //space
		indentSize: 2,
	}
}

// IsListLogger returns true when ListLogger.
func (item *LogEntry) IsListLogger() bool {
	return len(item.Entries) > 0
}

// String returns the entry as string.
func (item *LogEntry) String() string {
	return item.Logger.Printer.Format(item)
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
func (item *LogEntry) Append(level Level, msg string) {
	elem := LogEntry{
		Logger:  item.Logger,
		Level:   level,
		Message: msg,
		prefix:  item.prefix,
	}
	item.Entries = append(item.Entries, elem)
}

// Render prints log entry when ListLogger.
func (item *LogEntry) Render() {
	if item.IsListLogger() {
		item.Logger.log(item)
	}
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
	item.Level = DefaultLevel
	item.Message = msg
	item.Logger.log(item)
}

// Debug level message.
func (item *LogEntry) Debug(msg string) {
	item.Level = DebugLevel
	item.Message = msg
	item.Logger.log(item)
}

// Info level message.
func (item *LogEntry) Info(msg string) {
	item.Level = InfoLevel
	item.Message = msg
	item.Logger.log(item)
}

// Error level message.
func (item *LogEntry) Error(msg string) {
	item.Level = ErrorLevel
	item.Message = msg
	item.Logger.log(item)
}

// Warning level message.
func (item *LogEntry) Warning(msg string) {
	item.Level = WarningLevel
	item.Message = msg
	item.Logger.log(item)
}

// Success level message.
func (item *LogEntry) Success(msg string) {
	item.Level = SuccessLevel
	item.Message = msg
	item.Logger.log(item)
}

// Important level message.
func (item *LogEntry) Important(msg string) {
	item.Level = ImportantLevel
	item.Message = msg
	item.Logger.log(item)
}

// Fatal level message.
func (item *LogEntry) Fatal(msg string) {
	item.Level = FatalLevel
	item.Message = msg
	item.Logger.log(item)
	os.Exit(1)
}

// Plainf level formatted message.
func (item *LogEntry) Plainf(msg string, v ...interface{}) {
	item.Plain(fmt.Sprintf(msg, v...))
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
