DROP SCHEMA soccer;
CREATE SCHEMA soccer;
USE soccer;

CREATE TABLE competitions (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  competition_name VARCHAR(255)
);

CREATE TABLE teams (
  id INT AUTO_INCREMENT PRIMARY KEY,
  team_name VARCHAR(255),
  country VARCHAR(255)
);

CREATE TABLE competition_participation (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  competition_id INT(11),
  team_id INT NULL,
  FOREIGN KEY(team_id) REFERENCES teams(id)


);

CREATE TABLE players (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  team_id INT NULL,
  team_name VARCHAR(255),
  FOREIGN KEY(team_id) REFERENCES teams(id)
);


CREATE TABLE player_overall(
  id INT NULL,
  player_name VARCHAR(255) NULL,
  team_name VARCHAR(255) NULL,
  team_id INT NULL,
  nation VARCHAR(255) NULL,
  position VARCHAR(255) NULL,
  age INT NULL,
  birth_year INT NULL,
  matches_played INT NULL,
  minutes_played INT NULL,
  goals INT NULL,
  assists INT NULL,
  non_penalty_goals INT NULL,
  penalty_goals INT NULL,
  penalty_attempts INT NULL,
  yellow_cards INT NULL,
  red_cards INT NULL,
  FOREIGN KEY(id) REFERENCES players(id),
  FOREIGN KEY(team_id) REFERENCES teams(id)
);

CREATE TABLE team_overall(
  id INT NULL,
  team_name VARCHAR(255) NULL,
  goals INT NULL,
  assists INT NULL,
  non_penalty_goals INT NULL,
  penalty_goals INT NULL,
  penalty_attempts INT NULL,
  yellow_cards INT NULL,
  red_cards INT NULL,
  FOREIGN KEY(id) REFERENCES teams(id)
);

CREATE TABLE player_goalkeeping(
  id INT NULL,
  player_name VARCHAR(255) NULL,
  team_id INT NULL,
  team_name VARCHAR(255) NULL,
  goals_against INT NULL,
  saves INT NULL,
  save_percentage DECIMAL(16,2) NULL,
  clean_sheets INT NULL,
  clean_sheet_percentage DECIMAL(16,2) NULL,
  penalty_kicks_received INT NULL,
  penalty_kicks_lost INT NULL,
  penalty_kicks_saved INT NULL,
  FOREIGN KEY(id) REFERENCES players(id),
  FOREIGN KEY(team_id) REFERENCES teams(id)
);

CREATE TABLE team_goalkeeping(
  id INT NULL,
  team_name VARCHAR(255) NULL,
  goals_against INT NULL,
  saves INT NULL,
  save_percentage DECIMAL(16,2) NULL,
  clean_sheets INT NULL,
  clean_sheet_percentage DECIMAL(16,2) NULL,
  penalty_kicks_received INT NULL,
  penalty_kicks_lost INT NULL,
  penalty_kicks_saved INT NULL,
  FOREIGN KEY(id) REFERENCES teams(id)
);

CREATE TABLE player_shooting(
  id INT NULL,
  player_name VARCHAR(255) NULL,
  team_id INT NULL,
  team_name VARCHAR(255) NULL,
  shots INT NULL,
  shots_on_target INT NULL,
  goals_per_shot DECIMAL(16,2) NULL,
  goals_per_shot_on_target DECIMAL(16,2) NULL,
  average_shot_distance DECIMAL(16,2)  NULL,
  FOREIGN KEY(id) REFERENCES players(id),
  FOREIGN KEY(team_id) REFERENCES teams(id)
);

CREATE TABLE team_shooting(
  id INT NULL,
  team_name VARCHAR(255) NULL,
  shots INT NULL,
  shots_on_target INT NULL,
  goals_per_shot DECIMAL(16,2) NULL,
  goals_per_shot_on_target DECIMAL(16,2) NULL,
  average_shot_distance DECIMAL(16,2)  NULL,
  FOREIGN KEY(id) REFERENCES teams(id)
);

CREATE TABLE player_passing (
  id INT NULL,
  player_name VARCHAR(255) NULL,
  team_id INT NULL,
  team_name VARCHAR(255) NULL,
  total_completed_passes INT NULL,
  total_attempted_passes INT NULL,
  total_distance_passed DECIMAL(16,2) NULL,
  short_passes_attempted INT NULL,
  short_passes_completed INT NULL,
  medium_passes_attempted INT NULL,
  medium_passes_completed INT NULL,
  long_passes_attempted INT NULL,
  long_passes_completed INT NULL,
  key_passes INT NULL,
  passes_into_final_third INT NULL,
  FOREIGN KEY(id) REFERENCES players(id),
  FOREIGN KEY(team_id) REFERENCES teams(id)
);

CREATE TABLE team_passing(
  id INT NULL,
  team_name VARCHAR(255) NULL,
  total_completed_passes INT NULL,
  total_attempted_passes INT NULL,
  total_distance_passed DECIMAL(16,2) NULL,
  short_passes_attempted INT NULL,
  short_passes_completed INT NULL,
  medium_passes_attempted INT NULL,
  medium_passes_completed INT NULL,
  long_passes_attempted INT NULL,
  long_passes_completed INT NULL,
  key_passes INT NULL,
  passes_into_final_third INT NULL,
  FOREIGN KEY(id) REFERENCES teams(id)
);

CREATE TABLE player_defensive(
  id INT NULL,
  player_name VARCHAR(255) NULL,
  team_id INT NULL,
  team_name VARCHAR(255) NULL,
  tackles INT NULL,
  tackles_won INT NULL,
  tackles_d3 INT NULL,
  tackles_m3 INT NULL,
  tackles_a3 INT NULL,
  blocks INT NULL,
  shots_blocked INT NULL,
  passes_blocked INT NULL,
  interceptions INT NULL,
  clearances INT NULL,
  errors INT NULL,
  FOREIGN KEY(id) REFERENCES players(id),
  FOREIGN KEY(team_id) REFERENCES teams(id)
);

CREATE TABLE team_defensive(
  id INT NULL,
  team_name VARCHAR(255) NULL,
  tackles INT NULL,
  tackles_won INT NULL,
  tackles_d3 INT NULL,
  tackles_m3 INT NULL,
  tackles_a3 INT NULL,
  blocks INT NULL,
  shots_blocked INT NULL,
  passes_blocked INT NULL,
  interceptions INT NULL,
  clearances INT NULL,
  errors INT NULL,
  FOREIGN KEY(id) REFERENCES teams(id)
);

CREATE TABLE player_possession(
  id INT NULL,
  player_name VARCHAR(255) NULL,
  team_id INT NULL,
  team_name VARCHAR(255) NULL,
  touches INT NULL,
  touches_self_penalty INT NULL,
  touches_d3 INT NULL,
  touches_m3 INT NULL,
  touches_a3 INT NULL,
  touches_opponent_penalty INT NULL,
  dribbles INT NULL,
  dribbles_succeeded INT NULL,
  carries_final3 INT NULL,
  carries_opponent_penalty INT NULL,
  dispossesions INT NULL,
  miscontrols INT NULL,
  FOREIGN KEY(id) REFERENCES players(id),
  FOREIGN KEY(team_id) REFERENCES teams(id)
);

CREATE TABLE team_possession(
  id INT NULL,
  team_name VARCHAR(255) NULL,
  touches INT NULL,
  touches_self_penalty INT NULL,
  touches_d3 INT NULL,
  touches_m3 INT NULL,
  touches_a3 INT NULL,
  touches_opponent_penalty INT NULL,
  dribbles INT NULL,
  dribbles_succeeded INT NULL,
  carries_final3 INT NULL,
  carries_opponent_penalty INT NULL,
  dispossesions INT NULL,
  miscontrols INT NULL,
  FOREIGN KEY(id) REFERENCES teams(id)
);
