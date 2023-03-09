package Tables

import "time"

type CommentDB struct {
	Id        int64`gorm:"id"`
	VideoId   int64`gorm:"video_id"`
	UserId    int64`gorm:"user_id"`
	CommentText   string`gorm:"comment_text"`
	CreatDate time.Time`gorm:"create_time"`
}
func (CommentDB) TableName()string{
	return "comments"
}