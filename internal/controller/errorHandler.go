package controller

import "github.com/gin-gonic/gin"

func errHandler(err error) gin.H {
	return gin.H{"error": err.Error()}
}