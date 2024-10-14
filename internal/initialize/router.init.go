package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/longln/go-simplebank/global"
	"github.com/longln/go-simplebank/internal/routers"
)


func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.ServerConfig.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	// Logger
	// Cors
	// RateLimiter

	mainGroup := r.Group("/v1")
	accountGroup := routers.RouterGroupAccount.Account.AccountRouter
	{
		mainGroup.GET("checkStatus")
	}
	accountGroup.InitAccountRouter(mainGroup)
	return r
}