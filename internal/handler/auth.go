package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/go-graphql-ws-api/internal/utils"
)

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse represents the login response data
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

// User represents the user data in the response
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

// Login handles user authentication
func Login(c *gin.Context) {
	print(c)
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c.Writer, "Invalid request parameters", map[string][]string{
			"email":    {"The email field is required and must be a valid email"},
			"password": {"The password field is required and must be at least 6 characters"},
		})
		return
	}

	// TODO: Implement actual authentication logic
	// This is just a placeholder response
	response := LoginResponse{
		AccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
		User: User{
			ID:        "1",
			Email:     req.Email,
			Name:      "John Doe",
			CreatedAt: "2024-04-19T00:00:00Z",
		},
	}

	utils.Success(c.Writer, response, "Login successful")
}
