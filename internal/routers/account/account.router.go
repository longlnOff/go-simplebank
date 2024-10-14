package accounts

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-simplebank/internal/wire"
)


type AccountRouter struct {

}

func (ar *AccountRouter) InitAccountRouter(r *gin.RouterGroup) {
	accountController, err := wire.InitUserRouterHandler()

	if err != nil {
		panic(err)
	}

	accountRouterPublic := r.Group("/account")
	{
		accountRouterPublic.POST("/account", accountController.CreateAccount)
		accountRouterPublic.GET("/account/:id", accountController.GetAccount)
		accountRouterPublic.GET("/accounts", accountController.ListAccount)
	}
}