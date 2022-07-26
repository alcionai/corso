package exchange

import (
	"errors"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
	msmessage "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
	msitem "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item"
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

// optionsForMailFolders creates transforms the 'select' into a more dynamic call for MailFolders.
// var moreOps is a []string of options(e.g. "displayName", "isHidden")
// return is first call in MailFolders().GetWithRequestConfigurationAndResponseHandler(options, handler)
func OptionsForMailFolders(moreOps []string) (*msfolder.MailFoldersRequestBuilderGetRequestConfiguration, error) {
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

func OptionsForMessageSnapshot() *msmessage.MessagesRequestBuilderGetRequestConfiguration {
	selecting := []string{"id", "parentFolderId"}
	options := &msmessage.MessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &msmessage.MessagesRequestBuilderGetQueryParameters{
			Select: selecting,
		},
	}
	return options
}

func OptionsForSingleMessage(moreOps []string) (*msitem.MessageItemRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, messages)
	if err != nil {
		return nil, err
	}
	requestParams := &msitem.MessageItemRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &msitem.MessageItemRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParams,
	}
	return options, nil
}

func OptionsForMessages(moreOps []string) (*msmessage.MessagesRequestBuilderGetRequestConfiguration, error) {
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

// createMailFolder will create a mail folder iff a folder of the same name does not exit
func CreateMailFolder(service GraphService, user, folder string) (models.MailFolderable, error) {
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	isHidden := false
	requestBody.SetIsHidden(&isHidden)

	return service.Client.UsersById(user).MailFolders().Post(requestBody)
}

// deleteMailFolder removes the mail folder from the user's M365 Exchange account
func DeleteMailFolder(service GraphService, user, folderID string) error {
	return service.Client.UsersById(user).MailFoldersById(folderID).Delete()
}
