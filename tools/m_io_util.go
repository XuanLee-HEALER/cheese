package tools

import (
	"bufio"
	"os"
	"path/filepath"
)

func WriteStrToNewFile(fname string, content string) (int, error) {
	dirs := filepath.Dir(fname)
	err := os.MkdirAll(dirs, os.ModeDir|os.ModePerm)
	if err != nil {
		return 0, err
	}

	f, err := os.Create(fname)
	if err != nil {
		return 0, err
	}

	writer := bufio.NewWriter(f)
	defer f.Close()

	l, err := writer.WriteString(content)
	if err != nil {
		return 0, err
	}
	writer.Flush()

	return l, nil
}
