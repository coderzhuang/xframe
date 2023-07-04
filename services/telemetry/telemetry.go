package telemetry

//
//import (
//	"github.com/chenjiandongx/ginprom"
//	"github.com/coderzhuang/core"
//	"github.com/gin-gonic/gin"
//	"github.com/prometheus/client_golang/prometheus/promhttp"
//	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
//	b3prop "go.opentelemetry.io/contrib/propagators/b3"
//	"go.opentelemetry.io/otel"
//	"go.opentelemetry.io/otel/exporters/zipkin"
//	"go.opentelemetry.io/otel/sdk/resource"
//	sdktrace "go.opentelemetry.io/otel/sdk/trace"
//	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
//)
//
//func Init() core.Middle {
//	return func(e *gin.Engine) {
//		exporter, err := zipkin.New(core.Conf.Zipkin.Url)
//		if err != nil {
//			panic(err)
//		}
//		tp := sdktrace.NewTracerProvider(
//			sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
//			sdktrace.WithResource(resource.NewWithAttributes(
//				semconv.SchemaURL,
//				semconv.ServiceNameKey.String(core.Conf.HttpServer.Name),
//			)),
//		)
//		otel.SetTracerProvider(tp)
//		otel.SetTextMapPropagator(b3prop.New())
//
//		// trace
//		e.Use(otelgin.Middleware(core.Conf.HttpServer.Name, otelgin.WithPropagators(otel.GetTextMapPropagator())))
//
//		// prometheus
//		e.Use(ginprom.PromMiddleware(&ginprom.PromOpts{ExcludeRegexEndpoint: "^/(swagger|metrics)"}))
//		e.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
//	}
//}
