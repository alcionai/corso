package main

import (
	"context"
	"errors"
	"fmt"
	"os"
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
		fatal(ctx, "error while creating adapter", err)
	}

	var (
		client      = msgraphsdk.NewGraphServiceClient(adapter)
		testUser    = os.Getenv("CORSO_M365_TEST_USER_ID")
		testService = os.Getenv("SANITY_RESTORE_SERVICE")
		folder      = strings.TrimSpace(os.Getenv("SANITY_RESTORE_FOLDER"))
	)

	startTime, err := common.ExtractTime(folder)
	if err != nil {
		fatal(ctx, "error parsing start time", err)
	}

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
			fatal(ctx, "error getting the drive", err)
		}

		values := result.GetValue()

		for _, v := range values {
			var (
				itemID   = ptr.Val(v.GetId())
				itemName = ptr.Val(v.GetDisplayName())
				ictx     = clues.Add(ctx, "item_id", itemID, "item_name", itemName)
			)

			folderTime, err := common.ExtractTime(itemName)
			if err != nil && !errors.Is(err, common.ErrNoTimeString) {
				fatal(ctx, "extracting time from file name", err)
			}

			if !errors.Is(err, common.ErrNoTimeString) && startTime.Before(folderTime) {
				logger.Ctx(ictx).
					With("folder_time", folderTime).
					Info("skipping restore folder: not within time bound")

				continue
			}

			if itemName == folderName {
				restoreFolder = v
				continue
			}

			itemCount[itemName] = ptr.Val(v.GetTotalItemCount())
		}

		link, ok := ptr.ValOK(result.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemMailFoldersRequestBuilder(link, client.GetAdapter())
	}

	var (
		restoreFldID   = ptr.Val(restoreFolder.GetId())
		restoreFldName = ptr.Val(restoreFolder.GetDisplayName())
		user           = client.UsersById(testUser)
		folder         = user.MailFoldersById(restoreFldID)
	)

	ctx = clues.Add(
		ctx,
		"restore_folder_id", restoreFldID,
		"restore_folder_name", restoreFldName)

	childFolder, err := folder.ChildFolders().Get(ctx, nil)
	if err != nil {
		fatal(ctx, "error getting the drive", err)
	}

	for _, fld := range childFolder.GetValue() {
		var (
			fldID   = ptr.Val(fld.GetId())
			fldName = ptr.Val(fld.GetDisplayName())
			count   = ptr.Val(fld.GetTotalItemCount())
			ictx    = clues.Add(ctx,
				"child_folder_id", fldID,
				"child_folder_name", fldName,
				"expected_count", itemCount[fldName],
				"actual_count", count)
		)

		if itemCount[fldName] != count {
			logger.Ctx(ictx).Error("test failure: Restore item counts do not match")
			fmt.Println("Restore item counts do not match:")
			fmt.Println("-  expected:", itemCount[fldName])
			fmt.Println("-  actual:", count)
			fmt.Println("Folder:", fldName, ptr.Val(fld.GetId()))
			os.Exit(1)
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
		file = make(map[string]int64)
		// map itemID -> permission id -> []permission roles
		folderPermission = make(map[string]map[string][]string)
		restoreFolderID  = ""
	)

	drive, err := client.
		UsersById(testUser).
		Drive().
		Get(ctx, nil)
	if err != nil {
		fatal(ctx, "error getting the drive:", err)
	}

	var (
		driveID   = ptr.Val(drive.GetId())
		driveName = ptr.Val(drive.GetName())
	)

	ctx = clues.Add(ctx, "drive_id", driveID, "drive_name", driveName)

	response, err := client.
		DrivesById(driveID).
		Root().
		Children().
		Get(ctx, nil)
	if err != nil {
		fmt.Println(ctx, "Error getting drive by id:", err)
		os.Exit(1)
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

		if !startTime.IsZero() {
			folderTime, err := common.ExtractTime(itemName)
			if err != nil && !errors.Is(err, common.ErrNoTimeString) {
				fatal(ictx, "extracting time from file name", err)
			}

			if !errors.Is(err, common.ErrNoTimeString) && startTime.Before(folderTime) {
				logger.Ctx(ictx).
					With("folder_time", folderTime).
					Info("skipping restore folder: not within time bound")

				continue
			}
		}

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			file[itemName] = ptr.Val(driveItem.GetSize())
		}

		if driveItem.GetFolder() != nil {
			pcr, err := client.
				DrivesById(driveID).
				ItemsById(itemID).
				Permissions().
				Get(ictx, nil)
			if err != nil {
				fatal(ictx, "error getting item by id", err)
			}

			folderPermission[itemID] = permissionsIn(pcr, folderPermission[itemID])
		}
	}

	checkFileData(ctx, client, driveID, restoreFolderID, file, folderPermission)

	fmt.Println("Success")
}

func checkFileData(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	driveID,
	restoreFolderID string,
	file map[string]int64,
	folderPermission map[string]map[string][]string,
) {
	restored, err := client.
		DrivesById(driveID).
		ItemsById(restoreFolderID).
		Children().
		Get(ctx, nil)
	if err != nil {
		fatal(ctx, "error getting child folder", err)
	}

	for _, item := range restored.GetValue() {
		var (
			itemID   = ptr.Val(item.GetId())
			itemName = ptr.Val(item.GetName())
			itemSize = ptr.Val(item.GetSize())
		)

		if item.GetFile() != nil {
			if itemSize != file[itemName] {
				fmt.Println("File size does not match:")
				fmt.Println("-  expected:", file[itemName])
				fmt.Println("-  actual:", itemSize)
				fmt.Println("Item:", itemName, itemID)
				os.Exit(1)
			}

			continue
		}

		pcr, err := client.
			DrivesById(driveID).
			ItemsById(ptr.Val(item.GetId())).
			Permissions().
			Get(ctx, nil)
		if err != nil {
			fatal(ctx, "error getting permission", err)
		}

		var (
			expectItem = folderPermission[itemID]
			results    = permissionsIn(pcr, nil)
		)

		for pid, result := range results {
			expect := expectItem[pid]

			if !slices.Equal(expect, result) {
				fmt.Println("permissions are not equal")
				fmt.Println("-  expected: ", expect)
				fmt.Println("-  actual: ", result)
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
	fmt.Println(msg+":", err)
	os.Exit(1)
}

func permissionsIn(pcr models.PermissionCollectionResponseable, init map[string][]string) map[string][]string {
	result := map[string][]string{}

	if len(init) > 0 {
		maps.Copy(result, init)
	}

	// check if permission are correct on folder
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
