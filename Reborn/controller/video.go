package controller

import (
	"log"
	"net/http"
	"reborn/logic"
	"reborn/models/req"
	"reborn/models/resp"
	"reborn/tools/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//视频控制器

// feed 视频流功能
func Feed(c *gin.Context) {
	//获取请求参数

	reqPar := new(req.FeedReq)
	if err := c.ShouldBindQuery(reqPar); err != nil {
		//请求参数有误
		logger.SuggerLogger.Errorf("传入参数错误：%v", err) //记录错误
		c.JSON(http.StatusOK, resp.FeedResp{
			Resp: resp.Resp{StatusCode: 1, StatusMsg: "请求参数错误，获取视频流失败"},
		})
		return
	}
	//记录传入时间
	logger.SuggerLogger.Info("视频里请求传入时间：", zap.Int64("last_time", reqPar.LatestTime))
	//获取latest时间
	var latestTime time.Time
	if reqPar.LatestTime != 0 {
		latestTime = time.Unix(reqPar.LatestTime, 0)
	} else {
		latestTime = time.Now()
	}
	logger.SuggerLogger.Info("时间戳：\n", latestTime)
	//获取用户ID
	userId := c.GetInt64("userId") //得不到默认是0
	//输出到控制台
	log.Printf("用户id:%v请求视频流\n", userId)
	//处理视频流相关逻辑
	feed, err := logic.Feed(latestTime, userId)
	if err != nil {
		c.JSON(http.StatusOK, resp.FeedResp{
			Resp: resp.Resp{StatusCode: 1, StatusMsg: "数据库中没有视频"},
		})
		return
	}
	c.JSON(http.StatusOK, feed)
}
