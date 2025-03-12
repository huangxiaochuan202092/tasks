package routes

import (
	"net/http"
	"test/config"
	"test/handlers"
	"test/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	config.GetDB()
	utils.InitRedis()

	r.LoadHTMLGlob("templates/*")

	// 将 admin 路由移到最前面，避免被其他路由干扰
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{
			"title": "管理后台",
		})
	})

	r.GET("/send-code", func(c *gin.Context) {
		c.HTML(http.StatusOK, "send_code.html", nil)
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/user-actions", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user_actions.html", nil)
	})

	userGroup := r.Group("/user")
	{
		userGroup.POST("/send-code", handlers.SendCodeHandler)
		userGroup.POST("/register", handlers.RegisterUser)
		userGroup.POST("/login", handlers.LoginUser)
	}
	return r
}
