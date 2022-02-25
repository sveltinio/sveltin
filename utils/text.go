/**
 * Copyright © 2021 Mirco Veltri <github@mircoveltri.me>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

// Package utils ...
package utils

import (
	"strings"
	"time"
)

// ToMDFile returns a string with .md suffix
// example: ToMDFile("welcome") returns 'welcome.md'.
func ToMDFile(text string) string {
	return strings.ToUpper(text) + ".md"
}

// ToLibFilename returns a string a valid lib filename
// example: ToLibFilename("category") returns 'getCategory.js'.
func ToLibFilename(txt string) string {
	return `get` + strings.Title(txt) + `.js`
}

// ToTitle replace all '-' char with a white space and
// returns a copy of string s with all letters
// of string whose begin words mapped to their title case.
func ToTitle(txt string) string {
	cleanTitle := strings.ReplaceAll(txt, "-", " ")
	return strings.Title(cleanTitle)
}

// ToURL returns a trimmed string with '/' as prefix.
func ToURL(txt string) string {
	return "/" + Trimmed(txt)
}

// Trimmed strips away '"' from a string.
func Trimmed(txt string) string {
	return strings.Trim(txt, "\"")
}

// ToValidName returns a copy of string replacing '_' with '-'
// example: ToValidName("getting_started") returns 'getting-started'.
func ToValidName(txt string) string {
	return strings.ReplaceAll(txt, "_", "-")
}

// ToBasePath returns a copy of string replacing all occurrences
// for a string with trailing slash.
func ToBasePath(fullpath string, replace string) string {
	return strings.ReplaceAll(fullpath, replace+"/", "")
}

// ToPageVariableName returns a human readable string
// replacing '-' with white space and capitalizing
// first letters of each word.
func ToPageVariableName(txt string) string {
	var frags = strings.Split(txt, "-")
	for i := range frags {
		if i != 0 {
			frags[i] = strings.Title(frags[i])
		}
	}
	return strings.Join(frags, "")
}

// Today returns the current date as formatted string "DD-ShortMonth-YYYY".
func Today() string {
	return time.Now().Format("02-Jan-2006")
}

// CurrentYear returns the current calendar year as a string.
func CurrentYear() string {
	return time.Now().Format("2006")
}

// PlusOne adds one to the integer parameter.
func PlusOne(x int) int {
	return x + 1
}

// Sum adds two integer values.
func Sum(x int, y int) int {
	return x + y
}
