package graph

import (
	"context"
	"net/http"
	"sync"

	"github.com/alcionai/clues"
	khttp "github.com/microsoft/kiota-http-go"
	"golang.org/x/time/rate"

	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// Concurrency Limiter
// "how many calls at one time"
// ---------------------------------------------------------------------------

// concurrencyLimiter middleware limits the number of concurrent requests to graph API
type concurrencyLimiter struct {
	semaphore chan struct{}
}

var (
	once                  sync.Once
	concurrencyLim        *concurrencyLimiter
	maxConcurrentRequests = 4
)

func generateConcurrencyLimiter(capacity int) *concurrencyLimiter {
	if capacity < 1 || capacity > maxConcurrentRequests {
		capacity = maxConcurrentRequests
	}

	return &concurrencyLimiter{
		semaphore: make(chan struct{}, capacity),
	}
}

func InitializeConcurrencyLimiter(capacity int) {
	once.Do(func() {
		concurrencyLim = generateConcurrencyLimiter(capacity)
	})
}

func (cl *concurrencyLimiter) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	if cl == nil || cl.semaphore == nil {
		return nil, clues.New("nil concurrency limiter")
	}

	cl.semaphore <- struct{}{}
	defer func() {
		<-cl.semaphore
	}()

	return pipeline.Next(req, middlewareIndex)
}

//nolint:lll
// ---------------------------------------------------------------------------
// Rate Limiter
// "how many calls in a minute"
// https://learn.microsoft.com/en-us/sharepoint/dev/general-development/how-to-avoid-getting-throttled-or-blocked-in-sharepoint-online
// ---------------------------------------------------------------------------

// By default we try to keep calls below a 10k-per-10-minute threshold.
// (Different APIs may have different limitations.)
// 15 tokens every second nets 900 per minute.  That's 9000 every 10 minutes,
// which is a bit below the mark.
// But suppose we have a minute-long dry spell followed by a 10 minute tsunami.
// We'll have built up 900 tokens in reserve, so the first 900 calls go through
// immediately.  Over the next 10 minutes, we'll partition out the other calls
// at a rate of 900-per-minute, ending at a total of 9900.  Theoretically, if
// the volume keeps up after that, we'll always stay between 9000 and 9900 out
// of 10k.
const (
	defaultPerSecond = 15
	defaultMaxCap    = 900
	drivePerSecond   = 15
	driveMaxCap      = 1100
)

var (
	driveLimiter = rate.NewLimiter(defaultPerSecond, defaultMaxCap)
	// also used as the exchange service limiter
	defaultLimiter = rate.NewLimiter(defaultPerSecond, defaultMaxCap)
)

type LimiterCfg struct {
	Service path.ServiceType
}

type limiterCfgKey string

const limiterCfgCtxKey limiterCfgKey = "corsoGaphRateLimiterCfg"

func BindRateLimiterConfig(ctx context.Context, lc LimiterCfg) context.Context {
	return context.WithValue(ctx, limiterCfgCtxKey, lc)
}

func ctxLimiter(ctx context.Context) *rate.Limiter {
	lc, ok := extractRateLimiterConfig(ctx)
	if !ok {
		return defaultLimiter
	}

	switch lc.Service {
	case path.OneDriveService, path.SharePointService:
		return driveLimiter
	default:
		return defaultLimiter
	}
}

func extractRateLimiterConfig(ctx context.Context) (LimiterCfg, bool) {
	l := ctx.Value(limiterCfgCtxKey)
	if l == nil {
		return LimiterCfg{}, false
	}

	lc, ok := l.(LimiterCfg)

	return lc, ok
}

type limiterConsumptionKey string

const limiterConsumptionCtxKey limiterConsumptionKey = "corsoGraphRateLimiterConsumption"

const (
	defaultLimiterConsumption      = 1
	driveDefaultLimiterConsumption = 2
	// limit consumption rate for single-item GETs requests,
	// or delta-based multi-item GETs.
	SingleGetOrDeltaLC = 1
	// limit consumption rate for anything permissions related
	PermissionsLC = 5
)

// ConsumeNTokens ensures any calls using this context will consume
// n rate-limiter tokens.  Default is 1, and this value does not need
// to be established in the context to consume the default tokens.
// This should only get used on a per-call basis, to avoid cross-pollination.
func ConsumeNTokens(ctx context.Context, n int) context.Context {
	return context.WithValue(ctx, limiterConsumptionCtxKey, n)
}

func ctxLimiterConsumption(ctx context.Context, defaultConsumption int) int {
	l := ctx.Value(limiterConsumptionCtxKey)
	if l == nil {
		return defaultConsumption
	}

	lc, ok := l.(int)
	if !ok || lc < 1 {
		return defaultConsumption
	}

	return lc
}

// QueueRequest will allow the request to occur immediately if we're under the
// calls-per-minute rate.  Otherwise, the call will wait in a queue until
// the next token set is available.
func QueueRequest(ctx context.Context) {
	limiter := ctxLimiter(ctx)
	defaultConsumed := defaultLimiterConsumption

	if limiter == driveLimiter {
		defaultConsumed = driveDefaultLimiterConsumption
	}

	consume := ctxLimiterConsumption(ctx, defaultConsumed)

	if err := limiter.WaitN(ctx, consume); err != nil {
		logger.CtxErr(ctx, err).Error("graph middleware waiting on the limiter")
	}
}

// RateLimiterMiddleware is used to ensure we don't overstep per-min request limits.
type RateLimiterMiddleware struct{}

func (mw *RateLimiterMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	QueueRequest(req.Context())
	return pipeline.Next(req, middlewareIndex)
}
