package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func main() {
	router := gin.Default()

	router.GET("/health/", healthHandler)
	router.Run(":8000")
}
