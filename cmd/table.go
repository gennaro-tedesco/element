package cmd

import (
	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "print the periodic table in ascii format",
	Long:  "print the periodic table in ascii format",
	Run: func(cmd *cobra.Command, args []string) {
		printPeriodicTable()
	},
}

func init() {
	rootCmd.AddCommand(tableCmd)
}
