package main

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int, err error) {
	for name, val := range env {
		if val.NeedRemove {
			os.Unsetenv(name)
		} else {
			if _, ok := os.LookupEnv(name); ok {
				os.Unsetenv(name)
			}
			os.Setenv(name, val.Value)
		}
	}

	command := cmd[0]
	arg1 := cmd[1]
	arg2 := cmd[2]

	result, err := exec.Command(command, arg1, arg2).Output()
	if err != nil {
		return -1, err
	}

	fmt.Fprintln(os.Stdout, string(result))

	return 0, nil
}
