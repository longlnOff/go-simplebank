package initialize

import (
	"github.com/longln/go-simplebank/pkg/logger"
	"github.com/longln/go-simplebank/global"
)

func InitLogger() {
	logger := logger.NewLogger(global.Config.LogConfig)
	global.Logger = logger
}