package Tables

type FollowDB struct{
	Id int64`gorm:"id"`
	UserId int64`gorm:"user_id"`
	FollowerId int64`gorm:"follower_id"`
	IsFollow bool`gorm:"is_follow"`
}
func (FollowDB)TableName() string{
return "followRelations"
}