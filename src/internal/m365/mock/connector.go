package mock

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	kinject "github.com/alcionai/corso/src/internal/kopia/inject"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ inject.BackupProducer = &Controller{}

type Controller struct {
	Collections []data.BackupCollection
	Exclude     *prefixmatcher.StringSetMatcher

	Deets *details.Details

	Err error

	Stats data.CollectionStats

	ProtectedResourceID   string
	ProtectedResourceName string
	ProtectedResourceErr  error
}

func (ctrl Controller) ProduceBackupCollections(
	_ context.Context,
	_ inject.BackupProducerConfig,
	_ *count.Bus,
	_ *fault.Bus,
) (
	[]data.BackupCollection,
	prefixmatcher.StringSetReader,
	bool,
	error,
) {
	return ctrl.Collections, ctrl.Exclude, ctrl.Err == nil, ctrl.Err
}

func (ctrl *Controller) GetMetadataPaths(
	ctx context.Context,
	r kinject.RestoreProducer,
	base inject.ReasonAndSnapshotIDer,
	errs *fault.Bus,
) ([]path.RestorePaths, error) {
	return nil, clues.New("not implemented")
}

func (ctrl Controller) IsServiceEnabled(
	_ context.Context,
	_ path.ServiceType,
	_ string,
) (bool, error) {
	return true, ctrl.Err
}

func (ctrl Controller) Wait() *data.CollectionStats {
	return &ctrl.Stats
}

func (ctrl Controller) ConsumeRestoreCollections(
	_ context.Context,
	_ inject.RestoreConsumerConfig,
	_ []data.RestoreCollection,
	_ *fault.Bus,
	_ *count.Bus,
) (*details.Details, *data.CollectionStats, error) {
	return ctrl.Deets, &ctrl.Stats, ctrl.Err
}

func (ctrl Controller) CacheItemInfo(dii details.ItemInfo) {}

func (ctrl Controller) ProduceExportCollections(
	_ context.Context,
	_ int,
	_ control.ExportConfig,
	_ []data.RestoreCollection,
	_ *data.ExportStats,
	_ *fault.Bus,
) ([]export.Collectioner, error) {
	return nil, ctrl.Err
}

func (ctrl Controller) PopulateProtectedResourceIDAndName(
	ctx context.Context,
	protectedResource string, // input value, can be either id or name
	ins idname.Cacher,
) (idname.Provider, error) {
	return idname.NewProvider(ctrl.ProtectedResourceID, ctrl.ProtectedResourceName),
		ctrl.ProtectedResourceErr
}

func (ctrl Controller) SetRateLimiter(
	ctx context.Context,
	service path.ServiceType,
	options control.Options,
) context.Context {
	return ctx
}

var _ inject.RestoreConsumer = &RestoreConsumer{}

type RestoreConsumer struct {
	Deets *details.Details

	Err error

	Stats data.CollectionStats

	ProtectedResourceID   string
	ProtectedResourceName string
	ProtectedResourceErr  error
}

func (rc RestoreConsumer) IsServiceEnabled(
	context.Context,
	string,
) (bool, error) {
	return true, rc.Err
}

func (rc RestoreConsumer) PopulateProtectedResourceIDAndName(
	ctx context.Context,
	protectedResource string, // input value, can be either id or name
	ins idname.Cacher,
) (idname.Provider, error) {
	return idname.NewProvider(rc.ProtectedResourceID, rc.ProtectedResourceName),
		rc.ProtectedResourceErr
}

func (rc RestoreConsumer) CacheItemInfo(dii details.ItemInfo) {}

func (rc RestoreConsumer) ConsumeRestoreCollections(
	ctx context.Context,
	rcc inject.RestoreConsumerConfig,
	dcs []data.RestoreCollection,
	errs *fault.Bus,
	ctr *count.Bus,
) (*details.Details, *data.CollectionStats, error) {
	return rc.Deets, &rc.Stats, rc.Err
}
