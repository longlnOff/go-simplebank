package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/longln/go-simplebank/internal/database"
	"github.com/longln/go-simplebank/internal/token"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store db.Store
	router *gin.Engine
	tokenMaker token.Maker
}


func NewServer(store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{store: store, tokenMaker: tokenMaker}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	{
		router.POST("/accounts", server.createAccount)
		router.GET("/account/:id", server.getAccount)
		router.GET("accounts", server.listAccounts)
		router.POST("/transfers", server.createTransfer)

		router.POST("/users", server.createUser)
	}

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}