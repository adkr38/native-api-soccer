DROP SCHEMA soccer;
CREATE SCHEMA soccer;
USE soccer;

CREATE TABLE competitions (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  competition_name VARCHAR(255)
);

CREATE TABLE teams (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  team_name VARCHAR(255),
  country VARCHAR(255)
);

CREATE TABLE competition_participation (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  competition_id INT(11),
  team_id INT(11),
  FOREIGN KEY (competition_id) REFERENCES competitions (id),
  FOREIGN KEY (team_id) REFERENCES teams (id)
);

CREATE TABLE players (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  team_id INT,
  FOREIGN KEY (team_id) REFERENCES teams (id)
);


CREATE TABLE player_overall(
  player_id INT,
  nation VARCHAR(255),
  position VARCHAR(255),
  age INT,
  birth_year INT,
  matches_played INT,
  minutes_played INT,
  goals INT,
  assists INT,
  non_penalty_goals INT,
  penalty_goals INT,
  penalty_attempts INT,
  yellow_cards INT,
  red_cards INT,
  FOREIGN KEY (player_id) REFERENCES players(id)
);

CREATE TABLE team_overall(
  team_id INT,
  goals INT, 
  assists INT,
  non_penalty_goals INT,
  penalty_goals INT,
  penalty_attempts INT,
  yellow_cards INT,
  red_cards INT,
  FOREIGN KEY (team_id) REFERENCES teams(id)
);

CREATE TABLE player_goalkeeping(
  player_id INT,
  goals_against INT,
  saves INT,
  save_percentage DECIMAL(5,2),
  clean_sheets INT,
  clean_sheet_percentage DECIMAL(5,2),
  penalty_kicks_received INT,
  penalty_kicks_lost INT,
  penalty_kicks_saved INT,
  FOREIGN KEY (player_id) REFERENCES players(id)
);

CREATE TABLE team_goalkeeping(
  team_id INT,
  goals_against INT,
  saves INT,
  save_percentage DECIMAL(5,2),
  clean_sheets INT,
  clean_sheet_percentage DECIMAL(5,2),
  penalty_kicks_received INT,
  penalty_kicks_lost INT,
  penalty_kicks_saved INT,
  FOREIGN KEY (team_id) REFERENCES teams(id)
);

CREATE TABLE player_shooting(
  player_id INT,
  shots INT,
  shots_on_target INT,
  goals_per_shot DECIMAL(5,2),
  goals_per_shot_on_target DECIMAL(5,2),
  average_shot_distance DECIMAL(6,2),
  FOREIGN KEY (player_id) REFERENCES players(id)
);

CREATE TABLE team_shooting(
  team_id INT,
  shots INT,
  shots_on_target INT,
  goals_per_shot DECIMAL(5,2),
  goals_per_shot_on_target DECIMAL(5,2),
  average_shot_distance DECIMAL(6,2),
  FOREIGN KEY (team_id) REFERENCES teams(id)
);

CREATE TABLE player_passing (
  player_id INT,
  total_completed_passes INT,
  total_attempted_passes INT,
  total_distance_passed DECIMAL(6,2),
  short_passes_attempted INT,
  short_passes_completed INT,
  medium_passes_attempted INT,
  medium_passes_completed INT,
  long_passes_attempted INT,
  long_passes_completed INT,
  key_passes INT,
  passes_into_final_third INT,
  FOREIGN KEY (player_id) REFERENCES players(id)
);

CREATE TABLE team_passing(
  team_id INT,
  total_completed_passes INT,
  total_attempted_passes INT,
  total_distance_passed DECIMAL(6,2),
  short_passes_attempted INT,
  short_passes_completed INT,
  medium_passes_attempted INT,
  medium_passes_completed INT,
  long_passes_attempted INT,
  long_passes_completed INT,
  key_passes INT,
  passes_into_final_third INT,
  FOREIGN KEY (team_id) REFERENCES teams(id)
);

CREATE TABLE player_defense(
  player_id INT,
  tackles INT,
  tackles_won INT,
  tackles_d3 INT,
  tackles_m3 INT,
  tackles_a3 INT,
  blocks INT,
  shots_blocked INT,
  passes_blocked INT,
  interceptions INT,
  clearances INT,
  errors INT,
  FOREIGN KEY (player_id) REFERENCES players(id)
);

CREATE TABLE team_defense(
  team_id INT,
  tackles INT,
  tackles_won INT,
  tackles_d3 INT,
  tackles_m3 INT,
  tackles_a3 INT,
  blocks INT,
  shots_blocked INT,
  passes_blocked INT,
  interceptions INT,
  clearances INT,
  errors INT,
  FOREIGN KEY (team_id) REFERENCES teams(id)
);

CREATE TABLE player_possession(
  player_id INT,
  touches INT,
  touches_self_penalty INT,
  touches_d3 INT,
  touches_m3 INT,
  touches_a3 INT,
  touches_opponent_penalty INT,
  dribbles INT,
  dribbles_succeeded INT,
  carries_final3 INT,
  carries_opponent_penalty INT,
  dispossesions INT,
  miscontrols INT,
  FOREIGN KEY (player_id) REFERENCES players(id)
);

CREATE TABLE team_possession(
  team_id INT,
  touches INT,
  touches_self_penalty INT,
  touches_d3 INT,
  touches_m3 INT,
  touches_a3 INT,
  touches_opponent_penalty INT,
  dribbles INT,
  dribbles_succeeded INT,
  carries_final3 INT,
  carries_opponent_penalty INT,
  dispossesions INT,
  miscontrols INT,
  FOREIGN KEY (team_id) REFERENCES teams(id)
);
