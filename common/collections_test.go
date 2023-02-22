package common

import (
	"testing"

	"github.com/matryer/is"
)

func TestContains(t *testing.T) {
	tests := []struct {
		dictionary []string
		want       string
	}{
		{dictionary: []string{"svelte", "sveltekit", "framework"}, want: "svelte"},
	}

	for _, tc := range tests {
		is := is.New(t)
		got := Contains(tc.dictionary, tc.want)
		is.True(got)
	}
}

func TestNotValidContains(t *testing.T) {
	tests := []struct {
		dictionary []string
		want       string
	}{
		{dictionary: []string{"svelte", "sveltekit", "framework"}, want: "svelt"},
	}

	for _, tc := range tests {
		is := is.New(t)
		got := Contains(tc.dictionary, tc.want)
		is.Equal(false, got)

	}
}
