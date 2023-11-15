package mock

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	kinject "github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ inject.BackupProducer = &mockBackupProducer{}

type mockBackupProducer struct {
	colls                   []data.BackupCollection
	dcs                     data.CollectionStats
	injectNonRecoverableErr bool
}

func NewMockBackupProducer(
	colls []data.BackupCollection,
	dcs data.CollectionStats,
	injectNonRecoverableErr bool,
) mockBackupProducer {
	return mockBackupProducer{
		colls:                   colls,
		dcs:                     dcs,
		injectNonRecoverableErr: injectNonRecoverableErr,
	}
}

func (mbp *mockBackupProducer) ProduceBackupCollections(
	context.Context,
	inject.BackupProducerConfig,
	*count.Bus,
	*fault.Bus,
) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error) {
	if mbp.injectNonRecoverableErr {
		return nil, nil, false, clues.New("non-recoverable error")
	}

	return mbp.colls, nil, true, nil
}

func (mbp *mockBackupProducer) IsServiceEnabled(
	context.Context,
	path.ServiceType,
	string,
) (bool, error) {
	return true, nil
}

func (mbp *mockBackupProducer) Wait() *data.CollectionStats {
	return &mbp.dcs
}

func (mbp mockBackupProducer) GetMetadataPaths(
	ctx context.Context,
	r kinject.RestoreProducer,
	base inject.ReasonAndSnapshotIDer,
	errs *fault.Bus,
) ([]path.RestorePaths, error) {
	ctrl := m365.Controller{}
	return ctrl.GetMetadataPaths(ctx, r, base, errs)
}

func (mbp mockBackupProducer) SetRateLimiter(
	ctx context.Context,
	service path.ServiceType,
	options control.Options,
) context.Context {
	return ctx
}
