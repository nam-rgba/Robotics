// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: candicate.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const createCandicate = `-- name: CreateCandicate :one
INSERT INTO candicate (fullname, title, email, country, company, dateofbirth)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING can_id, fullname, title, email, country, ranklocal, rankworld, company, dateofbirth, coach_id
`

type CreateCandicateParams struct {
	Fullname    sql.NullString `json:"fullname"`
	Title       sql.NullString `json:"title"`
	Email       sql.NullString `json:"email"`
	Country     sql.NullString `json:"country"`
	Company     sql.NullString `json:"company"`
	Dateofbirth time.Time      `json:"dateofbirth"`
}

func (q *Queries) CreateCandicate(ctx context.Context, arg CreateCandicateParams) (Candicate, error) {
	row := q.db.QueryRowContext(ctx, createCandicate,
		arg.Fullname,
		arg.Title,
		arg.Email,
		arg.Country,
		arg.Company,
		arg.Dateofbirth,
	)
	var i Candicate
	err := row.Scan(
		&i.CanID,
		&i.Fullname,
		&i.Title,
		&i.Email,
		&i.Country,
		&i.Ranklocal,
		&i.Rankworld,
		&i.Company,
		&i.Dateofbirth,
		&i.CoachID,
	)
	return i, err
}

const getCandicate = `-- name: GetCandicate :one
SELECT can_id, fullname, title, email, country, ranklocal, rankworld, company, dateofbirth, coach_id FROM candicate
WHERE can_id = $1 LIMIT 1
`

func (q *Queries) GetCandicate(ctx context.Context, canID int64) (Candicate, error) {
	row := q.db.QueryRowContext(ctx, getCandicate, canID)
	var i Candicate
	err := row.Scan(
		&i.CanID,
		&i.Fullname,
		&i.Title,
		&i.Email,
		&i.Country,
		&i.Ranklocal,
		&i.Rankworld,
		&i.Company,
		&i.Dateofbirth,
		&i.CoachID,
	)
	return i, err
}

const listCandicates = `-- name: ListCandicates :many
SELECT can_id, fullname, title, email, country, ranklocal, rankworld, company, dateofbirth, coach_id FROM candicate
ORDER BY ranklocal
LIMIT $1
`

func (q *Queries) ListCandicates(ctx context.Context, limit int32) ([]Candicate, error) {
	rows, err := q.db.QueryContext(ctx, listCandicates, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Candicate
	for rows.Next() {
		var i Candicate
		if err := rows.Scan(
			&i.CanID,
			&i.Fullname,
			&i.Title,
			&i.Email,
			&i.Country,
			&i.Ranklocal,
			&i.Rankworld,
			&i.Company,
			&i.Dateofbirth,
			&i.CoachID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCandicate = `-- name: UpdateCandicate :exec
UPDATE candicate
SET fullname = $1, title = $2, email = $3, country = $4, company = $5, dateofbirth = $6
WHERE can_id = $7
`

type UpdateCandicateParams struct {
	Fullname    sql.NullString `json:"fullname"`
	Title       sql.NullString `json:"title"`
	Email       sql.NullString `json:"email"`
	Country     sql.NullString `json:"country"`
	Company     sql.NullString `json:"company"`
	Dateofbirth time.Time      `json:"dateofbirth"`
	CanID       int64          `json:"can_id"`
}

func (q *Queries) UpdateCandicate(ctx context.Context, arg UpdateCandicateParams) error {
	_, err := q.db.ExecContext(ctx, updateCandicate,
		arg.Fullname,
		arg.Title,
		arg.Email,
		arg.Country,
		arg.Company,
		arg.Dateofbirth,
		arg.CanID,
	)
	return err
}

const signCoach = `-- name: signCoach :exec
UPDATE candicate
SET coach_id = $1
WHERE can_id = $2
`

type signCoachParams struct {
	CoachID sql.NullInt32 `json:"coach_id"`
	CanID   int64         `json:"can_id"`
}

func (q *Queries) signCoach(ctx context.Context, arg signCoachParams) error {
	_, err := q.db.ExecContext(ctx, signCoach, arg.CoachID, arg.CanID)
	return err
}
