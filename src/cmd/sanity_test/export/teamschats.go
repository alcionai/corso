package export

import (
	"context"
	"io/fs"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func CheckTeamsChatsExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	sourceTree := populateChatsSanitree(
		ctx,
		ac,
		envs.UserID)

	fpTree := common.BuildFilepathSanitree(ctx, envs.RestoreContainer)

	comparator := func(
		ctx context.Context,
		expect *common.Sanitree[models.Chatable, models.Chatable],
		result *common.Sanitree[fs.FileInfo, fs.FileInfo],
	) {
		for key := range expect.Leaves {
			expect.Leaves[key].Size = 0 // chat sizes cannot be compared
		}

		updatedResultLeaves := map[string]*common.Sanileaf[fs.FileInfo, fs.FileInfo]{}

		for key, leaf := range result.Leaves {
			key = strings.TrimSuffix(key, ".json")
			leaf.Size = 0 // we cannot compare sizes
			updatedResultLeaves[key] = leaf
		}

		common.CompareLeaves(ctx, expect.Leaves, updatedResultLeaves, nil)
	}

	common.CompareDiffTrees(
		ctx,
		sourceTree,
		fpTree.Children[path.ChatsCategory.HumanString()],
		comparator)

	common.Infof(ctx, "Success")
}

func populateChatsSanitree(
	ctx context.Context,
	ac api.Client,
	userID string,
) *common.Sanitree[models.Chatable, models.Chatable] {
	root := &common.Sanitree[models.Chatable, models.Chatable]{
		ID:     userID,
		Name:   path.ChatsCategory.HumanString(),
		Leaves: map[string]*common.Sanileaf[models.Chatable, models.Chatable]{},
		// teamschat should not have child containers
	}

	cc := api.CallConfig{
		Expand: []string{"lastMessagePreview"},
	}

	chats, err := ac.Chats().GetChats(ctx, userID, cc)
	if err != nil {
		common.Fatal(ctx, "getting channels", err)
	}

	for _, chat := range chats {
		leaf := &common.Sanileaf[
			models.Chatable, models.Chatable,
		]{
			Parent: root,
			ID:     ptr.Val(chat.GetId()),
			Name:   ptr.Val(chat.GetId()),
		}

		root.Leaves[ptr.Val(chat.GetId())] = leaf
	}

	root.CountLeaves = len(root.Leaves)

	return root
}
