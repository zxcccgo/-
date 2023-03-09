package resp
//login resp
type UserLoginResp struct{
	Resp
	UserId int64`json:"user_id"`
	Token string`json:"token"`
}