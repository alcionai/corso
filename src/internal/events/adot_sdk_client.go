package events

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alcionai/corso/src/pkg/logger"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"

	"go.opentelemetry.io/otel/metric"

	metricSdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type collector struct {
	growCounter metric.Int64Counter
	meter       metric.Meter
}

func Newcollector(mp metric.MeterProvider) collector {
	rmc := collector{}

	rmc.meter = mp.Meter("corso-meter")
	rmc.registerCounter()

	return rmc
}

func (rmc *collector) RegisterMetricsClient(ctx context.Context, cfg Config) {
	go func() {
		for {
			rmc.updateCounter(ctx)
			time.Sleep(time.Second * 10)
		}
	}()
}

func (rmc *collector) registerCounter() {
	ctr, err := rmc.meter.Int64Counter(
		growCounter,
		metric.WithDescription("counter"),
		metric.WithUnit("count"),
	)
	if err != nil {
		fmt.Println(err)
	}

	rmc.growCounter = ctr
}

func (rmc *collector) updateCounter(ctx context.Context) {
	logger.Ctx(ctx).Infow("updateCounter")

	rmc.growCounter.Add(ctx, 20)
}

type Config struct {
	Host string
	Port string
}

func StartClient(ctx context.Context) (func(context.Context) error, error) {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("corso"),
	)
	if _, present := os.LookupEnv("OTEL_RESOURCE_ATTRIBUTES"); present {
		envResource, err := resource.New(ctx, resource.WithFromEnv())
		if err != nil {
			return nil, err
		}
		res = envResource
	}

	// Setup trace related
	tp, err := setupTraceProvider(ctx, res)
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(xray.Propagator{}) // Set AWS X-Ray propagator

	exp, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithEndpoint("0.0.0.0:4317"), otlpmetricgrpc.WithDialOption(grpc.WithBlock()))
	if err != nil {
		log.Fatalf("failed to create new OTLP metric exporter: %v", err)
	}

	meterProvider := metricSdk.NewMeterProvider(metricSdk.WithResource(res), metricSdk.WithReader(metricSdk.NewPeriodicReader(exp)))

	otel.SetMeterProvider(meterProvider)

	return func(context.Context) (err error) {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		// pushes any last exports to the receiver
		err = meterProvider.Shutdown(ctx)
		if err != nil {
			return err
		}
		err = tp.Shutdown(ctx)
		if err != nil {
			return err
		}
		return nil
	}, nil
}

// setupTraceProvider configures a trace exporter and an AWS X-Ray ID Generator.
func setupTraceProvider(ctx context.Context, res *resource.Resource) (*sdktrace.TracerProvider, error) {
	// INSECURE !! NOT TO BE USED FOR ANYTHING IN PRODUCTION
	// traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure())
	// Create and start new OTLP trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint("0.0.0.0:4317"), otlptracegrpc.WithDialOption(grpc.WithBlock()))
	if err != nil {
		log.Fatalf("failed to create new OTLP trace exporter: %v", err)
	}

	if err != nil {
		return nil, err
	}

	idg := xray.NewIDGenerator()

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(res),
		sdktrace.WithIDGenerator(idg),
	)
	return tp, nil
}
