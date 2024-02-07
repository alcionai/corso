package exchange

import (
	"context"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/logger"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
)

var (
	_ graph.ContainerResolver = &contactContainerCache{}
	_ containerRefresher      = &contactRefresher{}
)

type contactRefresher struct {
	getter containerGetter
	userID string
}

func (r *contactRefresher) refreshContainer(
	ctx context.Context,
	id string,
) (graph.CachedContainer, error) {
	c, err := r.getter.GetContainerByID(ctx, r.userID, id)
	if err != nil {
		return nil, clues.Stack(err)
	}

	f := graph.NewCacheFolder(c, nil, nil)

	return &f, nil
}

type contactContainerCache struct {
	*containerResolver
	enumer containersEnumerator[models.ContactFolderable]
	getter containerGetter
	userID string
}

func (cfc *contactContainerCache) init(
	ctx context.Context,
	baseNode string,
	baseContainerPath []string,
) error {
	if len(baseNode) == 0 {
		return clues.NewWC(ctx, "m365 folderID required for base contact folder")
	}

	if cfc.containerResolver == nil {
		cfc.containerResolver = newContainerResolver(&contactRefresher{
			userID: cfc.userID,
			getter: cfc.getter,
		})
	}

	return cfc.populateContactRoot(ctx, baseNode, baseContainerPath)
}

func (cfc *contactContainerCache) populateContactRoot(
	ctx context.Context,
	directoryID string,
	baseContainerPath []string,
) error {
	f, err := cfc.getter.GetContainerByID(ctx, cfc.userID, directoryID)
	if err != nil {
		return clues.Wrap(err, "fetching root folder")
	}

	temp := graph.NewCacheFolder(
		f,
		path.Builder{}.Append(ptr.Val(f.GetId())),   // path of IDs
		path.Builder{}.Append(baseContainerPath...)) // display location
	if err := cfc.addFolder(&temp); err != nil {
		return clues.WrapWC(ctx, err, "adding resolver dir")
	}

	return nil
}

// Populate is utility function for placing cache container
// objects into the Contact Folder Cache
// Function does NOT use Delta Queries as it is not supported
// as of (Oct-07-2022)
func (cfc *contactContainerCache) Populate(
	ctx context.Context,
	errs *fault.Bus,
	baseID string,
	baseContainerPath ...string,
) error {
	start := time.Now()

	logger.Ctx(ctx).Info("populating container cache")

	if err := cfc.init(ctx, baseID, baseContainerPath); err != nil {
		return clues.Wrap(err, "initializing")
	}

	el := errs.Local()

	containers, err := cfc.enumer.EnumerateContainers(
		ctx,
		cfc.userID,
		baseID)
	ctx = clues.Add(ctx, "num_enumerated_containers", len(containers))

	if err != nil {
		return clues.WrapWC(ctx, err, "enumerating containers")
	}

	for _, c := range containers {
		if el.Failure() != nil {
			return el.Failure()
		}

		cacheFolder := graph.NewCacheFolder(c, nil, nil)

		err := cfc.addFolder(&cacheFolder)
		if err != nil {
			errs.AddRecoverable(
				ctx,
				clues.StackWC(ctx, err).Label(fault.LabelForceNoBackupCreation))
		}
	}

	if err := cfc.populatePaths(ctx, errs); err != nil {
		return clues.Wrap(err, "populating paths")
	}

	logger.Ctx(ctx).Infow(
		"done populating container cache",
		"duration", time.Since(start))

	return el.Failure()
}
