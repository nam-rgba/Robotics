

-- name: GetTeamCandidates :many
SELECT * FROM team_candidate WHERE team_id = $1;

-- name: InviteByEmail :one
INSERT INTO team_candidate (team_id, can_id, invitation_status) VALUES ($1, $2, 'pending') RETURNING *;

-- name: CandidateResponse :exec
UPDATE team_candidate SET invitation_status = $3 WHERE team_id = $1 AND can_id = $2 RETURNING *;

-- name: RemoveCandidate :one
DELETE FROM team_candidate WHERE team_id = $1 AND can_id = $2 RETURNING *;

-- name: GetCandidates :many
SELECT * FROM team_candidate WHERE team_id = $1;




