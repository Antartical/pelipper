package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Ping -> handler for /ping route
*/
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "pong",
	})
	return
}
