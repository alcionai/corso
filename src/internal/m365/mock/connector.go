package mock

import (
	"context"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
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
	_ *fault.Bus,
) (
	[]data.BackupCollection,
	prefixmatcher.StringSetReader,
	bool,
	error,
) {
	return ctrl.Collections, ctrl.Exclude, ctrl.Err == nil, ctrl.Err
}

func (ctrl Controller) IsRunnable(
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
) (*details.Details, error) {
	return ctrl.Deets, ctrl.Err
}

func (ctrl Controller) CacheItemInfo(dii details.ItemInfo) {}

func (ctrl Controller) ProduceExportCollections(
	_ context.Context,
	_ int,
	_ selectors.Selector,
	_ control.ExportConfig,
	_ control.Options,
	_ []data.RestoreCollection,
	_ *fault.Bus,
) ([]export.Collection, error) {
	return nil, ctrl.Err
}

func (ctrl Controller) PopulateProtectedResourceIDAndName(
	ctx context.Context,
	protectedResource string, // input value, can be either id or name
	ins idname.Cacher,
) (string, string, error) {
	return ctrl.ProtectedResourceID,
		ctrl.ProtectedResourceName,
		ctrl.ProtectedResourceErr
}
