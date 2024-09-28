package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleUserRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/register", func(c *gin.Context) {
		var user struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	reqBody := `{"name": "John Doe", "email": "john@example.com", "password": "secret"}`
	req, _ := http.NewRequest("POST", "/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully")
}

func TestHandleUserLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": "mock_token"})
	})

	reqBody := `{"email": "john@example.com", "password": "secret"}`
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
