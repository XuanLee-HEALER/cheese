package tools

import (
	"bufio"
	"os"
)

func WriteStrToNewFile(fname string, content string) (int, error) {
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
