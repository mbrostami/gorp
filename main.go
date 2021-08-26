package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/c-bata/go-prompt"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

type completer struct {
	suggestions map[string][]prompt.Suggest
	prefix      string
}

func NewCompleter(prefix string) completer {
	c := new(completer)
	c.prefix = prefix
	c.suggestions = make(map[string][]prompt.Suggest)
	c.suggestions["git"] = []prompt.Suggest{
		{Text: "status"},
		{Text: "push"},
		{Text: "fetch"},
		{Text: "rebase "},
		{Text: "checkout -b "},
		{Text: "checkout master"},
		{Text: "checkout develop"},
		{Text: "commit -m "},
	}
	return *c
}

func (c completer) Completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(c.suggestions[c.prefix], d.GetWordBeforeCursor(), true)
}

var quit = prompt.KeyBind{
	Key: prompt.ControlC,
	Fn: func(b *prompt.Buffer) {
		os.Exit(0) // log.Fatal doesn't work, but panic somehow avoids this issue...
	},
}
var fquit = prompt.KeyBind{
	Key: prompt.ControlD,
	Fn: func(b *prompt.Buffer) {
		os.Exit(0) // log.Fatal doesn't work, but panic somehow avoids this issue...
	},
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("need at least one argument")
	}
	executionLoop(args[1:])
}

var removePrevWord = prompt.KeyBind{
	Key: prompt.ControlA,
	Fn: func(b *prompt.Buffer) {
		prompt.DeleteWord(b)
	},
}

func executionLoop(prefix []string) {
	name := prefix[0]
	var args []string
	args = append(args, prefix[0:]...)

	history := prompt.NewHistory()
	p := prompt.New(func(in string) {
		t := in
		var txt []string
		if t != "" {
			txt = strings.Split(t, " ")
			history.Add(t)
			args = append(args, txt...)
		}
		var argsNew []string
		argsNew = append(argsNew, "-c")
		argsNew = append(argsNew, strings.Join(args[:], " "))
		cmd := exec.Command("bash", argsNew...)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.Env = os.Environ()
		cmd.Dir, _ = os.Getwd()
		err := cmd.Run()
		if err != nil {
			fmt.Printf("error: "+ErrorColor+"\n", err)
		}
	}, NewCompleter(name).Completer,
		prompt.OptionPrefix(newline(args)),
		prompt.OptionPrefixTextColor(prompt.Purple),
		prompt.OptionAddKeyBind(quit),
		prompt.OptionAddKeyBind(fquit),
		prompt.OptionAddKeyBind(removePrevWord),
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x1b, 0x62},
			Fn:        prompt.GoLeftWord,
		}),
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x1b, 0x66},
			Fn:        prompt.GoRightWord,
		}),
		prompt.OptionAddASCIICodeBind(prompt.ASCIICodeBind{
			ASCIICode: []byte{0x1b, 0x7f},
			Fn:        prompt.DeleteWord,
		}),
	)
	p.Run()
}

func newline(args []string) string {
	return fmt.Sprintf("%s > ", strings.Join(args[:], " "))
}
