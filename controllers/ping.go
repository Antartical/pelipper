package controllers

/*
This modules contains Ping handler which will be use to healthcheck
the service
*/

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
