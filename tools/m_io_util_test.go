package tools_test

import (
	"cheese/tools"
	"testing"
)

func TestWriteToNewFile(t *testing.T) {
	wb, err := tools.WriteStrToNewFile("test/file", "test")
	if err != nil {
		t.Error(err)
	}
	t.Logf("write %d bytes\n", wb)
}
