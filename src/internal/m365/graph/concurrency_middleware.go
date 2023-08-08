package graph

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

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

// concurrencyLimiterMiddleware middleware limits the number of concurrent requests to graph API
type concurrencyLimiterMiddleware struct {
	semaphore chan struct{}
}

var (
	once                                sync.Once
	concurrencyLimitMiddlewareSingleton = &concurrencyLimiterMiddleware{}
	maxConcurrentRequests               = 4
)

func generateConcurrencyLimiter(capacity int) *concurrencyLimiterMiddleware {
	if capacity < 1 || capacity > maxConcurrentRequests {
		capacity = maxConcurrentRequests
	}

	return &concurrencyLimiterMiddleware{
		semaphore: make(chan struct{}, capacity),
	}
}

func InitializeConcurrencyLimiter(ctx context.Context, enable bool, capacity int) {
	once.Do(func() {
		switch enable {
		case true:
			logger.Ctx(ctx).Infow("turning on the concurrency limiter", "concurrency_limit", capacity)
			concurrencyLimitMiddlewareSingleton.semaphore = generateConcurrencyLimiter(capacity).semaphore
		case false:
			logger.Ctx(ctx).Info("turning off the concurrency limiter")
			concurrencyLimitMiddlewareSingleton = nil
		}
	})
}

func (cl *concurrencyLimiterMiddleware) Intercept(
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
	// 16 tokens every second nets 960 per minute.  That's 9600 every 10 minutes,
	// which is a bit below the mark.
	// If the bucket is full, we can push out 200 calls immediately, which brings
	// the total in the first 10 minutes to 9800.  We can toe that line if we want,
	// but doing so risks timeouts.  It's better to give the limits breathing room.
	defaultPerSecond = 16  // 16 * 60 * 10 = 9600
	defaultMaxCap    = 200 // real cap is 10k-per-10-minutes
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
	// https://learn.microsoft.com/en-us/sharepoint/dev/general-development
	// /how-to-avoid-getting-throttled-or-blocked-in-sharepoint-online#application-throttling
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

// ---------------------------------------------------------------------------
// global throttle fencing
// ---------------------------------------------------------------------------

// timedFence sets up a fence for a certain amount of time.
// the time can be extended arbitrarily.  All processes blocked at
// the fence will be let through when all timer extensions conclude.
type timedFence struct {
	mu     sync.Mutex
	c      chan struct{}
	timers map[int64]*time.Timer
}

func newTimedFence() *timedFence {
	return &timedFence{
		mu:     sync.Mutex{},
		c:      nil,
		timers: map[int64]*time.Timer{},
	}
}

// Block until the fence is let down.
// if no fence is up, return immediately.
// returns if the ctx deadlines before the fence is let down.
func (tf *timedFence) Block(ctx context.Context) error {
	// set to a local var to avoid race panics from tf.c
	// getting set to nil between the conditional check and
	// the read case.  If c gets closed between those two
	// points then the select case will exit immediately,
	// as if we didn't block at all.
	c := tf.c

	if c != nil {
		select {
		case <-ctx.Done():
			return clues.Wrap(ctx.Err(), "blocked on throttling fence")
		case <-c:
		}
	}

	return nil
}

// RaiseFence puts up a fence to block requests for the provided
// duration of time.  Seconds are always added to the current time.
// Multiple calls to RaiseFence are not additive. ie: calling
// `RaiseFence(5); RaiseFence(1)` will keep the fence up until
// now+5 seconds, not now+6 seconds.  When the last remaining fence
// is dropped, all currently blocked calls are allowed through.
func (tf *timedFence) RaiseFence(seconds time.Duration) {
	tf.mu.Lock()
	defer tf.mu.Unlock()

	if seconds < 1 {
		return
	}

	if tf.c == nil {
		tf.c = make(chan struct{})
	}

	timer := time.NewTimer(seconds)
	tid := time.Now().Add(seconds).UnixMilli()
	tf.timers[tid] = timer

	go func(c <-chan time.Time, id int64) {
		// wait for the timeout
		<-c

		tf.mu.Lock()
		defer tf.mu.Unlock()

		// remove the timer
		delete(tf.timers, id)

		// if no timers remain, close the channel to drop the fence
		// and set the fenc channel to nil
		if len(tf.timers) == 0 && tf.c != nil {
			close(tf.c)
			tf.c = nil
		}
	}(timer.C, tid)
}

// throttlingMiddleware is used to ensure we don't overstep per-min request limits.
type throttlingMiddleware struct {
	tf *timedFence
}

func (mw *throttlingMiddleware) Intercept(
	pipeline khttp.Pipeline,
	middlewareIndex int,
	req *http.Request,
) (*http.Response, error) {
	err := mw.tf.Block(req.Context())
	if err != nil {
		return nil, err
	}

	resp, err := pipeline.Next(req, middlewareIndex)
	if resp == nil || err != nil {
		return resp, err
	}

	seconds := getRetryAfterHeader(resp)
	if seconds < 1 {
		return resp, nil
	}

	// if all prior conditions pass, we need to add a fence that blocks
	// calls, globally, from progressing until the timeout retry-after
	// passes.
	mw.tf.RaiseFence(time.Duration(seconds) * time.Second)

	return resp, nil
}

func getRetryAfterHeader(resp *http.Response) int {
	if resp == nil || len(resp.Header) == 0 {
		return -1
	}

	if resp.StatusCode != http.StatusTooManyRequests {
		return -1
	}

	rah := resp.Header.Get(retryAfterHeader)
	if len(rah) == 0 {
		return -1
	}

	seconds, err := strconv.Atoi(rah)
	if err != nil {
		// the error itself is irrelevant, we only want
		// to wait if we have a clear length of time to wait until.
		return -1
	}

	return seconds
}
