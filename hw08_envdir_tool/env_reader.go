package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

var ErrEmptyDir = errors.New("empty dir path")

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	if len([]rune(dir)) == 0 {
		return nil, ErrEmptyDir
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	envs := make(Environment, len(files))

	for _, file := range files {
		bytes, err := ReadFirstLineFromFile(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}

		str := formatContent(bytes)

		env := new(EnvValue)

		if len(str) == 0 {
			env.NeedRemove = true
		} else {
			env.Value = str
			env.NeedRemove = false
		}

		name := strings.ReplaceAll(file.Name(), "=", "")
		envs[name] = *env
	}

	return envs, nil
}

func ReadFirstLineFromFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()

	if errors.Is(err, io.EOF) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return line, nil
}

func formatContent(content []byte) string {
	content = bytes.Replace(content, []byte("\x00"), []byte("\n"), 1)
	position := strings.Index(string(content), "\n")
	if position > -1 {
		content = content[:position]
	}
	trimStr := strings.TrimRight(string(content), "\t\v\r")
	trimStr = strings.TrimLeft(trimStr, "\t\v\r")
	trimStr = strings.ReplaceAll(trimStr, " ", "")
	return trimStr
}
