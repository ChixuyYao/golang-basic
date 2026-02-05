package logger

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

func example() {
	var logger Logger
	logger.Info("日志: %d", 111)
}

type Field struct {
	Key string
	Val any
}

// 其他形式的Logger结构声明
type LoggerV1 interface {
	Debug(msg string, args ...Field)
	Info(msg string, args ...Field)
	Warn(msg string, args ...Field)
	Error(msg string, args ...Field)
}

func exampleV1() {
	var logger LoggerV1
	logger.Info("日志:", Field{Key: "User Id", Val: "xxxx-xxxxxxxx-xxxx"})
}

type LoggerV2 interface {
	// 参考Logger,但不采用Field Struct定义,需要保证args参数数量为偶数
	Debug(msg string, args ...any)
	// ...
}
