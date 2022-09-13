package exchange

import (
	"context"
	"fmt"
	"strings"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
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
func GetAllMailFolders(gs graph.Service, user, nameContains string) ([]models.MailFolderable, error) {
	var (
		mfs = []models.MailFolderable{}
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

	cb := func(item any) bool {
		folder, ok := item.(models.MailFolderable)
		if !ok {
			err = errors.New("casting item to models.MailFolderable")
			return false
		}

		include := len(nameContains) == 0 ||
			(len(nameContains) > 0 && strings.Contains(*folder.GetDisplayName(), nameContains))
		if include {
			mfs = append(mfs, folder)
		}

		return true
	}

	if err := iter.Iterate(cb); err != nil {
		return nil, err
	}

	return mfs, err
}

// GetAllCalendars retrieves all event calendars for the specified user.
// If nameContains is populated, only returns calendars matching that property.
// Returns a slice of {ID, DisplayName} tuples.
func GetAllCalendars(gs graph.Service, user, nameContains string) ([]CalendarDisplayable, error) {
	var (
		cs  = []CalendarDisplayable{}
		err error
	)

	resp, err := GetAllCalendarNamesForUser(gs, user)
	if err != nil {
		return nil, err
	}

	iter, err := msgraphgocore.NewPageIterator(
		resp, gs.Adapter(), models.CreateCalendarCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}

	cb := func(item any) bool {
		cal, ok := item.(models.Calendarable)
		if !ok {
			err = errors.New("casting item to models.Calendarable")
			return false
		}

		include := len(nameContains) == 0 ||
			(len(nameContains) > 0 && strings.Contains(*cal.GetName(), nameContains))
		if include {
			cs = append(cs, *CreateCalendarDisplayable(cal))
		}

		return true
	}

	if err := iter.Iterate(cb); err != nil {
		return nil, err
	}

	return cs, err
}

// GetAllContactFolders retrieves all contacts folders for the specified user.
// If nameContains is populated, only returns folders matching that property.
// Returns a slice of {ID, DisplayName} tuples.
func GetAllContactFolders(gs graph.Service, user, nameContains string) ([]models.ContactFolderable, error) {
	var (
		cs  = []models.ContactFolderable{}
		err error
	)

	resp, err := GetAllContactFolderNamesForUser(gs, user)
	if err != nil {
		return nil, err
	}

	iter, err := msgraphgocore.NewPageIterator(
		resp, gs.Adapter(), models.CreateContactFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}

	cb := func(item any) bool {
		folder, ok := item.(models.ContactFolderable)
		if !ok {
			err = errors.New("casting item to models.ContactFolderable")
			return false
		}

		include := len(nameContains) == 0 ||
			(len(nameContains) > 0 && strings.Contains(*folder.GetDisplayName(), nameContains))
		if include {
			cs = append(cs, folder)
		}

		return true
	}

	if err := iter.Iterate(cb); err != nil {
		return nil, err
	}

	return cs, err
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

// parseCalendarIDFromEvent returns the M365 ID for a calendar
// @param reference: string from additionalData map of an event
// References should follow the form `https://... calendars('ID')/$ref`
// If the reference does not follow form an error is returned
func parseCalendarIDFromEvent(reference string) (string, error) {
	stringArray := strings.Split(reference, "calendars('")
	if len(stringArray) < 2 {
		return "", errors.New("calendarID not found")
	}

	temp := stringArray[1]
	stringArray = strings.Split(temp, "')/$ref")

	if len(stringArray) < 2 {
		return "", errors.New("calendarID not found")
	}

	calendarID := stringArray[0]

	if len(calendarID) == 0 {
		return "", errors.New("calendarID empty")
	}

	return calendarID, nil
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
			GetAllContactsForUser,
			IterateSelectAllDescendablesForCollections,
			nil
	}

	return nil, nil, nil, errors.New("exchange scope option not supported")
}

// GetRestoreContainer utility function to create
//  an unique folder for the restore process
// @param category: input from fullPath()[2]
// that defines the application the folder is created in.
func GetRestoreContainer(
	service graph.Service,
	user string,
	category path.CategoryType,
) (string, error) {
	name := fmt.Sprintf("Corso_Restore_%s", common.FormatNow(common.SimpleDateTimeFormat))
	option := categoryToOptionIdentifier(category)

	folderID, err := GetContainerID(service, name, user, option)
	if err == nil {
		return *folderID, nil
	}
	// Experienced error other than folder does not exist
	if !errors.Is(err, ErrFolderNotFound) {
		return "", support.WrapAndAppend(user, err, err)
	}

	switch option {
	case messages:
		fold, err := CreateMailFolder(service, user, name)
		if err != nil {
			return "", support.WrapAndAppend(user, err, err)
		}

		return *fold.GetId(), nil
	case contacts:
		fold, err := CreateContactFolder(service, user, name)
		if err != nil {
			return "", support.WrapAndAppend(user, err, err)
		}

		return *fold.GetId(), nil
	case events:
		calendar, err := CreateCalendar(service, user, name)
		if err != nil {
			return "", support.WrapAndAppend(user, err, err)
		}

		return *calendar.GetId(), nil
	default:
		return "", fmt.Errorf("category: %s not supported for folder creation", option)
	}
}

func RestoreExchangeObject(
	ctx context.Context,
	bits []byte,
	category path.CategoryType,
	policy control.CollisionPolicy,
	service graph.Service,
	destination, user string,
) error {
	if policy != control.Copy {
		return fmt.Errorf("restore policy: %s not supported", policy)
	}

	setting := categoryToOptionIdentifier(category)

	switch setting {
	case messages:
		return RestoreMailMessage(ctx, bits, service, control.Copy, destination, user)
	case contacts:
		return RestoreExchangeContact(ctx, bits, service, control.Copy, destination, user)
	case events:
		return RestoreExchangeEvent(ctx, bits, service, control.Copy, destination, user)
	default:
		return fmt.Errorf("type: %s not supported for exchange restore", category)
	}
}

// RestoreExchangeContact restores a contact to the @bits byte
// representation of M365 contact object.
// @destination M365 ID representing a M365 Contact_Folder
// Returns an error if the input bits do not parse into a models.Contactable object
// or if an error is encountered sending data to the M365 account.
// Post details: https://docs.microsoft.com/en-us/graph/api/user-post-contacts?view=graph-rest-1.0&tabs=go
func RestoreExchangeContact(
	ctx context.Context,
	bits []byte,
	service graph.Service,
	cp control.CollisionPolicy,
	destination, user string,
) error {
	contact, err := support.CreateContactFromBytes(bits)
	if err != nil {
		return err
	}

	response, err := service.Client().UsersById(user).ContactFoldersById(destination).Contacts().Post(contact)
	if err != nil {
		return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	if response == nil {
		return errors.New("msgraph contact post fail: REST response not received")
	}

	return nil
}

func RestoreExchangeEvent(
	ctx context.Context,
	bits []byte,
	service graph.Service,
	cp control.CollisionPolicy,
	destination, user string,
) error {
	event, err := support.CreateEventFromBytes(bits)
	if err != nil {
		return err
	}

	response, err := service.Client().UsersById(user).CalendarsById(destination).Events().Post(event)
	if err != nil {
		return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	if response == nil {
		return errors.New("msgraph event post fail: REST response not received")
	}

	return nil
}

// RestoreMailMessage utility function to place an exchange.Mail
// message into the user's M365 Exchange account.
// @param bits - byte array representation of exchange.Message from Corso backstore
// @param service - connector to M365 graph
// @param cp - collision policy that directs restore workflow
// @param destination - M365 Folder ID. Verified and sent by higher function. `copy` policy can use directly
func RestoreMailMessage(
	ctx context.Context,
	bits []byte,
	service graph.Service,
	cp control.CollisionPolicy,
	destination,
	user string,
) error {
	// Creates messageable object from original bytes
	originalMessage, err := support.CreateMessageFromBytes(bits)
	if err != nil {
		return errors.Wrapf(err, "restore mail message rcvd: %v", bits)
	}
	// Sets fields from original message from storage
	clone := support.ToMessage(originalMessage)
	valueID := RestorePropertyTag
	enableValue := RestoreCanonicalEnableValue

	// Set Extended Properties:
	// 1st: No transmission
	// 2nd: Send Date
	// 3rd: Recv Date
	sv1 := models.NewSingleValueLegacyExtendedProperty()
	sv1.SetId(&valueID)
	sv1.SetValue(&enableValue)

	sv2 := models.NewSingleValueLegacyExtendedProperty()
	sendPropertyValue := common.FormatLegacyTime(*clone.GetSentDateTime())
	sendPropertyTag := "SystemTime 0x0039"
	sv2.SetId(&sendPropertyTag)
	sv2.SetValue(&sendPropertyValue)

	sv3 := models.NewSingleValueLegacyExtendedProperty()
	recvPropertyValue := common.FormatLegacyTime(*clone.GetReceivedDateTime())
	recvPropertyTag := "SystemTime 0x0E06"
	sv3.SetId(&recvPropertyTag)
	sv3.SetValue(&recvPropertyValue)

	svlep := []models.SingleValueLegacyExtendedPropertyable{sv1, sv2, sv3}
	clone.SetSingleValueExtendedProperties(svlep)

	// Switch workflow based on collision policy
	switch cp {
	default:
		logger.Ctx(ctx).DPanicw("unrecognized restore policy; defaulting to copy",
			"policy", cp)
		fallthrough
	case control.Copy:
		return SendMailToBackStore(service, user, destination, clone)
	}
}

// SendMailToBackStore function for transporting in-memory messageable item to M365 backstore
// @param user string represents M365 ID of user within the tenant
// @param destination represents M365 ID of a folder within the users's space
// @param message is a models.Messageable interface from "github.com/microsoftgraph/msgraph-sdk-go/models"
func SendMailToBackStore(service graph.Service, user, destination string, message models.Messageable) error {
	sentMessage, err := service.Client().UsersById(user).MailFoldersById(destination).Messages().Post(message)
	if err != nil {
		return support.WrapAndAppend(": "+support.ConnectorStackErrorTrace(err), err, nil)
	}

	if sentMessage == nil {
		return errors.New("message not Sent: blocked by server")
	}

	if err != nil {
		return support.WrapAndAppend(": "+support.ConnectorStackErrorTrace(err), err, nil)
	}

	return nil
}
