package events

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	analytics "github.com/rudderlabs/analytics-go"

	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/storage"
)

// keys for ease of use
const (
	corsoVersion = "corso_version"
	repoID       = "repo_id"

	// Event Keys
	RepoInit     = "repo_init"
	BackupStart  = "backup_start"
	BackupEnd    = "backup_end"
	RestoreStart = "restore_start"
	RestoreEnd   = "restore_end"

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

	hash := repoHash(s, tenID)

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
			return Bus{}, errors.Wrap(err, "configuring event bus")
		}
	}

	return Bus{
		client:  client,
		repoID:  hash,
		version: "vTODO", // TODO: corso versioning implementation
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
		Set(corsoVersion, b.version)

	for k, v := range data {
		props.Set(k, v)
	}

	// need to setup identity when initializing a new repo
	if key == RepoInit {
		err := b.client.Enqueue(analytics.Identify{
			UserId: b.repoID,
			Traits: analytics.NewTraits().
				SetName(b.repoID),
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
