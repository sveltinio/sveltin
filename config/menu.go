/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

// MenuConfig is the struct representing a menu item.
type MenuConfig struct {
	Resources   []*ResourceItem
	Pages       []string
	WithContent bool
}

// ResourceItem is a struct representing a resource and its content as menu item.
type ResourceItem struct {
	name     string
	contents []string
}

// NewResourceItem returns a pointer to a ResourceItem struct.
func NewResourceItem(name string) *ResourceItem {
	return &ResourceItem{
		name:     name,
		contents: []string{},
	}
}

// GetName returns a string representing the resource item name.
func (r *ResourceItem) GetName() string {
	return r.name
}

// GetContents returns a slice of strings with contents for a resource.
func (r *ResourceItem) GetContents() []string {
	return r.contents
}

// AddChild appends an item to the slice of contents.
func (r *ResourceItem) AddChild(name string) {
	r.contents = append(r.contents, name)
}
