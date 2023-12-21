package site

import (
	"context"
	"fmt"
	stdpath "path"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// CollectLibraries constructs a onedrive Collections struct and Get()s
// all the drives associated with the site.
func CollectLibraries(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	bh drive.BackupHandler,
	tenantID string,
	ssmb *prefixmatcher.StringSetMatchBuilder,
	su support.StatusUpdater,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, bool, error) {
	logger.Ctx(ctx).Debug("creating SharePoint Library collections")

	var (
		collections = []data.BackupCollection{}
		colls       = drive.NewCollections(
			bh,
			tenantID,
			bpc.ProtectedResource,
			su,
			bpc.Options,
			counter)
	)

	msg := fmt.Sprintf(
		"%s (%s)",
		path.LibrariesCategory.HumanString(),
		stdpath.Base(bpc.ProtectedResource.Name()))

	pcfg := observe.ProgressCfg{
		Indent:            1,
		CompletionMessage: func() string { return fmt.Sprintf("(found %d items)", colls.NumItems) },
	}
	progressBar := observe.MessageWithCompletion(ctx, pcfg, msg)
	close(progressBar)

	odcs, canUsePreviousBackup, err := colls.Get(ctx, bpc.MetadataCollections, ssmb, errs)
	if err != nil {
		return nil, false, graph.Wrap(ctx, err, "getting library")
	}

	return append(collections, odcs...), canUsePreviousBackup, nil
}

// CollectPages constructs a sharepoint Collections struct and Get()s the associated
// M365 IDs for the associated Pages.
func CollectPages(
	ctx context.Context,
	bpc inject.BackupProducerConfig,
	creds account.M365Config,
	ac api.Client,
	scope selectors.SharePointScope,
	su support.StatusUpdater,
	counter *count.Bus,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	logger.Ctx(ctx).Debug("creating SharePoint Pages collections")

	var (
		el   = errs.Local()
		spcs = make([]data.BackupCollection, 0)
	)

	// make the betaClient
	// Need to receive From DataCollection Call
	adpt, err := graph.CreateAdapter(
		creds.AzureTenantID,
		creds.AzureClientID,
		creds.AzureClientSecret,
		counter)
	if err != nil {
		return nil, clues.Wrap(err, "creating azure client adapter")
	}

	betaService := betaAPI.NewBetaService(adpt)

	tuples, err := betaAPI.FetchPages(ctx, betaService, bpc.ProtectedResource.ID())
	if err != nil {
		return nil, err
	}

	for _, tuple := range tuples {
		if el.Failure() != nil {
			break
		}

		dir, err := path.Build(
			creds.AzureTenantID,
			bpc.ProtectedResource.ID(),
			path.SharePointService,
			path.PagesCategory,
			false,
			tuple.Name)
		if err != nil {
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "creating page collection path"))
		}

		collection := NewCollection(
			nil,
			dir,
			ac,
			scope,
			su,
			bpc.Options,
			false,
		)
		collection.SetBetaService(betaService)
		collection.AddItem(tuple.ID, time.Now())

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}

func CollectLists(
	ctx context.Context,
	bh backupHandler,
	bpc inject.BackupProducerConfig,
	ac api.Client,
	tenantID string,
	scope selectors.SharePointScope,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	logger.Ctx(ctx).Debug("Creating SharePoint List Collections")

	var (
		el   = errs.Local()
		spcs = make([]data.BackupCollection, 0)
		acc  = api.CallConfig{Select: idAnd("list", "lastModifiedDateTime")}
	)

	lists, err := bh.GetItems(ctx, acc)
	if err != nil {
		return nil, err
	}

	for _, list := range lists {
		if el.Failure() != nil {
			break
		}

		if api.SkipListTemplates.HasKey(ptr.Val(list.GetList().GetTemplate())) {
			continue
		}

		lmt := list.GetLastModifiedDateTime()

		_, err = dttm.ParseTime(lmt.String())
		validLmt := err == nil

		dir, err := path.Build(
			tenantID,
			bpc.ProtectedResource.ID(),
			path.SharePointService,
			path.ListsCategory,
			false,
			ptr.Val(list.GetId()))
		if err != nil {
			el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "creating list collection path"))
		}

		collection := NewCollection(
			bh,
			dir,
			ac,
			scope,
			su,
			bpc.Options,
			validLmt,
		)
		collection.AddItem(
			ptr.Val(list.GetId()),
			ptr.Val(lmt))

		spcs = append(spcs, collection)
	}

	return spcs, el.Failure()
}

func idAnd(ss ...string) []string {
	id := []string{"id"}

	if len(ss) == 0 {
		return id
	}

	return append(id, ss...)
}
