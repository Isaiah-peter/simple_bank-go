package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := RandomAccount(t)
	account2 := RandomAccount(t)

	// run n concurrent transfer transaction

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAcountID: account1.ID,
				ToAcountID:   account2.ID,
				Amount:       amount,
			})

			errs <- err

			results <- result
		}()
	}

	//check result

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		tran, err := store.GetTransfers(context.Background(), transfer.ID)
		require.NoError(t, err)
		fmt.Printf("%v\n", tran)

		//entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		entriesfrom, err := store.GetEntries(context.Background(), fromEntry.ID)
		require.NoError(t, err)
		fmt.Printf("%v\n", entriesfrom)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		entriesTo, err := store.GetEntries(context.Background(), toEntry.ID)
		require.NoError(t, err)
		fmt.Printf("%v\n", entriesTo)
		// check account

		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account1.ID, toAccount.ID)

		//TODO: update account
	}
}
