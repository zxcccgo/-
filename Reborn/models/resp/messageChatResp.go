package resp
//message chat resp
type MessageChatResp struct{
	Resp
	MessageList []Message `json:"message_list"`
}