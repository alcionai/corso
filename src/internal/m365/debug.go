package m365

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange"
	"github.com/alcionai/corso/src/internal/m365/collection/groups"
	teamschats "github.com/alcionai/corso/src/internal/m365/collection/teamchats"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/store"
)

func (ctrl *Controller) DeserializeMetadataFiles(
	ctx context.Context,
	colls []data.RestoreCollection,
) ([]store.MetadataFile, error) {
	if len(colls) == 0 {
		return []store.MetadataFile{}, nil
	}

	// assume all collections refer to the same service
	service := colls[0].FullPath().Service()

	switch service {
	case path.ExchangeService, path.ExchangeMetadataService:
		return exchange.DeserializeMetadataFiles(ctx, colls)
	case path.OneDriveService, path.OneDriveMetadataService:
		return drive.DeserializeMetadataFiles(ctx, colls, count.New())
	case path.SharePointService, path.SharePointMetadataService:
		return drive.DeserializeMetadataFiles(ctx, colls, count.New())
	case path.GroupsService, path.GroupsMetadataService:
		return groups.DeserializeMetadataFiles(ctx, colls)
	case path.TeamsChatsService, path.TeamsChatsMetadataService:
		return teamschats.DeserializeMetadataFiles(ctx, colls)
	default:
		return nil, clues.NewWC(ctx, "unrecognized service").With("service", service)
	}
}
