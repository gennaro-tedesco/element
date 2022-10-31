package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION number: change manually
const VERSION = "1.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print current version",
	Long:  "print current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
