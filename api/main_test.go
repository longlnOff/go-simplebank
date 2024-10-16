package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/longln/go-simplebank/internal/initialize"
)


func TestMain(m *testing.M) {
	// Connect to dbg
	initialize.LoadConfig()
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}