package exchange

import (
	"fmt"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msfolder "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders"
	msmessage "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages"
	msitem "github.com/microsoftgraph/msgraph-sdk-go/users/item/messages/item"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/selectors"
)

type optionIdentifier int

const mailCategory = "mail"

//go:generate stringer -type=optionIdentifier
const (
	unknown optionIdentifier = iota
	folders
	messages
	users
)

// GraphQuery represents functions which perform exchange-specific queries
// into M365 backstore.
//TODO: use selector or path for granularity into specific folders or specific date ranges
type GraphQuery func(graph.Service, []string) (absser.Parsable, error)

// GetAllMessagesForUser is a GraphQuery function for receiving all messages for a single user
func GetAllMessagesForUser(gs graph.Service, identities []string) (absser.Parsable, error) {
	selecting := []string{"id", "parentFolderId"}
	options, err := optionsForMessages(selecting)

	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(identities[0]).Messages().GetWithRequestConfigurationAndResponseHandler(options, nil)
}

// GetAllFolderDisplayNamesForUser is a GraphQuery function for getting FolderId and display
// names for Mail Folder. All other information for the MailFolder object is omitted.
func GetAllFolderNamesForUser(gs graph.Service, identities []string) (absser.Parsable, error) {
	options, err := optionsForMailFolders([]string{"id", "displayName"})

	if err != nil {
		return nil, err
	}

	return gs.Client().UsersById(identities[0]).MailFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)

}

// GraphIterateFuncs are iterate functions to be used with the M365 iterators (e.g. msgraphgocore.NewPageIterator)
// @returns a callback func that works with msgraphgocore.PageIterator.Iterate function
type GraphIterateFunc func(
	string,
	selectors.ExchangeScope,
	error,
	bool,
	account.M365Config,
	map[string]*Collection,
	chan<- *support.ConnectorOperationStatus,
) func(any) bool

// IterateSelectAllMessageForCollection utility function for
// Iterating through MessagesCollectionResponse
// During iteration, messages belonging to any folder are
// placed into a Collection based on the parent folder
func IterateSelectAllMessagesForCollections(
	tenant string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	return func(messageItem any) bool {
		// Defines the type of collection being created within the function
		collection_type := messages
		user := scope.Get(selectors.ExchangeUser)[0]

		message, ok := messageItem.(models.Messageable)
		if !ok {
			errs = support.WrapAndAppendf(user, errors.New("message iteration failure"), errs)
			return true
		}
		// Saving to messages to list. Indexed by folder
		directory := *message.GetParentFolderId()
		if _, ok = collections[directory]; !ok {
			service, err := createService(credentials, failFast)
			if err != nil {
				errs = support.WrapAndAppend(user, err, errs)
				return true
			}
			edc := NewCollection(
				user,
				[]string{tenant, user, mailCategory, directory},
				collection_type,
				service,
				statusCh,
			)
			collections[directory] = &edc
		}
		collections[directory].AddJob(*message.GetId())
		return true
	}
}

func IterateAndFilterMessagesForCollections(
	tenant string,
	scope selectors.ExchangeScope,
	errs error,
	failFast bool,
	credentials account.M365Config,
	collections map[string]*Collection,
	statusCh chan<- *support.ConnectorOperationStatus,
) func(any) bool {
	var isFilterSet bool
	return func(messageItem any) bool {
		user := scope.Get(selectors.ExchangeUser)[0]
		if !isFilterSet {

			err := AddMailFolderSpecificCollections(
				scope,
				tenant,
				user,
				collections,
				credentials,
				failFast,
				statusCh,
			)
			if err != nil {
				errs = support.WrapAndAppend(user, err, errs)
				return false
			}
			isFilterSet = true
		}

		message, ok := messageItem.(models.Messageable)
		if !ok {
			errs = support.WrapAndAppend(user, errors.New("message iteration failure"), errs)
			return true
		}
		// Saving only messages for the created directories
		directory := *message.GetParentFolderId()
		if _, ok = collections[directory]; !ok {
			return true
		}
		collections[directory].AddJob(*message.GetId())
		return true

	}
}

func AddMailFolderSpecificCollections(
	scope selectors.ExchangeScope,
	tenant string,
	user string,
	collections map[string]*Collection,
	credentials account.M365Config,
	failFast bool,
	statusCh chan<- *support.ConnectorOperationStatus,
) error {
	queryService, err := createService(credentials, failFast)
	if err != nil {
		return errors.New("unable to create a mail folder query service for " + user)
	}
	folderNames := scope.Get(selectors.ExchangeMailFolder)
	query, err := GetAllFolderNamesForUser(queryService, []string{user})
	if err != nil {
		return fmt.Errorf(
			"unable to query mail folder for %s: details: %s",
			user,
			support.ConnectorStackErrorTrace(err),
		)
	}
	// Iterator required to ensure all potential folders are inspected
	// when the breadth of the folder space is large
	pageIterator, err := msgraphgocore.NewPageIterator(
		query,
		&queryService.adapter,
		models.CreateMailFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return errors.Wrap(err, "unable to create iterator during mail folder query service")
	}
	var service graph.Service
	callbackFunc := func(pageItem any) bool {
		folder, ok := pageItem.(models.MailFolderable)
		if !ok {
			err = support.WrapAndAppend(user, errors.New("unable to transform folderable item"), err)
			return true
		}
		name := *folder.GetDisplayName()
		for _, permitted := range folderNames {
			if name != permitted {
				continue
			}
			directory := *folder.GetId()
			service, err = createService(credentials, failFast)
			if err != nil {
				err = support.WrapAndAppend(
					name,
					errors.New("unable to create service a folder query service for "+user),
					err,
				)
				return true
			}
			temp := NewCollection(
				user,
				[]string{tenant, user, mailCategory, directory},
				messages,
				service,
				statusCh,
			)
			collections[directory] = &temp
			break

		}
		return true
	}
	iterateFailure := pageIterator.Iterate(callbackFunc)
	if iterateFailure != nil {
		err = support.WrapAndAppend(user+" iterate failure", iterateFailure, err)
	}
	return err

}

//---------------------------------------------------
// exchange.Query Option Section
//------------------------------------------------

// optionsForMessages - used to select allowable options for exchange.Mail types
// @param moreOps is []string of options(e.g. "parentFolderId, subject")
// @return is first call in Messages().GetWithRequestConfigurationAndResponseHandler
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

// optionsForSingleMessage to select allowable option for a singular exchange.Mail object
// @params moreOps is []string of options (e.g. subject, content.Type)
// @return is first call in MessageById().GetWithRequestConfigurationAndResponseHandler
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

// optionsForMailFolders transforms the options into a more dynamic call for MailFolders.
// @param moreOps is a []string of options(e.g. "displayName", "isHidden")
// @return is first call in MailFolders().GetWithRequestConfigurationAndResponseHandler(options, handler)
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

// buildOptions - Utility Method for verifying if select options are valid for the m365 object type
// @return is a pair. The first is a string literal of allowable options based on the object type,
// the second is an error. An error is returned if an unsupported option or optionIdentifier was used
func buildOptions(options []string, optId optionIdentifier) ([]string, error) {
	var allowedOptions map[string]int

	fieldsForFolders := map[string]int{
		"displayName":    1,
		"isHidden":       2,
		"parentFolderId": 3,
		"id":             4,
	}

	fieldsForUsers := map[string]int{
		"birthday":       1,
		"businessPhones": 2,
		"city":           3,
		"companyName":    4,
		"department":     5,
		"displayName":    6,
		"employeeId":     7,
		"id":             8,
	}

	fieldsForMessages := map[string]int{
		"conservationId":    1,
		"conversationIndex": 2,
		"parentFolderId":    3,
		"subject":           4,
		"webLink":           5,
		"id":                6,
	}
	returnedOptions := []string{"id"}

	switch optId {
	case folders:
		allowedOptions = fieldsForFolders
	case users:
		allowedOptions = fieldsForUsers
	case messages:
		allowedOptions = fieldsForMessages
	case unknown:
		fallthrough
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
