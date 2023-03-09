package resp
//视频流resp
type FeedResp struct{
	Resp
	VideoList []Video`json:"video_list"`
	NextTime int64`json:"next_time"`
} 

