package DB

import (
	"reborn/models/Tables"
	"reborn/tools/logger"
)

//


func GetLikeByUserId(userId int64,videoId int64)(bool,error){
	var Like Tables.UserLikeVideo
 	if err := DB.Where("user_id=? and video_id=?", userId, videoId).Find(&Like).Error;err!=nil{
		logger.SuggerLogger.Errorf("用户视频关系查询失败：%v",err)
		return false,err
	}
	return Like.IsLike,nil


}