package resp
//follow list resp
type RelationListResp struct{
	Resp
	UserList []User `json:"user_list"`
}
