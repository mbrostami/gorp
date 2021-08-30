package completer

import "github.com/c-bata/go-prompt"

func gitSuggestions(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "status"},
		{Text: "push"},
		{Text: "pull"},
		{Text: "fetch"},
		{Text: "rebase "},
		{Text: "checkout -b "},
		{Text: "checkout master"},
		{Text: "checkout develop"},
		{Text: "commit -m "},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
