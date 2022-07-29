package exchange

import (
	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
	msmessage "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
)

type optionIdentifier int

//go:generate stringer -type=optionIdentifier
const (
	unknown optionIdentifier = iota
	folders
	messages
	users
)

type GraphQuery func(graph.Service, []string) (absser.Parsable, error)

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

// CheckOptions Utility Method for verifying if select options are valid the m365 object type
// returns a list of valid options
func buildOptions(options []string, selection optionIdentifier) ([]string, error) {
	var allowedOptions map[string]int

	fieldsForFolders := map[string]int{
		"displayName":    1,
		"isHidden":       2,
		"parentFolderId": 3,
	}

	fieldsForUsers := map[string]int{
		"birthday":       1,
		"businessPhones": 2,
		"city":           3,
		"companyName":    4,
		"department":     5,
		"displayName":    6,
		"employeeId":     7,
	}

	fieldsForMessages := map[string]int{
		"conservationId":    1,
		"conversationIndex": 2,
		"parentFolderId":    3,
		"subject":           4,
		"webLink":           5,
	}
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
		_, ok := allowedOptions[entry]
		if ok {
			returnedOptions = append(returnedOptions, entry)
		} else {
			return nil, errors.New("unsupported option")
		}
	}
	return returnedOptions, nil
}

func GetAllMessagesForUser(gs graph.Service, identities []string) (absser.Parsable, error) {
	selecting := []string{"id", "parentFolderId"}
	options, err := optionsForMessages(selecting)
	if err != nil {
		return nil, err
	}
	return gs.Client().UsersById(identities[0]).Messages().GetWithRequestConfigurationAndResponseHandler(options, nil)
}
