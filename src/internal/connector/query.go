package connector

import (
	"errors"

	"github.com/alcionai/corso/internal/connector/support"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
	msmessage "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
)

// TaskList is a a generic map of a list of items with a string index
type TaskList map[string][]string
type optionIdentifier int

//go:generate stringer -type=optionIdentifier
const (
	unknown optionIdentifier = iota
	folders
	messages
	users
)

// NewTaskList constructor for TaskList
func NewTaskList() TaskList {
	return make(map[string][]string, 0)
}

// AddTask helper method to ensure that keys and items are created properly
func (tl *TaskList) AddTask(key, value string) {
	aMap := *tl
	_, isCreated := aMap[key]
	if isCreated {
		aMap[key] = append(aMap[key], value)
	} else {
		aMap[key] = []string{value}
	}
}

// Contains is a helper method for verifying if element
// is contained within the slice
func Contains(elems []string, value string) bool {
	for _, s := range elems {
		if value == s {
			return true
		}
	}
	return false
}

// HasMailFolder helper function to see if MailFolder exists returns folderId
func HasMailFolder(name, user string, service graphService) (*string, error) {
	var errs error
	var isAvailable *string
	options, err := optionsForMailFolders([]string{"displayName"})
	if err != nil {
		return isAvailable, err
	}
	response, err := service.client.UsersById(user).MailFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil || response == nil {
		return isAvailable, err
	}
	pageIterator, err := msgraphgocore.NewPageIterator(response, &service.adapter, models.CreateMailFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return isAvailable, err
	}
	callbackFunc := func(folderItem any) bool {
		folder, ok := folderItem.(models.MailFolderable)
		if !ok {
			errs = support.WrapAndAppend(service.adapter.GetBaseUrl(), errors.New("HasFolder() iteration failure"), errs)
			return true
		}
		if *folder.GetDisplayName() == name {
			isAvailable = folder.GetId()
		}
		return true
	}
	iterateError := pageIterator.Iterate(callbackFunc)
	if iterateError != nil {
		errs = support.WrapAndAppend(service.adapter.GetBaseUrl(), iterateError, errs)
	}
	return isAvailable, errs

}

// optionsForMailFolders creates transforms the 'select' into a more dynamic call for MailFolders.
// var moreOps is a []string of options(e.g. "displayName", "isHidden")
// return is first call in MailFolders().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFolders(moreOps []string) (*msfolder.MailFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, folders)
	if err != nil {
		return nil, err
	}

	requestParameters := &msfolder.MailFoldersRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msfolder.MailFoldersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	return options, nil
}

func optionsForMessageSnapshot() *msmessage.MessagesRequestBuilderGetRequestConfiguration {
	selecting := []string{"id", "parentFolderId"}
	options := &msmessage.MessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &msmessage.MessagesRequestBuilderGetQueryParameters{
			Select: selecting,
		},
	}
	return options
}

func optionsForMessages(moreOps []string) (*msmessage.MessagesRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, messages)
	if err != nil {
		return nil, err
	}
	requestParameters := &msmessage.MessagesRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msmessage.MessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	return options, nil
}

// CheckOptions Utility Method for verifying if select options are valid the m365 object type
// returns a list of valid options
func buildOptions(options []string, selection optionIdentifier) ([]string, error) {
	var allowedOptions []string

	fieldsForFolders := []string{"displayName", "isHidden", "parentFolderId", "totalItemCount"}
	fieldsForUsers := []string{"birthday", "businessPhones", "city", "companyName", "department", "displayName", "employeeId"}
	fieldsForMessages := []string{"conservationId", "conversationIndex", "parentFolderId", "subject", "webLink"}
	returnedOptions := []string{"id"}

	switch selection {
	case folders:
		allowedOptions = fieldsForFolders
	case users:
		allowedOptions = fieldsForUsers
	case messages:
		allowedOptions = fieldsForMessages
	default:
		return nil, errors.New("unsupported option")
	}

	for _, entry := range options {
		result := Contains(allowedOptions, entry)
		if result {
			returnedOptions = append(returnedOptions, entry)
		} else {
			return nil, errors.New("unsupported option")
		}
	}
	return returnedOptions, nil
}
