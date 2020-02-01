package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher2 interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

type VisitedUrls struct {
	sync.RWMutex
	urls map[string]bool
}

var visitedUrls = VisitedUrls{
	urls: make(map[string]bool, 1000),
}

func (visitedUrls *VisitedUrls) put(key string, value bool) {
	visitedUrls.Lock()
	visitedUrls.urls[key] = value
	visitedUrls.Unlock()
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl2(url string, fetcher Fetcher2) {
	if _, ok := visitedUrls.urls[url]; ok {
		/* 已经访问过不再访问 */
		fmt.Printf("%s is visited\n", url)
		return
	}
	body, urls, err := fetcher.Fetch(url)
	visitedUrls.put(url, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl2(u, fetcher)
	}
	return
}

func main() {
	go Crawl2("http://golang.org/", fetcher2)
	time.Sleep(2 * time.Second)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher2 map[string]*fakeResult2

type fakeResult2 struct {
	body string
	urls []string
}

func (f fakeFetcher2) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher2 = fakeFetcher2{
	"http://golang.org/": &fakeResult2{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult2{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult2{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult2{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
