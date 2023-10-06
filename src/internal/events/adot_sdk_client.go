package events

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/alcionai/corso/src/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"

	"go.opentelemetry.io/otel/metric"

	metricSdk "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

var Ctr metric.Int64Counter
var AsyncCtr metric.Int64ObservableCounter
var token int64

type collector struct {
	meter metric.Meter
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
			time.Sleep(time.Second * 1)
		}
	}()

}

func (rmc *collector) registerCounter() {
	Ctr, _ = rmc.meter.Int64Counter(growCounter)
	AsyncCtr, _ = rmc.meter.Int64ObservableCounter("async_counter")

	cb := func(_ context.Context, o metric.Observer) error {
		logger.Ctx(context.Background()).Infow("Async counter callback")
		token += 100
		o.ObserveInt64(AsyncCtr, token)

		return nil
	}

	_, err := rmc.meter.RegisterCallback(
		cb,
		AsyncCtr,
	)

	if err != nil {
		log.Fatalf("failed to register callback: %v", err)
	}
}

func (rmc *collector) updateCounter(ctx context.Context) {
	logger.Ctx(ctx).Infow("updateCounter")

	Ctr.Add(ctx, 20)
}

type Config struct {
	Host string
	Port string
}

func StartClient(ctx context.Context) *metricSdk.MeterProvider {
	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("corso"),
	)
	if _, present := os.LookupEnv("OTEL_RESOURCE_ATTRIBUTES"); present {
		envResource, err := resource.New(ctx, resource.WithFromEnv())
		if err != nil {
			return nil
		}
		res = envResource
	}

	exp, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint("0.0.0.0:4317"),
		otlpmetricgrpc.WithDialOption(grpc.WithBlock()),
		otlpmetricgrpc.WithTemporalitySelector(metricSdk.DefaultTemporalitySelector),
	)
	if err != nil {
		log.Fatalf("failed to create new OTLP metric exporter: %v", err)
	}

	meterProvider := metricSdk.NewMeterProvider(
		metricSdk.WithReader(metricSdk.NewPeriodicReader(exp)),
		metricSdk.WithResource(res),
	)

	otel.SetMeterProvider(meterProvider)

	return meterProvider
}

// func deltaSelector(kind metricSdk.InstrumentKind) metricdata.Temporality {
// 	switch kind {
// 	case metricSdk.InstrumentKindCounter,
// 		metricSdk.InstrumentKindHistogram,
// 		metricSdk.InstrumentKindObservableGauge,
// 		metricSdk.InstrumentKindObservableCounter:
// 		return metricdata.DeltaTemporality
// 	case metricSdk.InstrumentKindUpDownCounter,
// 		metricSdk.InstrumentKindObservableUpDownCounter:
// 		return metricdata.CumulativeTemporality
// 	}

// 	panic("unknown instrument kind")
// }
