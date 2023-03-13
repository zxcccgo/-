package logic

import (
	"log"
	"reborn/DAO/DB"
	"reborn/models/Tables"
	"reborn/models/resp"
	"reborn/tools/logger"
	"time"
)

//逻辑层用与协调控制层需求和DAO层的服务，是业务逻辑的中间环节

//feed处理视频流逻辑
func Feed(latestTime time.Time,curUserId int64)(*resp.FeedResp,error){
	//进入函数，打印输出
	log.Println("feed function in logic layer is running.....")
	//根据时间参数，去取出预定数量的视频信息
	feedRes:=new(resp.FeedResp)
	DBvdList, err := DB.Feed(latestTime)
	if err!=nil{
		return feedRes,err
	}
	//组合视频列表
	VideoList:=make([]resp.Video,len(DBvdList))

	for i,videoTable:=range DBvdList{
		//补全视频信息
		video, err2 := CompleteVideoInfo(&videoTable,curUserId)
		if err2!=nil{
			logger.SuggerLogger.Errorf("视频组装失败：%v",err2)
		}
		//将补全的完整视频放入视频列表
		VideoList[i]=video
	}
	//组装视频流返回信息
	FeedResp:=resp.FeedResp{
		Resp: resp.Resp{
			StatusCode: 0,
			StatusMsg: "视频流请求成功",
		},
		VideoList: VideoList,
		NextTime: DBvdList[len(DBvdList)-1].PublishTime.Unix(),

	}


	return &FeedResp,nil
}

func CompleteVideoInfo(videoTable *Tables.VideosDB,curUserId int64)(resp.Video,error){
	CompeleteVD:=new(resp.Video)
	//完善视频信息
	CompeleteVD.Id=videoTable.Id
	CompeleteVD.PlayUrl=videoTable.PlayUrl
	CompeleteVD.CoverUrl=videoTable.CoverUrl
	CompeleteVD.FavoriteCnt=videoTable.FavoriteCount
	CompeleteVD.CommentCnt=videoTable.CommentCount
	CompeleteVD.Title=videoTable.Title
	//基本信息补全，完善作者（DB中只有作者ID）、点赞信息（当前用户与视频的关系）
	user,_:=GetUserByCurId(curUserId,videoTable.AuthorID)
	
	CompeleteVD.Author=user
	//是否点赞
	isLike,_:= GetLikeByUserId(curUserId, videoTable.Id)
	CompeleteVD.IsFavorite=isLike
	

	return *CompeleteVD,nil
}