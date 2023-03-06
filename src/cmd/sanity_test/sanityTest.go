package main

import (
	"context"
	"fmt"
	"os"
	"strings"

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
		os.Exit(1)
		return
	}

	testUser := os.Getenv("CORSO_M365_TEST_USER_ID")
	folder := strings.TrimSpace(os.Getenv("RESTORE_FOLDER"))

	client := msgraphsdk.NewGraphServiceClient(adapter)

	if os.Getenv("EXCHANGE_TEST") == "true" {
		checkEmailRestoration(client, testUser, folder)
		return
	}

	checkOnedriveRestoration(client, testUser, folder)
}

func checkEmailRestoration(client *msgraphsdk.GraphServiceClient, testUser, folderName string) {
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

	for _, r := range res {
		name := *r.GetDisplayName()

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
			fmt.Println("Restore was not succesfull for: ", *restore.GetDisplayName())
			os.Exit(1)
		}
	}
}

func checkOnedriveRestoration(client *msgraphsdk.GraphServiceClient, testUser, folderName string) {
	file := make(map[string]int64)
	restoreFolderID := ""

	drive, err := client.UsersById(testUser).Drive().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
		os.Exit(1)
	}

	response, _ := client.DrivesById(*drive.GetId()).Root().Children().Get(context.Background(), nil)
	for _, drive := range response.GetValue() {
		if *drive.GetName() == folderName {
			restoreFolderID = *drive.GetId()
			continue
		}

		size := *drive.GetSize()

		// check if file or folder
		if size > 0 {
			file[*drive.GetName()] = *drive.GetSize()
		}
	}

	restoreResponse, _ := client.
		DrivesById(*drive.GetId()).
		ItemsById(restoreFolderID).
		Children().
		Get(context.Background(), nil)

	for _, restoreResponse := range restoreResponse.GetValue() {
		if *restoreResponse.GetSize() != file[*restoreResponse.GetName()] {
			fmt.Printf("Size of file %s is different in drive %d and restored file: %d ",
				*restoreResponse.GetName(),
				file[*restoreResponse.GetName()],
				*restoreResponse.GetSize())
			os.Exit(1)
		}
	}

	fmt.Println("Success")
}
