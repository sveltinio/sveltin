/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package utils

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/list"
	jww "github.com/spf13/jwalterweatherman"
)

type LoggerWriter struct {
	content list.Writer
}

func NewLoggerWriter() LoggerWriter {
	return LoggerWriter{
		content: list.NewWriter(),
	}
}

func (w LoggerWriter) Reset() {
	w.content.Reset()
}

func (w LoggerWriter) Render() string {
	return w.content.Render()
}

func (w LoggerWriter) AppendItem(message string) {
	w.content.AppendItem(message)
}

func (w LoggerWriter) Indent() {
	w.content.Indent()
}

func (w LoggerWriter) UnIndent() {
	w.content.UnIndent()
}

type PrinterBuffer struct {
	buf bytes.Buffer
}

type PrinterContent struct {
	Title   string
	content string
}

func (p *PrinterContent) SetContent(content string) {
	p.content = content
}

func (pb PrinterBuffer) Print() {
	jww.FEEDBACK.Println(pb.buf.String())
}

func (pb PrinterBuffer) ToString() string {
	return pb.buf.String()
}

func PrettyPrinter(pc *PrinterContent) PrinterBuffer {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "\n%s:\n", pc.Title)

	buffer.WriteString(strings.Repeat("-", len(pc.Title)+1))
	buffer.WriteString("\n")
	if pc.content != "" {
		for _, line := range strings.Split(pc.content, "\n") {
			fmt.Fprintf(&buffer, "%s\n", line)
		}
	}
	return PrinterBuffer{
		buf: buffer,
	}
}
