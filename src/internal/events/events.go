package events

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"syscall"
	"time"

	"github.com/alcionai/clues"
	"github.com/armon/go-metrics"
	analytics "github.com/rudderlabs/analytics-go"

	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/storage"
)

// keys for ease of use
const (
	corsoVersion = "corso_version"
	repoID       = "repo_id"
	tenantID     = "m365_tenant_hash"

	// Event Keys
	CorsoStart   = "Corso Start"
	RepoInit     = "Repo Init"
	RepoConnect  = "Repo Connect"
	BackupStart  = "Backup Start"
	BackupEnd    = "Backup End"
	RestoreStart = "Restore Start"
	RestoreEnd   = "Restore End"

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
)

type Eventer interface {
	Event(context.Context, string, map[string]any)
	Close() error
}

// Bus handles all event communication into the events package.
type Bus struct {
	client analytics.Client

	repoID  string // one-way hash that uniquely identifies the repo.
	tenant  string // one-way hash that uniquely identifies the tenant.
	version string // the Corso release version
}

var (
	RudderStackWriteKey     string
	RudderStackDataPlaneURL string
)

func NewBus(ctx context.Context, s storage.Storage, tenID string, opts control.Options) (Bus, error) {
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
		client:  client,
		tenant:  tenantHash(tenID),
		version: version.Version,
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
		Set(corsoVersion, b.version)

	for k, v := range data {
		props.Set(k, v)
	}

	// need to setup identity when initializing a new repo
	if key == RepoInit {
		err := b.client.Enqueue(analytics.Identify{
			UserId: b.repoID,
			Traits: analytics.NewTraits().
				SetName(b.tenant).
				Set(tenantID, b.tenant),
		})
		if err != nil {
			logger.Ctx(ctx).Debugw("analytics event failure", "err", err)
		}
	}

	err := b.client.Enqueue(analytics.Track{
		Event:      key,
		UserId:     b.repoID,
		Timestamp:  time.Now().UTC(),
		Properties: props,
	})
	if err != nil {
		logger.Ctx(ctx).Info("analytics event failure", "err", err)
	}
}

func (b *Bus) SetRepoID(hash string) {
	b.repoID = hash
}

func tenantHash(tenID string) string {
	sum := md5.Sum([]byte(tenID))
	return fmt.Sprintf("%x", sum)
}

// ---------------------------------------------------------------------------
// metrics aggregation
// ---------------------------------------------------------------------------

type m string

// metrics collection bucket
const APICall m = "api_call"

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
		return ctx, func() { sig.Stop() }
	}

	stop := make(chan struct{})
	go dumpMetrics(ctx, stop, sig)

	flush := func() {
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
	tock := time.Tick(reportInterval)

	for {
		select {
		case <-tock:
			if err := syscall.Kill(syscall.Getpid(), metrics.DefaultSignal); err != nil {
				logger.CtxErr(ctx, err).Error("metrics interval signal")
			}
		case <-stop:
			return
		}
	}
}

// Inc increments the given category by 1.
func Inc(cat m, keys ...string) {
	cats := append([]string{string(cat)}, keys...)
	metrics.IncrCounter(cats, 1)
}

// Since records the duration between the provided time and now, in millis.
func Since(start time.Time, cat m, keys ...string) {
	cats := append([]string{string(cat)}, keys...)
	metrics.MeasureSince(cats, start)
}
