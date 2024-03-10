package controllers

import (
	model "github.com/celtic01/Light-Elearning/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (s *Server) Initialize() {
	s.DB = model.InitializeDB()
	s.Router = gin.Default()
	s.setRouters()
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}
