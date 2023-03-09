package resp

//comment list resp
type CommentListResp struct{
	Resp
	CommentList []Comment `json:"comment_list"`
}