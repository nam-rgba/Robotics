

-- name: GetCandidate :one
SELECT * FROM candidate
WHERE can_id = $1 LIMIT 1;

-- name: ListCandidates :many
SELECT * FROM candidate
ORDER BY ranklocal
LIMIT $1;

-- name: UpdateCandidate :exec
UPDATE candidate
SET fullname = $1, title = $2, country = $3, company = $4, dateofbirth = $5
WHERE can_id = $6;

-- name: SignCoach :exec
UPDATE candidate
SET coach_id = $1
WHERE can_id = $2;


-- name: RegisterCandidate :one
INSERT INTO candidate
(email, password) VALUES ($1, $2) RETURNING *;

-- name: GetCandidateByEmail :one
SELECT * FROM candidate
WHERE email = $1 LIMIT 1;

-- name: DeleteCandidate :exec
DELETE FROM candidate
WHERE can_id = $1;