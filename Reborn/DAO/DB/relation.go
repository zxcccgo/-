package DB

import (
	"reborn/models/Tables"
	"reborn/tools/logger"
)

//操作用户关系信息

//根据当前id和目标id获取是否关注
func GetRelationByCurId(curId,tarId int64)(bool,error){
	var relation Tables.FollowDB
	//查询失败
	if err := DB.Where("follower_id=? and user_id=?", curId, tarId).Find(&relation).Error;err!=nil{
		logger.SuggerLogger.Errorf("关系查找失败：%v",err)
		return false,err
	}
	//查询成功
	return relation.IsFollow,nil

}