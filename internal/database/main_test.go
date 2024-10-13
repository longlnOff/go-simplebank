package db

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/longln/go-simplebank/global"
	"github.com/longln/go-simplebank/internal/initialize"
)

var testQueries *Queries
// var testDB *sql.DB

func TestMain(m *testing.M) {
	initialize.Run()
	testQueries = New(global.TestDB)
	os.Exit(m.Run())
}