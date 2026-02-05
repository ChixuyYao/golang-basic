package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	InitViper()
	//initLogger()

	//tpCancel := ioc.InitOTEL()
	//defer func() {
	//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//	defer cancel()
	//	tpCancel(ctx)
	//}()

	InitPrometheus()
	server := InitWebServer()
	server.Run(":8080")
}

func InitViper() {
	file := pflag.String("config", "config/dev.yaml", "配置文件路径")
	pflag.Parse()
	// 设置读取信息选项
	viper.SetConfigType("yaml")
	viper.SetConfigFile(*file)
	// 读取配置
	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件加载错误")
	}
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

func InitPrometheus() {
	// http://localhost:8081/metrics
	go func() {
		// 提供Prometheus使用的端口,与Prometheus.yaml对应
		http.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			panic("Prometheus端口启动失败")
		}
	}()
}
