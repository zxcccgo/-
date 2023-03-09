package resp
//comment action resp

type CommentActionResp struct{
	Resp
	Comment Comment `json:"comment"`
}