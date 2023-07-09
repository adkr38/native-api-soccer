package main

import (
	"soccer-go/src/scrape"
)

func main()  {
    scrape.ScrapeMain()
    scrape.PostScrapingMySqlSetup()
}


