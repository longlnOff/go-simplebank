package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Create store object to extend db and queries operations

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error)
}



type SQLStore struct {
	DB *sql.DB		// to perform db transaction
	*Queries		// to perform individual query
} // --> so use Store, we can combine individual queries and db transaction


func NewStore(db *sql.DB) Store {
	return &SQLStore{
		DB: db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(q *Queries) error) error {
	// tx begin
	tx, err := store.DB.BeginTx(ctx, nil)  // we can set isolation level, default is read-committed in postgres
	if err != nil {
		return err
	}
	// perform query
	q := New(tx)
	err = fn(q)
	if err != nil {
		// if fail, rollback
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("rollback error: %s, err %s", rbErr, err)
		}
		return err
	}

	// if success, commit
	return tx.Commit()
}


type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`

}

type TransferTxResults struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error) {
	var result TransferTxResults

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// create transfer
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			return err
		}

		// create From entry
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		// create To entry
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}


		// TODO: update accounts' balance
		// account 1
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}
		if err != nil {
			return err
		}


		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccount(ctx, AddAccountParams{
		ID: accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccount(ctx, AddAccountParams{
		ID: accountID2,
		Amount: amount2,
	})
	if err != nil {
		return
	}

	return
}