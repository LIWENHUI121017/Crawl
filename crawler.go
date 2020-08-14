package main

import (
	"fmt"
	"github.com/jackdanger/collectlinks"
	"net/http"
)

//go常见的解析器xpath、jquery、
//正则都有，直接搜索即可
//直接用别人写好的轮子collectlinks，
//可以提取网页中所有的链接，下载方法go get -u github.com/jackdanger/collectlinks
func main() {
	url := "http://www.baidu.com"
	ch := make(chan string)
	go func() {
		ch <- url
	}()
	for uri := range ch {
		downLoad(uri, ch)
	}

}

//封装下载方法
func downLoad(url string, ch chan string) {
	client := &http.Client{}
	req, _ := http.NewRequest("Get", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}

	//函数结束后关闭相关链接
	defer res.Body.Close()

	links := collectlinks.All(res.Body)

	for _, link := range links {
		fmt.Println("parse url", link)
		go func() {
			ch <- link
		}()
	}
}
