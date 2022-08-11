/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package choose ...
package choose

import "github.com/charmbracelet/bubbles/list"

// GetItemsKeys return a slice of string representing the item's keys.
func GetItemsKeys(items []list.Item) []string {
	res := []string{}
	for _, v := range items {
		res = append(res, v.FilterValue())
	}
	return res
}

func (i *Item) String() string {
	return i.Name
}

// ToListItem converts a alice of strings in a slice of list.Item.
func ToListItem(items []string) []list.Item {
	res := []list.Item{}

	for _, v := range items {
		i := Item{
			Name: v,
			Desc: v,
		}
		res = append(res, i)
	}

	return res
}
