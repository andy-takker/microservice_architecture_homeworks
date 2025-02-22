package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthPath(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/health", healthHandler)

	t.Run("GET /health status should be OK", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/health", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("GET /health/ response should be status: OK", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/health", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, `{"status":"OK"}`, resp.Body.String())
	})

}

func TestGreeterPath(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/greeting/:name", greeterHandler)

	t.Run("GET /greeting status should be OK", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/greeting/vasya", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("GET /greeting response should be Hello, vasya!", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/greeting/vasya", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, `Hello, vasya!`, resp.Body.String())
	})

}
