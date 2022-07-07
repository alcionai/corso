package connector

import (
	"errors"

	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
	msmessage "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
)

// TaskList is a a generic map of a list of items with a string index
type TaskList struct {
	tasks map[string][]string
}

// NewTaskList constructor for TaskList
func NewTaskList() TaskList {
	taskList := &TaskList{
		tasks: make(map[string][]string, 0),
	}
	return *taskList
}

// AddTask helper method to ensure that keys and items are created properly
func (tl *TaskList) AddTask(key, value string) {
	_, isCreated := tl.tasks[key]
	if isCreated {
		tl.tasks[key] = append(tl.tasks[key], value)
	} else {
		tl.tasks[key] = []string{value}
	}
}

// GetTasks helper method for retrieving list by index
func (tl *TaskList) GetTasks(key string) []string {
	aList, ok := tl.tasks[key]
	if ok {
		return aList
	}
	return []string{}
}

func (tl *TaskList) GetKeys() []string {
	keys := make([]string, 0)
	for entry := range tl.tasks {
		keys = append(keys, entry)
	}
	return keys
}

// Length returns the amount of indexes within the struct
func (tl *TaskList) Length() int {
	return len(tl.tasks)
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
// var moreOps is a comma separated string of options(e.g. "displayName, isHidden")
// return is first call in MailFolders().GetWithRequestConfigurationAndResponseHandler(options, handler)
func optionsForMailFolders(moreOps []string) (*msfolder.MailFoldersRequestBuilderGetRequestConfiguration, error) {
	selecting, err := buildOptions(moreOps, 1)
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
	selecting, err := buildOptions(moreOps, 3)
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
func buildOptions(options []string, selection int) ([]string, error) {
	var allowedOptions []string

	fieldsForFolders := []string{"displayName", "isHidden", "parentFolderId", "totalItemCount"}
	fieldsForUsers := []string{"birthday", "businessPhones", "city", "companyName", "department", "displayName", "employeeId"}
	fieldsForMessages := []string{"conservationId", "conversationIndex", "parentFolderId", "subject", "webLink"}
	returnedOptions := []string{"id"}

	switch selection {
	case 1:
		allowedOptions = fieldsForFolders
	case 2:
		allowedOptions = fieldsForUsers
	case 3:
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
