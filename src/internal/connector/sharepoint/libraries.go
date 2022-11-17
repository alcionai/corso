package sharepoint

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
)

func CollectLibraries(
	ctx context.Context,
	serv graph.Service,
	tenantID string,
	siteIDs []string,
	scope selectors.SharePointScope,
	updater support.StatusUpdater,
	incrementWaitCount func(),
) ([]data.Collection, error) {
	var (
		collections = []data.Collection{}
		errs        error
	)

	for _, site := range scope.Get(selectors.SharePointSite) {
		logger.Ctx(ctx).With("site", site).Debug("Creating SharePoint Library collections")

		colls := onedrive.NewCollections(
			tenantID,
			site,
			onedrive.SharePointSource,
			folderMatcher{scope},
			serv,
			updater,
		)

		odcs, err := colls.Get(ctx)
		if err != nil {
			return nil, support.WrapAndAppend(site, err, errs)
		}

		collections = append(collections, odcs...)
	}

	for range collections {
		incrementWaitCount()
	}

	return collections, errs
}

type folderMatcher struct {
	scope selectors.SharePointScope
}

func (fm folderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.SharePointFolder)
}

func (fm folderMatcher) Matches(path string) bool {
	return fm.scope.Matches(selectors.SharePointFolder, path)
}
