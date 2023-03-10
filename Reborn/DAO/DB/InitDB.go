package DB

import (
	"fmt"
	"log"

	"reborn/models/Tables"
	"reborn/tools/logger"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
	viper.GetString("mysql.user"),
	viper.GetString("mysql.passwd"),
	viper.GetString("mysql.host"),
	viper.GetString("mysql.port"),
	viper.GetString("mysql.db"),
	viper.GetString("mysql.charset"),
	viper.GetString("mysql.parseTime"),
	viper.GetString("mysql.loc"))
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		logger.SuggerLogger.Errorf("DB初始化失败:",err)
		log.Fatalln("DB初始化失败:",err)


	}
	logger.SuggerLogger.Info("DB连接成功...")
	log.Println("DB 连接成功...")
	//第一次连接 创建表格
	CreateTables()
}

func CreateTables(){
	err := DB.AutoMigrate(&Tables.UserBasicDB{},
		&Tables.UserInfoDB{},
		&Tables.UserLikeVideo{},
		&Tables.UserShowDB{},
		&Tables.VideosDB{},
		&Tables.CommentDB{},
		&Tables.FollowDB{},
		&Tables.MessageDB{},
		&Tables.UserInfoDB{})
	if err!=nil{
		log.Println(err)
	}
	log.Println("表格初始化成功")
	logger.SuggerLogger.Info("表格初始化成功")

}
