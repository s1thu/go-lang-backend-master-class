package db

import (
	"backend-master-class/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	transfer := Transfer{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	createTransfer, err := testQueries.CreateTransfer(context.Background(), transfer.FromAccountID, transfer.ToAccountID, transfer.Amount)
	require.NoError(t, err)
	require.NotEmpty(t, createTransfer)

	require.Equal(t, transfer.FromAccountID, createTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, createTransfer.ToAccountID)
	require.Equal(t, transfer.Amount, createTransfer.Amount)
	require.NotZero(t, createTransfer.ID)
	require.NotZero(t, createTransfer.CreatedAt)

	return createTransfer
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	getTransfer, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, getTransfer)

	require.Equal(t, transfer1.ID, getTransfer.ID)
	require.Equal(t, transfer1.FromAccountID, getTransfer.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, getTransfer.ToAccountID)
	require.Equal(t, transfer1.Amount, getTransfer.Amount)
	require.NotZero(t, getTransfer.CreatedAt)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
	}

	transfers, err := testQueries.ListTransfers(context.Background(), account1.ID, account2.ID, 5, 5)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account2.ID)
	}
}
