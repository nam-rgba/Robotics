package sqlc

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nam-rgba/blv/db/util"
	"github.com/stretchr/testify/require"
)

func TestCreateCoach(t *testing.T) {
	arg := CreateCoachParams{
		Fullname: sql.NullString{
			String: util.RandomName(),
			Valid:  true,
		},
		Email: sql.NullString{
			String: util.RandomName(),
			Valid:  true,
		},
		Country: sql.NullString{
			String: util.RandomName(),
			Valid:  true,
		},
		Title: sql.NullString{
			String: util.RandomName(),
			Valid:  true,
		},
		Company: sql.NullString{
			String: util.RandomName(),
			Valid:  true,
		},
	}

	coach, err := testQueries.CreateCoach(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, coach)

	require.Equal(t, arg.Fullname.String, coach.Fullname.String)
	require.NotZero(t, coach.CoachID)

	require.Equal(t, arg.Email.String, coach.Email.String)
	require.Equal(t, arg.Country.String, coach.Country.String)
	require.Equal(t, arg.Title.String, coach.Title.String)
	require.Equal(t, arg.Company.String, coach.Company.String)
}
