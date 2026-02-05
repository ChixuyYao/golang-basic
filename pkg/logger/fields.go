package logger

var DEBUG = false

func Error(err error) Field {
	return Field{Key: "Error:", Val: err}
}

func SafeString(logger LoggerV1, key string, val string) Field {
	if DEBUG { // 开发环境
		return Field{Key: key, Val: val}
	} else { // 生产环境
		return Field{Key: key, Val: "********"}
	}
}
