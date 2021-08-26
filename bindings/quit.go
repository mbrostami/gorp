package bindings

import (
	"os"

	"github.com/c-bata/go-prompt"
)

func quit() prompt.KeyBind {
	return prompt.KeyBind{
		Key: ControlC,
		Fn: func(b *prompt.Buffer) {
			os.Exit(0) // log.Fatal doesn't work, but panic somehow avoids this issue...
		},
	}
}
