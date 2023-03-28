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
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/logger"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"golang.org/x/exp/slices"
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
		client           = msgraphsdk.NewGraphServiceClient(adapter)
		testUser         = os.Getenv("CORSO_M365_TEST_USER_ID")
		testService      = os.Getenv("SANITY_RESTORE_SERVICE")
		folder           = strings.TrimSpace(os.Getenv("SANITY_RESTORE_FOLDER"))
		startTime, _     = mustGetTimeFromName(ctx, folder)
		dataFolder       = os.Getenv("TEST_DATA")
		baseBackupFolder = os.Getenv("BASE_BACKUP")
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
		checkEmailRestoration(ctx, client, testUser, folder, dataFolder, baseBackupFolder, startTime)
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
	testUser, folderName, dataFolder, baseBackupFolder string,
	startTime time.Time,
) {
	var (
		restoreFolder    models.MailFolderable
		itemCount        = make(map[string]int32)
		restoreItemCount = make(map[string]int32)
		builder          = client.UsersById(testUser).MailFolders()
	)

	for {
		result, err := builder.Get(ctx, nil)
		if err != nil {
			fatal(ctx, "getting mail folders", err)
		}

		values := result.GetValue()

		for _, v := range values {
			itemName := ptr.Val(v.GetDisplayName())

			if itemName == folderName {
				restoreFolder = v
				continue
			}

			if itemName == dataFolder || itemName == baseBackupFolder {
				getAllSubFolder(ctx, client, testUser, v, itemName, dataFolder, itemCount)

				itemCount[itemName], _ = ptr.ValOK(v.GetTotalItemCount())
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
		UsersById(testUser).
		MailFoldersById(folderID).
		ChildFolders().
		Get(ctx, nil)
	if err != nil {
		fatal(ctx, "getting restore folder child folders", err)
	}

	for _, fld := range childFolder.GetValue() {
		restoreDisplayName, ok := ptr.ValOK(fld.GetDisplayName())
		if !ok {
			fmt.Println("display name not found. Will continue")
			continue
		}

		// check if folder is the data folder we loaded or the base backup to verify
		// the incremental backup worked fine
		if strings.EqualFold(restoreDisplayName, dataFolder) || strings.EqualFold(restoreDisplayName, baseBackupFolder) {
			count, _ := ptr.ValOK(fld.GetTotalItemCount())

			restoreItemCount[restoreDisplayName] = count
			checkAllSubFolder(ctx, client, testUser, fld, restoreDisplayName, dataFolder, restoreItemCount)
		}
	}

	verifyEmailData(ctx, restoreItemCount, itemCount)
}

func verifyEmailData(ctx context.Context, restoreMessageCount, messageCount map[string]int32) {
	for fldName, emailCount := range messageCount {
		if restoreMessageCount[fldName] != emailCount {
			logger.Ctx(ctx).Error("test failure: Restore item counts do not match")
			fmt.Println("Restore item counts do not match:")
			fmt.Println("*  expected:", emailCount)
			fmt.Println("*  actual:", restoreMessageCount[fldName])
			os.Exit(1)
		}
	}
}

// getAllSubFolder will recursively check for all subfolders and get the corresponding
// email count.
func getAllSubFolder(
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

		if strings.Contains(fullFolderName, dataFolder) {
			messageCount[fullFolderName] = ptr.Val(child.GetTotalItemCount())
			// recursively check for subfolders
			if childFolderCount > 0 {
				parentFolder := fullFolderName

				getAllSubFolder(ctx, client, testUser, child, parentFolder, dataFolder, messageCount)
			}
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
			//nolint:forbidigo
			fullFolderName = path.Join(parentFolder, childDisplayName)
		)

		if strings.Contains(fullFolderName, dataFolder) {
			childTotalCount, _ := ptr.ValOK(child.GetTotalItemCount())
			restoreMessageCount[fullFolderName] = childTotalCount
		}

		childFolderCount := ptr.Val(child.GetChildFolderCount())

		if childFolderCount > 0 {
			parentFolder := fullFolderName
			checkAllSubFolder(ctx, client, testUser, child, parentFolder, dataFolder, restoreMessageCount)
		}
	}
}

type permissionInfo struct {
	entityID string
	roles    []string
}

func checkOnedriveRestoration(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	testUser,
	folderName string,
	startTime time.Time,
) {
	var (
		// map itemID -> item size
		fileSizes = make(map[string]int64)
		// map itemID -> permission id -> []permission roles
		folderPermission        = make(map[string][]permissionInfo)
		restoreFile             = make(map[string]int64)
		restoreFolderPermission = make(map[string][]permissionInfo)
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

		folderTime, hasTime := mustGetTimeFromName(ictx, itemName)

		if !isWithinTimeBound(ctx, startTime, folderTime, hasTime) {
			continue
		}

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			fileSizes[itemName] = ptr.Val(driveItem.GetSize())
		}

		if driveItem.GetFolder() == nil && driveItem.GetPackage() == nil {
			continue
		}

		permissionIn(ctx, client, driveID, itemID, itemName, folderPermission)
		getOneDriveChildFolder(ctx, client, driveID, itemID, itemName, fileSizes, folderPermission)
	}

	getRestoreData(ctx, client, *drive.GetId(), restoreFolderID, restoreFile, restoreFolderPermission)

	for checkFolderName, checkfolderPer := range folderPermission {
		fmt.Printf("checking for folder: %s", checkFolderName)

		for i, orginalFolderPer := range checkfolderPer {
			if !(orginalFolderPer.entityID != restoreFolderPermission[checkFolderName][i].entityID) &&
				!slices.Equal(orginalFolderPer.roles, restoreFolderPermission[checkFolderName][i].roles) {
				fmt.Printf("permissions are not equal")
				fmt.Printf("*  expected role: %+v \n", orginalFolderPer.roles)
				fmt.Printf("*  actual:  %+v \n", restoreFolderPermission[checkFolderName][i].roles)
				fmt.Printf("* entitiy ID expected: %+v \n", orginalFolderPer.entityID)
				fmt.Printf("* entitiy ID actual:  %+v \n", restoreFolderPermission[checkFolderName][i].entityID)
				fmt.Println("Item:", checkFolderName)
				os.Exit(1)
			}
		}
	}

	for fileName, fileSize := range fileSizes {
		if fileSize != restoreFile[fileName] {
			fmt.Println("File size does not match:")
			fmt.Println("*  expected:", fileSize)
			fmt.Println("*  actual:", restoreFile[fileName])
			fmt.Println("Item:", fileName)
			os.Exit(1)
		}
	}

	fmt.Println("Success")
}

func getOneDriveChildFolder(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	driveID, itemID, parentName string,
	fileSizes map[string]int64,
	folderPermission map[string][]permissionInfo,
) {
	response, err := client.DrivesById(driveID).ItemsById(itemID).Children().Get(ctx, nil)
	if err != nil {
		fatal(ctx, "getting permission", err)
	}

	for _, driveItem := range response.GetValue() {
		var (
			itemID   = ptr.Val(driveItem.GetId())
			itemName = path.Join(parentName, ptr.Val(driveItem.GetName()))
		)

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			fileSizes[itemName] = ptr.Val(driveItem.GetSize())
		}

		if driveItem.GetFolder() == nil && driveItem.GetPackage() == nil {
			continue
		}

		permissionIn(ctx, client, driveID, itemID, itemName, folderPermission)
		getOneDriveChildFolder(ctx, client, driveID, itemID, itemName, fileSizes, folderPermission)
	}

}

func permissionIn(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	driveID, itemID, folderName string,
	perMap map[string][]permissionInfo,
) {
	perMap[folderName] = []permissionInfo{}

	pcr, err := client.
		DrivesById(driveID).
		ItemsById(itemID).
		Permissions().
		Get(ctx, nil)
	if err != nil {
		fatal(ctx, "getting permission", err)
	}

	for _, per := range pcr.GetValue() {
		if per.GetGrantedToV2() == nil {
			continue
		}

		var (
			gv2     = per.GetGrantedToV2()
			perInfo = permissionInfo{}
		)

		if gv2.GetUser() != nil {
			perInfo.entityID = ptr.Val(gv2.GetUser().GetId())
		} else if gv2.GetGroup() != nil {
			perInfo.entityID = ptr.Val(gv2.GetGroup().GetId())
		}

		perInfo.roles = per.GetRoles()

		slices.Sort(perInfo.roles)

		perMap[folderName] = append(perMap[folderName], perInfo)
	}
}

func getRestoreData(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	driveID,
	restoreFolderID string,
	restoreFile map[string]int64,
	restoreFolder map[string][]permissionInfo,
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
			restoreFile[itemName] = itemSize
			continue
		}

		if item.GetFolder() == nil && item.GetPackage() == nil {
			continue
		}

		permissionIn(ctx, client, driveID, itemID, itemName, restoreFolder)
		getOneDriveChildFolder(ctx, client, driveID, itemID, itemName, restoreFile, restoreFolder)

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

func mustGetTimeFromName(ctx context.Context, name string) (time.Time, bool) {
	t, err := common.ExtractTime(name)
	if err != nil && !errors.Is(err, common.ErrNoTimeString) {
		fatal(ctx, "extracting time from name: "+name, err)
	}

	return t, !errors.Is(err, common.ErrNoTimeString)
}

func isWithinTimeBound(ctx context.Context, bound, check time.Time, skip bool) bool {
	if skip {
		return true
	}

	if bound.Before(check) {
		logger.Ctx(ctx).
			With("boundary_time", bound, "check_time", check).
			Info("skipping restore folder: not older than time bound")

		return false
	}

	return true
}
