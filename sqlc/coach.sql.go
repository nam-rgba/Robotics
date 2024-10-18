// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: coach.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createCoach = `-- name: CreateCoach :one
INSERT INTO coach ( fullname, email, country, title, company)
VALUES ($1, $2, $3, $4, $5)
RETURNING coach_id, fullname, email, country, title, company, numberofcandidate
`

type CreateCoachParams struct {
	Fullname sql.NullString `json:"fullname"`
	Email    sql.NullString `json:"email"`
	Country  sql.NullString `json:"country"`
	Title    sql.NullString `json:"title"`
	Company  sql.NullString `json:"company"`
}

func (q *Queries) CreateCoach(ctx context.Context, arg CreateCoachParams) (Coach, error) {
	row := q.db.QueryRowContext(ctx, createCoach,
		arg.Fullname,
		arg.Email,
		arg.Country,
		arg.Title,
		arg.Company,
	)
	var i Coach
	err := row.Scan(
		&i.CoachID,
		&i.Fullname,
		&i.Email,
		&i.Country,
		&i.Title,
		&i.Company,
		&i.Numberofcandidate,
	)
	return i, err
}

const editCoach = `-- name: EditCoach :one
UPDATE coach
SET fullname = $2
WHERE coach_id = $1
RETURNING coach_id, fullname, email, country, title, company, numberofcandidate
`

type EditCoachParams struct {
	CoachID  int64          `json:"coach_id"`
	Fullname sql.NullString `json:"fullname"`
}

func (q *Queries) EditCoach(ctx context.Context, arg EditCoachParams) (Coach, error) {
	row := q.db.QueryRowContext(ctx, editCoach, arg.CoachID, arg.Fullname)
	var i Coach
	err := row.Scan(
		&i.CoachID,
		&i.Fullname,
		&i.Email,
		&i.Country,
		&i.Title,
		&i.Company,
		&i.Numberofcandidate,
	)
	return i, err
}

const getCoach = `-- name: GetCoach :one
SELECT coach_id, fullname, email, country, title, company, numberofcandidate FROM coach
WHERE coach_id = $1 LIMIT 1
`

func (q *Queries) GetCoach(ctx context.Context, coachID int64) (Coach, error) {
	row := q.db.QueryRowContext(ctx, getCoach, coachID)
	var i Coach
	err := row.Scan(
		&i.CoachID,
		&i.Fullname,
		&i.Email,
		&i.Country,
		&i.Title,
		&i.Company,
		&i.Numberofcandidate,
	)
	return i, err
}