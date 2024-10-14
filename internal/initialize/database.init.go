package initialize

import (
	"database/sql"
	"fmt"
	"time"
	"github.com/longln/go-simplebank/global"
	"go.uber.org/zap"
	_ "github.com/lib/pq"
)


func CheckErrorPanic(err error, errString string) {
    if err != nil {
        global.Logger.Error(errString, zap.Error(err))
    }
}


func InitDataBase() {
    dbConfig := global.Config.DatabaseConfig
    stringConfig := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
        dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DatabaseName)
    db, err := sql.Open(dbConfig.Driver, stringConfig)
    CheckErrorPanic(err, "cannot connect to database")
    global.Logger.Info("Postgres intialization success")
	global.TestDB = db
	SetPool()
}


func SetPool() {
    global.TestDB.SetConnMaxIdleTime(time.Duration(global.Config.DatabaseConfig.ConnMaxLifetime))
    global.TestDB.SetMaxIdleConns(global.Config.DatabaseConfig.MaxIdleConns)
    global.TestDB.SetMaxOpenConns(global.Config.DatabaseConfig.MaxOpenConns)
}