package Tables

import "time"

type VideosDB struct {
	VideoID       int64     `gorm:"id"`
	AuthorID      int64     `gorm:"author_id"`
	PlayUrl       string    `gorm:"play_url"`
	CoverUrl      string    `gorm:"cover_url"`
	PublishTime   time.Time `gorm:"publish_time"`
	FavoriteCount int64     `gorm:"favorite_count"`
	CommentCount  int64     `gorm:"comment_count"`
	Title         string    `gorm:"title"`
}
func(VideosDB) TableName()string{
	return "videos"
}