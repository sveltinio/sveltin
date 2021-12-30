package common

import (
	"strconv"
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

func TestValidCheckMinMaxArgs(t *testing.T) {
	tests := []struct {
		items []string
		min   int
		max   int
	}{
		{items: []string{"svelte", "sveltekit"}, min: 1, max: 2},
		{items: []string{"svelte", "sveltekit", "framework"}, min: 1, max: 4},
		{items: []string{"svelte"}, min: 1, max: 1},
	}

	for _, tc := range tests {
		is := is.New(t)

		err := CheckMinMaxArgs(tc.items, tc.min, tc.max)
		is.NoErr(err)
	}
}

func TestCheckMinMaxArgsBoundaries(t *testing.T) {
	tests := []struct {
		items []string
		min   int
		max   int
	}{
		{items: []string{"svelte"}, min: 2, max: 3},
	}

	for _, tc := range tests {
		is := is.New(t)

		err := CheckMinMaxArgs(tc.items, tc.min, tc.max)
		re := err.(*SveltinError)
		is.Equal(32, re.Code)
		is.Equal("SVELTIN NumOfArgsNotValidErrorWithMessage", re.Message)
		is.Equal(`SVELTIN NumOfArgsNotValidErrorWithMessage: This command expects at least `+strconv.Itoa(tc.min)+` argument.
Please check the help: sveltin [command] -h`, re.Error())
	}
}

func TestNotValidCheckMinMaxArgs(t *testing.T) {
	tests := []struct {
		items []string
		min   int
		max   int
	}{
		{items: []string{"svelte", "sveltekit", "framework"}, min: 1, max: 2},
	}

	for _, tc := range tests {
		is := is.New(t)

		err := CheckMinMaxArgs(tc.items, tc.min, tc.max)
		re := err.(*SveltinError)
		is.Equal(32, re.Code)
		is.Equal("SVELTIN NumOfArgsNotValidErrorWithMessage", re.Message)
		is.Equal(`SVELTIN NumOfArgsNotValidErrorWithMessage: This command expects maximum `+strconv.Itoa(tc.max)+` arguments.
Please check the help: sveltin [command] -h`, re.Error())
	}

}

func TestCheckMaxArgs(t *testing.T) {
	tests := []struct {
		items []string
		max   int
	}{
		{items: []string{"svelte", "sveltekit"}, max: 2},
		{items: []string{"svelte", "sveltekit", "framework"}, max: 4},
		{items: []string{"svelte"}, max: 1},
	}

	for _, tc := range tests {
		is := is.New(t)

		err := CheckMaxArgs(tc.items, tc.max)
		is.NoErr(err)
	}
}

func TestCheckMinMaxArgsWithMaxEqualsZero(t *testing.T) {
	tests := []struct {
		items []string
		max   int
	}{
		{items: []string{"svelte", "sveltekit", "framework"}, max: 0},
	}

	for _, tc := range tests {
		is := is.New(t)

		err := CheckMaxArgs(tc.items, tc.max)
		re := err.(*SveltinError)
		is.Equal(32, re.Code)
		is.Equal("SVELTIN NumOfArgsNotValidErrorWithMessage", re.Message)
		is.Equal(`SVELTIN NumOfArgsNotValidErrorWithMessage: This command expects no arguments. Please check the help: sveltin [command] -h`, re.Error())
	}
}

func TestCheckMinMaxArgsWithNumOfArgsGreaterThanMax(t *testing.T) {
	tests := []struct {
		items []string
		max   int
	}{
		{items: []string{"svelte", "sveltekit", "framework"}, max: 2},
	}

	for _, tc := range tests {
		is := is.New(t)

		err := CheckMaxArgs(tc.items, tc.max)
		re := err.(*SveltinError)
		is.Equal(32, re.Code)
		is.Equal("SVELTIN NumOfArgsNotValidErrorWithMessage", re.Message)
		is.Equal(`SVELTIN NumOfArgsNotValidErrorWithMessage: This command expects maximum `+strconv.Itoa(tc.max)+` arguments.
Please check the help: sveltin [command] -h`, re.Error())
	}
}
