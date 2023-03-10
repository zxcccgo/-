package viper
//viper 主要用于数据的初始化，方面统一修改并且支持运行时修改

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"reborn/tools/logger"
)

func InitViper() {
	viper.SetConfigFile("config\\conf.yaml")
	if err:=viper.ReadInConfig();err!=nil{
		log.Fatalln("viper init failed:",err)
		logger.SuggerLogger.Errorf("viper init failed:",err)

	}
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		log.Printf("viper.ReadInConfig failed,err:%v\n", err)
		logger.SuggerLogger.Errorf("viper.ReadInConfig failed,err:%v\n", err)

		return
	}
	viper.WatchConfig() //配置修改之后自动加载
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件需更改了...")
	})
	log.Println("viper 初始化成功...")
	logger.SuggerLogger.Info("viper 初始化成功...")
	
}