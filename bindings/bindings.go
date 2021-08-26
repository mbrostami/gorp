package bindings

import (
	"github.com/c-bata/go-prompt"
)

var ControlC = prompt.ControlC
var ControlD = prompt.ControlD
var OptionLeft = []byte{0x1b, 0x62}
var OptionRight = []byte{0x1b, 0x66}
var OptionBackspace = []byte{0x1b, 0x7f}

func KeyBindings() []prompt.KeyBind {
	var kb []prompt.KeyBind
	kb = append(kb, quit())
	kb = append(kb, forceQuit())
	return kb
}

func ASCIICodeBindings() []prompt.ASCIICodeBind {
	var ab []prompt.ASCIICodeBind
	ab = append(ab, prompt.ASCIICodeBind{
		ASCIICode: OptionLeft,
		Fn:        prompt.GoLeftWord,
	})
	ab = append(ab, prompt.ASCIICodeBind{
		ASCIICode: OptionBackspace,
		Fn:        prompt.DeleteWord,
	})
	ab = append(ab, prompt.ASCIICodeBind{
		ASCIICode: OptionRight,
		Fn:        prompt.GoRightWord,
	})
	return ab
}
