package telemetry

import (
	b3prop "go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"xframe/config"
)

func Init() *sdktrace.TracerProvider {
	exporter, err := zipkin.New(config.Conf.Zipkin.Url)
	if err != nil {
		panic(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.Conf.HttpServer.Name),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(b3prop.New())
	return tp
}
