package teamschats

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

var _ backupHandler[models.Chatable] = &usersChatsBackupHandler{}

type usersChatsBackupHandler struct {
	ac                  api.Chats
	protectedResourceID string
	tenantID            string
}

func NewUsersChatsBackupHandler(
	tenantID, protectedResourceID string,
	ac api.Chats,
) usersChatsBackupHandler {
	return usersChatsBackupHandler{
		ac:                  ac,
		protectedResourceID: protectedResourceID,
		tenantID:            tenantID,
	}
}

// chats have no containers.  Everything is stored at the root.
//
//lint:ignore U1000 required for interface compliance
func (bh usersChatsBackupHandler) getContainer(
	ctx context.Context,
	_ api.CallConfig,
) (container[models.Chatable], error) {
	return chatContainer(), nil
}

//lint:ignore U1000 required for interface compliance
func (bh usersChatsBackupHandler) getItemIDs(
	ctx context.Context,
	cc api.CallConfig,
) ([]models.Chatable, error) {
	return bh.ac.GetChats(
		ctx,
		bh.protectedResourceID,
		cc)
}

//lint:ignore U1000 required for interface compliance
func (bh usersChatsBackupHandler) includeItem(
	ch models.Chatable,
	scope selectors.TeamsChatsScope,
) bool {
	return scope.Matches(selectors.TeamsChatsChat, ptr.Val(ch.GetTopic()))
}

func (bh usersChatsBackupHandler) CanonicalPath() (path.Path, error) {
	return path.BuildPrefix(
		bh.tenantID,
		bh.protectedResourceID,
		path.TeamsChatsService,
		path.ChatsCategory)
}

//lint:ignore U1000 false linter issue due to generics
func (bh usersChatsBackupHandler) getItem(
	ctx context.Context,
	userID string,
	chatID string,
) (models.Chatable, *details.TeamsChatsInfo, error) {
	// FIXME: should retrieve and populate all messages in the chat.
	return nil, nil, clues.New("not implemented")
}

//lint:ignore U1000 false linter issue due to generics
func (bh usersChatsBackupHandler) augmentItemInfo(
	dgi *details.TeamsChatsInfo,
	c models.Chatable,
) {
	// no-op
}

func chatContainer() container[models.Chatable] {
	return container[models.Chatable]{
		storageDirFolders: path.Elements{},
		humanLocation:     path.Elements{},
	}
}
