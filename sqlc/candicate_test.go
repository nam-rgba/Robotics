package sqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/nam-rgba/blv/db/util"
	"github.com/stretchr/testify/require"
)

func TestCreateCandicate(t *testing.T) {
	arg := CreateCandidateParams{
		Fullname: sql.NullString{
			String: util.RandomName(),
			Valid:  true,
		},
		Title: sql.NullString{
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
		Company: sql.NullString{
			String: util.RandomName(),
			Valid:  true,
		},
		Dateofbirth: time.Now().UTC().Local(),
	}

	candicate, err := testQueries.CreateCandidate(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, candicate)

	require.Equal(t, arg.Fullname.String, candicate.Fullname.String)
	require.Equal(t, arg.Title.String, candicate.Title.String)
	require.Equal(t, arg.Email.String, candicate.Email.String)
	require.Equal(t, arg.Country.String, candicate.Country.String)
	require.Equal(t, arg.Company.String, candicate.Company.String)

	require.NotZero(t, candicate.CanID)
	require.NotZero(t, candicate.Email)

}
