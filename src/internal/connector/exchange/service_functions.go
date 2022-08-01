package exchange

import (
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
)

// CreateMailFolder makes a mail folder iff a folder of the same name does not exist
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-mailfolders?view=graph-rest-1.0&tabs=http
func CreateMailFolder(gs graph.Service, user, folder string) (models.MailFolderable, error) {
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	isHidden := false
	requestBody.SetIsHidden(&isHidden)

	return gs.Client().UsersById(user).MailFolders().Post(requestBody)
}

// DeleteMailFolder removes a mail folder with the corresponding M365 ID  from the user's M365 Exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/mailfolder-delete?view=graph-rest-1.0&tabs=http
func DeleteMailFolder(gs graph.Service, user, folderID string) error {
	return gs.Client().UsersById(user).MailFoldersById(folderID).Delete()
}

// GetMailFolderID query function to retrieve the M365 ID based on the folder's displayName.
// @param folderName the target folder's display name. Case sensitive
// @returns a *string if the folder exists, nil otherwise
func GetMailFolderID(service graph.Service, folderName, user string) (*string, error) {
	var errs error
	var folderId *string
	options, err := optionsForMailFolders([]string{"displayName"})
	if err != nil {
		return nil, err
	}
	response, err := service.Client().UsersById(user).MailFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, errors.New("mail folder query to m365 back store returned nil")
	}
	pageIterator, err := msgraphgocore.NewPageIterator(response, service.Adapter(), models.CreateMailFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	callbackFunc := func(folderItem any) bool {
		folder, ok := folderItem.(models.MailFolderable)
		if !ok {
			errs = support.WrapAndAppend(service.Adapter().GetBaseUrl(), errors.New("HasFolder() iteration failure"), errs)
			return true
		}
		if *folder.GetDisplayName() == folderName {
			folderId = folder.GetId()
			return false
		}
		return true
	}
	iterateError := pageIterator.Iterate(callbackFunc)
	if iterateError != nil {
		errs = support.WrapAndAppend(service.Adapter().GetBaseUrl(), iterateError, errs)
	}
	return folderId, errs

}
