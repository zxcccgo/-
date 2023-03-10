package logic

import (
	"log"
	"reborn/DAO/DB"
)

//点赞信息的相关操作

//根据用户id和视频id获取是否点赞信息
func GetLikeByUserId(userId int64,videoId int64)(bool,error){
	
	isLike, err := DB.GetLikeByUserId(userId, videoId)
	if err!=nil{
		log.Printf("用户视频关系查找失败：%v",err)
		return false,err
	}
	return isLike,nil

}