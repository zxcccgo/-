package Tables

//数据库中用于存储用户的基本信息
type UserBasicDB struct{
	UserId int64 `gorm:"user_id"`
	UserName string	`gorm:"username"`
	PassWD string	`gorm:"password"`
}
//数据库中存储用户的关注、粉丝、点赞数量
type UserInfoDB struct{
	UserId int64 `gorm:"user_id"`
	FollowNum int64	`gorm:"follow_num"`
	FansNum int64	`gorm:"fans_num"`
	Praise int64	`gorm:"praise_num"`
}
//数据库中存储用户的界面信息
type UserShowDB struct{
	UserId int64 `gorm:"user_id"`
	WorkCnt int64 `gorm:"work_count"`
	FavoriteCnt int64 `gorm:"favorite_count"`
}
//设置表名
func (UserBasicDB) TableName()string{
	return "userBasic"
}
func (UserInfoDB) TableName()string{
	return "userInfo"
}
func (UserShowDB) TableName() string{
	return "userShow"
}