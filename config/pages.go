/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package config

// Pages is the struct representing a public page and its content for a sveltin project.
type Pages struct {
	Content       string `mapstructure:"content"`
	Index         string `mapstructure:"index"`
	IndexEndpoint string `mapstructure:"index_pageload"`
	Slug          string `mapstructure:"slug"`
	SlugEndpoint  string `mapstructure:"slug_pageload"`
	SlugLayout    string `mapstructure:"slug_layout"`
}
