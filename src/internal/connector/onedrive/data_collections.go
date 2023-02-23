package onedrive

import (
	"context"
	"net/http"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
	"golang.org/x/exp/maps"
)

// ---------------------------------------------------------------------------
// OneDrive
// ---------------------------------------------------------------------------

type odFolderMatcher struct {
	scope selectors.OneDriveScope
}

func (fm odFolderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.OneDriveFolder)
}

func (fm odFolderMatcher) Matches(dir string) bool {
	return fm.scope.Matches(selectors.OneDriveFolder, dir)
}

// OneDriveDataCollections returns a set of DataCollection which represents the OneDrive data
// for the specified user
func DataCollections(
	ctx context.Context,
	selector selectors.Selector,
	metadata []data.RestoreCollection,
	tenant string,
	itemClient *http.Client,
	service graph.Servicer,
	su support.StatusUpdater,
	ctrlOpts control.Options,
) ([]data.BackupCollection, map[string]struct{}, error) {
	odb, err := selector.ToOneDriveBackup()
	if err != nil {
		return nil, nil, clues.Wrap(err, "parsing selector").WithClues(ctx)
	}

	var (
		user        = selector.DiscreteOwner
		collections = []data.BackupCollection{}
		allExcludes = map[string]struct{}{}
	)

	// for each scope that includes oneDrive items, get all
	for _, scope := range odb.Scopes() {
		logger.Ctx(ctx).With("user", user).Debug("Creating OneDrive collections")

		odcs, excludes, err := NewCollections(
			itemClient,
			tenant,
			user,
			OneDriveSource,
			odFolderMatcher{scope},
			service,
			su,
			ctrlOpts,
		).Get(ctx, metadata)
		if err != nil {
			return nil, nil, err
		}

		collections = append(collections, odcs...)

		maps.Copy(allExcludes, excludes)
	}

	return collections, allExcludes, nil
}
