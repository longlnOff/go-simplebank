package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/longln/go-simplebank/global"
	"github.com/longln/go-simplebank/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	// Connect to db
	LoadConfig()
	InitLogger()
	InitDataBase()
	SetPool()
	testQueries = New(global.TestDB)
	os.Exit(m.Run())
}

func LoadConfig() {
	viper := viper.New()

	viper.AddConfigPath("/home/longln/SourceCode/github.com/longln/go-simplebank/local")

	viper.SetConfigName("config")

	viper.SetConfigType("yaml")


	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
}

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

func InitLogger() {
	logger := logger.NewLogger(global.Config.LogConfig)
	global.Logger = logger
}