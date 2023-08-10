package routes

import (
	"github.com/charlygame/CatGameService/config"
	"github.com/charlygame/CatGameService/db"
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
	config.Init()
	db.Init()
	defer db.Disconnect()

	getRoutes()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(gin.Logger())
	router.Run(":8080")
}
