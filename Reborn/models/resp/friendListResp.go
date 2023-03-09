package resp
//friend list resp

type FriendListResp struct{
	Resp
	FriendList []FriendUser `json:"user_list"`
}

type FriendUser struct{
	User User `json:"user"`
	Message string `json:"message"`
	MsgType int64 `json:"msgType"`
}