-- name: CreateCandicate :one
INSERT INTO candicate (fullname, title, email, country, company, dateofbirth)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetCandicate :one
SELECT * FROM candicate
WHERE can_id = $1 LIMIT 1;

-- name: ListCandicates :many
SELECT * FROM candicate
ORDER BY ranklocal
LIMIT $1;

-- name: UpdateCandicate :exec
UPDATE candicate
SET fullname = $1, title = $2, email = $3, country = $4, company = $5, dateofbirth = $6
WHERE can_id = $7;

-- name: signCoach :exec
UPDATE candicate
SET coach_id = $1
WHERE can_id = $2;
