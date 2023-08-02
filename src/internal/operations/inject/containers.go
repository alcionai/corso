package inject

import (
	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// RestoreConsumerConfig is a container-of-things for holding options and
// configurations from various packages, all of which are widely used by
// restore consumers independent of service or data category.
type RestoreConsumerConfig struct {
	BackupVersion     int
	Options           control.Options
	ProtectedResource idname.Provider
	RestoreConfig     control.RestoreConfig
	Selector          selectors.Selector
}

// BackupProducerConfig is a container-of-things for holding options and
// configurations from various packages, all of which are widely used by
// backup producers independent of service or data category.
type BackupProducerConfig struct {
	LastBackupVersion   int
	MetadataCollections []data.RestoreCollection
	Options             control.Options
	ProtectedResource   idname.Provider
	Selector            selectors.Selector
}
