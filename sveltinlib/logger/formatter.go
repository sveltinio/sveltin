/**
 * Copyright © 2022 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package logger ...
package logger

// Formatter is the interface defining the methods to be implemented by a printer.
type Formatter interface {
	Format(*LogEntry, bool) string
	FormatList(*LogEntry) string
}
