package zap

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
var SuggerLogger *zap.SugaredLogger
func InitZAP() {
	writeSyncer := getLogWriter()
	encoder:=getEncoder()
	core:=zapcore.NewCore(encoder,writeSyncer,zapcore.DebugLevel)
	logger:=zap.New(core,zap.AddCaller())
	SuggerLogger=logger.Sugar()
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
