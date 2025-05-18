package main

import (
	"os"

	"code.byted.org/lixintao/workshop_demo/bookcode/ch8/crawl"
)

func main() {
	var myCrawl crawl.MyCrwal
	myCrawl = &crawl.CrawlImpl1{}
	//myCrawl = &crawl.CrawlImpl2{}
	myCrawl.Crawl(os.Args[1:])
}
