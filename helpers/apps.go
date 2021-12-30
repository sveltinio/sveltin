/*
Copyright © 2021 Mirco Veltri <github@mircoveltri.me>

Use of this source code is governed by Apache 2.0 license
that can be found in the LICENSE file.
*/
package helpers

import (
	"github.com/sveltinio/sveltin/config"
)

func InitAppTemplatesMap() map[string]config.AppTemplate {
	appTemplatesMap := make(map[string]config.AppTemplate)

	appTemplatesMap["starter"] = config.AppTemplate{
		Name: "sveltekit-static-starter",
		URL:  "https://github.com/sveltinio/sveltekit-static-starter",
	}
	return appTemplatesMap
}
