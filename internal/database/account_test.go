package db

import (
	"context"
	"database/sql"
	"testing"
	"time"
	"github.com/longln/go-simplebank/internal/utils"
	"github.com/stretchr/testify/require"
)



func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: utils.RandomOwner(),
		Balance: utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}


func TestGetAccount(t *testing.T) {
	// 1. Create random account
	randomAccount := createRandomAccount(t)

	// 2. Get account from db
	account, err := testQueries.GetAccount(context.Background(), randomAccount.ID)

	// 3. Compare two accounts
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, randomAccount.ID, account.ID)
	require.Equal(t, randomAccount.Owner, account.Owner)
	require.Equal(t, randomAccount.Balance, account.Balance)
	require.Equal(t, randomAccount.Currency, account.Currency)
	require.WithinDuration(t, randomAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestGetAccountForUpdate(t *testing.T) {
	// 1. Create random account
	randomAccount := createRandomAccount(t)

	// 2. Get account from db
	account, err := testQueries.GetAccountForUpdate(context.Background(), randomAccount.ID)

	// 3. Compare two accounts
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, randomAccount.ID, account.ID)
	require.Equal(t, randomAccount.Owner, account.Owner)
	require.Equal(t, randomAccount.Balance, account.Balance)
	require.Equal(t, randomAccount.Currency, account.Currency)
	require.WithinDuration(t, randomAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	// 1. Create random account
	randomAccount := createRandomAccount(t)

	// 2. Update that account
	arg := AddAccountParams{
		ID: randomAccount.ID,
		Amount: utils.RandomMoney(),
	}

	account, err := testQueries.AddAccount(context.Background(), arg)
	// 3. Compare information two accounts
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, randomAccount.ID, account.ID)
	require.Equal(t, randomAccount.Owner, account.Owner)
	require.Equal(t, arg.Amount + randomAccount.Balance, account.Balance)
}


func TestAddAccount(t *testing.T) {
	// 1. Create random account
	randomAccount := createRandomAccount(t)

	// 2. Update that account
	arg := UpdateAccountParams{
		ID: randomAccount.ID,
		Balance: utils.RandomMoney(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), arg)
	// 3. Compare information two accounts
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, randomAccount.ID, account.ID)
	require.Equal(t, randomAccount.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
}



func TestDeleteAccount(t *testing.T) {
	// 1. Create random account
	randomAccount := createRandomAccount(t)

	// 2. Delete account
	err := testQueries.DeleteAccount(context.Background(), randomAccount.ID)
	require.NoError(t, err)

	// 3. Get account, make sure it has been deleted
	account, err := testQueries.GetAccount(context.Background(), randomAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}

func TestListAccounts(t *testing.T) {
	// 1. Create several accounts
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	// 2. Get accounts
	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	// 3. Check requirements

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}