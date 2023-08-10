package user

import "github.com/gin-gonic/gin"

func AddUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	{
		userRoute.GET("/:user_id", Get)
		userRoute.POST("/", Insert)
		userRoute.PUT("/:user_id", Update)
		userRoute.GET("/wx_login/:code", WXLogin)
	}
}
