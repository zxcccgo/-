package resp

//通用结构

//resp
type Resp struct{
	StatusCode int32 `json:"status_code"`
	StatusMsg string `json:"status_msg"`
}
//general user
type User struct{
	Id int64`json:"id"`
	Name string`json:"name"`
	FollowCnt int64`json:"follow_count"`
	FollowerCnt int64`json:"follower_count"`
	IsFollow bool `json:"is_follow"`
	Avatar string`json:"avatar"`
	BackgroundImg string`json:"background_image"`
	Signature string`json:"signature"`
	TotalFavorited int64`json:"total_favorited"`
	WorkCnt int64`json:"work_count"`
	FavoriteCnt int64`json:"favorite_count"`
}
//general Video
type Video struct{
	Id int64 `json:"id"`
	Author User `json:"author"`
	PlayUrl string`json:"play_url"`
	CoverUrl string`json:"cover_url"`
	FavoriteCnt int64`json:"favorite_count"`
	CommentCnt int64`json:"comment_count"`
	IsFavorite bool`json:"is_favorite"`
	Title string`json:"title"`
}
//general comment
type Comment struct{
	Id int64`json:"id"`
	User User`json:"user"`
	Content string`json:"content"`
	CreatDate string`json:"create_date"`
}
//general message
type Message struct{
	Id int64`json:"id"`
	ToUserId int64 `json:"to_user_id"`
	FromUserId int64`json:"from_user_id"`
	Content string`json:"content"`
	CreatDate string`json:"create_date"`
}