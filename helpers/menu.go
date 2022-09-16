package helpers

import "github.com/sveltinio/sveltin/internal/tpltypes"

// NewMenuItems return a NoPageItems.
func NewMenuItems(resources []string, content map[string][]string) *tpltypes.MenuItems {
	r := new(tpltypes.MenuItems)
	r.Resources = resources
	r.Content = content
	return r
}
