package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
)

func main() {
	adapter, err := graph.CreateAdapter(
		os.Getenv("AZURE_TENANT_ID"),
		os.Getenv("AZURE_CLIENT_ID"),
		os.Getenv("AZURE_CLIENT_SECRET"),
	)
	if err != nil {
		fatal("error while creating adapter", err)
	}

	var (
		ctx              = context.Background()
		client           = msgraphsdk.NewGraphServiceClient(adapter)
		testUser         = os.Getenv("CORSO_M365_TEST_USER_ID")
		folder           = strings.TrimSpace(os.Getenv("RESTORE_FOLDER"))
		restoreStartTime = strings.SplitAfter(folder, "Corso_Restore_")[1]
	)

	startTime, err := time.Parse(time.RFC822, restoreStartTime)
	if err != nil {
		fatal("error parsing start time", err)
	}

	fmt.Println("Restore folder: ", folder)

	switch service := os.Getenv("RESTORE_SERVICE"); service {
	case "exchange":
		checkEmailRestoration(ctx, client, testUser, folder, startTime)
	default:
		checkOnedriveRestoration(ctx, client, testUser, folder, startTime)
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
			fatal("error getting the drive", err)
		}

		values := result.GetValue()

		for _, v := range values {
			itemName := ptr.Val(v.GetDisplayName())

			folderTime, err := common.ExtractTime(itemName)
			if err != nil && !errors.Is(err, common.ErrNoTimeString) {
				fatal("extracting time from file name", err)
			}

			if !errors.Is(err, common.ErrNoTimeString) && startTime.Before(folderTime) {
				fmt.Printf("skipping restore folder %s created after %s", itemName, folderName)
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
		user   = client.UsersById(testUser)
		folder = user.MailFoldersById(ptr.Val(restoreFolder.GetId()))
	)

	childFolder, err := folder.ChildFolders().Get(ctx, nil)
	if err != nil {
		fatal("error getting the drive", err)
	}

	for _, fld := range childFolder.GetValue() {
		var (
			fldName = ptr.Val(fld.GetDisplayName())
			count   = ptr.Val(fld.GetTotalItemCount())
		)

		if itemCount[fldName] != count {
			fmt.Println("Restore item counts do not match:")
			fmt.Println("-  expected:", itemCount[fldName])
			fmt.Println("-  actual:", count)
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

	drive, err := client.UsersById(testUser).Drive().Get(ctx, nil)
	if err != nil {
		fatal("error getting the drive:", err)
	}

	driveID := ptr.Val(drive.GetId())

	response, err := client.
		DrivesById(driveID).
		Root().
		Children().
		Get(ctx, nil)
	if err != nil {
		fmt.Println("Error getting drive by id:", err)
		os.Exit(1)
	}

	for _, driveItem := range response.GetValue() {
		var (
			itemID     = ptr.Val(driveItem.GetId())
			itemName   = ptr.Val(driveItem.GetName())
			rStartTime time.Time
		)

		if itemName == folderName {
			restoreFolderID = itemID
			continue
		}

		restoreStartTime := strings.SplitAfter(itemName, "Corso_Restore_")
		if len(restoreStartTime) > 1 {
			rStartTime, _ = time.Parse(time.RFC822, restoreStartTime[1])
			if startTime.Before(rStartTime) {
				fmt.Printf("The restore folder %s was created after %s. Will skip check.\n", itemName, folderName)
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
				Get(ctx, nil)
			if err != nil {
				fmt.Println("Error getting item by id:", err)
				os.Exit(1)
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
		fatal("error getting child folder", err)
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
			fatal("error getting permission", err)
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
			}
		}
	}
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func fatal(msg string, err error) {
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
