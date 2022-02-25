/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package builder ...
package builder

type director struct {
	builder iFileContentBuilder
}

// NewDirector create a director.
func NewDirector(b iFileContentBuilder) *director {
	return &director{
		builder: b,
	}
}

// SetBuilder set the Builder to be used.
func (d *director) SetBuilder(b iFileContentBuilder) {
	d.builder = b
}

// GetContent returns the Content struct used by the Builder.
func (d *director) GetContent() Content {
	d.builder.setContentType()
	if err := d.builder.setPathToTplFile(); err != nil {
		panic("something went wrong calling setPathToTplFile")
	}
	d.builder.setFuncs()
	return d.builder.GetContent()
}
