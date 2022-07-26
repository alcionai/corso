package connector

import (
	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// graphService internal struct that utilizes msgraph-sdk-go functions and interfaces
type graphService struct {
	client   msgraphsdk.GraphServiceClient
	adapter  msgraphsdk.GraphRequestAdapter
	failFast bool // if true service will exit sequence upon encountering an error
}

// createMailFolder will create a mail folder iff a folder of the same name does not exit
func createMailFolder(gc graphService, user, folder string) (models.MailFolderable, error) {
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	isHidden := false
	requestBody.SetIsHidden(&isHidden)

	return gc.client.UsersById(user).MailFolders().Post(requestBody)
}

// deleteMailFolder removes the mail folder from the user's M365 Exchange account
func deleteMailFolder(gc graphService, user, folderID string) error {
	return gc.client.UsersById(user).MailFoldersById(folderID).Delete()
}

func (gs *graphService) EnableFailFast() {
	gs.failFast = true
}

type GraphQuery func(graphService, []string) (absser.Parsable, error)

func GetAllMessagesForUser(gs graphService, identities []string) (absser.Parsable, error) {

	options := optionsForMessageSnapshot()
	return gs.client.UsersById(identities[0]).Messages().GetWithRequestConfigurationAndResponseHandler(options, nil)
}
