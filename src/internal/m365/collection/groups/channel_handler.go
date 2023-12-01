package groups

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

var _ backupHandler[models.Channelable, models.ChatMessageable] = &channelsBackupHandler{}

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

func (bh channelsBackupHandler) canMakeDeltaQueries() bool {
	return true
}

//lint:ignore U1000 required for interface compliance
func (bh channelsBackupHandler) getContainers(
	ctx context.Context,
	_ api.CallConfig,
) ([]container[models.Channelable], error) {
	chans, err := bh.ac.GetChannels(ctx, bh.protectedResource)
	results := make([]container[models.Channelable], 0, len(chans))

	for _, ch := range chans {
		results = append(results, channelContainer(ch))
	}

	return results, clues.Stack(err).OrNil()
}

func (bh channelsBackupHandler) getContainerItemIDs(
	ctx context.Context,
	containerPath path.Elements,
	prevDelta string,
	cc api.CallConfig,
) (pagers.AddedAndRemoved, error) {
	return bh.ac.GetChannelMessageIDs(
		ctx,
		bh.protectedResource,
		containerPath[0],
		prevDelta,
		cc)
}

//lint:ignore U1000 required for interface compliance
func (bh channelsBackupHandler) includeContainer(
	ch models.Channelable,
	scope selectors.GroupsScope,
) bool {
	return scope.Matches(selectors.GroupsChannel, ptr.Val(ch.GetDisplayName()))
}

func (bh channelsBackupHandler) canonicalPath(
	storageDirFolders path.Elements,
	tenantID string,
) (path.Path, error) {
	return storageDirFolders.
		Builder().
		ToDataLayerPath(
			tenantID,
			bh.protectedResource,
			path.GroupsService,
			path.ChannelMessagesCategory,
			false)
}

func (bh channelsBackupHandler) PathPrefix(tenantID string) (path.Path, error) {
	return path.Build(
		tenantID,
		bh.protectedResource,
		path.GroupsService,
		path.ChannelMessagesCategory,
		false)
}

func (bh channelsBackupHandler) GetItem(
	ctx context.Context,
	groupID string,
	containerIDs path.Elements,
	messageID string,
) (models.ChatMessageable, *details.GroupsInfo, error) {
	return bh.ac.GetChannelMessage(ctx, groupID, containerIDs[0], messageID)
}

func channelContainer(ch models.Channelable) container[models.Channelable] {
	return container[models.Channelable]{
		storageDirFolders:   path.Elements{ptr.Val(ch.GetId())},
		humanLocation:       path.Elements{ptr.Val(ch.GetDisplayName())},
		canMakeDeltaQueries: len(ptr.Val(ch.GetEmail())) > 0,
		container:           ch,
	}
}
