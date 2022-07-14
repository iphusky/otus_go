package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("copy full file", func(t *testing.T) {
		const pathFrom = "./testdata/input.txt"
		const pathTo = "./testdata/my_input.txt"

		err := Copy(pathFrom, pathTo, 0, 0)

		require.NoError(t, err)

		file, err := os.Open(pathTo)
		file.Close()

		os.Remove(pathTo)

		require.NoError(t, err)
	})
	t.Run("offset is bigger then filesize", func(t *testing.T) {
		const pathFrom = "./testdata/input.txt"
		const pathTo = "./testdata/my_input.txt"
		offset := 6618

		err := Copy(pathFrom, pathTo, int64(offset), 0)
		os.Remove(pathTo)

		require.Error(t, err, ErrOffsetExceedsFileSize)
	})
	t.Run("test copy with limit", func(t *testing.T) {
		var limit, offset int64

		const pathFrom = "./testdata/input.txt"
		const pathTo = "./testdata/my_input_limit10.txt"
		limit = 10
		offset = 0

		err := Copy(pathFrom, pathTo, offset, limit)
		require.Nil(t, err)

		fi, err := os.Stat(pathTo)
		require.Nil(t, err)

		require.Equal(t, limit, fi.Size())
		os.Remove(pathTo)
	})

	t.Run("test copy with offset", func(t *testing.T) {
		var limit, offset int64

		const pathFrom = "./testdata/input.txt"
		const pathTo = "./testdata/my_input_limit0_offset6000.txt"
		limit = 0
		offset = 6000

		err := Copy(pathFrom, pathTo, offset, limit)
		require.Nil(t, err)

		fiFrom, err := os.Stat(pathFrom)
		require.Nil(t, err)

		fiTo, err := os.Stat(pathTo)
		require.Nil(t, err)

		expected := fiFrom.Size() - offset

		require.Equal(t, expected, fiTo.Size())
		os.Remove(pathTo)
	})

	t.Run("test copy with offset and short file", func(t *testing.T) {
		var limit, offset int64

		const pathFrom = "./testdata/input.txt"
		const pathTo = "./testdata/my_input_limit0_offset6000.txt"
		limit = 0
		offset = 6000

		err := Copy(pathFrom, pathTo, offset, limit)
		require.Nil(t, err)

		fi, err := os.Stat(pathTo)
		require.Nil(t, err)

		require.Equal(t, int64(617), fi.Size())
		os.Remove(pathTo)
	})

	t.Run("test copy with limit and offset", func(t *testing.T) {
		var limit, offset int64

		const pathFrom = "./testdata/input.txt"
		const pathTo = "./testdata/my_input_limit10_offset1000.txt"
		limit = 10
		offset = 1000

		err := Copy(pathFrom, pathTo, offset, limit)
		require.Nil(t, err)

		fi, err := os.Stat(pathTo)
		require.Nil(t, err)

		require.Equal(t, limit, fi.Size())
		os.Remove(pathTo)
	})
}
