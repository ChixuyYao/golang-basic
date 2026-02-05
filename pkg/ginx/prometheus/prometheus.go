package prometheus

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type Builder struct {
	Namespace  string
	Subsystem  string
	Name       string
	Help       string
	InstanceId string // 实例ID,其他环境使用
}

func NewBuilder(namespace string, subsystem string, name string, help string, instanceId string) *Builder {
	return &Builder{Namespace: namespace, Subsystem: subsystem, Name: name, Help: help, InstanceId: instanceId}
}

// BuildResponseTime 接口请求时间
func (b *Builder) BuildResponseTime() gin.HandlerFunc {
	labels := []string{"method", "pattern", "status"}
	vector := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: b.Namespace,
		Subsystem: b.Subsystem,
		Name:      b.Name + "_response_time",
		Help:      b.Help,
		ConstLabels: map[string]string{
			"instance_id": b.InstanceId,
		},
		Objectives: map[float64]float64{
			0.5:   0.01,
			0.75:  0.01,
			0.9:   0.005,
			0.98:  0.002,
			0.99:  0.001,
			0.999: 0.0001,
		},
	}, labels)
	// 注册监控服务
	prometheus.MustRegister(vector)
	return func(ctx *gin.Context) {
		start := time.Now()
		defer func() {
			// 上传 Prometheus
			duration := time.Since(start).Milliseconds() // 请求时间的毫秒数
			method := ctx.Request.Method
			pattern := ctx.FullPath()
			status := ctx.Writer.Status()

			vector.WithLabelValues(method, pattern, strconv.Itoa(status)).Observe(float64(duration))
		}()
		ctx.Next()
	}
}

// BuildActiveRequest 活跃请求数量
func (b *Builder) BuildActiveRequest() gin.HandlerFunc {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: b.Namespace,
		Subsystem: b.Subsystem,
		Name:      b.Name + "_active_request",
		Help:      b.Help,
		ConstLabels: map[string]string{
			"instance_id": b.InstanceId,
		},
	})
	// 注册监控服务
	prometheus.MustRegister(gauge)
	return func(ctx *gin.Context) {
		gauge.Inc()
		defer gauge.Dec()
		ctx.Next()
	}
}
