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

//=============================================================================

type Logger interface {
	SetTitle(text string)
	GetTitle() string
	SetContent(text string)
	GetContent() string
	Reset()
}

type TextLogger struct {
	title   string
	content string
}

func NewTextLogger() *TextLogger {
	return &TextLogger{}
}

func (t *TextLogger) SetTitle(title string) {
	t.title = title
}

func (t *TextLogger) GetTitle() string {
	return t.title
}

func (t *TextLogger) SetContent(content string) {
	t.content = content
}

func (t *TextLogger) GetContent() string {
	return t.content
}

func (tl *TextLogger) Reset() {
	tl.title = ""
	tl.content = ""
}

//=============================================================================

type ListWriter struct {
	content list.Writer
}

func NewListWriter() ListWriter {
	return ListWriter{
		content: list.NewWriter(),
	}
}

func (lw ListWriter) Reset() {
	lw.content.Reset()
}

func (lw ListWriter) Render() string {
	return lw.content.Render()
}

func (lw ListWriter) AppendItem(message string) {
	lw.content.AppendItem(message)
}

func (lw ListWriter) Indent() {
	lw.content.Indent()
}

func (lw ListWriter) UnIndent() {
	lw.content.UnIndent()
}

func (lw ListWriter) SetStyle(style list.Style) {
	lw.content.SetStyle(style)
}

func (lw ListWriter) SetBulletTriangleStyle() {
	lw.content.SetStyle(list.StyleBulletTriangle)
}

//=============================================================================

type Printer interface {
	Print()
}

type PrinterBuffer struct {
	buf bytes.Buffer
}

func (pb PrinterBuffer) Print() {
	jww.FEEDBACK.Println(pb.buf.String())
}

func (pb PrinterBuffer) ToString() string {
	return pb.buf.String()
}

//=============================================================================

func PrettyPrinter(tl *TextLogger) PrinterBuffer {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "\n%s:\n", tl.GetTitle())

	buffer.WriteString(strings.Repeat("-", len(tl.GetTitle())+1))
	buffer.WriteString("\n")
	if tl.GetContent() != "" {
		for _, line := range strings.Split(tl.GetContent(), "\n") {
			fmt.Fprintf(&buffer, "%s\n", line)
		}
	}
	return PrinterBuffer{
		buf: buffer,
	}
}
