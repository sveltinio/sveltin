/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package main

import (
	_ "embed"

	"github.com/sveltinio/sveltin/cmd"
)

//go:embed resources/sveltin.yaml
var yamlConfig []byte

func main() {
	cmd.YamlConfig = yamlConfig
	cmd.Execute()
}
