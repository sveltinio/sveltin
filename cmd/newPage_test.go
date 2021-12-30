package cmd

import (
	"path/filepath"
	"testing"

	"github.com/matryer/is"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/sveltinio/sveltin/common"
	"github.com/sveltinio/sveltin/helpers"
)

func TestPageCmd(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		args     []string
		function func() func(cmd *cobra.Command, args []string)
	}{
		{
			args: []string{"page", "about", "contact", "-t", "svelte"},
			function: func() func(cmd *cobra.Command, args []string) {
				return func(cmd *cobra.Command, args []string) {
					err := common.CheckMaxArgs(args, 1)
					re := err.(*common.SveltinError)
					is.Equal(32, re.Code)
					is.Equal("SVELTIN NumOfArgsNotValidErrorWithMessage", re.Message)
				}
			},
		},
		{
			args: []string{"page", "about", "-t", "svelte"},
			function: func() func(cmd *cobra.Command, args []string) {
				return func(cmd *cobra.Command, args []string) {
					err := common.CheckMaxArgs(args, 1)
					is.NoErr(err)

					pageName, err := getPageName(args)
					is.Equal("about", pageName)
					is.NoErr(err)
				}
			},
		},
		{
			args: []string{"page", "about", "-t", "svelt"},
			function: func() func(cmd *cobra.Command, args []string) {
				return func(cmd *cobra.Command, args []string) {
					err := common.CheckMaxArgs(args, 1)
					is.NoErr(err)

					pageName, err := getPageName(args)
					is.Equal("about", pageName)
					is.NoErr(err)

					_, err = getPageType(pageType)
					re := err.(*common.SveltinError)
					is.Equal("SVELTIN PageTypeNotValidError", re.Message)
				}
			},
		},
		{
			args: []string{"page", "about", "-t", "svelte"},

			function: func() func(cmd *cobra.Command, args []string) {
				return func(cmd *cobra.Command, args []string) {
					memFS := afero.NewMemMapFs()
					memFS.MkdirAll("output", 0755)

					err := common.CheckMaxArgs(args, 1)
					is.NoErr(err)

					pageName, err := getPageName(args)
					is.Equal("about", pageName)
					is.NoErr(err)

					selectedPageType, err := getPageType(pageType)
					is.Equal("svelte", selectedPageType)
					is.NoErr(err)

					pathToFile := filepath.Join("output", "about.svelte")

					afero.WriteFile(memFS, pathToFile, []byte(""), 0644)

					exists, err := afero.Exists(memFS, pathToFile)
					is.NoErr(err)
					is.True(exists)
				}
			},
		},
	}

	for _, tc := range tt {
		sveltinCmd := &cobra.Command{Use: "sveltin"}
		pageCmd := &cobra.Command{Use: "page", Run: tc.function()}
		pageCmdFlags(pageCmd)
		sveltinCmd.AddCommand(pageCmd)

		c, err := helpers.ExecuteCommandC(sveltinCmd, tc.args...)

		is.Equal("page", c.Name())
		is.NoErr(err)
	}

}
