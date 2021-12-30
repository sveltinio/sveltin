/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package config

type API struct {
	Version  string      `mapstructure:"version"`
	Resource APIResource `mapstructure:"resource"`
	Metadata APIMetadata `mapstructure:"metadata"`
}

type APIResource struct {
	Filename string `mapstructure:"filename"`
	Public   string `mapstructure:"public"`
}

type APIMetadata struct {
	Filename string `mapstructure:"filename"`
	Public   string `mapstructure:"public"`
}
