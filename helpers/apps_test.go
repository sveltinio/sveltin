// Package helpers ...
package helpers

import (
	"testing"

	"github.com/matryer/is"
)

func TestInitAppTemplatesMap(t *testing.T) {
	is := is.New(t)

	appTemplatesMap := InitAppTemplatesMap()
	static := appTemplatesMap["starter"]

	is.Equal("sveltekit-static-starter", static.Name)
	is.Equal("https://github.com/sveltinio/sveltekit-static-starter", static.URL)
}
