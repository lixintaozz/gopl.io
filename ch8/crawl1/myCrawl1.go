package crawl

import (
	"fmt"
	"log"
	"sync"

	"code.byted.org/lixintao/workshop_demo/bookcode/ch5/links"
)

// 两种实现
type CrawlImpl1 struct{}
type CrawlImpl2 struct{}
type CrawlImpl3 struct{}

// 令牌
var tokens = make(chan struct{}, 20)
var wg sync.WaitGroup

// 版本1：使用sync.WaitGroup来保证程序正常结束（存在bug！目前只能处理url大于1个的情况）
func (*CrawlImpl1) Crawl(urls []string) {
	worklist := make(chan []string)

	wg.Add(1)
	go func() {
		worklist <- urls[1:]
	}()

	visited := make(map[string]bool)

	go func() {
		wg.Wait()
		close(worklist)
	}()

	go func() {
		worklist <- crawl1(urls[0])
	}()

	for list := range worklist {
		for _, link := range list {
			if !visited[link] {
				visited[link] = true
				wg.Add(1)
				go func(link string) {
					worklist <- crawl1(link)
				}(link)
			}
		}
	}
}

func crawl1(url string) []string {
	defer wg.Done()
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
