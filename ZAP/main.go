package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
var suggerLogger *zap.SugaredLogger

func main() {
	Initlogger()
	for i:=0;i<10;i++{
		suggerLogger.Infof("this is the %d cricle",i)
	}

}
func Initlogger() {
	writeSyncer := getLogWriter()
	encoder:=getEncoder()
	core:=zapcore.NewCore(encoder,writeSyncer,zapcore.DebugLevel)
	logger:=zap.New(core,zap.AddCaller())

	suggerLogger=logger.Sugar()
	
}
func getLogWriter() zapcore.WriteSyncer{

	f, _ := os.Create("./log.log")
	return zapcore.AddSync(f)
}
func getEncoder()zapcore.Encoder{
	encoderConfig:=zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel=zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
