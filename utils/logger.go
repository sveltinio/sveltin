/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package utils ...
package utils

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/list"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

//=============================================================================

// Logger interface declares the methods to set the loggers props.
type Logger interface {
	SetTitle(text string)
	GetTitle() string
	SetContent(text string)
	GetContent() string
	Reset()
}

// TextLogger is the struct used to write on the stdout as plain text.
type TextLogger struct {
	title   string
	content string
}

// NewTextLogger returns a pointer to a TextLogger.
func NewTextLogger() *TextLogger {
	return &TextLogger{}
}

// SetTitle sets the title for the TextLogger.
func (tl *TextLogger) SetTitle(title string) {
	tl.title = title
}

// GetTitle returns the TextLogger title as string.
func (tl *TextLogger) GetTitle() string {
	return tl.title
}

// SetContent sets the content for the TextLogger.
func (tl *TextLogger) SetContent(content string) {
	tl.content = content
}

// GetContent returns the TextLogger content as string.
func (tl *TextLogger) GetContent() string {
	return tl.content
}

// Reset the TextLogger.
func (tl *TextLogger) Reset() {
	tl.title = ""
	tl.content = ""
}

//=============================================================================

// ListWriter is the struct wrapping the list.Writer.
type ListWriter struct {
	content list.Writer
}

// NewListWriter creates a new ListWriter.
func NewListWriter() ListWriter {
	return ListWriter{
		content: list.NewWriter(),
	}
}

// Reset the ListWriter.
func (lw ListWriter) Reset() {
	lw.content.Reset()
}

// Render printout the ListWriter to the stadout.
func (lw ListWriter) Render() string {
	return lw.content.Render()
}

// AppendItem addend an item to the ListWriter.
func (lw ListWriter) AppendItem(message string) {
	lw.content.AppendItem(message)
}

// Indent applies indent the output.
func (lw ListWriter) Indent() {
	lw.content.Indent()
}

// UnIndent applies unindent to the outout.
func (lw ListWriter) UnIndent() {
	lw.content.UnIndent()
}

// SetStyle sets the list style to be applied on the ListWriter.
func (lw ListWriter) SetStyle(style list.Style) {
	lw.content.SetStyle(style)
}

// SetBulletTriangleStyle sets BulletTriangle as style to the ListWriter.
func (lw ListWriter) SetBulletTriangleStyle() {
	lw.content.SetStyle(list.StyleBulletTriangle)
}

//=============================================================================

// Printer interface declares just the single method for printing the buffer content.
type Printer interface {
	Print()
}

// PrinterBuffer is the struct representing the buffer.
type PrinterBuffer struct {
	buf bytes.Buffer
}

// Print prints out the buffer content to the stdout.
func (pb PrinterBuffer) Print() {
	jww.FEEDBACK.Println(pb.buf.String())
}

// ToString returns the buffer content as string.
func (pb PrinterBuffer) ToString() string {
	return pb.buf.String()
}

//=============================================================================

// PrettyPrinter returns a PrinterBuffer as formatted content ready to be printout to stdout.
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

//=============================================================================

// ProgressBar is the struct representing a progressbar instance.
type ProgressBar struct {
	progress *mpb.Progress
	bar      *mpb.Bar
}

// NewProgressBar returns a pointer to a ProgressBar struct.
func NewProgressBar(total int) *ProgressBar {
	p := mpb.New(mpb.WithWidth(64))
	// create a single bar, which will inherit container's width
	bar := p.AddBar(int64(total),
		mpb.PrependDecorators(
			decor.CountersNoUnit("%d / %d"),
			decor.OnComplete(
				// spinner decorator with default style
				decor.Spinner(nil, decor.WCSyncSpace), "done",
			),
		),
		mpb.AppendDecorators(
			// decor.DSyncWidth bit enables column width synchronization
			decor.Percentage(decor.WCSyncWidth),
		),
	)

	return &ProgressBar{
		progress: p,
		bar:      bar,
	}
}

// Increment increments progress by amount of n.
func (pb *ProgressBar) Increment() {
	pb.bar.Increment()
}

// Wait blocks until bar is completed or aborted.
func (pb *ProgressBar) Wait() {
	pb.progress.Wait()
}
