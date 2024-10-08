package db

import (
	"context"
	"testing"
	"time"
	"github.com/longln/go-simplebank/internal/utils"
	"github.com/stretchr/testify/require"
)


func createRandomTransfer(t *testing.T, from_account Account, to_account Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from_account.ID,
		ToAccountID: to_account.ID,
		Amount: utils.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	createRandomTransfer(t, from_account, to_account)
}


func TestGetTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	transfer := createRandomTransfer(t, from_account, to_account)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfer(t, from_account, to_account)
		createRandomTransfer(t, to_account, from_account)

	}

	arg := ListTransfersParams{
		FromAccountID: from_account.ID,
		ToAccountID: to_account.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}