/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package utils ...
package utils

import (
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

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
