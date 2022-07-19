package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotFound          = errors.New("file not found")
)

func getLimit(src *os.File, limitParam int64) (int64, error) {
	fi, err := src.Stat()
	if err != nil {
		return 0, err
	}

	if limitParam > fi.Size() || limitParam == 0 {
		limitParam = fi.Size()
	}

	return limitParam, nil
}

func openFileAndSeek(path string, offset int64) (*os.File, error) {
	fileFrom, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		return nil, err
	}

	fi, err := fileFrom.Stat()
	if err != nil {
		return nil, err
	}

	if fi.Size() == 0 {
		return nil, ErrUnsupportedFile
	}

	if fi.Size() < offset {
		return nil, ErrOffsetExceedsFileSize
	}

	_, err = fileFrom.Seek(offset, 0)

	if err != nil {
		return nil, err
	}

	return fileFrom, nil
}

func createFileAndPath(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o777); err != nil {
		return nil, err
	}

	return os.Create(path)
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileFrom, err := openFileAndSeek(fromPath, offset)
	if err != nil {
		return err
	}

	defer fileFrom.Close()

	fileTo, err := createFileAndPath(toPath)
	if err != nil {
		return err
	}

	defer fileTo.Close()

	limit, err = getLimit(fileFrom, limit)

	if err != nil {
		return err
	}

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(fileFrom)

	_, err = io.CopyN(fileTo, barReader, limit)

	if err != nil {
		switch {
		case errors.Is(err, io.EOF):
			return nil
		default:
			fmt.Printf("unexpected error: %s\n", err)
		}
		return err
	}

	bar.Finish()

	return nil
}
