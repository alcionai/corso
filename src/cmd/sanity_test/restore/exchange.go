package restore

import (
	"context"
	"fmt"
	stdpath "path"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/filters"
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
		restoreFolder    models.MailFolderable
		itemCount        = make(map[string]int32)
		restoreItemCount = make(map[string]int32)
	)

	fn := func(gcc graph.CachedContainer) error {
		mmf, ok := gcc.(models.MailFolderable)
		if !ok {
			return clues.New("mail folderable required")
		}

		itemName := ptr.Val(mmf.GetDisplayName())

		if itemName == envs.FolderName {
			restoreFolder = mmf
			return nil
		}

		if itemName == envs.DataFolder || itemName == envs.BaseBackupFolder {
			// otherwise, recursively aggregate all child folders.
			getAllMailSubFolders(
				ctx,
				ac,
				envs.UserID,
				mmf,
				itemName,
				envs.DataFolder,
				itemCount)

			itemCount[itemName] = ptr.Val(mmf.GetTotalItemCount())
		}

		return nil
	}

	err := ac.Mail().EnumerateContainers(
		ctx,
		envs.UserID,
		"",
		false,
		fn,
		fault.New(true))
	if err != nil {
		common.Fatal(ctx, "getting all mail folders", err)
	}

	folderID := ptr.Val(restoreFolder.GetId())
	folderName := ptr.Val(restoreFolder.GetDisplayName())
	ctx = clues.Add(
		ctx,
		"restore_folder_id", folderID,
		"restore_folder_name", folderName)

	childFolders, err := ac.Mail().GetContainerChildren(ctx, envs.UserID, folderID)
	if err != nil {
		common.Fatal(ctx, "getting restore folder child folders", err)
	}

	for _, fld := range childFolders {
		restoreDisplayName := ptr.Val(fld.GetDisplayName())

		// check if folder is the data folder we loaded or the base backup to verify
		// the incremental backup worked fine
		if strings.EqualFold(restoreDisplayName, envs.DataFolder) ||
			strings.EqualFold(restoreDisplayName, envs.BaseBackupFolder) {
			count, _ := ptr.ValOK(fld.GetTotalItemCount())

			restoreItemCount[restoreDisplayName] = count
			checkAllSubFolder(
				ctx,
				ac,
				fld,
				envs.UserID,
				restoreDisplayName,
				envs.DataFolder,
				restoreItemCount)
		}
	}

	verifyEmailData(ctx, restoreItemCount, itemCount)
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

// getAllSubFolder will recursively check for all subfolders and get the corresponding
// email count.
func getAllMailSubFolders(
	ctx context.Context,
	ac api.Client,
	testUser string,
	r models.MailFolderable,
	parentFolder,
	dataFolder string,
	messageCount map[string]int32,
) {
	folderID := ptr.Val(r.GetId())
	ctx = clues.Add(ctx, "parent_folder_id", folderID)

	childFolders, err := ac.Mail().GetContainerChildren(ctx, testUser, folderID)
	if err != nil {
		common.Fatal(ctx, "getting mail subfolders", err)
	}

	for _, child := range childFolders {
		var (
			childDisplayName = ptr.Val(child.GetDisplayName())
			childFolderCount = ptr.Val(child.GetChildFolderCount())
			//nolint:forbidigo
			fullFolderName = stdpath.Join(parentFolder, childDisplayName)
		)

		if filters.PathContains([]string{dataFolder}).Compare(fullFolderName) {
			messageCount[fullFolderName] = ptr.Val(child.GetTotalItemCount())
			// recursively check for subfolders
			if childFolderCount > 0 {
				parentFolder := fullFolderName

				getAllMailSubFolders(
					ctx,
					ac,
					testUser,
					child,
					parentFolder,
					dataFolder,
					messageCount)
			}
		}
	}
}

// checkAllSubFolder will recursively traverse inside the restore folder and
// verify that data matched in all subfolders
func checkAllSubFolder(
	ctx context.Context,
	ac api.Client,
	r models.MailFolderable,
	testUser, parentFolder, dataFolder string,
	restoreMessageCount map[string]int32,
) {
	folderID := ptr.Val(r.GetId())

	childFolders, err := ac.Mail().GetContainerChildren(ctx, testUser, folderID)
	if err != nil {
		common.Fatal(ctx, "getting mail subfolders", err)
	}

	for _, child := range childFolders {
		var (
			childDisplayName = ptr.Val(child.GetDisplayName())
			//nolint:forbidigo
			fullFolderName = stdpath.Join(parentFolder, childDisplayName)
		)

		if filters.PathContains([]string{dataFolder}).Compare(fullFolderName) {
			childTotalCount, _ := ptr.ValOK(child.GetTotalItemCount())
			restoreMessageCount[fullFolderName] = childTotalCount
		}

		childFolderCount := ptr.Val(child.GetChildFolderCount())

		if childFolderCount > 0 {
			parentFolder := fullFolderName
			checkAllSubFolder(
				ctx,
				ac,
				child,
				testUser,
				parentFolder,
				dataFolder,
				restoreMessageCount)
		}
	}
}
