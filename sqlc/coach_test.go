package sqlc

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nam-rgba/blv/db/util"
	"github.com/stretchr/testify/require"
)

func TestCreateCoach(t *testing.T) {
	arg := RegisterCoachParams{
		Email: sql.NullString{
			String: util.RandomName(),
			Valid:  true,
		},
		Password: "112233",
	}

	coach, err := testQueries.RegisterCoach(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, coach)

}
