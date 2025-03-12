package routes

import (
	"test/config"
	"test/handlers"
	"test/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	config.GetDB()
	utils.InitRedis()

	userGroup := r.Group("/user")
	{
		userGroup.POST("/send-code", handlers.SendCodeHandler)
		userGroup.POST("/register", handlers.RegisterUser)
		userGroup.POST("/login", handlers.LoginUser)
	}
	return r
}
