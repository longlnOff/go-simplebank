//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/longln/go-simplebank/internal/repository"
	"github.com/longln/go-simplebank/internal/service"
	"github.com/longln/go-simplebank/internal/controller"

)

func InitUserRouterHandler() (*controller.AccountController, error) {
	wire.Build(
		repository.NewAccountRepository,
		service.NewAccountService,
		controller.NewController,
	)

	return new(controller.AccountController), nil
}