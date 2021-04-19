package cmd

import (
	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	pt := getPeriodicTableData()
	for key, value := range pt {
		s = append(s, prompt.Suggest{Text: key, Description: value.(map[string]interface{})["Symbol"].(string)})
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
