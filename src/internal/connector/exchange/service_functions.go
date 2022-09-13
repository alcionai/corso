package exchange

import (
	"fmt"
	"strings"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var ErrFolderNotFound = errors.New("folder not found")

type exchangeService struct {
	client      msgraphsdk.GraphServiceClient
	adapter     msgraphsdk.GraphRequestAdapter
	failFast    bool // if true service will exit sequence upon encountering an error
	credentials account.M365Config
}

///------------------------------------------------------------
// Functions to comply with graph.Service Interface
//-------------------------------------------------------
func (es *exchangeService) Client() *msgraphsdk.GraphServiceClient {
	return &es.client
}

func (es *exchangeService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &es.adapter
}

func (es *exchangeService) ErrPolicy() bool {
	return es.failFast
}

// createService internal constructor for exchangeService struct returns an error
// iff the params for the entry are incorrect (e.g. len(TenantID) == 0, etc.)
// NOTE: Incorrect account information will result in errors on subsequent queries.
func createService(credentials account.M365Config, shouldFailFast bool) (*exchangeService, error) {
	adapter, err := graph.CreateAdapter(
		credentials.TenantID,
		credentials.ClientID,
		credentials.ClientSecret,
	)
	if err != nil {
		return nil, err
	}

	service := exchangeService{
		adapter:     *adapter,
		client:      *msgraphsdk.NewGraphServiceClient(adapter),
		failFast:    shouldFailFast,
		credentials: credentials,
	}

	return &service, err
}

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

type MailFolder struct {
	ID          string
	DisplayName string
}

// CreateCalendar makes an event Calendar with the name in the user's M365 exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-calendars?view=graph-rest-1.0&tabs=go
func CreateCalendar(gs graph.Service, user, calendarName string) (models.Calendarable, error) {
	requestbody := models.NewCalendar()
	requestbody.SetName(&calendarName)

	return gs.Client().UsersById(user).Calendars().Post(requestbody)
}

// DeleteCalendar removes calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func DeleteCalendar(gs graph.Service, user, calendarID string) error {
	return gs.Client().UsersById(user).CalendarsById(calendarID).Delete()
}

// CreateContactFolder makes a contact folder with the displayName of folderName.
// If successful, returns the created folder object.
func CreateContactFolder(gs graph.Service, user, folderName string) (models.ContactFolderable, error) {
	requestBody := models.NewContactFolder()
	requestBody.SetDisplayName(&folderName)

	return gs.Client().UsersById(user).ContactFolders().Post(requestBody)
}

// DeleteContactFolder deletes the ContactFolder associated with the M365 ID if permissions are valid.
// Errors returned if the function call was not successful.
func DeleteContactFolder(gs graph.Service, user, folderID string) error {
	return gs.Client().UsersById(user).ContactFoldersById(folderID).Delete()
}

// GetAllMailFolders retrieves all mail folders for the specified user.
// If nameContains is populated, only returns mail matching that property.
// Returns a slice of {ID, DisplayName} tuples.
func GetAllMailFolders(gs graph.Service, user, nameContains string) ([]MailFolder, error) {
	var (
		mfs = []MailFolder{}
		err error
	)

	resp, err := GetAllFolderNamesForUser(gs, user)
	if err != nil {
		return nil, err
	}

	iter, err := msgraphgocore.NewPageIterator(
		resp, gs.Adapter(), models.CreateMailFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}

	cb := func(folderItem any) bool {
		folder, ok := folderItem.(models.MailFolderable)
		if !ok {
			err = errors.New("HasFolder() iteration failure")
			return false
		}

		include := len(nameContains) == 0 ||
			(len(nameContains) > 0 && strings.Contains(*folder.GetDisplayName(), nameContains))
		if include {
			mfs = append(mfs, MailFolder{
				ID:          *folder.GetId(),
				DisplayName: *folder.GetDisplayName(),
			})
		}

		return true
	}

	if err := iter.Iterate(cb); err != nil {
		return nil, err
	}

	return mfs, err
}

// GetContainerID query function to retrieve a container's M365 ID.
// @param containerName is the target's name, user-readable and case sensitive
// @param category switches query and iteration to support  multiple exchange applications
// @returns a *string if the folder exists. If the folder does not exist returns nil, error-> folder not found
func GetContainerID(service graph.Service, containerName, user string, category optionIdentifier) (*string, error) {
	var (
		errs       error
		targetID   *string
		query      GraphQuery
		transform  absser.ParsableFactory
		isCalendar bool
	)

	switch category {
	case messages:
		query = GetAllFolderNamesForUser
		transform = models.CreateMailFolderCollectionResponseFromDiscriminatorValue
	case contacts:
		query = GetAllContactFolderNamesForUser
		transform = models.CreateContactFolderCollectionResponseFromDiscriminatorValue
	case events:
		query = GetAllCalendarNamesForUser
		transform = models.CreateCalendarCollectionResponseFromDiscriminatorValue
		isCalendar = true
	default:
		return nil, fmt.Errorf("unsupported category %s for GetContainerID()", category)
	}

	response, err := query(service, user)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"user %s M365 query: %s",
			user, support.ConnectorStackErrorTrace(err),
		)
	}

	pageIterator, err := msgraphgocore.NewPageIterator(
		response,
		service.Adapter(),
		transform,
	)
	if err != nil {
		return nil, err
	}

	callbackFunc := iterateFindContainerID(
		&targetID,
		containerName,
		service.Adapter().GetBaseUrl(),
		isCalendar,
		errs,
	)

	if err := pageIterator.Iterate(callbackFunc); err != nil {
		return nil, support.WrapAndAppend(service.Adapter().GetBaseUrl(), err, errs)
	}

	if targetID == nil {
		return nil, ErrFolderNotFound
	}

	return targetID, errs
}

// SetupExchangeCollectionVars is a helper function returns a sets
// Exchange.Type specific functions based on scope
func SetupExchangeCollectionVars(scope selectors.ExchangeScope) (
	absser.ParsableFactory,
	GraphQuery,
	GraphIterateFunc,
	error,
) {
	if scope.IncludesCategory(selectors.ExchangeMail) {
		if scope.IsAny(selectors.ExchangeMailFolder) {
			return models.CreateMessageCollectionResponseFromDiscriminatorValue,
				GetAllMessagesForUser,
				IterateSelectAllDescendablesForCollections,
				nil
		}

		return models.CreateMessageCollectionResponseFromDiscriminatorValue,
			GetAllMessagesForUser,
			IterateAndFilterMessagesForCollections,
			nil
	}

	if scope.IncludesCategory(selectors.ExchangeEvent) {
		return models.CreateCalendarCollectionResponseFromDiscriminatorValue,
			GetAllCalendarNamesForUser,
			IterateSelectAllEventsFromCalendars,
			nil
	}

	if scope.IncludesCategory(selectors.ExchangeContact) {
		return models.CreateContactFromDiscriminatorValue,
			GetAllContactFolderNamesForUser,
			IterateSelectAllContactsForCollections,
			nil
	}

	return nil, nil, nil, errors.New("exchange scope option not supported")
}
