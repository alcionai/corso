package events

import (
	"context"
	"io"
	"os"

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

const (
	APITokens   = "api_tokens"
	GrowCounter = "grow_counter"
	RLTokens    = "rate_limit_tokens"
)

// Array of metric keys
var metricKeys = []string{
	APITokens,
	GrowCounter,
	RLTokens,
}

// Map of metricsCategory to metric.Int64Counter
var data = map[string]metric.Int64Counter{}

func NewCollector(mp metric.MeterProvider) {
	rmc := collector{}

	rmc.meter = mp.Meter("corso-meter")

	for _, key := range metricKeys {
		data[key], _ = rmc.meter.Int64Counter(key)
	}
}

func NewMetrics(ctx context.Context, w io.Writer) (context.Context, func()) {
	mp := StartClient(ctx)
	NewCollector(mp)

	return ctx, func() {}
}

// Inc increments the given category by 1.
func Inc(ctx context.Context, cat string) {
	ctr := data[cat]
	ctr.Add(context.Background(), 1)
}

// IncN increments the given category by N.
func IncN(ctx context.Context, n int, cat string) {
	ctr := data[cat]
	ctr.Add(context.Background(), int64(n))
}

// func (rmc *collector) RegisterMetricsClient(ctx context.Context) {
// 	go func() {
// 		for {
// 			rmc.updateCounter(ctx)
// 			time.Sleep(time.Second * 1)
// 		}
// 	}()

// }

// func (rmc *collector) registerCounter() {
// 	Ctr, _ = rmc.meter.Int64Counter(growCounter)
// 	AsyncCtr, _ = rmc.meter.Int64ObservableCounter("async_counter")

// 	cb := func(_ context.Context, o metric.Observer) error {
// 		logger.Ctx(context.Background()).Infow("Async counter callback")
// 		token += 100
// 		o.ObserveInt64(AsyncCtr, token)

// 		return nil
// 	}

// 	_, err := rmc.meter.RegisterCallback(
// 		cb,
// 		AsyncCtr,
// 	)

// 	if err != nil {
// 		log.Fatalf("failed to register callback: %v", err)
// 	}
// }

// func (rmc *collector) updateCounter(ctx context.Context) {
// 	logger.Ctx(ctx).Infow("updateCounter")

// 	Ctr.Add(ctx, 20)
// }

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
		logger.CtxErr(ctx, err).Error("creating metrics exporter")
	}

	meterProvider := metricSdk.NewMeterProvider(
		metricSdk.WithReader(metricSdk.NewPeriodicReader(exp)),
		metricSdk.WithResource(res),
	)

	otel.SetMeterProvider(meterProvider)

	return meterProvider
}
