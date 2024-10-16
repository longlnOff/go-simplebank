package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/longln/go-simplebank/global"
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
	tokenMaker, err := token.NewPasetoMaker(global.Config.TokenConfig.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store: store, 
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.SetupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/account/:id", server.getAccount)
	authRoutes.GET("accounts", server.listAccounts)
	authRoutes.POST("/transfers", server.createTransfer)


	server.router = router
}