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

const (
	// Default goal is to keep calls below the 10k-per-10-minute threshold.
	// 14 tokens every second nets 840 per minute.  That's 8400 every 10 minutes,
	// which is a bit below the mark.
	// But suppose we have a minute-long dry spell followed by a 10 minute tsunami.
	// We'll have built up 750 tokens in reserve, so the first 750 calls go through
	// immediately.  Over the next 10 minutes, we'll partition out the other calls
	// at a rate of 840-per-minute, ending at a total of 9150.  Theoretically, if
	// the volume keeps up after that, we'll always stay between 8400 and 9150 out
	// of 10k.  Worst case scenario, we have an extra minute of padding to allow
	// up to 9990.
	defaultPerSecond = 14  // 14 * 60 = 840
	defaultMaxCap    = 750 // real cap is 10k-per-10-minutes
	// since drive runs on a per-minute, rather than per-10-minute bucket, we have
	// to keep the max cap equal to the per-second cap.  A large maxCap pool (say,
	// 1200, similar to the per-minute cap) would allow us to make a flood of 2400
	// calls in the first minute, putting us over the per-minute limit.  Keeping
	// the cap at the per-second burst means we only dole out a max of 1240 in one
	// minute (20 cap + 1200 per minute + one burst of padding).
	drivePerSecond = 20 // 20 * 60 = 1200
	driveMaxCap    = 20 // real cap is 1250-per-minute
)

var (
	driveLimiter = rate.NewLimiter(drivePerSecond, driveMaxCap)
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
	defaultLC      = 1
	driveDefaultLC = 2
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
	defaultConsumed := defaultLC

	if limiter == driveLimiter {
		defaultConsumed = driveDefaultLC
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
