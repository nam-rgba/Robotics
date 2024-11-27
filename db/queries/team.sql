-- name: CreateTeam :one
INSERT INTO team (teamname, coach_id ) VALUES ($1, $2) RETURNING *;

-- name: GetTeam :many
SELECT * FROM team
WHERE coach_id = $1;

-- name: GetTeamById :one
SELECT * FROM team
WHERE team_id = $1;

-- name: GetMaxTeamId :one
SELECT maxteam FROM team WHERE team_id = $1;

-- name: GetNumberOfCandidates :one
SELECT count(can_id) FROM team_candidate WHERE team_id = $1;



