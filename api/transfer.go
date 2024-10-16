package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/longln/go-simplebank/internal/database"
	"github.com/longln/go-simplebank/internal/token"
	"github.com/longln/go-simplebank/internal/vo"
)

func (server *Server) createTransfer(ctx *gin.Context) {
	var params vo.TransferRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}
	from_account, ok := server.validAccount(ctx, params.FromAccountID, params.Currency)
	if !ok {
		return
	}

	_, ok = server.validAccount(ctx, params.ToAccountID, params.Currency)
	if !ok {
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if from_account.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errHandler(err))
		return
	}
	arg := db.TransferTxParams{
		FromAccountID: params.FromAccountID,
		ToAccountID: params.ToAccountID,
		Amount: params.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)

	if err != nil {	
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err :=server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errHandler(err))
			return account, false
		}
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return account, false
	}
	return account, true
}