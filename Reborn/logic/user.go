package logic

import (
	"log"
	"reborn/DAO/DB"
	"reborn/models/resp"
	"strconv"

	"github.com/spf13/viper"
)

//处理用户操作相关逻辑

//根据当前用户id和目标用户id，返回用户对象和关系数据
func GetUserByCurId(curId,tarId int64)(resp.User,error){
	user:=new(resp.User)
	//补全用户与关系信息
	user.Id=tarId
	user.Avatar=viper.GetString("app.address")+strconv.Itoa(viper.GetInt("app.port"))+"/static/"+"profile.jpg"//头像暂时用一样的
	user.BackgroundImg=viper.GetString("app.address")+strconv.Itoa(viper.GetInt("app.port"))+"/static/"+"background.jpg"//背景暂时用一样的
	user.Signature="是故意的还是不小心的？"//先用一样的
	//获取目标姓名
	name, err := DB.GetNameById(tarId)
	if err!=nil{
		log.Printf("id 查询name失败：%v",err)
	}
	user.Name=name
	//获取目标关注、粉丝、获赞数量
	followCnt, fansCnt, likeCnt, err2 := DB.GetUserInfoById(tarId)
	if err2!=nil{
		log.Printf("id 查询用户相关信息失败：%v",err)
	}
	user.FollowCnt=followCnt
	user.FollowerCnt=fansCnt
	user.FavoriteCnt=likeCnt
	//当前用户是否关注过tarId
	var isFollow bool
	
	if curId==tarId{
		user.IsFollow=true
	}else{
		isFollow, err = DB.GetRelationByCurId(curId, tarId)
		if err!=nil{
			log.Println("关系不存在")

		}
	}
	user.IsFollow=isFollow

	//tar发布的作品数量和获赞数量
	workCnt, GetLiked, err3 := DB.GetPubAndFavCntById(tarId)
	if err3!=nil{
		log.Printf("获取用户视频及点赞信息失败：%v",err3)
	}
	user.WorkCnt=workCnt
	user.TotalFavorited=GetLiked


	return *user,nil
}