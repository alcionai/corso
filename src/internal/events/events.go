package events

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"github.com/alcionai/clues"
	"github.com/armon/go-metrics"
	analytics "github.com/rudderlabs/analytics-go"

	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
)

// keys for ease of use
const (
	corsoVersion       = "corso_version"
	repoID             = "repo_id"
	tenantID           = "m365_tenant_hash"
	tenantIDDeprecated = "m365_tenant_hash_deprecated"

	// Event Keys
	CorsoStart       = "Corso Start"
	RepoInit         = "Repo Init"
	RepoConnect      = "Repo Connect"
	BackupStart      = "Backup Start"
	BackupEnd        = "Backup End"
	RestoreStart     = "Restore Start"
	RestoreEnd       = "Restore End"
	MaintenanceStart = "Maintenance Start"
	MaintenanceEnd   = "Maintenance End"

	// Event Data Keys
	BackupCreateTime = "backup_creation_time"
	BackupID         = "backup_id"
	DataRetrieved    = "data_retrieved"
	DataStored       = "data_stored"
	Duration         = "duration"
	EndTime          = "end_time"
	ItemsRead        = "items_read"
	ItemsWritten     = "items_written"
	Resources        = "resources"
	RestoreID        = "restore_id"
	Service          = "service"
	StartTime        = "start_time"
	Status           = "status"
	RepoID           = "not_found"
)

const (
	sha256OutputLength  = 64
	truncatedHashLength = 32
)

type Eventer interface {
	Event(context.Context, string, map[string]any)
	Close() error
}

// Bus handles all event communication into the events package.
type Bus struct {
	client analytics.Client

	repoID           string // one-way hash that uniquely identifies the repo.
	tenant           string // one-way hash that uniquely identifies the tenant.
	tenantDeprecated string // one-way hash that uniquely identified the tenand (old hashing algo for continuity).
	version          string // the Corso release version
}

var (
	RudderStackWriteKey     string
	RudderStackDataPlaneURL string
)

func NewBus(ctx context.Context, tenID string, opts control.Options) (Bus, error) {
	if opts.DisableMetrics {
		return Bus{}, nil
	}

	envWK := os.Getenv("RUDDERSTACK_CORSO_WRITE_KEY")
	if len(envWK) > 0 {
		RudderStackWriteKey = envWK
	}

	envDPU := os.Getenv("RUDDERSTACK_CORSO_DATA_PLANE_URL")
	if len(envDPU) > 0 {
		RudderStackDataPlaneURL = envDPU
	}

	var client analytics.Client

	if len(RudderStackWriteKey) > 0 && len(RudderStackDataPlaneURL) > 0 {
		var err error
		client, err = analytics.NewWithConfig(
			RudderStackWriteKey,
			RudderStackDataPlaneURL,
			analytics.Config{
				Logger: logger.WrapCtx(ctx, logger.ForceDebugLogLevel()),
			})

		if err != nil {
			return Bus{}, clues.Wrap(err, "configuring event bus").WithClues(ctx)
		}
	}

	return Bus{
		client:           client,
		tenant:           sha256Truncated(tenID),
		tenantDeprecated: tenantHash(tenID),
		version:          version.Version,
	}, nil
}

func (b Bus) Close() error {
	if b.client == nil {
		return nil
	}

	return b.client.Close()
}

func (b Bus) Event(ctx context.Context, key string, data map[string]any) {
	if b.client == nil {
		return
	}

	props := analytics.
		NewProperties().
		Set(repoID, b.repoID).
		Set(tenantID, b.tenant).
		Set(tenantIDDeprecated, b.tenantDeprecated).
		Set(corsoVersion, b.version)

	for k, v := range data {
		props.Set(k, v)
	}

	// need to setup identity when initializing or connecting to a repo
	if key == RepoInit || key == RepoConnect {
		err := b.client.Enqueue(analytics.Identify{
			UserId: b.tenant,
			Traits: analytics.NewTraits().
				SetName(b.tenant).
				Set(tenantID, b.tenant).
				Set(tenantIDDeprecated, b.tenantDeprecated).
				Set(repoID, b.repoID),
		})
		if err != nil {
			logger.CtxErr(ctx, err).Debug("analytics event failure: repo identity")
		}
	}

	err := b.client.Enqueue(analytics.Track{
		Event:      key,
		UserId:     b.tenant,
		Timestamp:  time.Now().UTC(),
		Properties: props,
	})
	if err != nil {
		logger.CtxErr(ctx, err).Info("analytics event failure: tracking event")
	}
}

func (b *Bus) SetRepoID(hash string) {
	b.repoID = hash
}

func sha256Truncated(tenID string) string {
	outputLength := int(math.Min(truncatedHashLength, sha256OutputLength))

	hash := sha256.Sum256([]byte(tenID))
	hexHash := fmt.Sprintf("%x", hash)

	return hexHash[0:outputLength]
}

func tenantHash(tenID string) string {
	sum := md5.Sum([]byte(tenID))
	return fmt.Sprintf("%x", sum)
}

// ---------------------------------------------------------------------------
// metrics aggregation
// ---------------------------------------------------------------------------

type metricsCategory string

// metrics collection bucket
const (
	APICall metricsCategory = "api_call"
)

// configurations
const (
	reportInterval    = 1 * time.Minute
	retentionDuration = 2 * time.Minute
)

// NewMetrics embeds a metrics bus into the provided context.  The bus can be
// utilized with calls like Inc and Since.
func NewMetrics(ctx context.Context, w io.Writer) (context.Context, func()) {
	var (
		// report interval time-bounds metrics into buckets.  Retention
		// controls how long each interval sticks around.  Neither one controls
		// logging rates; that's handled by dumpMetrics().
		sink = metrics.NewInmemSink(reportInterval, retentionDuration)
		cfg  = metrics.DefaultConfig("corso")
		sig  = metrics.NewInmemSignal(sink, metrics.DefaultSignal, w)
	)

	cfg.EnableHostname = false
	cfg.EnableRuntimeMetrics = false

	gm, err := metrics.NewGlobal(cfg, sink)
	if err != nil {
		logger.CtxErr(ctx, err).Error("metrics bus constructor")
		sig.Stop()

		return ctx, func() {}
	}

	stop := make(chan struct{})
	go dumpMetrics(ctx, stop, sig)

	flush := func() {
		signalDump(ctx)
		time.Sleep(500 * time.Millisecond)
		close(stop)
		sig.Stop()
		gm.Shutdown()
	}

	// return context.WithValue(ctx, sinkCtxKey, sink), flush
	return ctx, flush
}

// dumpMetrics runs a loop that sends a os signal (SIGUSR1 on linux/mac, SIGBREAK on windows)
// every logging interval.  This syscall getts picked up by the metrics inmem signal and causes
// it to dump metrics to the provided writer (which should be the logger).
// Expectation is for users to call this in a goroutine.  Any signal or close() on the stop chan
// will exit the loop.
func dumpMetrics(ctx context.Context, stop <-chan struct{}, sig *metrics.InmemSignal) {
	tock := time.NewTicker(reportInterval)

	for {
		select {
		case <-tock.C:
			signalDump(ctx)
		case <-stop:
			return
		}
	}
}

// Inc increments the given category by 1.
func Inc(cat metricsCategory, keys ...string) {
	cats := append([]string{string(cat)}, keys...)
	metrics.IncrCounter(cats, 1)
}

// IncN increments the given category by N.
func IncN(n int, cat metricsCategory, keys ...string) {
	cats := append([]string{string(cat)}, keys...)
	metrics.IncrCounter(cats, float32(n))
}

// Since records the duration between the provided time and now, in millis.
func Since(start time.Time, cat metricsCategory, keys ...string) {
	cats := append([]string{string(cat)}, keys...)
	metrics.MeasureSince(cats, start)
}
