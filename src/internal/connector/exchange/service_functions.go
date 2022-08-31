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

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/pkg/account"
	"github.com/alcionai/corso/pkg/control"
	"github.com/alcionai/corso/pkg/logger"
	"github.com/alcionai/corso/pkg/selectors"
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
		transform = models.CreateContactFolderFromDiscriminatorValue
	case events:
		query = GetAllCalendarNamesForUser
		transform = models.CreateCalendarCollectionResponseFromDiscriminatorValue
		isCalendar = true
	default:
		return nil, fmt.Errorf("unsupported category %s for GetFolderID()", category)
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

	callbackFunc := iterateFindFolderID(
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
		return models.CreateEventCollectionResponseFromDiscriminatorValue,
			GetAllEventsForUser,
			IterateSelectAllEventsForCollections,
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

// GetRestoreFolder utility function to create
//  an unique folder for the restore process
// @param category: input from fullPath()[2]
// that defines the application the folder is created in.
func GetRestoreFolder(
	service graph.Service,
	user, category string,
) (string, error) {
	newFolder := fmt.Sprintf("Corso_Restore_%s", common.FormatNow(common.SimpleDateTimeFormat))

	switch category {
	case mailCategory, contactsCategory:
		return establishFolder(service, newFolder, user, categoryToOptionIdentifier(category))
	default:
		return "", fmt.Errorf("%s category not supported", category)
	}
}

func establishFolder(
	service graph.Service,
	folderName, user string,
	optID optionIdentifier,
) (string, error) {
	folderID, err := GetContainerID(service, folderName, user, optID)
	if err == nil {
		return *folderID, nil
	}
	// Experienced error other than folder does not exist
	if !errors.Is(err, ErrFolderNotFound) {
		return "", support.WrapAndAppend(user, err, err)
	}

	switch optID {
	case messages:
		fold, err := CreateMailFolder(service, user, folderName)
		if err != nil {
			return "", support.WrapAndAppend(user, err, err)
		}

		return *fold.GetId(), nil
	case contacts:
		fold, err := CreateContactFolder(service, user, folderName)
		if err != nil {
			return "", support.WrapAndAppend(user, err, err)
		}

		return *fold.GetId(), nil
	default:
		return "", fmt.Errorf("category: %s not supported for folder creation", optID)
	}
}

func RestoreExchangeObject(
	ctx context.Context,
	bits []byte,
	category string,
	policy control.CollisionPolicy,
	service graph.Service,
	destination, user string,
) error {
	var setting optionIdentifier

	switch category {
	case mailCategory, contactsCategory:
		setting = categoryToOptionIdentifier(category)
	default:
		return fmt.Errorf("type: %s not supported for exchange restore", category)
	}

	if policy != control.Copy {
		return fmt.Errorf("restore policy: %s not supported", policy)
	}

	switch setting {
	case messages:
		return RestoreMailMessage(ctx, bits, service, control.Copy, destination, user)
	case contacts:
		return RestoreExchangeContact(ctx, bits, service, control.Copy, destination, user)
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
		return err
	}
	// Sets fields from original message from storage
	clone := support.ToMessage(originalMessage)
	valueID := RestorePropertyTag
	enableValue := RestoreCanonicalEnableValue
	sv := models.NewSingleValueLegacyExtendedProperty()
	sv.SetId(&valueID)
	sv.SetValue(&enableValue)
	svlep := []models.SingleValueLegacyExtendedPropertyable{sv}
	clone.SetSingleValueExtendedProperties(svlep)

	draft := false
	clone.SetIsDraft(&draft)

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

	return nil
}
