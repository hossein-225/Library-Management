package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	userpb "github.com/hossein-225/Library-Management/user-service/proto"
	"google.golang.org/grpc"
)

// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param name formData string true "User name"
// @Param email formData string true "User email"
// @Param password formData string true "User password"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /register [post]
func HandleUserRegister(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	err := registerUser(c.Request.Context(), name, email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func registerUser(ctx context.Context, name, email, password string) error {
	conn, err := grpc.NewClient("user-service:50052", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	req := &userpb.RegisterUserRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	_, err = client.RegisterUser(ctx, req)
	return err
}

// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param email formData string true "User email"
// @Param password formData string true "User password"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func HandleUserLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	token, err := loginUser(c.Request.Context(), email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func loginUser(ctx context.Context, email, password string) (string, error) {
	conn, err := grpc.NewClient("user-service:50052", grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	req := &userpb.AuthenticateUserRequest{
		Email:    email,
		Password: password,
	}

	res, err := client.AuthenticateUser(ctx, req)
	if err != nil {
		return "", err
	}

	return res.Token, nil
}
