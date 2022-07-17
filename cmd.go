package main

import (
	"cheese/tools"
	"context"
	"fmt"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/html"
)

// 角色信息
const roleUrl = "https://bbs.mihoyo.com/ys/obc/channel/map/189/25"

func main() {
	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()
	// fetchRoleHtml(ctx)
	f, err := os.Open("role.html")
	if err != nil {
		log.Fatal("open file error")
	}
	root, err := html.Parse(f)
	if err != nil {
		log.Fatal("read file error")
	}

	// tools.IterateAllNode(root)
	content, err := tools.QueryBySelector(root, "ul", []string{"channel-list"})
	if err != nil {
		log.Fatal("query error")
	}
	content, err = tools.QueryBySelector(content, "div", []string{"collection-avatar"})
	if err != nil {
		log.Fatal("query error")
	}
	log.Info("find div ", content.Data)
	avatars := tools.QueryAllBySelector(content, "a", []string{})
	fmt.Println("role number: ", len(avatars))

	roleUrl := make(map[string]string)
	for _, avatar := range avatars {
		titleNode, err := tools.QueryBySelector(avatar, "div", []string{"collection-avatar__title"})
		if err != nil {
			log.Fatal("query error")
		}
		name := tools.FetchTextFrom(titleNode)
		url := tools.FetchAttrVal(avatar, "href")
		roleUrl[name] = url
	}
	for k, v := range roleUrl {
		fmt.Println("role: ", k, " url: ", v)
	}
}

func fetchRoleHtml(ctx context.Context) {
	var htmlContent string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(roleUrl),
		chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery),
	); err != nil {
		log.Fatal(err)
	}

	i, err := tools.WriteStrToNewFile("role.html", htmlContent)
	if err != nil {
		log.Fatal("write file error!")
	}
	log.Info("html content length is ", i)
}
