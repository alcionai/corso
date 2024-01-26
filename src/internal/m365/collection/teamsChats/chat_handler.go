package teamschats

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

var _ backupHandler[models.Chatable] = &usersChatsBackupHandler{}

type usersChatsBackupHandler struct {
	ac api.Chats
	qp graph.QueryParams
}

func NewUsersChatsBackupHandler(
	qp graph.QueryParams,
	ac api.Chats,
) usersChatsBackupHandler {
	return usersChatsBackupHandler{
		ac: ac,
		qp: qp,
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
) ([]models.Chatable, error) {
	cc := api.CallConfig{
		Expand: []string{"lastMessagePreview"},
	}

	return bh.ac.GetChats(
		ctx,
		bh.qp.ProtectedResource.ID(),
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
		bh.qp.TenantID,
		bh.qp.ProtectedResource.ID(),
		path.TeamsChatsService,
		path.ChatsCategory)
}

//lint:ignore U1000 false linter issue due to generics
func (bh usersChatsBackupHandler) getItem(
	ctx context.Context,
	chat models.Chatable,
) (models.Chatable, *details.TeamsChatsInfo, error) {
	if chat == nil {
		return nil, nil, clues.Stack(core.ErrNotFound)
	}

	chatID := ptr.Val(chat.GetId())

	msgs, err := bh.ac.GetChatMessages(ctx, chatID, api.CallConfig{})
	if err != nil {
		return nil, nil, clues.Stack(err)
	}

	chat.SetMessages(msgs)

	members, err := bh.ac.GetChatMembers(ctx, chatID, api.CallConfig{})
	if err != nil {
		return nil, nil, clues.Stack(err)
	}

	chat.SetMembers(members)

	return chat, api.TeamsChatInfo(chat), nil
}

func chatContainer() container[models.Chatable] {
	return container[models.Chatable]{
		storageDirFolders: path.Elements{},
		humanLocation:     path.Elements{},
	}
}
