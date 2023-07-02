package main

import (
	"soccer-go/src/scrape"
	"soccer-go/src/statenums"
	"soccer-go/src/utils"
	"time"
)

func main()  {
    utils.TimeFunc(scrape.Scrape,statenums.TEAMS,scrape.ConvertTeamHeader)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.PLAYERS,scrape.ConvertPlayerHeader)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.PLAYER_OVERALL,scrape.ConvertPlayerOverallStatHeader)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.PLAYER_PASSING,scrape.ConvertPassingStatHeaderPlayer)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.PLAYER_SHOOTING,scrape.ConvertShootingStatHeaderPlayer)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.PLAYER_POSSESSION,scrape.ConvertPossesionHeaderPlayer)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.PLAYER_DEFENSIVE,scrape.ConvertDefensiveStatHeaderPlayer)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.TEAM_OVERALL,scrape.ConvertTeamOverallStatHeader)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.TEAM_PASSING,scrape.ConvertPassingStatHeaderTeam)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.TEAM_SHOOTING,scrape.ConvertShootingStatHeaderTeam)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.TEAM_POSSESSION,scrape.ConvertPossesionHeaderTeam)
    time.Sleep(time.Duration(2))
    utils.TimeFunc(scrape.Scrape,statenums.TEAM_DEFENSIVE,scrape.ConvertDefensiveStatHeaderTeam)
}


