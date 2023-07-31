package user

import "github.com/gin-gonic/gin"

func AddUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	{
		userRoute.GET("/:user_id", Get)
	}
}
