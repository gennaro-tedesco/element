package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cwd, _ = os.Getwd()

var rootCmd = &cobra.Command{
	Use:   "boilit",
	Args:  cobra.ExactArgs(1),
	Short: "the periodic table on the command line",
	Long:  `the periodic table on the command line`,
	Run: func(cmd *cobra.Command, args []string) {
		pt := getPeriodicTableData()
		printElementData(args[0], pt)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
element: the periodic table on the command line

Arguments:
  name    the name of the element to show information of

Usage:
  element name [flag]

Available Commands:
  random  print information of a randomly chosen element

Flags:
  -h, --help   help for boilit
`
}
