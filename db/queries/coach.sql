
-- name: GetCoach :one
SELECT * FROM coach
WHERE coach_id = $1 LIMIT 1;

-- name: EditCoach :one
UPDATE coach
SET fullname = $2
WHERE coach_id = $1
RETURNING *;

-- name: RegisterCoach :one
INSERT INTO coach
(email, password) VALUES ($1, $2) RETURNING *;

-- name: GetCoachByEmail :one
SELECT * FROM coach
WHERE email = $1 LIMIT 1;

