package main

import (
	"os"
	"os/exec"
	"strings"
)

func execute(args []string) error {
	var argsNew []string
	argsNew = append(argsNew, "-c")
	argsNew = append(argsNew, strings.Join(args[:], " "))
	cmd := exec.Command("bash", argsNew...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	cmd.Env = os.Environ()
	cmd.Dir, _ = os.Getwd()
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
