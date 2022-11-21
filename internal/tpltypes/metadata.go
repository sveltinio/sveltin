/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package tpltypes

// MetadataData is the struct representing the user selection for the metadata.
type MetadataData struct {
	Name     string
	Resource string
	Type     string
}

// NewMetadataData creates a pointer to a NewMetadataData struct.
func NewMetadataData(name, mdResource, mdType string) *MetadataData {
	return &MetadataData{
		Name:     name,
		Resource: mdResource,
		Type:     mdType,
	}
}
