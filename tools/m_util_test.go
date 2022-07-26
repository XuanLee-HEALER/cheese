package tools_test

import (
	"cheese/tools"
	"testing"
)

// func TestWriteToNewFile(t *testing.T) {
// 	wb, err := tools.WriteStrToNewFile("test/file", "test")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	t.Logf("write %d bytes\n", wb)
// }

func TestFindAll(t *testing.T) {
	all := tools.FindAll("AToB", `[A-Z][a-z]*`)
	t.Logf("find result: %v\n", all)
	if len(all) != 3 {
		t.Error("find wrong")
	}
}

func TestFindUrlSuffix(t *testing.T) {
	matched := tools.FindUrlSuffix("https://uploadstatic.mihoyo.com/ys-obc/2021/07/10/4328207/9638f8183e94151cc362689193064d52_2311649614603673169.png")
	t.Logf("find result: %s\n", matched)
	if matched != "" {
		t.Error("find wrong")
	}
}
