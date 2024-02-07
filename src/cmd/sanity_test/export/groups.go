package export

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/cmd/sanity_test/common"
	"github.com/alcionai/canario/src/cmd/sanity_test/driveish"
	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
)

func CheckGroupsExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	// assumes we only need to sanity check the default site.
	// should we expand this to check all sites in the group?
	// are we backing up / restoring more than the default site?
	site, err := ac.Sites().GetByID(ctx, envs.TeamSiteID, api.CallConfig{})
	if err != nil {
		common.Fatal(ctx, "getting the drive:", err)
	}

	drive, err := ac.Sites().GetDefaultDrive(ctx, envs.TeamSiteID)
	if err != nil {
		common.Fatal(ctx, "getting the drive:", err)
	}

	checkChannelMessagesExport(
		ctx,
		ac,
		envs)

	envs.RestoreContainer = filepath.Join(
		envs.RestoreContainer,
		"Libraries",
		ptr.Val(site.GetName()),
		"Documents") // check in default loc
	driveish.CheckExport(
		ctx,
		ac,
		drive,
		envs)
}

func removeDeletedMessages(
	ctx context.Context,
	tree *common.Sanitree[models.Channelable, models.ChatMessageable],
) {
	for _, child := range tree.Children {
		for _, leaf := range child.Leaves {
			if leaf.Self.GetDeletedDateTime() != nil {
				delete(child.Leaves, leaf.ID)
				child.CountLeaves--
			}
		}
	}
}

func checkChannelMessagesExport(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	sourceTree := populateMessagesSanitree(
		ctx,
		ac,
		envs.GroupID)

	fpTree := common.BuildFilepathSanitree(ctx, envs.RestoreContainer)

	comparator := func(
		ctx context.Context,
		expect *common.Sanitree[models.Channelable, models.ChatMessageable],
		result *common.Sanitree[fs.FileInfo, fs.FileInfo],
	) {
		for key := range expect.Leaves {
			expect.Leaves[key].Size = 0 // msg sizes cannot be compared
		}

		updatedResultLeaves := map[string]*common.Sanileaf[fs.FileInfo, fs.FileInfo]{}

		for key, leaf := range result.Leaves {
			key = strings.TrimSuffix(key, ".json")
			leaf.Size = 0 // we cannot compare sizes
			updatedResultLeaves[key] = leaf
		}

		common.CompareLeaves(ctx, expect.Leaves, updatedResultLeaves, nil)
	}

	removeDeletedMessages(ctx, sourceTree)

	common.CompareDiffTrees(
		ctx,
		sourceTree,
		fpTree.Children["Messages"],
		comparator)

	common.Infof(ctx, "Success")
}

func populateMessagesSanitree(
	ctx context.Context,
	ac api.Client,
	groupID string,
) *common.Sanitree[models.Channelable, models.ChatMessageable] {
	root := &common.Sanitree[models.Channelable, models.ChatMessageable]{
		ID:   groupID,
		Name: path.ChannelMessagesCategory.HumanString(),
		// group should not have leaves
		Children: map[string]*common.Sanitree[models.Channelable, models.ChatMessageable]{},
	}

	channels, err := ac.Channels().GetChannels(ctx, groupID)
	if err != nil {
		common.Fatal(ctx, "getting channels", err)
	}

	for _, ch := range channels {
		child := &common.Sanitree[
			models.Channelable, models.ChatMessageable,
		]{
			Parent: root,
			ID:     ptr.Val(ch.GetId()),
			Name:   ptr.Val(ch.GetDisplayName()),
			Leaves: map[string]*common.Sanileaf[models.Channelable, models.ChatMessageable]{},
			// no children in channels
		}

		msgs, err := ac.Channels().GetChannelMessages(
			ctx,
			groupID,
			ptr.Val(ch.GetId()),
			api.CallConfig{
				// include all nessage replies in each message
				Expand: []string{"replies"},
			})
		if err != nil {
			common.Fatal(ctx, "getting channel messages", err)
		}

		filteredMsgs := []models.ChatMessageable{}

		for _, msg := range msgs {
			// filter out system messages (we don't really work with them)
			if api.IsNotSystemMessage(msg) {
				filteredMsgs = append(filteredMsgs, msg)
			}
		}

		if len(filteredMsgs) == 0 {
			common.Infof(ctx, "skipped empty channel: %s", ptr.Val(ch.GetDisplayName()))
			continue
		}

		for _, msg := range filteredMsgs {
			child.CountLeaves++
			child.Leaves[ptr.Val(msg.GetId())] = &common.Sanileaf[
				models.Channelable,
				models.ChatMessageable,
			]{
				Self: msg,
				ID:   ptr.Val(msg.GetId()),
				Name: ptr.Val(msg.GetId()),         // channel messages have no display name
				Size: int64(len(msg.GetReplies())), // size is the count of replies
			}
		}

		root.Children[ptr.Val(ch.GetDisplayName())] = child
	}

	return root
}
