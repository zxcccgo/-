package router

import (
	"github.com/gin-gonic/gin"
)

//初始化路由
func InitRouter() {
	r := gin.Default()
	RouterMAP(r)
	r.Run()
	
}
//路由功能
func RouterMAP(r *gin.Engine){
	r.Static("./static","./public")//静态文件映射
	
	//建立路由组

	//基本功能
	
	//互动功能
	//社交功能



}

