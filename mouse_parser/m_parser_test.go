package mouseparser_test

import (
	mouseparser "cheese/mouse_parser"
	"testing"

	"github.com/labstack/gommon/log"
)

const testHtml = "../roledetail/宵宫.html"

var testParser *mouseparser.MouseParser

func init() {
	localParser, err := mouseparser.NewParser(testHtml)
	if err != nil {
		log.Error(err)
	}

	testParser = localParser
}

func TestReadTag(t *testing.T) {
	testAttrs := make(map[string]string)
	testAttrs["class"] = "obc-tmp-character__box--title"
	n, err := testParser.ReadTag("p", testAttrs)
	t.Log("n.tagName ", n.Data)
	t.Log("n.TextContent ", n.FirstChild.Data)
	if err != nil || n.FirstChild.Data != "宵宫" {
		t.Error(err)
	}
}
