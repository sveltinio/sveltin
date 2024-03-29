package resources

import (
	"testing"

	"github.com/matryer/is"
)

func TestSveltinFilesFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/statics/files/mdsvex.config.js", SveltinFilesFS["mdsvex"])
}

func TestSveltinStaticsFS(t *testing.T) {
	is := is.New(t)
	is.Equal("internal/statics/images/dummy.jpeg", SveltinImagesFS["dummy"])
}
