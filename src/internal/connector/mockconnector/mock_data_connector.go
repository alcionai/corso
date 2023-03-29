package mockconnector

import (
	"context"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type GraphConnector struct {
	Collections []data.BackupCollection
	Exclude     map[string]map[string]struct{}

	Deets *details.Details

	Err error

	Stats data.CollectionStats
}

func (gc GraphConnector) ProduceBackupCollections(
	_ context.Context,
	_, _ string,
	_ selectors.Selector,
	_ []data.RestoreCollection,
	_ control.Options,
	_ *fault.Bus,
) (
	[]data.BackupCollection,
	map[string]map[string]struct{},
	error,
) {
	return gc.Collections, gc.Exclude, gc.Err
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
