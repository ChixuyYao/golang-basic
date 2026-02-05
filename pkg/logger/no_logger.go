package logger

// 用于生产环境切换,不需要输出任何日志的情况
type NoLogger struct {
}

func NewNoLogger() *NoLogger {
	return &NoLogger{}
}

func (n *NoLogger) Info(msg string, args ...Field)  {}
func (n *NoLogger) Warn(msg string, args ...Field)  {}
func (n *NoLogger) Error(msg string, args ...Field) {}
func (n *NoLogger) Debug(msg string, args ...Field) {}
