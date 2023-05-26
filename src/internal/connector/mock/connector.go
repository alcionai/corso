package mock

import (
	"context"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var _ inject.BackupProducer = &GraphConnector{}

type GraphConnector struct {
	Collections []data.BackupCollection
	Exclude     *prefixmatcher.StringSetMatcher

	Deets *details.Details

	Err error

	Stats data.CollectionStats
}

func (gc GraphConnector) ProduceBackupCollections(
	_ context.Context,
	_ idname.Provider,
	_ selectors.Selector,
	_ []data.RestoreCollection,
	_ int,
	_ control.Options,
	_ *fault.Bus,
) (
	[]data.BackupCollection,
	prefixmatcher.StringSetReader,
	error,
) {
	return gc.Collections, gc.Exclude, gc.Err
}

func (gc GraphConnector) IsBackupRunnable(
	_ context.Context,
	_ path.ServiceType,
	_ string,
) (bool, error) {
	return true, gc.Err
}

func (gc GraphConnector) Wait() *data.CollectionStats {
	return &gc.Stats
}

func (gc GraphConnector) ConsumeRestoreCollections(
	_ context.Context,
	_ int,
	_ account.Account,
	_ selectors.Selector,
	_ control.RestoreDestination,
	_ control.Options,
	_ []data.RestoreCollection,
	_ *fault.Bus,
) (*details.Details, error) {
	return gc.Deets, gc.Err
}
