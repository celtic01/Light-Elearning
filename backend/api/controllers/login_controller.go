package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/celtic01/Light-Elearning/api/auth"
	model "github.com/celtic01/Light-Elearning/api/models"
	"github.com/celtic01/Light-Elearning/api/security"
	"github.com/gin-gonic/gin"
)

func (server *Server) Login(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"Unable to get request": err.Error()})
		return
	}
	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(400, gin.H{"Cannot unmarshal body": err.Error()})
		return
	}
	user.Prepare()
	userData, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		c.JSON(400, gin.H{"Login Failed": err.Error()})
		return
	}
	c.JSON(200, userData)
}

func (server *Server) SignIn(email, password string) (map[string]interface{}, error) {
	var err error
	userData := make(map[string]interface{})
	user := model.User{}
	err = server.DB.Model(&model.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	err = security.VerifyPassword(user.Password, password)
	if err != nil {
		return nil, err
	}
	token, err := auth.CreateToken(user.ID)
	fmt.Printf("this is the token: %s and secret %s", token, os.Getenv("API_SECRET"))
	if err != nil {
		fmt.Println("this is the error creating the token: ", err)
		return nil, err
	}
	userData["token"] = token

	userData["id"] = user.ID

	userData["email"] = user.Email
	return userData, nil
}
