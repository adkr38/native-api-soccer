package statenums

type StatEnum string

const(
  TEAM_OVERALL StatEnum = "https://fbref.com/en/comps/Big5/stats/squads/Big-5-European-Leagues-Stats"
  TEAM_GOALKEEPING StatEnum = "https://fbref.com/en/comps/Big5/keepers/squads/Big-5-European-Leagues-Stats"
  TEAM_SHOOTING StatEnum = "https://fbref.com/en/comps/Big5/shooting/squads/Big-5-European-Leagues-Stats"
  TEAM_PASSING StatEnum = "https://fbref.com/en/comps/Big5/passing/squads/Big-5-European-Leagues-Stats"
  TEAM_DEFENSIVE StatEnum = "https://fbref.com/en/comps/Big5/defense/squads/Big-5-European-Leagues-Stats"
  TEAM_POSSESSION StatEnum = "https://fbref.com/en/comps/Big5/possession/squads/Big-5-European-Leagues-Stats"
  PLAYER_OVERALL StatEnum = "https://fbref.com/en/comps/Big5/stats/players/Big-5-European-Leagues-Stats"
  PLAYER_GOALKEEPING StatEnum = "https://fbref.com/en/comps/Big5/keepers/players/Big-5-European-Leagues-Stats"
  PLAYER_SHOOTING StatEnum = "https://fbref.com/en/comps/Big5/shooting/players/Big-5-European-Leagues-Stats"
  PLAYER_PASSING StatEnum = "https://fbref.com/en/comps/Big5/passing/players/Big-5-European-Leagues-Stats"
  PLAYER_DEFENSIVE StatEnum = "https://fbref.com/en/comps/Big5/defense/players/Big-5-European-Leagues-Stats"
  PLAYER_POSSESSION StatEnum = "https://fbref.com/en/comps/Big5/possession/players/Big-5-European-Leagues-Stats"
  COMPETITIONS StatEnum =  "COMPETITIONS"
  TEAMS StatEnum =  "TEAMS"
  COMPETITION_PARTICIPATION StatEnum =  "COMPETITION_PARTICIPATIONç"
  PLAYERS StatEnum = "PLAYERS"
  )

var statEnumMap =  map[StatEnum]string{
    TEAM_OVERALL:          "team_overall",
    TEAM_GOALKEEPING:      "team_goalkeeping",
    TEAM_SHOOTING:         "team_shooting",
    TEAM_PASSING:          "team_passing",
    TEAM_DEFENSIVE:        "team_defensive",
    TEAM_POSSESSION:       "team_possession",
    PLAYER_OVERALL:        "player_overall",
    PLAYER_GOALKEEPING:    "player_goalkeeping",
    PLAYER_SHOOTING:       "player_shooting",
    PLAYER_PASSING:        "player_passing",
    PLAYER_DEFENSIVE:      "player_defensive",
    PLAYER_POSSESSION:     "player_possession",
    COMPETITIONS:          "competitions",
    TEAMS:                 "teams",
    COMPETITION_PARTICIPATION: "competition_participation",
    PLAYERS:               "players",

}
func GetEnumName(s StatEnum) string {
  return  statEnumMap[s]
}

// func GetEnumName(s StatEnum) string{
//   fmt.Println(fmt.Sprintf("%s",s))
//   return fmt.Sprintf("%s",s)
//    // reflectVal := reflect.ValueOf(s)
//    // return reflect.TypeOf(s).Field(int(reflectVal.Int())).Name
//  }
//
// // func GetEnumName(s StatEnum) string{
// //   return string(s)
// //
// // }
