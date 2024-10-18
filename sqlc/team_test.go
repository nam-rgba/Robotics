package sqlc

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nam-rgba/blv/db/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTeam(t *testing.T) {

	arg := CreateTeamParams{
		Teamname: sql.NullString{
			String: util.RandomName(),
			Valid:  true,
		},
		CoachID: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
		JoinCode: sql.NullString{
			String: util.RandomString(6),
			Valid:  true,
		},
	}

	team, err := testQueries.CreateTeam(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, team)

	require.Equal(t, arg.Teamname.String, team.Teamname.String)
	require.NotZero(t, team.TeamID)

}
