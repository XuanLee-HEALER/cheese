package tools

import (
	"errors"
	"fmt"

	"golang.org/x/net/html"
)

func FetchTextFrom(node *html.Node) string {
	return ""
}

func QueryAllBySelector(node *html.Node, nTag string, nClass []string) []*html.Node {
	res := []*html.Node{}

	if node.Type == html.ElementNode && node.Data == nTag && classMatch(node, nClass) {
		res = append(res, node)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		eles := QueryAllBySelector(c, nTag, nClass)
		if len(res) != 0 {
			res = append(res, eles...)
		}
	}

	return res
}

func QueryBySelector(node *html.Node, nTag string, nClass []string) (*html.Node, error) {
	if node.Type == html.ElementNode && node.Data == nTag && classMatch(node, nClass) {
		return node, nil
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		tn, _ := QueryBySelector(c, nTag, nClass)
		if tn != nil {
			return tn, nil
		}
	}
	return nil, errors.New("query error")
}

func classMatch(node *html.Node, nClass []string) bool {
	if len(nClass) == 0 {
		return true
	}

	find := func(strs []string, str string) bool {
		for _, s := range strs {
			if s == str {
				return true
			}
		}
		return false
	}
	for _, attr := range node.Attr {
		if attr.Key == "class" && find(nClass, attr.Val) {
			return true
		}
	}
	return false
}

func IterateAllNode(node *html.Node) {
	if node != nil {
		IterateNode(node)
		for sib := node.NextSibling; sib != nil; sib = sib.NextSibling {
			IterateNode(sib)
		}
	}
}

func IterateNode(node *html.Node) {
	if node.Data == "ul" {
		fmt.Println(node.Type, " ", node.Data, " ", node.Attr)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		IterateNode(c)
	}
}
