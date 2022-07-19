package mouseparser

import (
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

type MouseParserError struct {
	desc string
}

func NewMouseParserError(desc string) MouseParserError {
	return MouseParserError{
		desc: desc,
	}
}

func (err MouseParserError) Error() string {
	return err.desc
}

type MouseParser struct {
	filename string

	root *html.Node
	cur  *html.Node
}

// NewParser 根据给定的文件路径新建一个Parser
// path html文件路径
// *MouseParser 返回Parser的指针
// error 可能是打开文件的错误或者解析错误
func NewParser(path string) (*MouseParser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	n, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	var p MouseParser = MouseParser{
		filename: filepath.Base(path),
		root:     n,
		cur:      n,
	}
	return &p, nil
}

func (p *MouseParser) Reset() {
	p.cur = p.root
}

func (p *MouseParser) ReadTag(tagName string, attrs map[string]string) (*html.Node, error) {
	if p.cur.Type == html.ElementNode && p.cur.Data == tagName && attrsMatch(attrs, p.cur) {
		return p.cur, nil
	}
	for c := p.cur.FirstChild; c != nil; c = c.NextSibling {
		p.cur = c
		tn, err := p.ReadTag(tagName, attrs)
		if err != nil {
			continue
		}
		p.cur = tn
		return p.cur, nil
	}
	return nil, NewMouseParserError("tag not found.")
}

func (p *MouseParser) ReadTags(tagName string, attrs map[string]string) []*html.Node {
	return nil
}

func (p *MouseParser) ReadTagFrom(pTagName string, pAttrs map[string]string, tagName string, attrs map[string]string) *html.Node {
	return nil
}

func (p *MouseParser) ReadTagsFrom(pTagName string, pAttrs map[string]string, tagName string, attrs map[string]string) []*html.Node {
	return nil
}

func (p *MouseParser) ReadTextFrom(tagName string, attrs map[string]string) string {
	return ""
}

func attrsMatch(attrs map[string]string, node *html.Node) bool {
	if len(attrs) == 0 {
		return true
	}

	dic := extractAttrsFromNode(node)
	for k, v := range attrs {
		if t, ok := dic[k]; ok && t == v {
			return true
		}
	}
	return false
}

func extractAttrsFromNode(node *html.Node) map[string]string {
	var res map[string]string = make(map[string]string)
	for _, attr := range node.Attr {
		res[attr.Key] = attr.Val
	}
	return res
}
