package api

import (
	"os"
	"testing"
	"github.com/gin-gonic/gin"
)


func TestMain(m *testing.M) {
	// Connect to db
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}