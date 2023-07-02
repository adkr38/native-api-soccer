package main

import (
	"soccer-go/src/scrape"
	"soccer-go/src/statenums"
	"soccer-go/src/utils"
)

func main()  {
    utils.TimeFunc(scrape.Scrape,statenums.TEAMS,scrape.ConvertTeamHeader)
    utils.TimeFunc(scrape.Scrape,statenums.PLAYERS,scrape.ConvertPlayerHeader)
    utils.TimeFunc(scrape.Scrape,statenums.TEAM_OVERALL,scrape.ConvertTeamOverallStatHeader)
    utils.TimeFunc(scrape.Scrape,statenums.PLAYER_OVERALL,scrape.ConvertPlayerOverallStatHeader)
}


