package main

import (
	"cheese/entity"
	mouseparser "cheese/mouse_parser"
	"cheese/tools"
	"context"

	"github.com/chromedp/chromedp"
	"github.com/labstack/gommon/log"
)

// 角色信息
const roleUrl = "https://bbs.mihoyo.com/ys/obc/channel/map/189/25"

func main() {
	// const basePath = "./roledetail/"
	// ru, err := ori.DbInst.SelectAllRoleUrl()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rChan := make(chan entity.Role, len(ru))
	// for _, e := range ru {
	// 	go ParseRolePage(basePath+e.RoleName+".html", rChan)
	// }
	// for e := range rChan {
	// 	fmt.Println(e)
	// }
	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()
	// fetchRoleHtml(ctx)

	// roles, err := ori.DbInst.SelectAllRoleUrl()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, role := range roles {
	// 	fetchRoleDetailHtml(role.RoleName, role.RoleUrl, ctx)
	// 	time.Sleep(time.Second * 3)
	// }

	// f, err := os.Open("role.html")
	// if err != nil {
	// 	log.Fatal("open file error")
	// }
	// root, err := html.Parse(f)
	// if err != nil {
	// 	log.Fatal("read file error")
	// }

	// tools.IterateAllNode(root)
	// content, err := tools.QueryBySelector(root, "ul", []string{"channel-list"})
	// if err != nil {
	// 	log.Fatal("query error")
	// }
	// content, err = tools.QueryBySelector(content, "div", []string{"collection-avatar"})
	// if err != nil {
	// 	log.Fatal("query error")
	// }
	// log.Info("find div ", content.Data)
	// avatars := tools.QueryAllBySelector(content, "a", []string{})
	// fmt.Println("role number: ", len(avatars))

	// roleUrl := make(map[string]string)
	// for _, avatar := range avatars {
	// 	titleNode, err := tools.QueryBySelector(avatar, "div", []string{"collection-avatar__title"})
	// 	if err != nil {
	// 		log.Fatal("query error")
	// 	}
	// 	name := tools.FetchTextFrom(titleNode)
	// 	url := tools.FetchAttrVal(avatar, "href")
	// 	roleUrl[name] = url
	// }
	// for k, v := range roleUrl {
	// 	_, err = ori.DbInst.InsertRoleUrl(k, v)
	// 	if err != nil {
	// 		log.Errorf("insert k{%s} v{%s} error{%s}\n", k, v, err)
	// 	}
	// }
}

func ParseRolePage(filename string, ch chan<- entity.Role) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Error("parse ", filename, " error occured! ", err)
	// 		ch <- entity.Role{}
	// 		close(ch)
	// 	}
	// }()
	exData := make(map[string]string)

	p, err := mouseparser.NewParser(filename)
	if err != nil {
		log.Error("init parser error")
		panic(err)
	}

	attrs := make(map[string]string)
	attrs["class"] = "obc-tmp-character__box"
	cr, err := p.ReadTag("div", attrs)
	if err != nil {
		log.Error("read role div error")
		panic(err)
	}

	attrs["class"] = "obc-tmp-character__box--title"
	c, err := p.ReadTag("p", attrs)
	if err != nil {
		log.Error("read name p error")
		panic(err)
	}
	str, err := p.ReadTextFrom(c)
	if err != nil {
		log.Error("read name text error")
		panic(err)
	}
	exData["姓名"] = str

	p.SetHead(cr)
	attrs["class"] = "obc-tmp-character__property"
	_, err = p.ReadTag("div", attrs)
	if err != nil {
		log.Error("read role info div error")
		panic(err)
	}

	attrs["class"] = "obc-tmp-character__key"
	ns, err := p.ReadTags("div", attrs)
	if err != nil {
		log.Error("read role attr div error")
		panic(err)
	}
	keys := make([]string, len(ns))
	for _, n := range ns {
		t, _ := p.ReadTextFrom(n)
		keys = append(keys, t)
	}
	// fmt.Printf("%v", keys)

	attrs["class"] = "obc-tmp-character__value"
	ns, err = p.ReadTags("div", attrs)
	if err != nil {
		log.Error("read role attr value div error")
		panic(err)
	}
	values := make([]string, len(ns))
	for _, n := range ns {
		t, _ := p.ReadTextFrom(n)
		values = append(values, t)
	}

	for i, e := range keys {
		exData[e] = values[i]
	}

	// fmt.Printf("%v", exData)
	ch <- entity.Role{
		Name:    exData["姓名"],
		Birth:   entity.NewRoleBirth(exData["生日"]),
		From:    exData["所属"],
		Feature: exData["定位"],
		Weapon:  entity.FromWeapon(exData["武器类型"]),
		Destiny: exData["命之座"],
		Dub:     exData["称号"],
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

func fetchRoleDetailHtml(fName string, url string, ctx context.Context) {
	const baseUrl = "https://bbs.mihoyo.com"
	var fetchUrl string = baseUrl + url
	var htmlContent string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(fetchUrl),
		chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery),
	); err != nil {
		log.Fatal(err)
	}

	i, err := tools.WriteStrToNewFile("roledetail/"+fName+".html", htmlContent)
	if err != nil {
		log.Fatal("write file error!")
	}
	log.Info("html content length is ", i)
}
