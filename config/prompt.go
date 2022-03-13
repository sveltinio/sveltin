/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

// PromptContent represents an ask to the user and containes the relative error message.
type PromptContent struct {
	ErrorMsg string
	Label    string
}

// PromptObject represents an item to display inside the list by promptui.
type PromptObject struct {
	Id   string
	Name string
}

func (po *PromptObject) String() string {
	return po.Id
}
