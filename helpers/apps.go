/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package helpers

import (
	"github.com/sveltinio/sveltin/config"
)

// InitStartersTemplatesMap creates a map[string]string containining project name and repo url
// used by the `sveltin new` command to clone the starter project.
func InitStartersTemplatesMap() map[string]config.StarterTemplate {
	starterTemplatesMap := make(map[string]config.StarterTemplate)

	starterTemplatesMap["starter"] = config.StarterTemplate{
		Name: "sveltekit-static-starter",
		URL:  "https://github.com/sveltinio/sveltekit-static-starter",
	}

	starterTemplatesMap["blog-theme-starter"] = config.StarterTemplate{
		Name: "blog-theme-starter",
		URL:  "https://github.com/sveltinio/blog-theme-starter",
	}
	return starterTemplatesMap
}
