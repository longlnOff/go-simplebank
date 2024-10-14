package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	db "github.com/longln/go-simplebank/internal/database"
	"github.com/longln/go-simplebank/internal/service"
	"github.com/longln/go-simplebank/internal/vo"
)


type AccountController struct {
	accountService service.IAccountService
}

func NewController(accountService service.IAccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

func (a *AccountController) CreateAccount(ctx *gin.Context) {
	var params vo.CreateAccountRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: params.Owner,
		Balance: 0,
		Currency: params.Currency,
	}

	account, err := a.accountService.CreateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (a *AccountController) GetAccount(ctx *gin.Context) {
	var params vo.GetAccountRequest

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}

	account, err := a.accountService.GetAccount(ctx, params.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (a *AccountController) ListAccount(ctx *gin.Context) {
	var params vo.ListAccountRequest

	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}

	accounts, err := a.accountService.ListAccounts(ctx, db.ListAccountsParams{
		Limit: params.PageID,
		Offset: (params.PageID-1)*params.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}