package Tables
type UserLikeVideo struct{
	Id int64 `gorm:"id"`
	VideoId int64`gorm:"video_id"`
	IsLike bool`gorm:"is_like"`
}

func (UserLikeVideo)TableName() string{
	return "isLike"
}