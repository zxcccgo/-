package DB

import (
	"reborn/models/Tables"
	"reborn/tools/logger"
)

//用户相关数据库操作

func GetNameById(userId int64)(string,error){
	var user Tables.UserBasicDB
	if err := DB.Where("user_id=?", userId).Find(&user).Error;err!=nil{
		logger.SuggerLogger.Errorf("用户id:%d找不到name",userId)
		return "",err
	}


	return user.UserName,nil
}

func GetUserInfoById(userId int64)(int64,int64,int64,error){
	var userInfo Tables.UserInfoDB
	if err := DB.Where("user_id=?", userId).Find(&userInfo).Error;err!=nil{
		logger.SuggerLogger.Errorf("用户id：%d找不到userinfo",userId)
		return -1,-1,-1,err
	}
	return userInfo.FollowNum,userInfo.FansNum,userInfo.Praise,nil

}