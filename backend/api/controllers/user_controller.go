package controllers

import (
	"encoding/json"
	"io"

	model "github.com/celtic01/Light-Elearning/api/models"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (s *Server) CreateUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	userRequest := UserCreateRequest{}

	if err != nil {
		c.JSON(400, gin.H{"Unable to get request": err.Error()})
		return
	}

	err = json.Unmarshal(body, &userRequest)
	if err != nil {
		c.JSON(400, gin.H{"Cannot unmarshal body": err.Error()})
		return
	}

	user := model.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
	user.Prepare()

	userCreated, err := user.SaveUser(s.DB)
	if err != nil {
		c.JSON(500, gin.H{"Unable to create user": err.Error()})
		return
	}

	c.JSON(200, userCreated)

}

func (s *Server) test(c *gin.Context) {
	// Read the request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Unable to read request body", "details": err.Error()})
		return
	}

	// Process the request body as needed
	// For example, you can parse it into a struct
	var requestBody map[string]interface{}
	if err := json.Unmarshal(body, &requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Unable to parse request body", "details": err.Error()})
		return
	}

	// Generate a response based on the processed data
	responseData := gin.H{
		"message": "Received and processed request body",
		"body":    requestBody,
	}

	// Return the response
	c.JSON(200, responseData)
}
