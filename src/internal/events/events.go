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
	"github.com/alcionai/corso/src/pkg/storage"
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
	ExportStart      = "Export Start"
	ExportEnd        = "Export End"
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
	ExportID         = "export_id"
	Service          = "service"
	StartTime        = "start_time"
	Status           = "status"

	// default values for keys
	RepoIDNotFound = "not_found"
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

func NewBus(ctx context.Context, s storage.Storage, tenID string, co control.Options) (Bus, error) {
	if co.DisableMetrics {
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
	APICall     = "api_call"
	growCounter = "grow_counter"
)

// configurations
const (
// reportInterval    = 1 * time.Minute
// retentionDuration = 2 * time.Minute
)

func NewMetrics(ctx context.Context, w io.Writer) (context.Context, func()) {
	mp := StartClient(ctx)

	//mpx := otel.GetMeterProvider()
	rmc := Newcollector(mp)
	rmc.RegisterMetricsClient(ctx, Config{})

	return ctx, func() {}
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
