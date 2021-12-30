package utils

import (
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestTextUtils(t *testing.T) {
	is := is.New(t)

	is.Equal("Getting Started", ToTitle("getting-started"))
	is.Equal("/getting-started", ToURL("getting-started"))
	is.Equal(time.Now().Format("02-Jan-2006"), Today())
	is.Equal("2021", CurrentYear())
	is.Equal(2, PlusOne(1))
	is.Equal(3, Sum(1, 2))
}
