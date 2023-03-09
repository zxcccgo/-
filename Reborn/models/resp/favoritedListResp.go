package resp
//favorited list
type FavoriteListResp struct{
	Resp
	VideoList []Video `json:"video_list"`
}