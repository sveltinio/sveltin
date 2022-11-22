package utils

import (
	"testing"

	"github.com/matryer/is"
)

func TestNumberUtils(t *testing.T) {
	is := is.New(t)

	is.Equal(2, PlusOne(1))
	is.Equal(3, Sum(1, 2))
	is.Equal(4, MinusOne(5))
}
