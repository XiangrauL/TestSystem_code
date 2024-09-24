package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go_get_score/controller"
	"time"
)

func RegisterRouter(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许的来源
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/grades", controller.GetGrades)
	router.POST("/change_password", controller.ChangePassword)

	// 监听并启动服务
	err := router.Run(":7028")
	if err != nil {
		panic("启动HTTP Server失败: " + err.Error())
	}
}
