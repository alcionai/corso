package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"golang.org/x/exp/slices"
)

func main() {
	adapter, err := graph.CreateAdapter(
		os.Getenv("AZURE_TENANT_ID"),
		os.Getenv("AZURE_CLIENT_ID"),
		os.Getenv("AZURE_CLIENT_SECRET"))
	if err != nil {
		fmt.Println("error while creating adapter: ", err)
		os.Exit(1)

		return
	}

	testUser := "HenriettaM@10rqc2.onmicrosoft.com"
	folder := "Corso_Restore_28-Mar-2023_07-13-04"
	// testUser := os.Getenv("CORSO_M365_TEST_USER_ID")
	// folder := strings.TrimSpace(os.Getenv("RESTORE_FOLDER"))
	dataFolder := os.Getenv("TEST_DATA")
	baseBackupFolder := os.Getenv("BASE_BACKUP")

	restoreStartTime := strings.SplitAfter(folder, "Corso_Restore_")[1]
	startTime, _ := time.Parse(time.RFC822, restoreStartTime)

	fmt.Println("Restore folder: ", folder)

	client := msgraphsdk.NewGraphServiceClient(adapter)

	switch service := os.Getenv("RESTORE_SERVICE"); service {
	case "exchange":
		checkEmailRestoration(client, testUser, folder, dataFolder, baseBackupFolder, startTime)
	default:
		checkOnedriveRestoration(client, testUser, folder, startTime)
	}
}

// checkEmailRestoration verifies that the emails count in restored folder is equivalent to
// emails in actual m365 account
func checkEmailRestoration(
	client *msgraphsdk.GraphServiceClient,
	testUser,
	folderName,
	dataFolder,
	baseBackupFolder string,
	startTime time.Time,
) {
	var (
		messageCount        = make(map[string]int32)
		restoreFolder       models.MailFolderable
		restoreMessageCount = make(map[string]int32)
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
				fmt.Println("display name not found. Will continue")
				continue
			}

			if name == folderName {
				restoreFolder = r
				continue
			}

			if name == dataFolder || name == baseBackupFolder {
				getAllSubFolder(client, testUser, r, name, dataFolder, messageCount)

				messageCount[name], _ = ptr.ValOK(r.GetTotalItemCount())
			}
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
			fmt.Println("display name not found. Will continue")
			continue
		}

		// check if folder is the data folder we loaded or the base backup to verify
		// the incremental backup worked fine
		if strings.EqualFold(restoreDisplayName, dataFolder) || strings.EqualFold(restoreDisplayName, baseBackupFolder) {
			restoreItemCount, _ := ptr.ValOK(restore.GetTotalItemCount())

			restoreMessageCount[restoreDisplayName] = restoreItemCount
			checkAllSubFolder(client, testUser, restore, restoreDisplayName, dataFolder, restoreMessageCount)
		}
	}

	verifyEmailData(restoreMessageCount, messageCount)
}

func verifyEmailData(restoreMessageCount, messageCount map[string]int32) {
	for folderName, emailCount := range messageCount {
		fmt.Println("verifying message count for ", folderName)

		if restoreMessageCount[folderName] != emailCount {
			fmt.Println("Restore was not succesfull for: ", folderName,
				"Folder count: ", emailCount,
				"Restore count: ", restoreMessageCount[folderName])
			os.Exit(1)
		}
	}
}

// getAllSubFolder will recursively check for all subfolders and get the corresponding
// email count.
func getAllSubFolder(
	client *msgraphsdk.GraphServiceClient,
	testUser string,
	r models.MailFolderable,
	parentFolder,
	dataFolder string,
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

		if strings.Contains(fullFolderName, dataFolder) {
			messageCount[fullFolderName], _ = ptr.ValOK(child.GetTotalItemCount())
			childFolderCount, _ := ptr.ValOK(child.GetChildFolderCount())

			// recursively check for subfolders
			if childFolderCount > 0 {
				parentFolder := fullFolderName

				getAllSubFolder(client, testUser, child, parentFolder, dataFolder, messageCount)
			}
		}
	}
}

// checkAllSubFolder will recursively traverse inside the restore folder and
// verify that data matched in all subfolders
func checkAllSubFolder(
	client *msgraphsdk.GraphServiceClient,
	testUser string,
	r models.MailFolderable,
	parentFolder,
	dataFolder string,
	restoreMessageCount map[string]int32,
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

		if strings.Contains(fullFolderName, dataFolder) {
			childTotalCount, _ := ptr.ValOK(child.GetTotalItemCount())
			restoreMessageCount[fullFolderName] = childTotalCount
		}

		childFolderCount, _ := ptr.ValOK(child.GetChildFolderCount())

		if childFolderCount > 0 {
			parentFolder := fullFolderName
			checkAllSubFolder(client, testUser, child, parentFolder, dataFolder, restoreMessageCount)
		}
	}
}

type permissionInfo struct {
	entityID string
	roles    []string
}

func checkOnedriveRestoration(client *msgraphsdk.GraphServiceClient, testUser, folderName string, startTime time.Time) {
	var (
		file            = make(map[string]int64)
		perMap          = make(map[string][]permissionInfo)
		restoreFolderID = ""
		restoreFile     = make(map[string]int64)
		restorePerMap   = make(map[string][]permissionInfo)
		rStartTime      time.Time
	)

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

			checkPermission(permission, perMap, *driveItem.GetName())

			continue
		}
	}

	getRestoreData(client, *drive.GetId(), restoreFolderID, restoreFile, restorePerMap)

	for checkFolderName, checkfolderPer := range perMap {
		for i, orginalFolderPer := range checkfolderPer {
			if !(orginalFolderPer.entityID != restorePerMap[checkFolderName][i].entityID) &&
				!slices.Equal(orginalFolderPer.roles, restorePerMap[checkFolderName][i].roles) {
				fmt.Println("permissions are not equal")
				fmt.Printf("*  expected role: %+v \n", orginalFolderPer.roles)
				fmt.Printf("*  actual:  %+v \n", restorePerMap[checkFolderName][i].roles)
				fmt.Printf("* entitiy ID expected: %+v \n", orginalFolderPer.entityID)
				fmt.Printf("* entitiy ID actual:  %+v \n", restorePerMap[checkFolderName][i].entityID)
				fmt.Println("Item:", checkFolderName)
				os.Exit(1)
			}
		}
	}

	for fileName, fileSize := range file {
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

func checkPermission(
	permissionColl models.PermissionCollectionResponseable,
	perMap map[string][]permissionInfo,
	folderName string,
) {
	perMap[folderName] = []permissionInfo{}

	for _, per := range permissionColl.GetValue() {
		perInfo := permissionInfo{}

		if per.GetGrantedToV2() == nil {
			continue
		}

		gv2 := per.GetGrantedToV2()

		if gv2.GetUser() != nil {
			perInfo.entityID = ptr.Val(gv2.GetUser().GetId())
		} else if gv2.GetGroup() != nil {
			perInfo.entityID = ptr.Val(gv2.GetGroup().GetId())
		}

		perInfo.roles = per.GetRoles()

		perMap[folderName] = append(perMap[folderName], perInfo)
	}
}

func getRestoreData(
	client *msgraphsdk.GraphServiceClient,
	driveID,
	restoreFolderID string,
	restoreFile map[string]int64,
	restoreFolder map[string][]permissionInfo,
) {
	itemBuilder := client.DrivesById(driveID).ItemsById(restoreFolderID)

	restoreResponses, err := itemBuilder.Children().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting child folder: %v\n", err)
		os.Exit(1)
	}

	for _, restoreData := range restoreResponses.GetValue() {
		restoreName := ptr.Val(restoreData.GetName())

		if restoreData.GetFile() != nil {
			restoreFile[restoreName] = *restoreData.GetSize()
			continue
		}

		itemBuilder := client.DrivesById(driveID).ItemsById(*restoreData.GetId())
		if restoreData.GetFolder() != nil {
			permissionColl, err := itemBuilder.Permissions().Get(context.TODO(), nil)
			if err != nil {
				fmt.Printf("Error getting permission: %v\n", err)
				os.Exit(1)
			}

			checkPermission(permissionColl, restoreFolder, restoreName)
		}
	}
}
