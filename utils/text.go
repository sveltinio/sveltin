/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package utils

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// IsEmpty returns true if a string is empty.
func IsEmpty(txt string) bool {
	return len(txt) == 0
}

// ToMDFile returns a string with .md extension
// example: ToMDFile("getting started", false) returns 'getting-started.md'.
// example: ToMDFile("getting started", true) returns 'GETTING-STARTED.md'.
func ToMDFile(txt string, uppercase bool) string {
	slug := ToSlug(txt)
	if uppercase {
		slug = strings.ToUpper(slug)
	}
	return fmt.Sprintf("%s.md", slug)
}

// ToLibFile returns a string a valid lib filename
// example: ToLibFile("category") returns 'apiCategory.ts'.
func ToLibFile(txt string) string {
	vName := ToVariableName(txt)
	return fmt.Sprintf("load%s.ts", ToTitle(vName))
}

// ToTitle replace all '-' char with a white space and
// returns a copy of string s with all letters
// of string whose begin words mapped to their title case.
func ToTitle(txt string) string {
	cleanTitle := strings.ReplaceAll(txt, "-", " ")
	c := cases.Title(language.Und, cases.NoLower)
	return c.String(cleanTitle)
}

// ToURL returns a trimmed string with '/' as prefix.
func ToURL(txt string) string {
	return "/" + Trimmed(txt)
}

// Trimmed strips away '"' from a string.
func Trimmed(txt string) string {
	return strings.Trim(txt, "\"")
}

// ToSlug returns a copy of string with lowercase
// replacing "_" and whitespaces with "-"
// example: ToSlug("New Resource") returns new-resource.
func ToSlug(txt string) string {
	cleanString := strings.ToLower(txt)
	cleanString = Trimmed(cleanString)
	cleanString = strings.ReplaceAll(cleanString, " ", "-")
	cleanString = strings.ReplaceAll(cleanString, "_", "-")
	return cleanString
}

// ToSnakeCase returns a copy of string with lowercase
// replacing "-" and whitespaces with "_"
// example: ToSnakeCase("New Resource") returns new_resource.
func ToSnakeCase(txt string) string {
	cleanString := strings.ToLower(txt)
	cleanString = strings.ReplaceAll(cleanString, " ", "_")
	cleanString = strings.ReplaceAll(cleanString, "-", "_")
	return cleanString
}

// ToBasePath returns a copy of string replacing all occurrences
// for a string with trailing slash.
func ToBasePath(fullpath string, replace string) string {
	return strings.ReplaceAll(fullpath, replace+"/", "")
}

// ToVariableName returns a copy of string to be used as variable name.
func ToVariableName(txt string) string {
	slug := ToSlug(txt)
	var frags = strings.Split(slug, "-")
	for i := range frags {
		if i != 0 {
			frags[i] = ToTitle(frags[i])
		}
	}
	return strings.Join(frags, "")
}

// ReplaceIfNested returns a copy of string to be used as variable name when a resource name is a nested one.
func ReplaceIfNested(txt string) string {
	return strings.ReplaceAll(txt, "/", "_")
}

// Today returns the current date as formatted string "DD-ShortMonth-YYYY".
func Today() string {
	return time.Now().Format("02-Jan-2006")
}

// CurrentYear returns the current calendar year as a string.
func CurrentYear() string {
	return time.Now().Format("2006")
}
