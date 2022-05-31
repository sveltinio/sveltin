// Package helpers ...
package helpers

import (
	"testing"

	"github.com/matryer/is"
)

func TestInitStartersTemplatesMap(t *testing.T) {
	is := is.New(t)

	startersTemplatesMap := InitStartersTemplatesMap()
	static := startersTemplatesMap["starter"]

	is.Equal("sveltekit-static-starter", static.Name)
	is.Equal("https://github.com/sveltinio/sveltekit-static-starter", static.URL)
}
