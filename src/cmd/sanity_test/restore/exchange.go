package restore

import (
	"context"
	"fmt"
	stdpath "path"
	"strings"
	"time"

	"github.com/alcionai/clues"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/filters"
)

// CheckEmailRestoration verifies that the emails count in restored folder is equivalent to
// emails in actual m365 account
func CheckEmailRestoration(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	testUser, folderName, dataFolder, baseBackupFolder string,
	startTime time.Time,
) {
	var (
		restoreFolder    models.MailFolderable
		itemCount        = make(map[string]int32)
		restoreItemCount = make(map[string]int32)
		builder          = client.Users().ByUserId(testUser).MailFolders()
	)

	for {
		result, err := builder.Get(ctx, nil)
		if err != nil {
			common.Fatal(ctx, "getting mail folders", err)
		}

		values := result.GetValue()

		for _, v := range values {
			itemName := ptr.Val(v.GetDisplayName())

			if itemName == folderName {
				restoreFolder = v
				continue
			}

			if itemName == dataFolder || itemName == baseBackupFolder {
				// otherwise, recursively aggregate all child folders.
				getAllMailSubFolders(ctx, client, testUser, v, itemName, dataFolder, itemCount)

				itemCount[itemName] = ptr.Val(v.GetTotalItemCount())
			}
		}

		link, ok := ptr.ValOK(result.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemMailFoldersRequestBuilder(link, client.GetAdapter())
	}

	folderID := ptr.Val(restoreFolder.GetId())
	folderName = ptr.Val(restoreFolder.GetDisplayName())
	ctx = clues.Add(
		ctx,
		"restore_folder_id", folderID,
		"restore_folder_name", folderName)

	childFolder, err := client.
		Users().
		ByUserId(testUser).
		MailFolders().
		ByMailFolderId(folderID).
		ChildFolders().
		Get(ctx, nil)
	if err != nil {
		common.Fatal(ctx, "getting restore folder child folders", err)
	}

	for _, fld := range childFolder.GetValue() {
		restoreDisplayName := ptr.Val(fld.GetDisplayName())

		// check if folder is the data folder we loaded or the base backup to verify
		// the incremental backup worked fine
		if strings.EqualFold(restoreDisplayName, dataFolder) || strings.EqualFold(restoreDisplayName, baseBackupFolder) {
			count, _ := ptr.ValOK(fld.GetTotalItemCount())

			restoreItemCount[restoreDisplayName] = count
			checkAllSubFolder(ctx, client, fld, testUser, restoreDisplayName, dataFolder, restoreItemCount)
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
	client *msgraphsdk.GraphServiceClient,
	testUser string,
	r models.MailFolderable,
	parentFolder,
	dataFolder string,
	messageCount map[string]int32,
) {
	var (
		folderID       = ptr.Val(r.GetId())
		count    int32 = 99
		options        = &users.ItemMailFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMailFoldersItemChildFoldersRequestBuilderGetQueryParameters{
				Top: &count,
			},
		}
	)

	ctx = clues.Add(ctx, "parent_folder_id", folderID)

	childFolder, err := client.
		Users().
		ByUserId(testUser).
		MailFolders().
		ByMailFolderId(folderID).
		ChildFolders().
		Get(ctx, options)
	if err != nil {
		common.Fatal(ctx, "getting mail subfolders", err)
	}

	for _, child := range childFolder.GetValue() {
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

				getAllMailSubFolders(ctx, client, testUser, child, parentFolder, dataFolder, messageCount)
			}
		}
	}
}

// checkAllSubFolder will recursively traverse inside the restore folder and
// verify that data matched in all subfolders
func checkAllSubFolder(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	r models.MailFolderable,
	testUser,
	parentFolder,
	dataFolder string,
	restoreMessageCount map[string]int32,
) {
	var (
		folderID       = ptr.Val(r.GetId())
		count    int32 = 99
		options        = &users.ItemMailFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMailFoldersItemChildFoldersRequestBuilderGetQueryParameters{
				Top: &count,
			},
		}
	)

	childFolder, err := client.
		Users().
		ByUserId(testUser).
		MailFolders().
		ByMailFolderId(folderID).
		ChildFolders().
		Get(ctx, options)
	if err != nil {
		common.Fatal(ctx, "getting mail subfolders", err)
	}

	for _, child := range childFolder.GetValue() {
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
			checkAllSubFolder(ctx, client, child, testUser, parentFolder, dataFolder, restoreMessageCount)
		}
	}
}
