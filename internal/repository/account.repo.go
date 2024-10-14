package repository

import (
	"context"

	"github.com/longln/go-simplebank/global"
	db "github.com/longln/go-simplebank/internal/database"
)


type IAccountRepository interface {
	AddAccount(ctx context.Context, arg db.AddAccountParams) (db.Account, error)
	CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (db.Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (db.Account, error)
	UpdateAccount(ctx context.Context, arg db.UpdateAccountParams) (db.Account, error)
	ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error)
}

type accountRepository struct {
	store *db.Store
}

func (ar *accountRepository) AddAccount(ctx context.Context, arg db.AddAccountParams) (db.Account, error) {
	return ar.store.AddAccount(ctx, arg)
}

func (ar *accountRepository) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	return ar.store.CreateAccount(ctx, arg)
}

func (ar *accountRepository) DeleteAccount(ctx context.Context, id int64) error {
	return ar.store.DeleteAccount(ctx, id)
}

func (ar *accountRepository) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	return ar.store.GetAccount(ctx, id)
}

func (ar *accountRepository) GetAccountForUpdate(ctx context.Context, id int64) (db.Account, error) {
	return ar.store.GetAccountForUpdate(ctx, id)
}

func (ar *accountRepository) UpdateAccount(ctx context.Context, arg db.UpdateAccountParams) (db.Account, error) {
	return ar.store.UpdateAccount(ctx, arg)
}

func (ar *accountRepository) ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error) {
	return ar.store.ListAccounts(ctx, arg)
}

func NewAccountRepository() IAccountRepository {
	return &accountRepository{
		store: db.NewStore(global.TestDB),
	}
}