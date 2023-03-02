package main

import (
	"context"
	"fmt"
	"os"

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

	testUser := os.Getenv("CORSO_M365_LOAD_TEST_USER_ID")

	client := msgraphsdk.NewGraphServiceClient(adapter)

	if os.Getenv("EXCHANGE_TEST") == "true" {
		checkEmailRestoration(client, testUser)
		return
	}

}

func checkEmailRestoration(client *msgraphsdk.GraphServiceClient, testUser string) {
	result, err := client.UsersById(testUser).MailFolders().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting the drive: %v\n", err)
		os.Exit(1)
	}

	res := result.GetValue()

	var messageCount = make(map[string]int32)
	var restoreFolder models.MailFolderable
	for _, r := range res {
		name := *r.GetDisplayName()

		if name == os.Getenv("RESTORE_FOLDER") {
			restoreFolder = r
			continue
		}

		messageCount[*r.GetDisplayName()] = *r.GetTotalItemCount()
	}

	user := client.UsersById(testUser)
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
