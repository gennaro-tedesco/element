package cmd

import (
	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "show a random element",
	Long:  `show a random element`,
	Run: func(cmd *cobra.Command, args []string) {
		pt := getPeriodicTableData()
		element := getRandomKey(pt)
		printElementData(element, pt)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
