package db

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/longln/go-simplebank/global"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
// var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	global.TestDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}
	testQueries = New(global.TestDB)
	os.Exit(m.Run())
}