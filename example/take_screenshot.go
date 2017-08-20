package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nladuo/go-splash"
)

func main() {
	client, err := splash.NewSplashClient("localhost", "8050")
	if err != nil {
		panic(err)
	}
	response, err := client.Get("https://www.baidu.com/s?wd=github", &splash.Option{Png: true})
	if err != nil {
		panic(err)
	}
	fmt.Println("Title:", response.Title)
	fmt.Println("Url:", response.Url)
	fmt.Println("RequestedUrl:", response.RequestedUrl)
	err = response.SavedPng("./baidu_github.png")
	if err != nil {
		panic(err)
	}
	//select search results by goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response.Html))
	if err != nil {
		panic(err)
	}
	fmt.Println("Results:")
	doc.Find(".c-container h3 a").Each(func(i int, contentSelection *goquery.Selection) {
		fmt.Println(i+1, "-->", contentSelection.Text())
	})
}
