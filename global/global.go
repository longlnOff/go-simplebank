package global

import (
	"database/sql"
	"github.com/longln/go-simplebank/pkg/logger"
	"github.com/longln/go-simplebank/pkg/setting"
)


var (
	Config setting.Config
	Logger *logger.LoggerZap
	TestDB *sql.DB
)