package cmd

import (
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

var cwd, _ = os.Getwd()

var rootCmd = &cobra.Command{
	Use:   "boilit",
	Args:  cobra.MaximumNArgs(1),
	Short: "the periodic table on the command line",
	Long:  `the periodic table on the command line`,
	Run: func(cmd *cobra.Command, args []string) {
		pt := getPeriodicTableData()
		if len(args) > 0 {
			printElementData(args[0], pt)
		} else {
			t := prompt.Input("select element: ",
				completer,
				prompt.OptionSuggestionTextColor(prompt.Color(prompt.ControlD)),
				prompt.OptionSuggestionBGColor(prompt.Color(prompt.ControlA)),
				prompt.OptionSelectedSuggestionTextColor(prompt.Color(prompt.LightGray)),
				prompt.OptionSelectedSuggestionBGColor(prompt.DarkGray),
				prompt.OptionPreviewSuggestionTextColor(prompt.Color(prompt.ControlD)),
			)
			printElementData(t, pt)
		}
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
