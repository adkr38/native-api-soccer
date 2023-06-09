package scrape

var teamConversions = map[string]string{
  "squad":"team_name",
  "comp":"country",
}

func ConvertTeamHeader(str string) string{
  return teamConversions[str]
}


var playerConversions= map[string]string{
  "player":"name",
  "squad":"team_name",
}


func ConvertPlayerHeader(str string) string{
  return playerConversions[str]
}

var teamOverallConversions= map[string]string{
  "squad":"team_name",
  "nation":"nation",
  "pos":"position",
  "age":"age",
  "born":"birth_year",
  "mp":"matches_played",
  "min":"minutes_played",
  "gls":"goals",
  "ast":"assists",
  "g-pk":"non_penalty_goals",
  "pk":"penalty_goals",
  "pkatt":"penalty_attempts",
  "crdy":"yellow_cards",
  "crdr":"red_cards",
}

var playerOverallConversions= map[string]string{
  "player":"player_name",
  "squad":"team_name",
  "comp":"competition",
  "nation":"nation",
  "pos":"position",
  "age":"age",
  "born":"birth_year",
  "mp":"matches_played",
  "min":"minutes_played",
  "gls":"goals",
  "ast":"assists",
  "g-pk":"non_penalty_goals",
  "pk":"penalty_goals",
  "pkatt":"penalty_attempts",
  "crdy":"yellow_cards",
  "crdr":"red_cards",
}


func ConvertPlayerOverallStatHeader(str string) string{
  return playerOverallConversions[str]
}

func ConvertTeamOverallStatHeader(str string) string{
  return teamOverallConversions[str]
}

var goalkeepingConversions= map[string]string{
  "player":"player_name",
  "squad":"team_name",
  "nation":"nation",
  "pos":"position",
  "age":"age",
  "born":"birth_year",
  "mp":"matches_played",
  "min":"minutes_played",
  "ga":"goals_against",
  "saves":"minutes",
  "gls":"goals",
  "ast":"assists",
  "g-pk":"non_penalty_goals",
  "pk":"penalty_goals",
  "pkatt":"penalty_attempts",
  "crdy":"yellow_cards",
  "crdr":"red_cards",
}


func ConvertGoalkeepingStatHeader(str string) string{
  return goalkeepingConversions[str]
}


var shootingConversionsPlayer= map[string]string{
  "player":"player_name",
  "squad":"team_name",
  "sh":"shots",
  "sot":"shots_on_target",
  "g/sh":"goals_per_shot",
  "g/sot":"goals_per_shot_on_target",
  "dist" :"average_shot_distance",
}

var shootingConversionsTeam= map[string]string{
  "player":"player_name",
  "squad":"team_name",
  "sh":"shots",
  "sot":"shots_on_target",
  "g/sh":"goals_per_shot",
  "g/sot":"goals_per_shot_on_target",
  "dist" :"average_shot_distance",
}

func ConvertShootingStatHeaderPlayer(str string) string{
  return shootingConversionsPlayer[str]
}

func ConvertShootingStatHeaderTeam(str string) string{
  return shootingConversionsTeam[str]
}

// Have to add an incrementor on duplicate names on thead > th
var passingConversionsPlayer = map[string]string{
  "player":"player_name",
  "squad":"team_name",
  "cmp":"total_completed_passes",
  "att":"total_attempted_passes",
  "totdist":"total_distance_passed",
  "cmp1":"short_passes_completed",
  "att1":"short_passes_attempted",
  "cmp2":"medium_passes_completed",
  "att2":"medium_passes_attempted",
  "cmp3":"long_passes_completed",
  "att3":"long_passes_attempted",
  "kp":"key_passes",
  "1/3":"passes_into_final_third",
}

var passingConversionsTeam= map[string]string{
  "player":"player_name",
  "squad":"team_name",
  "cmp":"total_completed_passes",
  "att":"total_attempted_passes",
  "totdist":"total_distance_passed",
  "cmp1":"short_passes_completed",
  "att1":"short_passes_attempted",
  "cmp2":"medium_passes_completed",
  "att2":"medium_passes_attempted",
  "cmp3":"long_passes_completed",
  "att3":"long_passes_attempted",
  "kp":"key_passes",
  "1/3":"passes_into_final_third",
}

func ConvertPassingStatHeaderPlayer(str string) string{
  return passingConversionsPlayer[str]
}

func ConvertPassingStatHeaderTeam(str string) string{
  return passingConversionsTeam[str]
}

var defensiveConversionsPlayer= map[string]string{
    "player":"player_name",
  "squad":"team_name",
  "tkl":"tackles",
  "def 3rd":"tackles_d3",
  "mid 3rd":"tackles_m3",
  "att 3rd":"tackles_a3",
  "blocks":"blocks",
  "sh":"shots_blocked",
  "pass":"passes_blocked",
  "int":"interceptions",
  "clr":"clearances",
  "err":"errors",
}


var defensiveConversionsTeam= map[string]string{
    "squad":"team_name",
  "tkl":"tackles",
  "def 3rd":"tackles_d3",
  "mid 3rd":"tackles_m3",
  "att 3rd":"tackles_a3",
  "blocks":"blocks",
  "sh":"shots_blocked",
  "pass":"passes_blocked",
  "int":"interceptions",
  "clr":"clearances",
  "err":"errors",
}

func ConvertDefensiveStatHeaderPlayer(str string) string{
  return defensiveConversionsPlayer[str]
}

func ConvertDefensiveStatHeaderTeam(str string) string{
  return defensiveConversionsTeam[str]
}

var possesionConversionsPlayer= map[string]string{
    "player":"player_name",
  "squad":"team_name",
  "touches":"touches",
  "def pen":"touches_self_penalty",
  "def 3rd":"touches_d3",
  "mid 3rd":"touches_m3",
  "att 3rd":"touches_a3",
  "att pen":"touches_opponent_penalty",
  "att":"dribbles",
  "succ":"dribbles_succeeded",
  "1/3":"carries_final3",
  "cpa":"carries_opponent_penalty",
  "mis":"miscontrols",
  "dis":"dispossesions",
}

var possesionConversionsTeam= map[string]string{
    "squad":"team_name",
  "touches":"touches",
  "def pen":"touches_self_penalty",
  "def 3rd":"touches_d3",
  "mid 3rd":"touches_m3",
  "att 3rd":"touches_a3",
  "att pen":"touches_opponent_penalty",
  "att":"dribbles",
  "succ":"dribbles_succeeded",
  "1/3":"carries_final3",
  "cpa":"carries_opponent_penalty",
  "mis":"miscontrols",
  "dis":"dispossesions",
}

func ConvertPossesionHeaderPlayer(str string) string{
  return possesionConversionsPlayer[str]
}

func ConvertPossesionHeaderTeam(str string) string{
  return possesionConversionsTeam[str]
}

