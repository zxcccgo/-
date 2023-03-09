package Tables

import "time"

type MessageDB struct {
	Id         int64 `gorm:"id"`
	ToUserId   int64`gorm:"to_user_id"`
	FromUserId int64`gorm:"from_user_id"`
	Content    string`gorm:"content"`
	CreateTime time.Time`gorm:"create_time"`
}
func (MessageDB)TableName()string{
	return "messages"
}