package main

import (
	"reborn/DAO/DB"
	"reborn/router"
	"reborn/tools/viper"

	"reborn/tools/zap"
)

func main() {
	zap.InitZAP()
	viper.InitViper()
	DB.InitDB()
	router.InitRouter()
}
