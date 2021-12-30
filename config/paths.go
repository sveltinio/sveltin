/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package config

type Paths struct {
	Build   string `mapstructure:"build"`
	Config  string `mapstructure:"config"`
	Content string `mapstructure:"content"`
	Static  string `mapstructure:"static"`
	Src     string `mapstructure:"src"`
	Routes  string `mapstructure:"routes"`
	Lib     string `mapstructure:"lib"`
	API     string `mapstructure:"api"`
	Themes  string `mapstructure:"themes"`
}
