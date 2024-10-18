package sqlc

import (
	"context"
	"database/sql"
)

type AcceptCandidateTxParams struct {
	CoachID     int64 `json:"coach_id"`
	CandidateID int64 `json:"can_id"`
	TeamID      int64 `json:"team_id"`
}
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTX(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	return tx.Commit()
}

type AcceptCandidateResult struct {
	TeamCandidate TeamCandidate `json:"team_candidate"`
}

func (store *Store) AcceptCandidate(ctx context.Context, arg AcceptCandidateTxParams) (AcceptCandidateResult, error) {
	var result AcceptCandidateResult
	err := store.execTX(ctx, func(q *Queries) error {
		var err error
		// 1. Set candidate's team status
		result.TeamCandidate, err = q.AcceptCandidate(ctx, AcceptCandidateParams{
			TeamID: arg.TeamID,
			CanID:  arg.CandidateID,
		})
		if err != nil {
			return err
		}

		// 2. Set candidate's coach_id
		err = q.SignCoach(ctx, SignCoachParams{
			CanID:   arg.CandidateID,
			CoachID: sql.NullInt64{Int64: arg.CoachID, Valid: true},
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
