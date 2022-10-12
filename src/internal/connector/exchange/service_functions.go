package exchange

import (
	"context"
	"fmt"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
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
		return nil, err
	}

	for _, c := range resolver.GetCacheFolders() {
		temp, _ := c.Path().ToDataLayerExchangePathForCategory(
			"not",
			"used",
			path.EmailCategory,
			false,
		)
		directories := temp.Folders()

		if qp.Scope.Matches(selectors.ExchangeMailFolder, directories[len(directories)-1]) {
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
		return nil, err
	}

	for _, c := range resolver.GetCacheFolders() {
		temp, _ := c.Path().ToDataLayerExchangePathForCategory(
			"not",
			"used",
			path.EventsCategory,
			false,
		)
		directories := temp.Folders()

		if qp.Scope.Matches(selectors.ExchangeEventCalendar, directories[len(directories)-1]) {
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
	containers := make([]graph.CachedContainer, 0)

	resolver, err := PopulateExchangeContainerResolver(ctx, qp, path.ContactsCategory)
	if err != nil {
		return nil, err
	}

	for _, c := range resolver.GetCacheFolders() {
		temp := c.Path()
		p, _ := temp.ToDataLayerExchangePathForCategory(
			"not",
			"used",
			path.ContactsCategory,
			false)
		directories := p.Folders()

		if qp.Scope.Matches(selectors.ExchangeContactFolder, directories[len(directories)-1]) {
			containers = append(containers, c)
		}
	}

	return containers, err
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
			IterateAndFilterDescendablesForCollections,
			nil
	}

	if scope.IncludesCategory(selectors.ExchangeContact) {
		return models.CreateContactFolderCollectionResponseFromDiscriminatorValue,
			GetAllContactFolderNamesForUser,
			IterateSelectAllContactsForCollections,
			nil
	}

	if scope.IncludesCategory(selectors.ExchangeEvent) {
		return models.CreateCalendarCollectionResponseFromDiscriminatorValue,
			GetAllCalendarNamesForUser,
			IterateSelectAllEventsFromCalendars,
			nil
	}

	return nil, nil, nil, errors.New("exchange scope option not supported")
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

	if err := res.Populate(ctx, cacheRoot, cacheRoot); err != nil {
		return nil, errors.Wrap(err, "populating directory resolver")
	}

	return res, nil
}

func resolveCollectionPath(
	ctx context.Context,
	resolver graph.ContainerResolver,
	tenantID, user, folderID string,
	category path.CategoryType,
) (path.Path, error) {
	if resolver == nil {
		// Allows caller to default to old-style path.
		return nil, errors.WithStack(errNilResolver)
	}

	p, err := resolver.IDToPath(ctx, folderID)
	if err != nil {
		return nil, errors.Wrap(err, "resolving folder ID")
	}

	return p.ToDataLayerExchangePathForCategory(
		tenantID,
		user,
		category,
		false,
	)
}

func getCollectionPath(
	ctx context.Context,
	qp graph.QueryParams,
	resolver graph.ContainerResolver,
	directory string,
	category path.CategoryType,
) (path.Path, error) {
	returnPath, err := resolveCollectionPath(
		ctx,
		resolver,
		qp.Credentials.TenantID,
		qp.User,
		directory,
		category,
	)
	if err == nil {
		return returnPath, nil
	}

	aPath, err1 := path.Builder{}.Append(directory).
		ToDataLayerExchangePathForCategory(
			qp.Credentials.TenantID,
			qp.User,
			category,
			false,
		)
	if err1 == nil {
		return aPath, nil
	}

	return nil,
		support.WrapAndAppend(
			fmt.Sprintf(
				"both path generate functions failed for %s:%s:%s",
				qp.User,
				category,
				directory),
			err,
			err1,
		)
}
