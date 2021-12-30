/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package config

type MenuConfig struct {
	Resources   []*ResourceItem
	Pages       []string
	WithContent bool
}

type ResourceItem struct {
	name     string
	contents []string
}

func NewResourceItem(name string) *ResourceItem {
	return &ResourceItem{
		name:     name,
		contents: []string{},
	}
}

func (r *ResourceItem) GetName() string {
	return r.name
}

func (r *ResourceItem) GetContents() []string {
	return r.contents
}

func (r *ResourceItem) AddChild(name string) {
	r.contents = append(r.contents, name)
}
