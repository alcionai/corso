package events

import (
	"crypto/md5"
	"fmt"
	"os"
	"time"

	analytics "github.com/rudderlabs/analytics-go"
)

// keys for ease of use
const (
	corsoVersion = "corso-version"
	repoID       = "repo-id"
	payload      = "payload"

	// Event Keys
	RepoInit    = "repo-init"
	BackupStart = "backup-start"
	BackupEnd   = "backup-end"

	// Event Data Keys
	BackupID              = "backup-id"
	ExchangeResources     = "exchange-resources"
	ExchangeDataRetrieved = "exchange-data-retrieved"
	ExchangeDataStored    = "exchange-data-stored"
	EndTime               = "end-time"
	StartTime             = "start-time"
	Duration              = "duration"
	Status                = "status"
)

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

func NewBus(repoProvider, bucket, prefix, tenantID string) Bus {
	hash := repoHash(repoProvider, bucket, prefix, tenantID)

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

func (b Bus) Event(key string, data map[string]any) {
	if b.client == nil {
		return
	}

	props := analytics.
		NewProperties().
		Set(repoID, b.repoID).
		Set(corsoVersion, b.version).
		Set(payload, data)

	b.client.Enqueue(analytics.Track{
		Event:      key,
		Timestamp:  time.Now().UTC(),
		Properties: props,
	})
}

func repoHash(repoProvider, bucket, prefix, tenantID string) string {
	sum := md5.Sum(
		[]byte(repoProvider + bucket + prefix + tenantID),
	)

	return fmt.Sprintf("%x", sum)
}
