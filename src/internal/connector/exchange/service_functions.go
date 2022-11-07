package exchange

import (
	"context"
	"fmt"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
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
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret,
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating microsoft graph service for exchange")
	}

	service := exchangeService{
		adapter:     *adapter,
		client:      *msgraphsdk.NewGraphServiceClient(adapter),
		failFast:    shouldFailFast,
		credentials: credentials,
	}

	return &service, nil
}

// CreateMailFolder makes a mail folder iff a folder of the same name does not exist
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-mailfolders?view=graph-rest-1.0&tabs=http
func CreateMailFolder(ctx context.Context, gs graph.Service, user, folder string) (models.MailFolderable, error) {
	isHidden := false
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	requestBody.SetIsHidden(&isHidden)

	return gs.Client().UsersById(user).MailFolders().Post(ctx, requestBody, nil)
}

func CreateMailFolderWithParent(
	ctx context.Context,
	gs graph.Service,
	user, folder, parentID string,
) (models.MailFolderable, error) {
	isHidden := false
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	requestBody.SetIsHidden(&isHidden)

	return gs.Client().
		UsersById(user).
		MailFoldersById(parentID).
		ChildFolders().
		Post(ctx, requestBody, nil)
}

// DeleteMailFolder removes a mail folder with the corresponding M365 ID  from the user's M365 Exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/mailfolder-delete?view=graph-rest-1.0&tabs=http
func DeleteMailFolder(ctx context.Context, gs graph.Service, user, folderID string) error {
	return gs.Client().UsersById(user).MailFoldersById(folderID).Delete(ctx, nil)
}

// CreateCalendar makes an event Calendar with the name in the user's M365 exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-calendars?view=graph-rest-1.0&tabs=go
func CreateCalendar(ctx context.Context, gs graph.Service, user, calendarName string) (models.Calendarable, error) {
	requestbody := models.NewCalendar()
	requestbody.SetName(&calendarName)

	return gs.Client().UsersById(user).Calendars().Post(ctx, requestbody, nil)
}

// DeleteCalendar removes calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func DeleteCalendar(ctx context.Context, gs graph.Service, user, calendarID string) error {
	return gs.Client().UsersById(user).CalendarsById(calendarID).Delete(ctx, nil)
}

// CreateContactFolder makes a contact folder with the displayName of folderName.
// If successful, returns the created folder object.
func CreateContactFolder(
	ctx context.Context,
	gs graph.Service,
	user, folderName string,
) (models.ContactFolderable, error) {
	requestBody := models.NewContactFolder()
	temp := folderName
	requestBody.SetDisplayName(&temp)

	return gs.Client().UsersById(user).ContactFolders().Post(ctx, requestBody, nil)
}

// DeleteContactFolder deletes the ContactFolder associated with the M365 ID if permissions are valid.
// Errors returned if the function call was not successful.
func DeleteContactFolder(ctx context.Context, gs graph.Service, user, folderID string) error {
	return gs.Client().UsersById(user).ContactFoldersById(folderID).Delete(ctx, nil)
}

// GetAllMailFolders retrieves all mail folders for the specified user.
// If nameContains is populated, only returns mail matching that property.
// Returns a slice of {ID, DisplayName} tuples.
func GetAllMailFolders(
	ctx context.Context,
	qp graph.QueryParams,
	gs graph.Service,
) ([]graph.CachedContainer, error) {
	containers := make([]graph.CachedContainer, 0)

	resolver, err := PopulateExchangeContainerResolver(ctx, qp, path.EmailCategory)
	if err != nil {
		return nil, errors.Wrap(err, "building directory resolver in GetAllMailFolders")
	}

	for _, c := range resolver.Items() {
		directory := c.Path().String()
		if len(directory) == 0 {
			continue
		}

		if qp.Scope.Matches(selectors.ExchangeMailFolder, directory) {
			containers = append(containers, c)
		}
	}

	return containers, nil
}

// GetAllCalendars retrieves all event calendars for the specified user.
// If nameContains is populated, only returns calendars matching that property.
// Returns a slice of {ID, DisplayName} tuples.
func GetAllCalendars(
	ctx context.Context,
	qp graph.QueryParams,
	gs graph.Service,
) ([]graph.CachedContainer, error) {
	containers := make([]graph.CachedContainer, 0)

	resolver, err := PopulateExchangeContainerResolver(ctx, qp, path.EventsCategory)
	if err != nil {
		return nil, errors.Wrap(err, "building calendar resolver in GetAllCalendars")
	}

	for _, c := range resolver.Items() {
		directory := c.Path().String()

		if qp.Scope.Matches(selectors.ExchangeEventCalendar, directory) {
			containers = append(containers, c)
		}
	}

	return containers, err
}

// GetAllContactFolders retrieves all contacts folders with a unique display
// name for the specified user. If multiple folders have the same display name
// the result is undefined. TODO: Replace with Cache Usage
// https://github.com/alcionai/corso/issues/1122
func GetAllContactFolders(
	ctx context.Context,
	qp graph.QueryParams,
	gs graph.Service,
) ([]graph.CachedContainer, error) {
	var query string

	containers := make([]graph.CachedContainer, 0)

	resolver, err := PopulateExchangeContainerResolver(ctx, qp, path.ContactsCategory)
	if err != nil {
		return nil, errors.Wrap(err, "building directory resolver in GetAllContactFolders")
	}

	for _, c := range resolver.Items() {
		directory := c.Path().String()

		if len(directory) == 0 {
			query = DefaultContactFolder
		} else {
			query = directory
		}

		if qp.Scope.Matches(selectors.ExchangeContactFolder, query) {
			containers = append(containers, c)
		}
	}

	return containers, err
}

func GetContainers(
	ctx context.Context,
	qp graph.QueryParams,
	gs graph.Service,
) ([]graph.CachedContainer, error) {
	category := qp.Scope.Category().PathType()

	switch category {
	case path.ContactsCategory:
		return GetAllContactFolders(ctx, qp, gs)
	case path.EmailCategory:
		return GetAllMailFolders(ctx, qp, gs)
	case path.EventsCategory:
		return GetAllCalendars(ctx, qp, gs)
	default:
		return nil, fmt.Errorf("path.Category %s not supported", category)
	}
}

// PopulateExchangeContainerResolver gets a folder resolver if one is available for
// this category of data. If one is not available, returns nil so that other
// logic in the caller can complete as long as they check if the resolver is not
// nil. If an error occurs populating the resolver, returns an error.
func PopulateExchangeContainerResolver(
	ctx context.Context,
	qp graph.QueryParams,
	category path.CategoryType,
) (graph.ContainerResolver, error) {
	var (
		res          graph.ContainerResolver
		cacheRoot    string
		service, err = createService(qp.Credentials, qp.FailFast)
	)

	if err != nil {
		return nil, err
	}

	switch category {
	case path.EmailCategory:
		res = &mailFolderCache{
			userID: qp.User,
			gs:     service,
		}
		cacheRoot = rootFolderAlias

	case path.ContactsCategory:
		res = &contactFolderCache{
			userID: qp.User,
			gs:     service,
		}
		cacheRoot = DefaultContactFolder

	case path.EventsCategory:
		res = &eventCalendarCache{
			userID: qp.User,
			gs:     service,
		}
		cacheRoot = DefaultCalendar

	default:
		return nil, fmt.Errorf("ContainerResolver not present for %s type", category)
	}

	if err := res.Populate(ctx, cacheRoot); err != nil {
		return nil, errors.Wrap(err, "populating directory resolver")
	}

	return res, nil
}

func pathAndMatch(qp graph.QueryParams, category path.CategoryType, c graph.CachedContainer) (path.Path, bool) {
	var (
		directory string
		pb        = c.Path()
	)

	// Clause ensures that DefaultContactFolder is inspected properly
	if category == path.ContactsCategory && *c.GetDisplayName() == DefaultContactFolder {
		pb = c.Path().Append(DefaultContactFolder)
	}

	dirPath, err := pb.ToDataLayerExchangePathForCategory(
		qp.Credentials.AzureTenantID,
		qp.User,
		category,
		false,
	)
	if err != nil {
		return nil, false
	}

	if dirPath == nil && category == path.EmailCategory {
		return nil, false // Only true for root mail folder
	}

	directory = pb.String()

	switch category {
	case path.EmailCategory:
		return dirPath, qp.Scope.Matches(selectors.ExchangeMailFolder, directory)
	case path.ContactsCategory:
		return dirPath, qp.Scope.Matches(selectors.ExchangeContactFolder, directory)
	case path.EventsCategory:
		return dirPath, qp.Scope.Matches(selectors.ExchangeEventCalendar, directory)
	default:
		return nil, false
	}
}
