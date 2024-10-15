package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/longln/go-simplebank/internal/database"
	"github.com/longln/go-simplebank/internal/vo"
)

func (server *Server) createTransfer(ctx *gin.Context) {
	var params vo.TransferRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}

	if !server.validAccount(ctx, params.FromAccountID, params.Currency) {
		return
	}

	if !server.validAccount(ctx, params.ToAccountID, params.Currency) {
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

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err :=server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errHandler(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return false
	}
	return true
}