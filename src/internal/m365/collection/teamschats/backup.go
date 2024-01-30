package teamschats

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

func CreateCollections[I chatsItemer](
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	bh backupHandler[I],
	tenantID string,
	scope selectors.TeamsChatsScope,
	statusUpdater support.StatusUpdater,
	useLazyReader bool,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	var (
		category = scope.Category().PathType()
		qp       = graph.QueryParams{
			Category:          category,
			ProtectedResource: bpc.ProtectedResource,
			TenantID:          tenantID,
		}
	)

	cc := api.CallConfig{
		CanMakeDeltaQueries: false,
	}

	container, err := bh.getContainer(ctx, cc)
	if err != nil {
		return nil, false, clues.Stack(err)
	}

	counter.Add(count.Containers, 1)

	collection, err := populateCollection[I](
		ctx,
		qp,
		bh,
		statusUpdater,
		container,
		scope,
		useLazyReader,
		bpc.Options,
		counter,
		errs)
	if err != nil {
		return nil, false, clues.Wrap(err, "filling collections")
	}

	collections := []data.BackupCollection{collection}

	metadataPrefix, err := path.BuildMetadata(
		qp.TenantID,
		qp.ProtectedResource.ID(),
		path.TeamsChatsService,
		qp.Category,
		false)
	if err != nil {
		return nil, false, clues.WrapWC(ctx, err, "making metadata path prefix").
			Label(count.BadPathPrefix)
	}

	metadataCollection, err := graph.MakeMetadataCollection(
		metadataPrefix,
		// no deltas or previousPaths are used here; we store empty files instead
		[]graph.MetadataCollectionEntry{
			graph.NewMetadataEntry(metadata.PreviousPathFileName, map[string]string{}),
			graph.NewMetadataEntry(metadata.DeltaURLsFileName, map[string]string{}),
		},
		statusUpdater,
		counter.Local())
	if err != nil {
		return nil, false, clues.WrapWC(ctx, err, "making metadata collection")
	}

	collections = append(collections, metadataCollection)

	// no deltas involved in this category, so canUsePrevBackups is always true.
	return collections, true, nil
}

func populateCollection[I chatsItemer](
	ctx context.Context,
	qp graph.QueryParams,
	bh backupHandler[I],
	statusUpdater support.StatusUpdater,
	container container[I],
	scope selectors.TeamsChatsScope,
	useLazyReader bool,
	ctrlOpts control.Options,
	counter *count.Bus,
	errs *fault.Bus,
) (data.BackupCollection, error) {
	var (
		cl         = counter.Local()
		collection data.BackupCollection
		err        error
	)

	ctx = clues.AddLabelCounter(ctx, cl.PlainAdder())
	cc := api.CallConfig{
		CanMakeDeltaQueries: false,
	}

	items, err := bh.getItemIDs(ctx, cc)
	if err != nil {
		errs.AddRecoverable(ctx, clues.Stack(err))
		return collection, clues.Stack(errs.Failure()).OrNil()
	}

	// Only create a collection if the path matches the scope.
	includedItems := []I{}

	for _, item := range items {
		if !bh.includeItem(item, scope) {
			cl.Inc(count.SkippedItems)
			continue
		}

		includedItems = append(includedItems, item)
	}

	cl.Add(count.ItemsAdded, int64(len(includedItems)))

	p, err := bh.CanonicalPath()
	if err != nil {
		err = clues.StackWC(ctx, err).Label(count.BadCollPath)
		errs.AddRecoverable(ctx, err)

		return collection, clues.Stack(errs.Failure()).OrNil()
	}

	collection = NewCollection(
		data.NewBaseCollection(
			p,
			p,
			container.humanLocation.Builder(),
			ctrlOpts,
			false,
			cl),
		bh,
		qp.ProtectedResource.ID(),
		includedItems,
		container,
		statusUpdater)

	return collection, clues.Stack(errs.Failure()).OrNil()
}
