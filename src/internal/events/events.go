package events

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"time"

	analytics "github.com/rudderlabs/analytics-go"

	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/storage"
)

// keys for ease of use
const (
	corsoVersion = "corso-version"
	repoID       = "repo-id"
	payload      = "payload"

	// Event Keys
	RepoInit     = "repo-init"
	BackupStart  = "backup-start"
	BackupEnd    = "backup-end"
	RestoreStart = "restore-start"
	RestoreEnd   = "restore-end"

	// Event Data Keys
	BackupID              = "backup-id"
	ExchangeResources     = "exchange-resources"
	ExchangeDataRetrieved = "exchange-data-retrieved"
	ExchangeDataStored    = "exchange-data-stored"
	EndTime               = "end-time"
	StartTime             = "start-time"
	Duration              = "duration"
	Status                = "status"
	ItemsRead             = "items-read"
	ItemsWritten          = "items-written"
)

type Eventer interface {
	Event(context.Context, string, map[string]any)
	Close() error
}

// Bus handles all event communication into the events package.
type Bus struct {
	client analytics.Client

	repoID  string // one-way hash that uniquely identifies the repo.
	version string // the Corso release version
}

var (
	WriteKey     string
	DataPlaneURL string
)

func NewBus(s storage.Storage, tenID string, opts control.Options) Bus {
	if opts.DisableMetrics {
		return Bus{}
	}

	hash := repoHash(s, tenID)

	envWK := os.Getenv("RUDDERSTACK_CORSO_WRITE_KEY")
	if len(envWK) > 0 {
		WriteKey = envWK
	}

	envDPU := os.Getenv("RUDDERSTACK_CORSO_DATA_PLANE_URL")
	if len(envDPU) > 0 {
		DataPlaneURL = envDPU
	}

	var client analytics.Client
	if len(WriteKey) > 0 && len(DataPlaneURL) > 0 {
		client = analytics.New(WriteKey, DataPlaneURL)
	}

	return Bus{
		client:  client,
		repoID:  hash,
		version: "vTODO", // TODO: corso versioning implementation
	}
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
		Set(corsoVersion, b.version)

	if len(data) > 0 {
		props.Set(payload, data)
	}

	err := b.client.Enqueue(analytics.Track{
		Event:      key,
		UserId:     b.repoID,
		Timestamp:  time.Now().UTC(),
		Properties: props,
	})
	if err != nil {
		logger.Ctx(ctx).Debugw("analytics event failure", "err", err)
	}
}

func storageID(s storage.Storage) string {
	id := s.Provider.String()

	switch s.Provider {
	case storage.ProviderS3:
		s3, err := s.S3Config()
		if err != nil {
			return id
		}

		id += s3.Bucket + s3.Prefix
	}

	return id
}

func repoHash(s storage.Storage, tenID string) string {
	sum := md5.Sum(
		[]byte(storageID(s) + tenID),
	)

	return fmt.Sprintf("%x", sum)
}
