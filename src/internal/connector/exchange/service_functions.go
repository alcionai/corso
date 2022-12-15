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
	credentials account.M365Config
}

// ------------------------------------------------------------
// Functions to comply with graph.Servicer Interface
// ------------------------------------------------------------

func (es *exchangeService) Client() *msgraphsdk.GraphServiceClient {
	return &es.client
}

func (es *exchangeService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &es.adapter
}

// createService internal constructor for exchangeService struct returns an error
// iff the params for the entry are incorrect (e.g. len(TenantID) == 0, etc.)
// NOTE: Incorrect account information will result in errors on subsequent queries.
func createService(credentials account.M365Config) (*exchangeService, error) {
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
		credentials: credentials,
	}

	return &service, nil
}

// CreateMailFolder makes a mail folder iff a folder of the same name does not exist
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-mailfolders?view=graph-rest-1.0&tabs=http
func CreateMailFolder(ctx context.Context, gs graph.Servicer, user, folder string) (models.MailFolderable, error) {
	isHidden := false
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	requestBody.SetIsHidden(&isHidden)

	return gs.Client().UsersById(user).MailFolders().Post(ctx, requestBody, nil)
}

func CreateMailFolderWithParent(
	ctx context.Context,
	gs graph.Servicer,
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
func DeleteMailFolder(ctx context.Context, gs graph.Servicer, user, folderID string) error {
	return gs.Client().UsersById(user).MailFoldersById(folderID).Delete(ctx, nil)
}

// CreateCalendar makes an event Calendar with the name in the user's M365 exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-calendars?view=graph-rest-1.0&tabs=go
func CreateCalendar(ctx context.Context, gs graph.Servicer, user, calendarName string) (models.Calendarable, error) {
	requestbody := models.NewCalendar()
	requestbody.SetName(&calendarName)

	return gs.Client().UsersById(user).Calendars().Post(ctx, requestbody, nil)
}

// DeleteCalendar removes calendar from user's M365 account
// Reference: https://docs.microsoft.com/en-us/graph/api/calendar-delete?view=graph-rest-1.0&tabs=go
func DeleteCalendar(ctx context.Context, gs graph.Servicer, user, calendarID string) error {
	return gs.Client().UsersById(user).CalendarsById(calendarID).Delete(ctx, nil)
}

// CreateContactFolder makes a contact folder with the displayName of folderName.
// If successful, returns the created folder object.
func CreateContactFolder(
	ctx context.Context,
	gs graph.Servicer,
	user, folderName string,
) (models.ContactFolderable, error) {
	requestBody := models.NewContactFolder()
	temp := folderName
	requestBody.SetDisplayName(&temp)

	return gs.Client().UsersById(user).ContactFolders().Post(ctx, requestBody, nil)
}

// DeleteContactFolder deletes the ContactFolder associated with the M365 ID if permissions are valid.
// Errors returned if the function call was not successful.
func DeleteContactFolder(ctx context.Context, gs graph.Servicer, user, folderID string) error {
	return gs.Client().UsersById(user).ContactFoldersById(folderID).Delete(ctx, nil)
}

// populateExchangeContainerResolver gets a folder resolver if one is available for
// this category of data. If one is not available, returns nil so that other
// logic in the caller can complete as long as they check if the resolver is not
// nil. If an error occurs populating the resolver, returns an error.
func populateExchangeContainerResolver(
	ctx context.Context,
	qp graph.QueryParams,
) (graph.ContainerResolver, error) {
	var (
		res          graph.ContainerResolver
		cacheRoot    string
		service, err = createService(qp.Credentials)
	)

	if err != nil {
		return nil, err
	}

	switch qp.Category {
	case path.EmailCategory:
		res = &mailFolderCache{
			userID: qp.ResourceOwner,
			gs:     service,
		}
		cacheRoot = rootFolderAlias

	case path.ContactsCategory:
		res = &contactFolderCache{
			userID: qp.ResourceOwner,
			gs:     service,
		}
		cacheRoot = DefaultContactFolder

	case path.EventsCategory:
		res = &eventCalendarCache{
			userID: qp.ResourceOwner,
			gs:     service,
		}
		cacheRoot = DefaultCalendar

	default:
		return nil, fmt.Errorf("ContainerResolver not present for %s type", qp.Category)
	}

	if err := res.Populate(ctx, cacheRoot); err != nil {
		return nil, errors.Wrap(err, "populating directory resolver")
	}

	return res, nil
}

func pathAndMatch(
	qp graph.QueryParams,
	c graph.CachedContainer,
	scope selectors.ExchangeScope,
) (path.Path, bool) {
	var (
		category  = scope.Category().PathType()
		directory string
		pb        = c.Path()
	)

	if c.Deleted() {
		return nil, true
	}

	// Clause ensures that DefaultContactFolder is inspected properly
	if category == path.ContactsCategory && *c.GetDisplayName() == DefaultContactFolder {
		pb = c.Path().Append(DefaultContactFolder)
	}

	dirPath, err := pb.ToDataLayerExchangePathForCategory(
		qp.Credentials.AzureTenantID,
		qp.ResourceOwner,
		category,
		false,
	)
	// Containers without a path (e.g. Root mail folder) always err here.
	if err != nil {
		return nil, false
	}

	directory = pb.String()

	switch category {
	case path.EmailCategory:
		return dirPath, scope.Matches(selectors.ExchangeMailFolder, directory)
	case path.ContactsCategory:
		return dirPath, scope.Matches(selectors.ExchangeContactFolder, directory)
	case path.EventsCategory:
		return dirPath, scope.Matches(selectors.ExchangeEventCalendar, directory)
	default:
		return nil, false
	}
}
