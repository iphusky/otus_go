package main

import (
	"errors"
	"io"
	"os"
)

var ErrNotEnoughArgs = errors.New("not enough arguments")

func main() {
	if err := run(); err != nil {
		io.WriteString(os.Stderr, err.Error())
	}
}

func run() error {
	args := os.Args

	if len(args) < 4 {
		return ErrNotEnoughArgs
	}

	env, err := ReadDir(args[1])
	if err != nil {
		return err
	}

	_, err = RunCmd(args[2:], env)

	if err != nil {
		return err
	}

	return nil
}
