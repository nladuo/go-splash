package main

import (
	// "fmt"
	// "github.com/PuerkitoBio/goquery"
	"github.com/nladuo/go-splash"
	// "strings"
)

func main() {
	// js_script := "function(){document.getElementById('kw').setAttribute('value', 'github');document.getElementById('su').click();}"
	request, _ := splash.NewSplashRequest("localhost", "8050")
	request.Get("http://baidu.com", nil)
}
