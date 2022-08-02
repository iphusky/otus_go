package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("simple run", func(t *testing.T) {
		path, err := ioutil.TempDir(".", "testdata")
		require.NoError(t, err)

		defer os.RemoveAll(path)

		file, err := ioutil.TempFile(path, "TESTENV")
		require.NoError(t, err)

		defer os.Remove(file.Name())

		_, err = file.Write([]byte("HELLO WORLD!"))
		require.NoError(t, err)
		err = file.Close()
		require.NoError(t, err)

		envs, err := ReadDir(path)
		require.NoError(t, err)

		var args []string

		args = append(args, "ls", "-l", "-a")

		code, err := RunCmd(args, envs)
		require.NoError(t, err)

		testEnv, ok := os.LookupEnv(filepath.Base(file.Name()))
		require.True(t, ok)
		require.Equal(t, "HELLO WORLD!", testEnv)
		require.Equal(t, 0, code)
	})
}
