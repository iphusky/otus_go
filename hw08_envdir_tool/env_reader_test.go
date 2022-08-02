package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
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

		for name, val := range envs {
			require.Equal(t, filepath.Base(file.Name()), name)
			require.False(t, val.NeedRemove)
			require.Equal(t, "HELLO WORLD!", val.Value)
		}
	})

	t.Run("need delete", func(t *testing.T) {
		path, err := ioutil.TempDir(".", "testdata")
		require.NoError(t, err)

		defer os.RemoveAll(path)

		file, err := ioutil.TempFile(path, "TESTENV")
		require.NoError(t, err)

		defer os.Remove(file.Name())

		_, err = file.Write([]byte(""))
		require.NoError(t, err)
		err = file.Close()
		require.NoError(t, err)

		envs, err := ReadDir(path)
		require.NoError(t, err)

		for name, val := range envs {
			require.Equal(t, filepath.Base(file.Name()), name)
			require.True(t, val.NeedRemove)
			require.Equal(t, "", val.Value)
		}
	})
}
