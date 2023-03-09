package main

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlinfo struct {
	Address   string `mapstructure:"address"`
	Port      int    `mapstructure:"port"`
	Protocal  string `mapstructure:"protocal"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Charset   string `mapstructure:"charset"`
	ParseTime string `mapstructure:"parseTime"`
	Loc       string `mapstructure:"loc"`
	Database  string `mapstructure:"database"`
}
var DB *gorm.DB
func main() {
	ViperInit()
}
func ViperInit() {
	v := viper.New()
	v.SetConfigFile("./conf.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dsn:="root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err2 error
	DB, err2 = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err2!=nil{
		panic("db 连接失败")
	}
	fmt.Println("连接成功")
	DB.AutoMigrate(&mysqlinfo{})
}
