package restore

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// CheckEmailRestoration verifies that the emails count in restored folder is equivalent to
// emails in actual m365 account
func CheckEmailRestoration(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	var (
		folderNameToItemCount        = make(map[string]int32)
		folderNameToRestoreItemCount = make(map[string]int32)
	)

	restoredTree := buildSanitree(ctx, ac, envs.UserID, envs.FolderName)
	dataTree := buildSanitree(ctx, ac, envs.UserID, envs.DataFolder)

	ctx = clues.Add(
		ctx,
		"restore_folder_id", restoredTree.ContainerID,
		"restore_folder_name", restoredTree.ContainerName,
		"original_folder_id", dataTree.ContainerID,
		"original_folder_name", dataTree.ContainerName)

	verifyEmailData(ctx, folderNameToRestoreItemCount, folderNameToItemCount)

	common.AssertEqualTrees[models.MailFolderable](
		ctx,
		dataTree,
		restoredTree.Children[envs.DataFolder])
}

func verifyEmailData(ctx context.Context, restoreMessageCount, messageCount map[string]int32) {
	for fldName, expected := range messageCount {
		got := restoreMessageCount[fldName]

		common.Assert(
			ctx,
			func() bool { return expected == got },
			fmt.Sprintf("Restore item counts do not match: %s", fldName),
			expected,
			got)
	}
}

func buildSanitree(
	ctx context.Context,
	ac api.Client,
	userID, folderName string,
) *common.Sanitree[models.MailFolderable] {
	gcc, err := ac.Mail().GetContainerByName(
		ctx,
		userID,
		api.MsgFolderRoot,
		folderName)
	if err != nil {
		common.Fatal(
			ctx,
			fmt.Sprintf("finding folder by name %q", folderName),
			err)
	}

	mmf, ok := gcc.(models.MailFolderable)
	if !ok {
		common.Fatal(
			ctx,
			"mail folderable required",
			clues.New("casting "+*gcc.GetDisplayName()+" to models.MailFolderable"))
	}

	root := &common.Sanitree[models.MailFolderable]{
		Container:     mmf,
		ContainerID:   ptr.Val(mmf.GetId()),
		ContainerName: ptr.Val(mmf.GetDisplayName()),
		ContainsItems: int(ptr.Val(mmf.GetTotalItemCount())),
		Children:      map[string]*common.Sanitree[models.MailFolderable]{},
	}

	recurseSubfolders(ctx, ac, root, userID)

	return root
}

func recurseSubfolders(
	ctx context.Context,
	ac api.Client,
	parent *common.Sanitree[models.MailFolderable],
	userID string,
) {
	childFolders, err := ac.Mail().GetContainerChildren(
		ctx,
		userID,
		parent.ContainerID)
	if err != nil {
		common.Fatal(ctx, "getting subfolders", err)
	}

	for _, child := range childFolders {
		c := &common.Sanitree[models.MailFolderable]{
			Container:     child,
			ContainerID:   ptr.Val(child.GetId()),
			ContainerName: ptr.Val(child.GetDisplayName()),
			ContainsItems: int(ptr.Val(child.GetTotalItemCount())),
			Children:      map[string]*common.Sanitree[models.MailFolderable]{},
		}

		parent.Children[c.ContainerName] = c

		if ptr.Val(child.GetChildFolderCount()) > 0 {
			recurseSubfolders(ctx, ac, c, userID)
		}
	}
}
