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
// @param folderName the target folder's display name
// @returns a *string if the folder exists, nil otherwise
func GetMailFolderID(service graph.Service, name, user string) (*string, error) {
	var errs error
	var folderId *string
	options, err := OptionsForMailFolders([]string{"displayName"})
	if err != nil {
		return folderId, err
	}
	response, err := service.Client().UsersById(user).MailFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil || response == nil {
		return folderId, err
	}
	pageIterator, err := msgraphgocore.NewPageIterator(response, service.Adapter(), models.CreateMailFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return folderId, err
	}
	callbackFunc := func(folderItem any) bool {
		folder, ok := folderItem.(models.MailFolderable)
		if !ok {
			errs = support.WrapAndAppend(service.Adapter().GetBaseUrl(), errors.New("HasFolder() iteration failure"), errs)
			return true
		}
		if *folder.GetDisplayName() == name {
			folderId = folder.GetId()
		}
		return true
	}
	iterateError := pageIterator.Iterate(callbackFunc)
	if iterateError != nil {
		errs = support.WrapAndAppend(service.Adapter().GetBaseUrl(), iterateError, errs)
	}
	return folderId, errs

}
