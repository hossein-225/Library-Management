package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	userpb "github.com/hossein-225/Library-Management/user-service/proto"
	userpbMock "github.com/hossein-225/Library-Management/user-service/proto/mock"
	"github.com/stretchr/testify/assert"
)

func TestHandleUserRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserClient := userpbMock.NewMockUserServiceClient(ctrl)

	mockUserClient.EXPECT().RegisterUser(gomock.Any(), gomock.Any()).Return(&userpb.RegisterUserResponse{}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/users/register", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")

		err := registerUser(c.Request.Context(), mockUserClient, name, email, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	reqBody := `name=John Doe&email=john@example.com&password=secret`
	req, _ := http.NewRequest("POST", "/users/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully")
}

func TestHandleUserLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserClient := userpbMock.NewMockUserServiceClient(ctrl)

	mockUserClient.EXPECT().AuthenticateUser(gomock.Any(), gomock.Any()).Return(&userpb.AuthenticateUserResponse{Token: "mock_token"}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/users/login", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		token, err := loginUser(c.Request.Context(), mockUserClient, email, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	reqBody := `email=john@example.com&password=secret`
	req, _ := http.NewRequest("POST", "/users/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "mock_token")
}

func TestHandleGetUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserClient := userpbMock.NewMockUserServiceClient(ctrl)

	mockUserClient.EXPECT().GetUserProfile(gomock.Any(), gomock.Any()).Return(&userpb.GetUserProfileResponse{
		User: &userpb.User{Name: "John Doe", Email: "john@example.com", Role: "user"},
	}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/users/profile", func(c *gin.Context) {
		email := "john@example.com"
		profile, err := getUserProfile(c.Request.Context(), mockUserClient, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name":  profile.Name,
			"email": profile.Email,
			"role":  profile.Role,
		})
	})

	req, _ := http.NewRequest("POST", "/users/profile", nil)
	req.Header.Set("Authorization", "Bearer valid_token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestHandleUpdateUserProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserClient := userpbMock.NewMockUserServiceClient(ctrl)

	mockUserClient.EXPECT().UpdateUserProfile(gomock.Any(), gomock.Any()).Return(&userpb.UpdateUserProfileResponse{}, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.PUT("/users/profile", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		err := updateUserProfile(c.Request.Context(), mockUserClient, name, email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
	})

	reqBody := `name=John Updated&email=john_updated@example.com`
	req, _ := http.NewRequest("PUT", "/users/profile", strings.NewReader(reqBody))
	req.Header.Set("Authorization", "Bearer valid_token")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Profile updated successfully")
}
