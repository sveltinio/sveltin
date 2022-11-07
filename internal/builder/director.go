/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package builder implements the Builder design pattern used to generate all sveltin's artifacts (project, page, metadata, etc.)
package builder

// Director defines the order in which to execute the building steps.
type Director struct {
	builder IFileContentBuilder
}

// NewDirector create a director.
func NewDirector(b IFileContentBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// SetBuilder set the Builder to be used.
func (d *Director) SetBuilder(b IFileContentBuilder) {
	d.builder = b
}

// GetContent returns the Content struct used by the Builder.
func (d *Director) GetContent() Content {
	d.builder.setContentType()
	if err := d.builder.setPathToTplFile(); err != nil {
		panic("something went wrong calling setPathToTplFile")
	}
	d.builder.setFuncs()
	return d.builder.GetContent()
}
