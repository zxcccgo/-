package DB

import (
	"reborn/models/Tables"
	"reborn/tools/logger"
)

//publish 相关的一些操作

//根据用户id获取发布的视频数量和点赞数量
func GetPubAndFavCntById(userId int64)(int64,int64,error){
	var userShow Tables.UserShowDB
	if err := DB.Where("id=?", userId).Find(&userShow).Error;err!=nil{
		logger.SuggerLogger.Errorf("用户发布视频及获赞信息获取失败：%v",err)
		return -1,-1,err
	}
	//查询成功
	return userShow.WorkCnt,userShow.FavoriteCnt,nil
}