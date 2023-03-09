package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/alcionai/corso/src/internal/connector/graph"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
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
		// since multiple test cases are running in test env, the test case fails
		checkCalendarsRestoration(client, testUser, folder, startTime)
	default:
		checkOnedriveRestoration(client, testUser, folder, startTime)
	}
}

// TODO: since multiple test cases are running in test env, the test case fails. We will need another test account
func checkCalendarsRestoration(
	client *msgraphsdk.GraphServiceClient,
	testUser,
	folderName string,
	startTime time.Time,
) {
	user := client.UsersById(testUser)
	calendar := user.Calendars()

	result, err := calendar.Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
		os.Exit(1)
	}

	totalRestoreEvent := 0
	totalEvent := 0

	for _, r := range result.GetValue() {
		calendarItem, err := user.CalendarsById(*r.GetId()).Events().Get(context.TODO(), nil)
		if err != nil {
			fmt.Printf("Error calendar by id: %v\n", err)
			os.Exit(1)
		}

		restoreStartTime := strings.SplitAfter(*r.GetName(), "Corso_Restore_")[1]
		rStartTime, _ := time.Parse(time.RFC822, restoreStartTime)

		if startTime.Before(rStartTime) {
			fmt.Printf("The restore folder %s was created after %s. Will skip check.", *r.GetName(), folderName)
			continue
		}

		if strings.Contains(*r.GetName(), folderName) {
			totalRestoreEvent = len(calendarItem.GetValue())
			fmt.Printf("Calendar restore folder:  %s with events: %d \n",
				*r.GetName(),
				totalRestoreEvent)

			continue
		}

		eventCount := len(calendarItem.GetValue())
		fmt.Printf("Calendar folder: %s with %d \n", *r.GetName(), eventCount)
		totalEvent = totalEvent + eventCount
	}

	if totalRestoreEvent != totalEvent {
		fmt.Printf("Restore was not successful total events: %d restored events: %d", totalEvent, totalRestoreEvent)
		os.Exit(1)
	}
}

func checkEmailRestoration(client *msgraphsdk.GraphServiceClient, testUser, folderName string, startTime time.Time) {
	var (
		messageCount  = make(map[string]int32)
		restoreFolder models.MailFolderable
	)

	user := client.UsersById(testUser)
	mail := user.MailFolders()

	result, err := mail.Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
		os.Exit(1)
	}

	res := result.GetValue()

	// Hack: have added this to check if we are facing issues due to timeout or similar issue
	if len(res) < 1 {
		result, err = mail.Get(context.Background(), nil)
		if err != nil {
			fmt.Printf("Error getting the drive: %v\n", err)
			os.Exit(1)
		}

		res = result.GetValue()
	}

	for _, r := range res {
		name := *r.GetDisplayName()

		restoreStartTime := strings.SplitAfter(name, "Corso_Restore_")[1]
		rStartTime, _ := time.Parse(time.RFC822, restoreStartTime)

		if startTime.Before(rStartTime) {
			fmt.Printf("The restore folder %s was created after %s. Will skip check.", name, folderName)
			continue
		}

		if name == folderName {
			restoreFolder = r
			continue
		}

		messageCount[*r.GetDisplayName()] = *r.GetTotalItemCount()
	}

	user = client.UsersById(testUser)
	folder := user.MailFoldersById(*restoreFolder.GetId())

	childFolder, err := folder.ChildFolders().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
		os.Exit(1)
	}

	for _, restore := range childFolder.GetValue() {
		if messageCount[*restore.GetDisplayName()] != *restore.GetTotalItemCount() {
			fmt.Println("Restore was not succesfull for: ",
				*restore.GetDisplayName(),
				"Folder count: ",
				messageCount[*restore.GetDisplayName()],
				"Restore count: ",
				*restore.GetTotalItemCount())
			os.Exit(1)
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

		restoreStartTime := strings.SplitAfter(*driveItem.GetName(), "Corso_Restore_")[1]
		rStartTime, _ := time.Parse(time.RFC822, restoreStartTime)

		if startTime.Before(rStartTime) {
			fmt.Printf("The restore folder %s was created after %s. Will skip check.", *driveItem.GetName(), folderName)
			continue
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
