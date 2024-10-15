package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/longln/go-simplebank/global"
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

	resp := newUserResponse(user)

	ctx.JSON(http.StatusOK, resp)
}

func newUserResponse(user db.User) vo.UserResponse {
	return vo.UserResponse{
		UserName: user.UserName,
		FullName: user.FullName,
		Email: user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) loginUser(ctx *gin.Context) {
	var params vo.LoginUserRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errHandler(err))
		return
	}

	user, err := server.store.GetUser(ctx, params.UserName)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errHandler(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	err = utils.CheckPasswordHash(params.Password, user.HashedPassword)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errHandler(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.UserName, global.Config.TokenConfig.AccessTokenDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errHandler(err))
		return
	}

	resp := vo.LoginUserResponse{
		AccessToken: accessToken,
		User: newUserResponse(user),
	}


	ctx.JSON(http.StatusOK, resp)
}