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

func greeterHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, %s!", c.Param("name"))
}

func main() {
	router := gin.Default()

	router.GET("/health", healthHandler)
	router.GET("/greeting/:name", greeterHandler)
	router.Run(":8000")
}
