package cmd

import (
	"fmt"

	"github.com/pydio/packr"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints packr version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(packr.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
