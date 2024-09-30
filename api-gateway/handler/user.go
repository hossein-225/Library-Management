package handler

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	authpb "github.com/hossein-225/Library-Management/auth-service/proto"
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
// @Router /users/register [post]
func HandleUserRegister(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	conn, err := grpc.NewClient("user-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	err = registerUser(c.Request.Context(), client, name, email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func registerUser(ctx context.Context, client userpb.UserServiceClient, name, email, password string) error {
	req := &userpb.RegisterUserRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}

	_, err := client.RegisterUser(ctx, req)
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
// @Router /users/login [post]
func HandleUserLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	conn, err := grpc.NewClient("user-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	token, err := loginUser(c.Request.Context(), client, email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func loginUser(ctx context.Context, client userpb.UserServiceClient, email, password string) (string, error) {
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

// @Summary Get a user's profile
// @Description Retrieves the profile information for the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param email formData string true "User email"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{} "Profile retrieved successfully"
// @Failure 400 {string} string "Email cannot be empty"
// @Failure 401 {string} string "Invalid or missing token"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users/profile [post]
func HandleGetUserProfile(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	authConn, err := grpc.NewClient("auth-service:50054", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer authConn.Close()

	authClient := authpb.NewAuthServiceClient(authConn)

	email, _, err := authenticateUser(c.Request.Context(), authClient, token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	conn, err := grpc.NewClient("user-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	profile, err := getUserProfile(c.Request.Context(), client, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":  profile.Name,
		"email": profile.Email,
		"role":  profile.Role,
	})
}

func getUserProfile(ctx context.Context, client userpb.UserServiceClient, email string) (*userpb.User, error) {
	req := &userpb.GetUserProfileRequest{
		Email: email,
	}

	res, err := client.GetUserProfile(ctx, req)
	if err != nil {
		return nil, err
	}

	return res.User, nil
}

// @Summary Update a user's profile
// @Description Update profile information for the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Param name formData string true "New name"
// @Param email formData string true "New email"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} map[string]interface{} "Profile updated successfully"
// @Failure 400 {string} string "Email, name, or email cannot be empty"
// @Failure 401 {string} string "Invalid or missing token"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal server error"
// @Router /users/profile [put]
func HandleUpdateUserProfile(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
		return
	}

	token = strings.TrimPrefix(token, "Bearer ")

	authConn, err := grpc.NewClient("auth-service:50054", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer authConn.Close()

	authClient := authpb.NewAuthServiceClient(authConn)

	email, _, err := authenticateUser(c.Request.Context(), authClient, token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	name := c.PostForm("name")

	if name == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Email are required"})
		return
	}

	conn, err := grpc.NewClient("user-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	err = updateUserProfile(c.Request.Context(), client, name, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func updateUserProfile(ctx context.Context, client userpb.UserServiceClient, name, email string) error {
	req := &userpb.UpdateUserProfileRequest{
		Name:  name,
		Email: email,
	}

	_, err := client.UpdateUserProfile(ctx, req)
	return err
}
