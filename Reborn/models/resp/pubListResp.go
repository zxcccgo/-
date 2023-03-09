package resp
//Public list resp
type PublishListResp struct{
	Resp
	VideoList []Video `json:"video_list"`
}