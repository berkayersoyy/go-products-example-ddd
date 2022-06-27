package jaeger

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
	"log"
	"os"
)

func InitJaeger(ctx *gin.Context, spanName string, httpMethod string) (opentracing.Tracer, opentracing.Span) {
	cfg := jaegercfg.Configuration{
		ServiceName: "go-products-example-ddd",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: os.Getenv("JAEGER_HOST"),
		},
	}

	jLogger := jaegerlog.StdLogger
	tracer, _, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(metrics.NullFactory),
	)
	if err != nil {
		log.Fatalf("ERROR: cannot init Jaeger: %v\n", err)
	}
	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan(spanName)

	url := ctx.Request.Host + ctx.Request.URL.String()
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, httpMethod)

	return tracer, span
}
