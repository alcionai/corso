package mock

import (
	"context"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	kinject "github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type mockBackupProduer struct{}

func NewMockBackupProducer() inject.BackupProducer {
	return mockBackupProduer{}
}

func (mbp mockBackupProduer) ProduceBackupCollections(
	context.Context,
	inject.BackupProducerConfig,
	*fault.Bus,
) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error) {
	panic("unimplemented")
}
func (mbp mockBackupProduer) Wait() *data.CollectionStats { panic("unimplemented") }
func (mbp mockBackupProduer) IsServiceEnabled(context.Context, path.ServiceType, string) (bool, error) {
	panic("unimplemented")
}

func (mbp mockBackupProduer) CollectMetadata(
	ctx context.Context,
	r kinject.RestoreProducer,
	man kopia.ManifestEntry,
	errs *fault.Bus,
) ([]data.RestoreCollection, error) {
	// Since the controller does not need anything special, we can
	// directly use it
	ctrl := m365.Controller{}
	return ctrl.CollectMetadata(ctx, r, man, errs)
}
