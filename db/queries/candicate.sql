-- name: CreateCandidate :one
INSERT INTO candidate (fullname, title, email, country, company, dateofbirth)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetCandidate :one
SELECT * FROM candidate
WHERE can_id = $1 LIMIT 1;

-- name: ListCandidates :many
SELECT * FROM candidate
ORDER BY ranklocal
LIMIT $1;

-- name: UpdateCandidate :exec
UPDATE candidate
SET fullname = $1, title = $2, email = $3, country = $4, company = $5, dateofbirth = $6
WHERE can_id = $7;

-- name: SignCoach :exec
UPDATE candidate
SET coach_id = $1
WHERE can_id = $2;


