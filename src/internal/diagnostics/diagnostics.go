package diagnostics

import (
	"context"
	"os"
	"runtime/trace"
	"strings"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/logger"
)

var localRun bool

/*
Currently using AWS x-ray for observability:
https://docs.aws.amazon.com/xray/latest/devguide/xray-concepts.html#xray-concepts-annotations

runtime/trace is also collected for load_test metrics gathering.
*/

// Initialize trace and span collection and emission.
// Should only be called as an initialization step in the context of a local run
// (such as load testing).  SDK users need not initialize the xray, as they should
// already be running their own collector.
func InitCollector() error {
	cfg := xray.Config{
		DaemonAddr:     "0.0.0.0:2000",
		ServiceVersion: "3.3.5",
	}

	if err := xray.Configure(cfg); err != nil {
		return errors.Wrap(err, "initializing observability tooling")
	}

	// TODO: feed in the corso logger
	xray.SetLogger(xraylog.NewDefaultLogger(os.Stderr, xraylog.LogLevelInfo))

	return nil
}

// Start kicks off a parent segment for tracking.  Start should only be called
// internally, and only once per corso execution.  SDK users may provide contexts
// with existing segments rather calling Start.
// The returned context will contain the parent segment (you should only ever have
// one) for all child spans to collect within; adding a span to any context besides
// this one (and its descendants) will slice that span from observation.
// The returned func closes and flushes the parent segment.
func Start(ctx context.Context, name string) (context.Context, func()) {
	ctx, seg := xray.BeginSegment(ctx, name)
	seg.TraceID = xray.NewTraceID()

	rgn := trace.StartRegion(ctx, name)
	localRun = true

	return ctx, func() {
		seg.Close(nil)
		rgn.End()
	}
}

type extender interface {
	extend(context.Context, *xray.Segment)
}

type annotation struct {
	k string
	v any
}

func (a annotation) extend(ctx context.Context, span *xray.Segment) {
	if err := span.AddAnnotation(a.k, a.v); err != nil {
		logger.Ctx(ctx).Errorw(
			"diagnostics annotation addition",
			"error", err,
			"key", a.k,
			"value", a.v,
			"span", span.Name)
	}
}

// Index annotates spans with filterable, groupable properties.
// Index values must be of type key, value, or boolean.
func Index(k string, v any) extender {
	return annotation{k, v}
}

type metadata struct {
	k string
	v any
}

func (m metadata) extend(ctx context.Context, span *xray.Segment) {
	if err := span.AddMetadata(m.k, m.v); err != nil {
		logger.Ctx(ctx).Errorw(
			"diagnostics metadata addition",
			"error", err,
			"key", m.k,
			"value", m.v,
			"span", span.Name)
	}
}

// Label tags spans with non-filterable, purely informational data.
// Label values can be any type.
func Label(k string, v any) extender {
	return metadata{k, v}
}

// Adds a Span to the given context.  Spans may be extended with indexes
// for filtering and grouping, or with labels for contextual info.
// Named variable returns are necessary here to prevent nil responses
// during panic handling.
func Span(ctx context.Context, name string, ext ...extender) (_ctx context.Context, _fn func()) {
	// no-op if no parent segment exists
	if xray.GetSegment(ctx) == nil {
		return ctx, func() {}
	}

	// spans created without an existing parent segment in the ctx will panic.
	defer func() {
		if r := recover(); r != nil {
			_ctx = ctx
			_fn = func() {}

			var rmsg string

			if s, ok := r.(string); ok {
				rmsg = s
			} else if e, ok := r.(error); ok {
				rmsg = e.Error()
			}

			if strings.Contains(rmsg, "segment cannot be found") {
				return
			}

			logger.Ctx(ctx).Errorw("diagnostics.Span", "panic", r)
		}
	}()

	ctx, span := xray.BeginSubsegment(ctx, name)
	rgn := trace.StartRegion(ctx, name)

	for _, e := range ext {
		e.extend(ctx, span)
	}

	_ctx = ctx
	_fn = func() {
		rgn.End()

		if span == nil {
			return
		}

		// during a local run we always deliver segment info to the daemon
		if localRun {
			span.CloseAndStream(nil)
			return
		}

		span.Close(nil)
	}

	return _ctx, _fn
}
