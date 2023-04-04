package test

import (
	"context"
	"testing"

	helper "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/helpers"
	repository "github.com/pedrovitorrs/gofinances-backend/internal/api/v1/repository/sqlc"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) repository.User {
	arg := repository.CreateUserParams{
		Username: helper.RandomString(6),
		Password: helper.RandomString(12),
		Email:    helper.RandomEmail(8),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.NotEmpty(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	userRandom := createRandomUser(t)
	userDbGenerated, err := testQueries.GetUser(context.Background(), userRandom.Username)
	require.NoError(t, err)
	require.NotEmpty(t, userDbGenerated)

	require.Equal(t, userRandom.Username, userDbGenerated.Username)
	require.Equal(t, userRandom.Password, userDbGenerated.Password)
	require.Equal(t, userRandom.Email, userDbGenerated.Email)
	require.NotEmpty(t, userDbGenerated.CreatedAt)
}

func TestGetUserById(t *testing.T) {
	userRandom := createRandomUser(t)
	userDbGenerated, err := testQueries.GetUserById(context.Background(), userRandom.ID)
	require.NoError(t, err)
	require.NotEmpty(t, userDbGenerated)

	require.Equal(t, userRandom.Username, userDbGenerated.Username)
	require.Equal(t, userRandom.Password, userDbGenerated.Password)
	require.Equal(t, userRandom.Email, userDbGenerated.Email)
	require.NotEmpty(t, userDbGenerated.CreatedAt)
}
