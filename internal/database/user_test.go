package db

import (
	"context"
	"testing"
	"time"
	"github.com/longln/go-simplebank/internal/utils"
	"github.com/stretchr/testify/require"
)


func createRandomUser(t *testing.T) User {
	hasedPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		UserName: utils.RandomOwner(),
		HashedPassword: hasedPassword,
		FullName: utils.RandomOwner(),
		Email: utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}


func TestGetUser(t *testing.T) {
	// 1. Create random user
	randomUser := createRandomUser(t)

	// 2. Get user from db
	user, err := testQueries.GetUser(context.Background(), randomUser.UserName)

	// 3. Compare two accounts
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, randomUser.UserName, user.UserName)
	require.Equal(t, randomUser.HashedPassword, user.HashedPassword)
	require.Equal(t, randomUser.FullName, user.FullName)
	require.Equal(t, randomUser.Email, user.Email)
	require.WithinDuration(t, randomUser.CreatedAt, user.CreatedAt, time.Second)
	require.WithinDuration(t, randomUser.PasswordChangedAt, user.PasswordChangedAt, time.Second)
}