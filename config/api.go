/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package config ...
package config

// API represents the API folder structure in a Sveltin project.
type API struct {
	Version  string      `mapstructure:"version"`
	Resource APIResource `mapstructure:"resource"`
	Metadata APIMetadata `mapstructure:"metadata"`
}

// APIResource represents the folder name and its file for a resource API in a sveltin project.
type APIResource struct {
	Filename string `mapstructure:"filename"`
	Public   string `mapstructure:"public"`
}

// APIMetadata represents the folder name and its file for a metadata API in a sveltin project.
type APIMetadata struct {
	Filename string `mapstructure:"filename"`
	Public   string `mapstructure:"public"`
}
