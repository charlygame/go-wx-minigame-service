package user

import (
	"github.com/charlygame/CatGameService/utils"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {

	userPrivate := rg.Group("/user")
	userPrivate.Use(utils.JwtAuthMiddleware())
	{
		userPrivate.GET("/", Get)
		userPrivate.PUT("/", Update)
		userPrivate.GET("/rank-list", GetRankList)
	}

	userPublic := rg.Group("/user")
	{
		userPublic.GET("/wx_login/:code", WXLogin)
	}
}
