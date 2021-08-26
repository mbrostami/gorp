package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mbrostami/gorp/bindings"

	"github.com/mbrostami/gorp/completer"

	"github.com/c-bata/go-prompt"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
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
				fmt.Printf("error: "+ErrorColor+"\n", err)
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
