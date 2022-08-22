package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	// CliVersion is the current sveltin cli version number.
	CliVersion string = "0.9.0-pre.1"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Sveltin",
	Long:  `This is Sveltin's version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(CliVersion)
	},
}
