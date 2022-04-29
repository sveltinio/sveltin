/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package helpers ...
package helpers

import (
	"github.com/sveltinio/sveltin/config"
)

//InitAppTemplatesMap creates a map[string]string containining project name and repo url
// used by the `sveltin new` command to clone the starter project.
func InitAppTemplatesMap() map[string]config.AppTemplate {
	appTemplatesMap := make(map[string]config.AppTemplate)

	appTemplatesMap["starter"] = config.AppTemplate{
		Name: "sveltekit-static-starter",
		URL:  "https://github.com/sveltinio/sveltekit-static-starter",
	}

	appTemplatesMap["blog-theme-starter"] = config.AppTemplate{
		Name: "blog-theme-starter",
		URL:  "https://github.com/sveltinio/blog-theme-starter",
	}
	return appTemplatesMap
}
