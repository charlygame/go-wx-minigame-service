package test

import (
	"github.com/charlygame/CatGameService/config"
	"github.com/charlygame/CatGameService/db"
	"github.com/gin-gonic/gin"
)

func Init(env_file string) {
	gin.SetMode(gin.TestMode)

	if env_file == "" {
		env_file = "../test.env"
	}
	config.InitWithEnvFile(env_file)
	db.Init()
}

func Clear() {
	db.ClearDB()
	db.Disconnect()
}
