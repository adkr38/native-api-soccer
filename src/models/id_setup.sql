USE soccer;

UPDATE players p
JOIN teams t ON p.team_name = t.team_name
SET p.team_id = t.id;

UPDATE player_overall ps
JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
SET ps.id = p.id,
ps.team_id = p.team_id;

UPDATE player_defensive ps
JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
SET ps.id = p.id,
ps.team_id = p.team_id;

UPDATE player_shooting ps
JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
SET ps.id = p.id,
ps.team_id = p.team_id;

UPDATE player_passing ps
JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
SET ps.id = p.id,
ps.team_id = p.team_id;

UPDATE player_possession ps
JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
SET ps.id = p.id,
ps.team_id = p.team_id;

ALTER TABLE player_overall
DROP COLUMN team_name,
DROP COLUMN player_name;

ALTER TABLE player_defensive
DROP COLUMN team_name,
DROP COLUMN player_name;

ALTER TABLE player_shooting
DROP COLUMN team_name,
DROP COLUMN player_name;

ALTER TABLE player_passing
DROP COLUMN team_name,
DROP COLUMN player_name;

ALTER TABLE player_possession
DROP COLUMN team_name,
DROP COLUMN player_name;

ALTER TABLE players
DROP COLUMN team_name;

UPDATE team_overall tt
JOIN teams t ON tt.team_name = t.team_name
SET tt.id = t.id,
tt.id = t.id;

UPDATE team_defensive tt
JOIN teams t ON tt.team_name = t.team_name
SET tt.id = t.id,
tt.id = t.id;

UPDATE team_shooting tt
JOIN teams t ON tt.team_name = t.team_name
SET tt.id = t.id,
tt.id = t.id;

UPDATE team_possession tt
JOIN teams t ON tt.team_name = t.team_name
SET tt.id = t.id,
tt.id = t.id;


UPDATE team_passing tt
JOIN teams t ON tt.team_name = t.team_name
SET tt.id = t.id,
tt.id = t.id;


ALTER TABLE team_overall
DROP COLUMN team_name;

ALTER TABLE team_defensive
DROP COLUMN team_name;

ALTER TABLE team_shooting
DROP COLUMN team_name;

ALTER TABLE team_passing
DROP COLUMN team_name;

ALTER TABLE team_possession
DROP COLUMN team_name;
