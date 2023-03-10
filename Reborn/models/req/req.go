package req
//req结构体用来接收前端参数，防止参数错误

// RequestSignUp 用户注册请求的结构体
type SignUpReq struct {
	Username string `form:"username"`
	Password string `form:"password"` //这里前端传的是query 这里需要添加form标签
}

// RequestLogin 用户登录请求的结构体，虽然一样内容，但是为了以后的可扩展性，还是用不一样的
type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"` //这里前端传的是query 这里需要添加form标签
}

// RequestUserInfo 获取用户信息的请求
type UserInfoReq struct {
	UserID int64  `form:"user_id"`
	Token  string `form:"token"`
}

// RequestFeed 视频流请求
type FeedReq struct {
	LatestTime int64  `form:"latest_time"`
	Token      string `form:"token"`
}

// RequestPublish 发布视频请求
type PublishReq struct {
	Token string `form:"token"`
	Data  []byte `form:"data"`
	Title string `form:"title"`
}

// RequestFavorite 视频点赞和取消点赞请求
type FavoriteReq struct {
	Token      string `form:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int32  `form:"action_type"`
}

// RequestCommentAction 视频评论操作
type CommonActionReq struct {
	Token       string `form:"token"`
	VideoId     int64  `form:"video_id"`
	ActionType  int32  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

// RequestRelation 关注操作
type RelationReq struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	ActionType int32  `form:"action_type"`
}