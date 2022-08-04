package exchange

import (
	"context"
	"fmt"
	"strings"
	"time"

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

func (es *exchangeService) Client() *msgraphsdk.GraphServiceClient {
	return &es.client
}

func (es *exchangeService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &es.adapter
}

func (es *exchangeService) ErrPolicy() bool {
	return es.failFast
}

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

// GetAllMailFolders retrieves all mail folders for the specified user.
// If nameContains is populated, only returns mail matching that property.
// Returns a slice of {ID, DisplayName} tuples.
func GetAllMailFolders(gs graph.Service, user, nameContains string) ([]MailFolder, error) {
	var (
		mfs = []MailFolder{}
		err error
	)

	opts, err := optionsForMailFolders([]string{"id", "displayName"})
	if err != nil {
		return nil, err
	}

	resp, err := gs.Client().UsersById(user).MailFolders().GetWithRequestConfigurationAndResponseHandler(opts, nil)
	if err != nil {
		return nil, err
	}

	iter, err := msgraphgocore.NewPageIterator(resp, gs.Adapter(), models.CreateMailFolderCollectionResponseFromDiscriminatorValue)
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

// GetMailFolderID query function to retrieve the M365 ID based on the folder's displayName.
// @param folderName the target folder's display name. Case sensitive
// @returns a *string if the folder exists. If the folder does not exist returns nil, error-> folder not found
func GetMailFolderID(service graph.Service, folderName, user string) (*string, error) {
	var errs error
	var folderID *string
	options, err := optionsForMailFolders([]string{"displayName"})
	if err != nil {
		return nil, err
	}
	response, err := service.Client().UsersById(user).MailFolders().GetWithRequestConfigurationAndResponseHandler(options, nil)
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, errors.New("mail folder query to m365 back store returned nil")
	}
	pageIterator, err := msgraphgocore.NewPageIterator(response, service.Adapter(), models.CreateMailFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, err
	}
	callbackFunc := func(folderItem any) bool {
		folder, ok := folderItem.(models.MailFolderable)
		if !ok {
			errs = support.WrapAndAppend(service.Adapter().GetBaseUrl(), errors.New("HasFolder() iteration failure"), errs)
			return true
		}
		if *folder.GetDisplayName() == folderName {
			folderID = folder.GetId()
			return false
		}
		return true
	}
	iterateError := pageIterator.Iterate(callbackFunc)
	if iterateError != nil {
		errs = support.WrapAndAppend(service.Adapter().GetBaseUrl(), iterateError, errs)
	} else if folderID == nil {
		return nil, ErrFolderNotFound
	}
	return folderID, errs

}

// SetupExchangeCollectionVars is a helper function returns a sets
// Exchange.Type specific functions based on scope
func SetupExchangeCollectionVars(scope selectors.ExchangeScope) (
	absser.ParsableFactory,
	GraphQuery,
	GraphIterateFunc,
) {
	if scope.IncludesCategory(selectors.ExchangeMail) {

		return models.CreateMessageCollectionResponseFromDiscriminatorValue,
			GetAllMessagesForUser,
			IterateSelectAllMessagesForCollections
	}
	return nil, nil, nil

}

// GetCopyRestoreFolder utility function to create an unique folder for the restore process
func GetCopyRestoreFolder(service graph.Service, user string) (*string, error) {
	now := time.Now().UTC()
	newFolder := fmt.Sprintf("Corso_Restore_%s", common.FormatSimpleDateTime(now))
	isFolder, err := GetMailFolderID(service, newFolder, user)
	if err != nil {
		// Verify unique folder was not found
		if errors.Is(err, ErrFolderNotFound) {

			fold, err := CreateMailFolder(service, user, newFolder)
			if err != nil {
				return nil, support.WrapAndAppend(user, err, err)
			}
			return fold.GetId(), nil
		}

		return nil, err

	}
	return isFolder, nil
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
