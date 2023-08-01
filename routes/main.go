package routes

import (
	"github.com/charlygame/CatGameService/resources/user"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func getRoutes() {
	rg := router.Group("/v1")
	{
		user.AddUserRoutes(rg)
	}
}

func Run() {
	getRoutes()
	router.Use(gin.Logger(), gin.Recovery())
	router.Run(":8080")
}
