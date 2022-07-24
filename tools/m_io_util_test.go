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

func TestFindAll(t *testing.T) {
	all := tools.FindAll("AToB", `[A-Z][a-z]*`)
	t.Logf("find result: %v\n", all)
	if len(all) != 3 {
		t.Error("find wrong")
	}
}
