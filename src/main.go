package main

import (
	"soccer-go/src/scrape"
	"soccer-go/src/statenums"
	"soccer-go/src/utils"
)

func main()  {
  utils.TimeFunc(scrape.Scrape,statenums.PASSING_PLAYER)
}


