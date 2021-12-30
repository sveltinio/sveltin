package utils

import (
	"strings"
	"time"
)

func ToMDFile(text string) string {
	return strings.ToUpper(text) + ".md"
}

func ToLibFilename(txt string) string {
	return `get` + strings.Title(txt) + `.js`
}

func ToTitle(txt string) string {
	cleanTitle := strings.ReplaceAll(txt, "-", " ")
	return strings.Title(cleanTitle)
}

func ToURL(txt string) string {
	return "/" + Trimmed(txt)
}

func Trimmed(txt string) string {
	return strings.Trim(txt, "\"")
}

func Today() string {
	return time.Now().Format("02-Jan-2006")
}

func CurrentYear() string {
	return time.Now().Format("2006")
}

func PlusOne(x int) int {
	return x + 1
}

func Sum(x int, y int) int {
	return x + y
}
