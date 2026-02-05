package prometheus

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestCounter(t *testing.T) {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "命名空间",
		Subsystem: "子系统",
		Name:      "名字",
	})
	prometheus.MustRegister(counter)
	counter.Inc() // 计数加一
	counter.Add(2)
}

func TestGauge(t *testing.T) {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "namespace",
		Subsystem: "sub_system",
		Name:      "name",
	})
	prometheus.MustRegister(gauge)
}

func TestHistogram(t *testing.T) {
	histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "namespace",
		Subsystem: "sub_system",
		Name:      "name",
		// 分桶
		Buckets: []float64{10, 50, 100, 200, 500, 10000, 100000},
	})
	prometheus.MustRegister(histogram)
	histogram.Observe(12)
}

func TestSummary(t *testing.T) {
	summary := prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "namespace",
		Subsystem: "sub_system",
		Name:      "name",
		// 数据观测
		Objectives: map[float64]float64{
			// KEY:百分比,VAL:误差
			0.5:   0.01, // 区间: 0.49 - 0.51
			0.75:  0.01,
			0.9:   0.005,
			0.98:  0.002,
			0.99:  0.001,
			0.999: 0.0001,
		},
	})
	prometheus.MustRegister(summary)
	summary.Observe(12)
}

func TestVector(t *testing.T) {
	summaryVec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:      "name",
		Subsystem: "sub_system",
		ConstLabels: map[string]string{
			"server":   "localhost:9091",
			"env":      "test",
			"app_name": "test_app",
		},
		Help: "HELP",
	}, []string{"pattern", "method", "status"})
	// 请求:pattern=/user/:id, method=POST, status=200时,响应时间128
	summaryVec.WithLabelValues("/user/:id", "POST", "200").Observe(128)
}
