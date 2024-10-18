-- name: CreateCompetition :one
INSERT INTO competition 
(name, decription ) 
VALUES ($1, $2) RETURNING *;