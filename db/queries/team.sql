-- name: CreateTeam :one
INSERT INTO team (teamname, coach_id, join_code ) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTeam :many
SELECT * FROM team
WHERE coach_id = $1;

-- name: GetTeamByCode :one
SELECT * FROM team
WHERE join_code = $1;



