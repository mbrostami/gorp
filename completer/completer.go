package completer

import "github.com/c-bata/go-prompt"

func New(command string) func(d prompt.Document) []prompt.Suggest {
	switch command {
	case "git":
		return gitSuggestions
	}
	return func(d prompt.Document) []prompt.Suggest { return nil }
}
