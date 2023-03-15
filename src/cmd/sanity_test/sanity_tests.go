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
		os.Getenv("AZURE_CLIENT_SECRET"))
	if err != nil {
		fatal("error while creating adapter", err)
	}

	var (
		ctx         = context.Background()
		client      = msgraphsdk.NewGraphServiceClient(adapter)
		testUser    = os.Getenv("CORSO_M365_TEST_USER_ID")
		testService = os.Getenv("SANITY_RESTORE_SERVICE")
		folder      = strings.TrimSpace(os.Getenv("SANITY_RESTORE_FOLDER"))
	)

	startTime, err := common.ExtractTime(folder)
	if err != nil {
		fatal("error parsing start time", err)
	}

	fmt.Println("Restore folder: ", folder)

	switch testService {
	case "exchange":
		checkEmailRestoration(ctx, client, testUser, folder, startTime)
	case "onedrive":
		checkOnedriveRestoration(ctx, client, testUser, folder, startTime)
	default:
		fatal("no service specified", nil)
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
				fmt.Printf("skipping restore folder %s created after %s\n", itemName, folderName)
				continue
			}

			if itemName == folderName {
				restoreFolder = v
				continue
			}

			getAllSubFolder(client, testUser, v, itemName, itemCount)

			itemCount[itemName] = ptr.Val(v.GetTotalItemCount())
		}

		link, ok := ptr.ValOK(result.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemMailFoldersRequestBuilder(link, client.GetAdapter())
	}

	var (
		user     = client.UsersById(testUser)
		folderID = ptr.Val(restoreFolder.GetId())
		folder   = user.MailFoldersById(folderID)
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
			fmt.Println("Folder:", fldName, ptr.Val(fld.GetId()))
			os.Exit(1)
		}

		checkAllSubFolder(client, testUser, fld, fldName, itemCount)
	}
}

// getAllSubFolder will recursively check for all subfolders and get the corresponding
// email count.
func getAllSubFolder(
	client *msgraphsdk.GraphServiceClient,
	testUser string,
	r models.MailFolderable,
	parentFolder string,
	messageCount map[string]int32,
) {
	folderID, ok := ptr.ValOK(r.GetId())

	if !ok {
		fmt.Println("unable to get sub folder ID")
		return
	}

	user := client.UsersById(testUser)
	folder := user.MailFoldersById(folderID)

	var count int32 = 99

	childFolder, err := folder.ChildFolders().Get(
		context.Background(),
		&users.ItemMailFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMailFoldersItemChildFoldersRequestBuilderGetQueryParameters{
				Top: &count,
			},
		})
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
		os.Exit(1)
	}

	for _, child := range childFolder.GetValue() {
		childDisplayName, _ := ptr.ValOK(child.GetDisplayName())

		fullFolderName := parentFolder + "/" + childDisplayName

		messageCount[fullFolderName], _ = ptr.ValOK(child.GetTotalItemCount())

		childFolderCount, _ := ptr.ValOK(child.GetChildFolderCount())

		// recursively check for subfolders
		if childFolderCount > 0 {
			parentFolder := fullFolderName

			getAllSubFolder(client, testUser, child, parentFolder, messageCount)
		}
	}
}

// checkAllSubFolder will recursively traverse inside the restore folder and
// verify that data matched in all subfolders
func checkAllSubFolder(
	client *msgraphsdk.GraphServiceClient,
	testUser string,
	r models.MailFolderable,
	parentFolder string,
	messageCount map[string]int32,
) {
	folderID, ok := ptr.ValOK(r.GetId())

	if !ok {
		fmt.Println("unable to get sub folder ID")
		return
	}

	user := client.UsersById(testUser)
	folder := user.MailFoldersById(folderID)

	var count int32 = 99

	childFolder, err := folder.ChildFolders().Get(
		context.Background(),
		&users.ItemMailFoldersItemChildFoldersRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemMailFoldersItemChildFoldersRequestBuilderGetQueryParameters{
				Top: &count,
			},
		})
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
		os.Exit(1)
	}

	for _, child := range childFolder.GetValue() {
		childDisplayName, _ := ptr.ValOK(child.GetDisplayName())

		fullFolderName := parentFolder + "/" + childDisplayName

		childTotalCount, _ := ptr.ValOK(child.GetTotalItemCount())

		if messageCount[fullFolderName] != childTotalCount {
			fmt.Println("Restore was not succesfull for: ", fullFolderName,
				"Folder count: ", messageCount[fullFolderName],
				"Restore count: ", childTotalCount)
			os.Exit(1)
		}

		childFolderCount, _ := ptr.ValOK(child.GetChildFolderCount())

		if childFolderCount > 0 {
			parentFolder := fullFolderName

			checkAllSubFolder(client, testUser, child, parentFolder, messageCount)
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
				fmt.Printf("skipping restore folder %s created after %s\n", itemName, folderName)
				continue
			}
		}

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			file[itemName] = ptr.Val(driveItem.GetSize())
		}

		if driveItem.GetFolder() != nil {
			folderPermission[itemID] = permissionsIn(ctx, client, driveID, itemID, folderPermission[itemID])
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
		fatal("error getting item children", err)
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

		if item.GetFolder() == nil {
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

func fatal(msg string, err error) {
	fmt.Println(msg+":", err)
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
		fatal("error getting permission", err)
	}

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
