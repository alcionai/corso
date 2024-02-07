package teamschats

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/pkg/backup/details"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/selectors"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
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
	// corner case: many Topics are empty, and empty inputs are automatically
	// set to non-matching in the selectors code.  This allows us to include
	// everything without needing to check the topic value in that case.
	if scope.IsAny(selectors.TeamsChatsChat) {
		return true
	}

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
