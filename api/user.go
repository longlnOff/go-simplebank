package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/longln/go-simplebank/internal/database"
	"github.com/longln/go-simplebank/internal/utils"
	"github.com/longln/go-simplebank/internal/vo"
)

func (server *Server) createUser(ctx *gin.Context) {
	var params vo.CreateUserRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}

	hashedPassword, err := utils.HashPassword(params.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	arg := db.CreateUserParams{
		UserName: params.UserName,
		HashedPassword: hashedPassword,
		FullName: params.FullName,
		Email: params.Email,
	}


	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {	
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errHandler(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	resp := vo.CreateUserResponse{
		UserName: user.UserName,
		FullName: user.FullName,
		Email: user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, resp)
}