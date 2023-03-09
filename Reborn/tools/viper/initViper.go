package viper

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"reborn/tools/zap"
)

func InitViper() {
	viper.SetConfigFile("config\\conf.yaml")
	if err:=viper.ReadInConfig();err!=nil{
		log.Fatalln("viper init failed:",err)
		zap.SuggerLogger.Errorf("viper init failed:",err)

	}
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		log.Printf("viper.ReadInConfig failed,err:%v\n", err)
		zap.SuggerLogger.Errorf("viper.ReadInConfig failed,err:%v\n", err)

		return
	}
	viper.WatchConfig() //配置修改之后自动加载
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件需更改了...")
	})
	log.Println("viper 初始化成功...")
	zap.SuggerLogger.Info("viper 初始化成功...")
	
}