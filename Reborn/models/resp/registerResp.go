package resp
//register Resp

type UserRegisterResp struct{
	Resp
	UserId int64`json:"user_id"`
	Token string`json:"token"`
}