/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package tpltypes

const (
	// Blank represents the fontmatter-only template id used when generating the content file.
	Blank string = "blank"
	// Sample represents the sample-content template id used when generating the content file.
	Sample string = "sample"
)

// ContentData is the struct representing the user selection for new content.
type ContentData struct {
	Name     string
	Resource string
	Type     string
}

// NewContentData creates a pointer to a ContentData struct.
func NewContentData(name, cResource string, isSample bool) *ContentData {
	cType := Blank
	if isSample {
		cType = Sample
	}
	return &ContentData{
		Name:     name,
		Resource: cResource,
		Type:     cType,
	}
}
