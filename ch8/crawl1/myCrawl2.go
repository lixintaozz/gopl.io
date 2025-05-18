package crawl

import (
	"fmt"
	"log"

	"code.byted.org/lixintao/workshop_demo/bookcode/ch5/links"
)

// 版本2：使用n来统计worklist中的待爬取url
func (*CrawlImpl2) Crawl(urls []string) {
	worklist := make(chan []string)

	go func() {
		worklist <- urls
	}()

	visited := make(map[string]bool)
	for n := 1; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !visited[link] {
				visited[link] = true
				n++
				go func(link string) {
					worklist <- crawl2(link)
				}(link)
			}
		}
	}
}

func crawl2(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
