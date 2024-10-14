package service

import (
	"context"

	db "github.com/longln/go-simplebank/internal/database"
	"github.com/longln/go-simplebank/internal/repository"
)

type IAccountService interface {
	CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error)
	GetAccount(ctx context.Context, id int64) (db.Account, error)
	ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error)
}

type accountService struct {
	accountRepo repository.IAccountRepository
}

// CreateAccount implements IAccountService.
func (a *accountService) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {
	return a.accountRepo.CreateAccount(ctx, arg)
}

func (a *accountService) GetAccount(ctx context.Context, id int64) (db.Account, error) {
	return a.accountRepo.GetAccount(ctx, id)
}

func (a *accountService) ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.Account, error) {
	return a.accountRepo.ListAccounts(ctx, arg)
}

func NewAccountService(accountRepo repository.IAccountRepository) IAccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
}

