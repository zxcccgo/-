package DB

import (
	"errors"
	"log"
	"reborn/models/Tables"
	"reborn/tools/logger"
	"time"

	"github.com/spf13/viper"
)

//用于处理DB中的视频信息

//根据latest_time取出配置中规定数量的视频
func Feed(latestTime time.Time)([]Tables.VideosDB,error){
	//获取配置文件中的视频数量
	videoCnt:=viper.GetInt("video.count")
	videos:=make([]Tables.VideosDB,videoCnt)
	if err:=DB.Limit(videoCnt).Order("publish_time desc").Find(&videos).Error;err!=nil{
		logger.SuggerLogger.Errorf("feed流查询错误：%v",err)
		return videos,err
	}
	//查询成功
	if len(videos)==0{
		logger.SuggerLogger.Info("没有新的视频了")
		log.Println("没有新的视频了")
		return videos,errors.New("DB is empty")
	}
	return videos,nil
}