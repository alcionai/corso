package groups

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ backupHandler = &channelsBackupHandler{}

type channelsBackupHandler struct {
	ac                api.Channels
	protectedResource string
}

func NewChannelBackupHandler(
	protectedResource string,
	ac api.Channels,
) channelsBackupHandler {
	return channelsBackupHandler{
		ac:                ac,
		protectedResource: protectedResource,
	}
}

func (bh channelsBackupHandler) getChannels(
	ctx context.Context,
) ([]models.Channelable, error) {
	return bh.ac.GetChannels(ctx, bh.protectedResource)
}

func (bh channelsBackupHandler) getChannelMessagesDelta(
	ctx context.Context,
	channelID, prevDelta string,
) ([]models.ChatMessageable, api.DeltaUpdate, error) {
	return bh.ac.GetChannelMessagesDelta(ctx, bh.protectedResource, channelID, prevDelta)
}

func (bh channelsBackupHandler) includeContainer(
	ctx context.Context,
	qp graph.QueryParams,
	ch models.Channelable,
	scope selectors.GroupsScope,
) bool {
	return scope.Matches(selectors.GroupsChannel, ptr.Val(ch.GetDisplayName()))
}

func (bh channelsBackupHandler) canonicalPath(
	folders *path.Builder,
	tenantID string,
) (path.Path, error) {
	return folders.
		ToDataLayerPath(
			tenantID,
			bh.protectedResource,
			path.GroupsService,
			path.ChannelMessagesCategory,
			false)
}
