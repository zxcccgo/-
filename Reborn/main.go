package main

import (
	"reborn/DAO/DB"
	"reborn/router"
	"reborn/tools/viper"

	"reborn/tools/logger"
)

func main() {
	logger.InitZAP()
	viper.InitViper()
	DB.InitDB()
	router.InitRouter()
}
