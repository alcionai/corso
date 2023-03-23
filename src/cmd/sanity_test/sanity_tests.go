package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
)

func main() {
	adapter, err := graph.CreateAdapter(
		os.Getenv("AZURE_TENANT_ID"),
		os.Getenv("AZURE_CLIENT_ID"),
		os.Getenv("AZURE_CLIENT_SECRET"),
	)
	if err != nil {
		fmt.Println("error while creating adapter: ", err)
		os.Exit(1)

		return
	}

	testUser := os.Getenv("CORSO_M365_TEST_USER_ID")
	folder := strings.TrimSpace(os.Getenv("RESTORE_FOLDER"))

	restoreStartTime := strings.SplitAfter(folder, "Corso_Restore_")[1]
	startTime, _ := time.Parse(time.RFC822, restoreStartTime)

	fmt.Println("Restore folder: ", folder)

	client := msgraphsdk.NewGraphServiceClient(adapter)

	switch service := os.Getenv("RESTORE_SERVICE"); service {
	case "exchange":
		checkEmailRestoration(client, testUser, folder, startTime)
	default:
		checkOnedriveRestoration(client, testUser, folder, startTime)
	}
}

// checkEmailRestoration verifies that the emails count in restored folder is equivalent to
// emails in actual m365 account
func checkEmailRestoration(
	client *msgraphsdk.GraphServiceClient,
	testUser,
	folderName string,
	startTime time.Time,
) {
	var (
		messageCount  = make(map[string]int32)
		restoreFolder models.MailFolderable
	)

	user := client.UsersById(testUser)
	builder := user.MailFolders()

	for {
		result, err := builder.Get(context.Background(), nil)
		if err != nil {
			fmt.Printf("Error getting the drive: %v\n", err)
			os.Exit(1)
		}

		res := result.GetValue()

		for _, r := range res {
			name, ok := ptr.ValOK(r.GetDisplayName())
			if !ok {
				continue
			}

			var rStartTime time.Time

			restoreStartTime := strings.SplitAfter(name, "Corso_Restore_")
			if len(restoreStartTime) > 1 {
				rStartTime, _ = time.Parse(time.RFC822, restoreStartTime[1])
				if startTime.Before(rStartTime) {
					fmt.Printf("The restore folder %s was created after %s. Will skip check.", name, folderName)
					continue
				}
			}

			if name == folderName {
				restoreFolder = r
				continue
			}

			getAllSubFolder(client, testUser, r, name, messageCount)

			messageCount[name], _ = ptr.ValOK(r.GetTotalItemCount())
		}

		link, ok := ptr.ValOK(result.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemMailFoldersRequestBuilder(link, client.GetAdapter())
	}

	folderID, ok := ptr.ValOK(restoreFolder.GetId())
	if !ok {
		fmt.Printf("can't find ID of restore folder")
		os.Exit(1)
	}

	folder := user.MailFoldersById(folderID)

	childFolder, err := folder.ChildFolders().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
		os.Exit(1)
	}

	for _, restore := range childFolder.GetValue() {
		restoreDisplayName, ok := ptr.ValOK(restore.GetDisplayName())
		if !ok {
			continue
		}

		restoreItemCount, _ := ptr.ValOK(restore.GetTotalItemCount())

		if messageCount[restoreDisplayName] != restoreItemCount {
			fmt.Println("Restore was not succesfull for: ", restoreDisplayName,
				"Folder count: ", messageCount[restoreDisplayName],
				"Restore count: ", restoreItemCount)
			os.Exit(1)
		}

		checkAllSubFolder(client, testUser, restore, restoreDisplayName, messageCount)
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

func checkOnedriveRestoration(client *msgraphsdk.GraphServiceClient, testUser, folderName string, startTime time.Time) {
	file := make(map[string]int64)
	folderPermission := make(map[string][]string)
	restoreFolderID := ""

	drive, err := client.UsersById(testUser).Drive().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
		os.Exit(1)
	}

	response, err := client.DrivesById(*drive.GetId()).Root().Children().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting drive by id: %v\n", err)
		os.Exit(1)
	}

	for _, driveItem := range response.GetValue() {
		if *driveItem.GetName() == folderName {
			restoreFolderID = *driveItem.GetId()
			continue
		}

		var rStartTime time.Time

		restoreStartTime := strings.SplitAfter(*driveItem.GetName(), "Corso_Restore_")
		if len(restoreStartTime) > 1 {
			rStartTime, _ = time.Parse(time.RFC822, restoreStartTime[1])
			if startTime.Before(rStartTime) {
				fmt.Printf("The restore folder %s was created after %s. Will skip check.", *driveItem.GetName(), folderName)
				continue
			}
		}

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			file[*driveItem.GetName()] = *driveItem.GetSize()
		}

		if driveItem.GetFolder() != nil {
			permission, err := client.
				DrivesById(*drive.GetId()).
				ItemsById(*driveItem.GetId()).
				Permissions().
				Get(context.TODO(), nil)
			if err != nil {
				fmt.Printf("Error getting item by id: %v\n", err)
				os.Exit(1)
			}

			// check if permission are correct on folder
			for _, permission := range permission.GetValue() {
				folderPermission[*driveItem.GetName()] = permission.GetRoles()
			}

			continue
		}
	}

	checkFileData(client, *drive.GetId(), restoreFolderID, file, folderPermission)

	fmt.Println("Success")
}

func checkFileData(
	client *msgraphsdk.GraphServiceClient,
	driveID,
	restoreFolderID string,
	file map[string]int64,
	folderPermission map[string][]string,
) {
	itemBuilder := client.DrivesById(driveID).ItemsById(restoreFolderID)

	restoreResponses, err := itemBuilder.Children().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting child folder: %v\n", err)
		os.Exit(1)
	}

	for _, restoreData := range restoreResponses.GetValue() {
		restoreName := *restoreData.GetName()

		if restoreData.GetFile() != nil {
			if *restoreData.GetSize() != file[restoreName] {
				fmt.Printf("Size of file %s is different in drive %d and restored file: %d ",
					restoreName,
					file[restoreName],
					*restoreData.GetSize())
				os.Exit(1)
			}

			continue
		}

		itemBuilder := client.DrivesById(driveID).ItemsById(*restoreData.GetId())

		if restoreData.GetFolder() != nil {
			permissionColl, err := itemBuilder.Permissions().Get(context.TODO(), nil)
			if err != nil {
				fmt.Printf("Error getting permission: %v\n", err)
				os.Exit(1)
			}

			userPermission := []string{}

			for _, perm := range permissionColl.GetValue() {
				userPermission = perm.GetRoles()
			}

			if !reflect.DeepEqual(folderPermission[restoreName], userPermission) {
				fmt.Printf("Permission mismatch for %s.", restoreName)
				os.Exit(1)
			}
		}
	}
}
