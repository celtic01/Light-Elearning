package api

import "github.com/celtic01/Light-Elearning/api/controllers"

var Server = controllers.Server{}

func Run() {
	Server.Initialize()
	Server.Run(":8080")
}
