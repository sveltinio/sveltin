/**
 * Copyright © 2021-present Sveltin contributors <github@sveltin.io>
 *
 * Use of this source code is governed by Apache 2.0 license
 * that can be found in the LICENSE file.
 */

package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/resources"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// IsEmpty returns true if the string is empty.
func IsEmpty(txt string) bool {
	return len(txt) == 0
}

// IsEmptySlice returns true if the string slice is empty.
func IsEmptySlice(txt []string) bool {
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

// ConvertJSStringToStringArray returns a JS/TS array of string from a comma separated JS/TS string as input.
func ConvertJSStringToStringArray(value string) string {
	res := strings.Trim(value, " ")
	res = strings.ReplaceAll(res, "'", "")
	res = strings.ReplaceAll(res, "\"", "")
	res1 := strings.Split(res, ",")
	res1 = common.RemoveEmpty(res1)

	var newKeywords []string
	for _, v := range res1 {
		_v := strings.Trim(v, " ")
		newKeywords = append(newKeywords, fmt.Sprintf("'%s'", _v))

	}
	return fmt.Sprintf("[%s]", strings.Join(newKeywords, ", "))
}

// MakeCmdLongMsg returns the string used for the command long help message.
func MakeCmdLongMsg(text string) string {
	return fmt.Sprintf(`%s
%s`, resources.GetASCIIArt(), text)
}

// ReplaceLast returns a copy of the given string with the lastnon-overlapping instances of the old string replaced with the new one.
func ReplaceLast(s, old, new string) string {
	idx := strings.LastIndex(s, old)
	if idx == -1 {
		return s
	}
	return s[:idx] + new + s[idx+len(old):]
}

// ToFloat64 converts the given string to a float64.
func ToFloat64(s string) (float64, error) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// SemVersionToFloat converts a given semantic version string as float64.
func SemVersionToFloat(v string) float64 {
	if IsEmpty(v) {
		return 0
	}
	// to make semversion string looks like a float
	vCleaned := ReplaceLast(v, ".", "")
	// convert to float64
	f, err := ToFloat64(vCleaned)
	if err != nil {
		// Light: red-500, Dark: red-500
		redColor := lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"}
		redText := lipgloss.NewStyle().Foreground(redColor).Render
		fmt.Println(redText(err.Error()))
	}
	return f
}
