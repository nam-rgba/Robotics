package sqlc

import (
	"context"
	"database/sql"
	"errors"
)

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

// Func to invite candidate to join team
// 1. Check if candidate is already in a team
// 2. Check slot of team
// 3. Add candidate to team, set status to pending
// 4. Set candidate is in team
type InviteCandidateTxParams struct {
	CoachID     int64 `json:"coach_id"`
	CandidateID int64 `json:"can_id"`
	TeamID      int64 `json:"team_id"`
}

type InviteCandidateResult struct {
	TeamCandidate TeamCandidate `json:"team_candidate"`
}

func (store *Store) InviteCandidate(ctx context.Context, arg InviteCandidateTxParams) (InviteCandidateResult, error) {
	var result InviteCandidateResult
	err := store.execTX(ctx, func(q *Queries) error {
		var err error
		//  Check if candidate is already in a team
		isInTeam, err := q.GetInTeam(ctx, arg.CandidateID)
		if err != nil {
			return err
		}
		if isInTeam.Bool {
			return errors.New("candidate is already in a team")
		}

		// Check slot of team
		teamSlot, err := q.GetNumberOfCandidates(ctx, arg.TeamID)
		if err != nil {
			return err
		}
		if teamSlot >= 5 {
			return errors.New("team is full")
		}

		// Add candidate to team, set status to pending
		result.TeamCandidate, err = q.InviteByEmail(ctx, InviteByEmailParams{
			arg.TeamID,
			arg.CandidateID,
		})
		if err != nil {
			return err
		}

		// Set candidate is in team
		err = q.SetInTeam(ctx, SetInTeamParams{
			sql.NullBool{Bool: true, Valid: true},
			arg.CandidateID,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

// Func to response invitation, just set status
type ResponseInvitationTxParams struct {
	CandidateID int64  `json:"can_id"`
	TeamID      int64  `json:"team_id"`
	Status      string `json:"status"`
}

func (store *Store) ResponseInvitation(ctx context.Context, arg ResponseInvitationTxParams) error {
	return store.execTX(ctx, func(q *Queries) error {
		err := q.CandidateResponse(ctx, CandidateResponseParams{
			arg.TeamID,
			arg.CandidateID,
			sql.NullString{String: arg.Status, Valid: true},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
