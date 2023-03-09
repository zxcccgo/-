package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	RouterMAP(r)
	r.Run()
	
}
func RouterMAP(r *gin.Engine){
	access := r.Group("./douyin")

	//basic
	access.GET("/feed", echo)
}
func echo(r *gin.Context){
	r.JSON(http.StatusOK,map[string]string{
	"name":"zx",
	})
}