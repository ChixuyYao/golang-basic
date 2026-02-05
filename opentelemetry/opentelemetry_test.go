package opentelemetry

import (
	"context"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func TestServer(t *testing.T) {
	res, err := newResource("DEMO", "v0.0.1")
	require.NoError(t, err)

	prop := newPropagator()
	// 客户端,服务端之间传递Tracing信息
	otel.SetTextMapPropagator(prop)

	// 初始化Trace Provider, 后续打点使用

	tp, err := newTraceProvider(res)
	require.NoError(t, err)

	defer tp.Shutdown(context.Background())

	otel.SetTracerProvider(tp)

	server := gin.Default()
	server.GET("/test", func(c *gin.Context) {

	})
}

func newResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(serviceVersion),
		),
	)
}

// 跨端传递问题,链路元数据传递(如:AB测试标记位,全链路压力测试标记)
func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider(res *resource.Resource) (*trace.TracerProvider, error) {
	exporter, err := zipkin.New("http://localhost:9411/api/v2/spans")
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter, trace.WithBatchTimeout(time.Second)),
		trace.WithResource(res),
	)
	return traceProvider, nil
}
