package sqlc

import (
	"context"
	"database/sql"
	"testing"

	"github.com/nam-rgba/blv/db/util"
	"github.com/stretchr/testify/require"
)

func TestCreateCandicate(t *testing.T) {
	arg := RegisterCandidateParams{
		Email: sql.NullString{
			String: util.RandomString(20),
			Valid:  true,
		},
		Password: "112233",
	}

	candicate, err := testQueries.RegisterCandidate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, candicate)

	require.NotZero(t, candicate.CanID)
	require.NotZero(t, candicate.Email)

}
