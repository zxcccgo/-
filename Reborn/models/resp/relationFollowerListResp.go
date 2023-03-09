package resp
// follower list resp
type FollowerListResp struct{
	Resp
	UserList []User `json:"user_list"`
}