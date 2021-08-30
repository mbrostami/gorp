package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mbrostami/gorp/bindings"
	"github.com/mbrostami/gorp/completer"

	"github.com/c-bata/go-prompt"
	log "github.com/sirupsen/logrus"
)

func main() {
	oArgs := os.Args
	if len(oArgs) < 2 {
		log.Fatal("need at least one argument")
	}
	args := oArgs[1:]
	baseCommand := args[0]

	history := prompt.NewHistory()
	p := prompt.New(
		func(in string) {
			newArgs := args[0:]
			if in != "" {
				newArgs = append(newArgs, strings.Split(in, " ")...)
				history.Add(in)
			}
			if err := execute(newArgs); err != nil {
				log.Error(err)
			}
		},
		completer.New(baseCommand),
		prompt.OptionPrefix(newline(args)),
		prompt.OptionPrefixTextColor(prompt.Purple),
		prompt.OptionAddKeyBind(bindings.KeyBindings()...),
		prompt.OptionAddASCIICodeBind(bindings.ASCIICodeBindings()...),
	)
	p.Run()
}

func newline(args []string) string {
	return fmt.Sprintf("%s > ", strings.Join(args[:], " "))
}
