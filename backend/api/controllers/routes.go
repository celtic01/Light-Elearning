package controllers

import (
	middlewares "github.com/celtic01/Light-Elearning/api/middlewares"
)

func (s *Server) setRouters() {
	// User Routes
	v1 := s.Router.Group("/api/v1")
	{
		v1.POST("/users", s.CreateUser)
		v1.POST("/login", s.Login)
		v1.POST("/test", middlewares.TokenAuthMiddleware(), s.test)

	}

	// s.Router.GET("/users/:id", s.GetUser)
	// s.Router.PUT("/users/:id", s.UpdateUser)
	// s.Router.DELETE("/users/:id", s.DeleteUser)
}
