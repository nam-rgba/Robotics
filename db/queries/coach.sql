-- name: CreateCoach :one
INSERT INTO coach (coach_id, fullname, email, country, title, company)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetCoach :one
SELECT * FROM coach
WHERE coach_id = $1 LIMIT 1;

-- name: EditCoach :one
UPDATE coach
SET fullname = $2
WHERE coach_id = $1
RETURNING *;
