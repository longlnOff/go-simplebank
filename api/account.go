package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/longln/go-simplebank/internal/database"
	"github.com/longln/go-simplebank/internal/token"
	"github.com/longln/go-simplebank/internal/vo"
)


func (server *Server) createAccount(ctx *gin.Context) {
	var params vo.CreateAccountRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	account, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner: authPayload.Username,
		Balance: 0,
		Currency: params.Currency,
	})

	if err != nil {	
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errHandler(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func errHandler(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) getAccount(ctx *gin.Context) {
	var params vo.GetAccountRequest

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	account, err := server.store.GetAccount(ctx, params.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errHandler(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var params vo.ListAccountRequest

	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	accounts, err := server.store.ListAccounts(ctx, db.ListAccountsParams{
		Owner: authPayload.Username,
		Limit: params.PageID,
		Offset: (params.PageID-1)*params.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}