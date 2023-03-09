package resp
//userInfoResp

type UserResponse struct{
	Resp
	User User `json:"user"`
}