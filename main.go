package main

import "go.uber.org/zap"

func main() {
	initLogger()

	server := InitWebServer()
	server.Run(":8080")
}

func initLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
	/* zap日志模块使用
	zap.L().Error("msg", ...fields)
	zap.L().Debug("msg", ...fields)
	*/

}
