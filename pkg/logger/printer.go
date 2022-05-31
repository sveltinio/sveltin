/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package logger ...
package logger

// Printer is the interface defining the methods to be implemented by a printer.
type Printer interface {
	Print(string)
	SetPrinterOptions(options *PrinterOptions)
	Formatter
}

// PrinterOptions represents the valid options for a printer.
type PrinterOptions struct {
	Timestamp       bool
	TimestampFormat string
	Colors          bool
	Labels          bool
	Icons           bool
}
