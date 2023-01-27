/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package migrations

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/tidwall/gjson"
)

func isEqual(s1, s2 string) bool {
	return s1 == s2
}

func getDevDependency(content []byte, name string) (string, bool) {
	value := gjson.GetBytes(content, fmt.Sprintf("devDependencies.%s", name))
	if value.Exists() {
		return value.Str, true
	}
	return "", false
}

func versionAsNum(text string) (float64, error) {
	re := regexp.MustCompile(`(\d+\.\d+)`)
	match := re.FindStringSubmatch(text)

	return strconv.ParseFloat(match[1], 64)
}
