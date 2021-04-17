package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cwd, _ = os.Getwd()

var rootCmd = &cobra.Command{
	Use:   "boilit",
	Short: "the periodic table on the command line",
	Long:  `the periodic table on the command line`,
	Run: func(cmd *cobra.Command, args []string) {
		m := getPeriodicTableData()
		fmt.Println(m["Hydrogen"].(map[string]interface{})["FirstIonization"])
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetHelpTemplate(getRootHelp())
	// rootCmd.Flags().StringP("path", "p", cwd, "root path of plugin directory")
}

func getRootHelp() string {
	return `
element: the periodic table on the command line

Arguments:
  name    the name of the element to show information of

Usage:
  element name [flag]

Flags:
  -h, --help   help for boilit
`
}
