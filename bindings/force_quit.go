package bindings

import (
	"os"

	"github.com/c-bata/go-prompt"
)

func forceQuit() prompt.KeyBind {
	return prompt.KeyBind{
		Key: ControlD,
		Fn: func(b *prompt.Buffer) {
			os.Exit(0) // log.Fatal doesn't work, but panic somehow avoids this issue...
		},
	}
}
