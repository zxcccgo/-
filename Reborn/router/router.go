package router

import (
	"log"
	"reborn/controller"
	"reborn/tools/logger"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() {
	r := gin.Default()
	RouterMAP(r)
	if err := r.Run(); err != nil {
		logger.SuggerLogger.Errorf("服务开启失败：%w", err)
	}

}

// 路由功能
func RouterMAP(r *gin.Engine) {
	logger.SuggerLogger.Info("开启路由...")
	log.Println("开启路由")
	r.Static("/static", "./public") //静态文件映射

	//建立路由组
	api := r.Group("/douyin")
	//基本功能
	api.GET("/feed/", controller.Feed)
	//互动功能
	//社交功能

}
