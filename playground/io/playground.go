package main

import (
	"bufio"
	"os"

	"github.com/labstack/gommon/log"
)

func main() {
	i, err := writeStrToNewFile("./test", "test")
	if err != nil {
		log.Error("something wrong!")
	}
	log.Info("write to %d bytes.", i)
}

func writeStrToNewFile(fname string, content string) (int, error) {
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
