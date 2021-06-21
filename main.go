package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("need at least one argument")
	}

	executionLoop(args[1:])
}

func executionLoop(prefix []string) {
	name := prefix[0]
	var args []string
	if len(prefix) > 1 {
		args = prefix[1:]
	}

	cmdScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(WarningColor+" "+InfoColor+" > ", name, args)
	for cmdScanner.Scan() {
		if cmdScanner.Text() == "\\q" {
			return
		}
		var txt []string
		if cmdScanner.Text() != "" {
			txt = strings.Split(cmdScanner.Text(), " ")
		}
		argsNew := append(args, txt...)
		cmd := exec.Command(name, argsNew...)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.Env = os.Environ()
		err := cmd.Run()
		if err != nil {
			fmt.Printf("error: "+ErrorColor+"\n", err)
		}
		fmt.Printf(WarningColor+" "+InfoColor+" > ", name, args)
	}
}
