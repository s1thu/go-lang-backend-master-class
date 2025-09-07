package db

import (
	"backend-master-class/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := Account{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(),
		args.Owner,
		args.Balance,
		args.Currency,
	)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreate(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	fetchedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)

	require.Equal(t, account.ID, fetchedAccount.ID)
	require.Equal(t, account.Owner, fetchedAccount.Owner)
	require.Equal(t, account.Balance, fetchedAccount.Balance)
	require.Equal(t, account.Currency, fetchedAccount.Currency)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	args := Account{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(),
		args.ID,
		args.Balance)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, args.Balance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	fetchedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, fetchedAccount)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestList(t *testing.T) {
	for range 5 {
		createRandomAccount(t)
	}

	fetchedAccounts, err := testQueries.ListAccount(context.Background(), 5, 5)
	require.NoError(t, err)
	require.Len(t, fetchedAccounts, 5)

	for _, v := range fetchedAccounts {
		require.NotEmpty(t, v)
	}
}
