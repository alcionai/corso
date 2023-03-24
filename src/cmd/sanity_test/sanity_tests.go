package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/alcionai/clues"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/logger"
)

func main() {
	ctx, log := logger.Seed(context.Background(), "info", logger.GetLogFile(""))
	defer func() {
		_ = log.Sync() // flush all logs in the buffer
	}()

	adapter, err := graph.CreateAdapter(
		os.Getenv("AZURE_TENANT_ID"),
		os.Getenv("AZURE_CLIENT_ID"),
		os.Getenv("AZURE_CLIENT_SECRET"))
	if err != nil {
		fatal(ctx, "creating adapter", err)
	}

	var (
		client      = msgraphsdk.NewGraphServiceClient(adapter)
		testUser    = os.Getenv("CORSO_M365_TEST_USER_ID")
		testService = os.Getenv("SANITY_RESTORE_SERVICE")
		folder      = strings.TrimSpace(os.Getenv("SANITY_RESTORE_FOLDER"))
		startTime   = mustGetTimeFromName(ctx, folder)
	)

	ctx = clues.Add(
		ctx,
		"resource_owner", testUser,
		"service", testService,
		"sanity_restore_folder", folder,
		"start_time", startTime.Format(time.RFC3339Nano))

	logger.Ctx(ctx).Info("starting sanity test check")

	switch testService {
	case "exchange":
		checkEmailRestoration(ctx, client, testUser, folder, startTime)
	case "onedrive":
		checkOnedriveRestoration(ctx, client, testUser, folder, startTime)
	default:
		fatal(ctx, "no service specified", nil)
	}
}

// checkEmailRestoration verifies that the emails count in restored folder is equivalent to
// emails in actual m365 account
func checkEmailRestoration(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	testUser, folderName string,
	startTime time.Time,
) {
	var (
		itemCount     = make(map[string]int32)
		restoreFolder models.MailFolderable
		builder       = client.UsersById(testUser).MailFolders()
	)

	for {
		result, err := builder.Get(ctx, nil)
		if err != nil {
			fatal(ctx, "getting mail folders", err)
		}

		values := result.GetValue()

		// recursive restore folder discovery before proceeding with tests
		for _, v := range values {
			var (
				itemID     = ptr.Val(v.GetId())
				itemName   = ptr.Val(v.GetDisplayName())
				ictx       = clues.Add(ctx, "item_id", itemID, "item_name", itemName)
				folderTime = mustGetTimeFromName(ctx, itemName)
			)

			// only test against folders within the test boundary time
			if !errors.Is(err, common.ErrNoTimeString) && startTime.Before(folderTime) {
				logger.Ctx(ictx).
					With("folder_time", folderTime).
					Info("skipping restore folder: not within time bound")

				continue
			}

			// if we found the folder to testt against, back out of this loop.
			if itemName == folderName {
				restoreFolder = v
				continue
			}

			// otherwise, recursively aggregate all child folders.
			getAllSubFolder(ctx, client, testUser, v, itemName, itemCount)

			itemCount[itemName] = ptr.Val(v.GetTotalItemCount())
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
		UsersById(testUser).
		MailFoldersById(folderID).
		ChildFolders().
		Get(ctx, nil)
	if err != nil {
		fatal(ctx, "getting restore folder child folders", err)
	}

	for _, fld := range childFolder.GetValue() {
		var (
			fldID   = ptr.Val(fld.GetId())
			fldName = ptr.Val(fld.GetDisplayName())
			count   = ptr.Val(fld.GetTotalItemCount())
			ictx    = clues.Add(
				ctx,
				"child_folder_id", fldID,
				"child_folder_name", fldName,
				"expected_count", itemCount[fldName],
				"actual_count", count)
		)

		if itemCount[fldName] != count {
			logger.Ctx(ictx).Error("test failure: Restore item counts do not match")
			fmt.Println("Restore item counts do not match:")
			fmt.Println("*  expected:", itemCount[fldName])
			fmt.Println("*  actual:", count)
			fmt.Println("Folder:", fldName, ptr.Val(fld.GetId()))
			os.Exit(1)
		}

		checkAllSubFolder(ctx, client, testUser, fld, fldName, itemCount)
	}
}

// getAllSubFolder will recursively check for all subfolders and get the corresponding
// email count.
func getAllSubFolder(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	testUser string,
	r models.MailFolderable,
	parentFolder string,
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
		UsersById(testUser).
		MailFoldersById(folderID).
		ChildFolders().
		Get(ctx, options)
	if err != nil {
		fatal(ctx, "getting mail subfolders", err)
	}

	for _, child := range childFolder.GetValue() {
		var (
			childDisplayName = ptr.Val(child.GetDisplayName())
			childFolderCount = ptr.Val(child.GetChildFolderCount())
			fullFolderName   = parentFolder + "/" + childDisplayName
		)

		messageCount[fullFolderName], _ = ptr.ValOK(child.GetTotalItemCount())

		// recursively check for subfolders
		if childFolderCount > 0 {
			parentFolder := fullFolderName

			getAllSubFolder(ctx, client, testUser, child, parentFolder, messageCount)
		}
	}
}

// checkAllSubFolder will recursively traverse inside the restore folder and
// verify that data matched in all subfolders
func checkAllSubFolder(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	testUser string,
	r models.MailFolderable,
	parentFolder string,
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

	childFolder, err := client.
		UsersById(testUser).
		MailFoldersById(folderID).
		ChildFolders().
		Get(ctx, options)
	if err != nil {
		fatal(ctx, "getting mail subfolders", err)
	}

	for _, child := range childFolder.GetValue() {
		var (
			childDisplayName = ptr.Val(child.GetDisplayName())
			childTotalCount  = ptr.Val(child.GetTotalItemCount())
			//nolint:forbidigo
			fullFolderName = path.Join(parentFolder, childDisplayName)
		)

		if messageCount[fullFolderName] != childTotalCount {
			fmt.Println("Message count doesn't match:")
			fmt.Println("*  expected:", messageCount[fullFolderName])
			fmt.Println("*  actual:", childTotalCount)
			fmt.Println("Item:", fullFolderName, folderID)
			os.Exit(1)
		}

		childFolderCount := ptr.Val(child.GetChildFolderCount())

		if childFolderCount > 0 {
			checkAllSubFolder(ctx, client, testUser, child, fullFolderName, messageCount)
		}
	}
}

func checkOnedriveRestoration(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	testUser, folderName string,
	startTime time.Time,
) {
	var (
		// map itemID -> item size
		fileSizes = make(map[string]int64)
		// map itemID -> permission id -> []permission roles
		folderPermission = make(map[string]map[string][]string)
	)

	drive, err := client.
		UsersById(testUser).
		Drive().
		Get(ctx, nil)
	if err != nil {
		fatal(ctx, "getting the drive:", err)
	}

	var (
		driveID         = ptr.Val(drive.GetId())
		driveName       = ptr.Val(drive.GetName())
		restoreFolderID string
	)

	ctx = clues.Add(ctx, "drive_id", driveID, "drive_name", driveName)

	response, err := client.
		DrivesById(driveID).
		Root().
		Children().
		Get(ctx, nil)
	if err != nil {
		fatal(ctx, "getting drive by id", err)
	}

	for _, driveItem := range response.GetValue() {
		var (
			itemID   = ptr.Val(driveItem.GetId())
			itemName = ptr.Val(driveItem.GetName())
			ictx     = clues.Add(ctx, "item_id", itemID, "item_name", itemName)
		)

		if itemName == folderName {
			restoreFolderID = itemID
			continue
		}

		folderTime := mustGetTimeFromName(ictx, itemName)

		if !errors.Is(err, common.ErrNoTimeString) && startTime.Before(folderTime) {
			logger.Ctx(ictx).
				With("folder_time", folderTime).
				Info("skipping restore folder: not within time bound")

			continue
		}

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			fileSizes[itemName] = ptr.Val(driveItem.GetSize())
		}

		folderPermission[itemID] = permissionsIn(ctx, client, driveID, itemID, folderPermission[itemID])
	}

	checkFileData(ctx, client, driveID, restoreFolderID, fileSizes, folderPermission)

	fmt.Println("Success")
}

func checkFileData(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	driveID,
	restoreFolderID string,
	fileSizes map[string]int64,
	folderPermission map[string]map[string][]string,
) {
	restored, err := client.
		DrivesById(driveID).
		ItemsById(restoreFolderID).
		Children().
		Get(ctx, nil)
	if err != nil {
		fatal(ctx, "getting child folder", err)
	}

	for _, item := range restored.GetValue() {
		var (
			itemID   = ptr.Val(item.GetId())
			itemName = ptr.Val(item.GetName())
			itemSize = ptr.Val(item.GetSize())
		)

		if item.GetFile() != nil {
			if itemSize != fileSizes[itemName] {
				fmt.Println("File size does not match:")
				fmt.Println("*  expected:", fileSizes[itemName])
				fmt.Println("*  actual:", itemSize)
				fmt.Println("Item:", itemName, itemID)
				os.Exit(1)
			}

			continue
		}

		if item.GetFolder() == nil && item.GetPackage() == nil {
			continue
		}

		var (
			expectItem = folderPermission[itemID]
			results    = permissionsIn(ctx, client, driveID, itemID, nil)
		)

		for pid, result := range results {
			expect := expectItem[pid]

			if !slices.Equal(expect, result) {
				fmt.Println("permissions are not equal")
				fmt.Println("*  expected: ", expect)
				fmt.Println("*  actual: ", result)
				fmt.Println("Item:", itemName, itemID)
				fmt.Println("Permission:", pid)
				os.Exit(1)
			}
		}
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func fatal(ctx context.Context, msg string, err error) {
	logger.CtxErr(ctx, err).Error("test failure: " + msg)
	fmt.Println(msg+": ", err)
	os.Exit(1)
}

func permissionsIn(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	driveID, itemID string,
	init map[string][]string,
) map[string][]string {
	result := map[string][]string{}

	pcr, err := client.
		DrivesById(driveID).
		ItemsById(itemID).
		Permissions().
		Get(ctx, nil)
	if err != nil {
		fatal(ctx, "getting permission", err)
	}

	if len(init) > 0 {
		maps.Copy(result, init)
	}

	for _, p := range pcr.GetValue() {
		var (
			pid   = ptr.Val(p.GetId())
			roles = p.GetRoles()
		)

		slices.Sort(roles)

		result[pid] = roles
	}

	return result
}

func mustGetTimeFromName(ctx context.Context, name string) time.Time {
	t, err := common.ExtractTime(name)
	if err != nil && !errors.Is(err, common.ErrNoTimeString) {
		fatal(ctx, "extracting time from name: "+name, err)
	}

	return t
}
