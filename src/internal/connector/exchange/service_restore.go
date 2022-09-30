package exchange

import (
	"bytes"
	"context"
	"fmt"
	"runtime/trace"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// GetRestoreContainer utility function to create
//  an unique folder for the restore process
// @param category: input from fullPath()[2]
// that defines the application the folder is created in.
func GetRestoreContainer(
	ctx context.Context,
	service graph.Service,
	user string,
	category path.CategoryType,
	name string,
) (string, error) {
	option := categoryToOptionIdentifier(category)

	folderID, err := GetContainerID(ctx, service, name, user, option)
	if err == nil {
		return *folderID, nil
	}
	// Experienced error other than folder does not exist
	if !errors.Is(err, ErrFolderNotFound) {
		return "", support.WrapAndAppend(user+": lookup failue during GetContainerID", err, err)
	}

	switch option {
	case messages:
		fold, err := CreateMailFolder(ctx, service, user, name)
		if err != nil {
			return "", support.WrapAndAppend(fmt.Sprintf("creating folder %s for user %s", name, user), err, err)
		}

		return *fold.GetId(), nil
	case contacts:
		fold, err := CreateContactFolder(ctx, service, user, name)
		if err != nil {
			return "", support.WrapAndAppend(user+"failure during CreateContactFolder during restore Contact", err, err)
		}

		return *fold.GetId(), nil
	case events:
		calendar, err := CreateCalendar(ctx, service, user, name)
		if err != nil {
			return "", support.WrapAndAppend(user+"failure during CreateCalendar during restore Event", err, err)
		}

		return *calendar.GetId(), nil
	default:
		return "", fmt.Errorf("category: %s not supported for folder creation: GetRestoreContainer", option)
	}
}

// RestoreExchangeObject directs restore pipeline towards restore function
// based on the path.CategoryType. All input params are necessary to perform
// the type-specific restore function.
func RestoreExchangeObject(
	ctx context.Context,
	bits []byte,
	category path.CategoryType,
	policy control.CollisionPolicy,
	service graph.Service,
	destination, user string,
) error {
	if policy != control.Copy {
		return fmt.Errorf("restore policy: %s not supported for RestoreExchangeObject", policy)
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
		return fmt.Errorf("type: %s not supported for RestoreExchangeObject", category)
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
		return errors.Wrap(err, "failure to create contact from bytes: RestoreExchangeContact")
	}

	response, err := service.Client().UsersById(user).ContactFoldersById(destination).Contacts().Post(ctx, contact, nil)
	if err != nil {
		name := *contact.GetGivenName()

		return errors.Wrap(
			err,
			"failure to create Contact during RestoreExchangeContact: "+name+" "+
				support.ConnectorStackErrorTrace(err),
		)
	}

	if response == nil {
		return errors.New("msgraph contact post fail: REST response not received")
	}

	return nil
}

// RestoreExchangeEvent restores a contact to the @bits byte
// representation of M365 event object.
// @param destination is the M365 ID representing Calendar that will receive the event.
// Returns an error if input byte array doesn't parse into models.Eventable object
// or if an error occurs during sending data to M365 account.
// Post details: https://docs.microsoft.com/en-us/graph/api/user-post-events?view=graph-rest-1.0&tabs=http
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

	transformedEvent := support.ToEventSimplified(event)

	response, err := service.Client().UsersById(user).CalendarsById(destination).Events().Post(ctx, transformedEvent, nil)
	if err != nil {
		return errors.Wrap(err,
			fmt.Sprintf(
				"failure to event creation failure during RestoreExchangeEvent: %s",
				support.ConnectorStackErrorTrace(err)),
		)
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
		return err
	}
	// Sets fields from original message from storage
	clone := support.ToMessage(originalMessage)
	valueID := MailRestorePropertyTag
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
	sendPropertyTag := MailSendDateTimeOverrideProperty
	sv2.SetId(&sendPropertyTag)
	sv2.SetValue(&sendPropertyValue)

	sv3 := models.NewSingleValueLegacyExtendedProperty()
	recvPropertyValue := common.FormatLegacyTime(*clone.GetReceivedDateTime())
	recvPropertyTag := MailReceiveDateTimeOverriveProperty
	sv3.SetId(&recvPropertyTag)
	sv3.SetValue(&recvPropertyValue)

	svlep := []models.SingleValueLegacyExtendedPropertyable{sv1, sv2, sv3}
	clone.SetSingleValueExtendedProperties(svlep)

	// Switch workflow based on collision policy
	switch cp {
	default:
		logger.Ctx(ctx).DPanicw("restoreMailMessage received unrecognized restore policy; defaulting to copy",
			"policy", cp)
		fallthrough
	case control.Copy:
		return SendMailToBackStore(ctx, service, user, destination, clone)
	}
}

// SendMailToBackStore function for transporting in-memory messageable item to M365 backstore
// @param user string represents M365 ID of user within the tenant
// @param destination represents M365 ID of a folder within the users's space
// @param message is a models.Messageable interface from "github.com/microsoftgraph/msgraph-sdk-go/models"
func SendMailToBackStore(
	ctx context.Context,
	service graph.Service,
	user, destination string,
	message models.Messageable,
) error {
	var (
		attached []models.Attachmentable
		errs     error
	)

	if *message.GetHasAttachments() {
		attached = message.GetAttachments()
		message.SetAttachments([]models.Attachmentable{})
	}

	sentMessage, err := service.Client().UsersById(user).MailFoldersById(destination).Messages().Post(ctx, message, nil)
	if err != nil {
		return errors.Wrap(err,
			user+": failure sendMailAPI:  "+support.ConnectorStackErrorTrace(err),
		)
	}

	if sentMessage == nil {
		return errors.New("message not Sent: blocked by server")
	}

	if len(attached) > 0 {
		id := *sentMessage.GetId()
		for _, attachment := range attached {
			_, err = service.Client().
				UsersById(user).
				MailFoldersById(destination).
				MessagesById(id).
				Attachments().
				Post(ctx, attachment, nil)
			if err != nil {
				errs = support.WrapAndAppend(id,
					err,
					errs,
				)
			}
		}

		return errs
	}

	return nil
}

// RestoreExchangeDataCollections restores M365 objects in data.Collection to MSFT
// store through GraphAPI.
// @param dest:  container destination to M365
func RestoreExchangeDataCollections(
	ctx context.Context,
	gs graph.Service,
	dest control.RestoreDestination,
	dcs []data.Collection,
) (*support.ConnectorOperationStatus, error) {
	var (
		pathCounter         = map[string]bool{}
		rootFolder          string
		attempts, successes int
		errs                error
		// TODO policy to be updated from external source after completion of refactoring
		policy = control.Copy
	)

	errUpdater := func(id string, err error) {
		errs = support.WrapAndAppend(id, err, errs)
	}

	for _, dc := range dcs {
		a, s, root, canceled := restoreCollection(ctx, gs, dc, rootFolder, pathCounter, dest, policy, errUpdater)
		attempts += a
		successes += s
		rootFolder = root

		if canceled {
			break
		}
	}

	status := support.CreateStatus(ctx, support.Restore, attempts, successes, len(pathCounter), errs)

	return status, errs
}

// restoreCollection handles restoration of an individual collection.
func restoreCollection(
	ctx context.Context,
	gs graph.Service,
	dc data.Collection,
	rootFolder string,
	pathCounter map[string]bool,
	dest control.RestoreDestination,
	policy control.CollisionPolicy,
	errUpdater func(string, error),
) (int, int, string, bool) {
	defer trace.StartRegion(ctx, "gc:exchange:restoreCollection").End()
	trace.Log(ctx, "gc:exchange:restoreCollection", dc.FullPath().String())

	var (
		attempts, successes int
		folderID            string
		err                 error
		items               = dc.Items()
		directory           = dc.FullPath()
		service             = directory.Service()
		category            = directory.Category()
		user                = directory.ResourceOwner()
		directoryCheckFunc  = generateRestoreContainerFunc(gs, user, category, dest.ContainerName)
	)

	folderID, root, err := directoryCheckFunc(ctx, err, directory.String(), rootFolder, pathCounter)
	if err != nil { // assuming FailFast
		errUpdater(directory.String(), err)
		return 0, 0, rootFolder, false
	}

	for {
		select {
		case <-ctx.Done():
			errUpdater("context cancelled", ctx.Err())
			return attempts, successes, root, true

		case itemData, ok := <-items:
			if !ok {
				return attempts, successes, root, false
			}
			attempts++

			trace.Log(ctx, "gc:exchange:restoreCollection:item", itemData.UUID())

			buf := &bytes.Buffer{}

			_, err := buf.ReadFrom(itemData.ToReader())
			if err != nil {
				errUpdater(itemData.UUID()+": byteReadError during RestoreDataCollection", err)
				continue
			}

			err = RestoreExchangeObject(ctx, buf.Bytes(), category, policy, gs, folderID, user)
			if err != nil {
				//  More information to be here
				errUpdater(
					itemData.UUID()+": failed to upload RestoreExchangeObject: "+service.String()+"-"+category.String(),
					err)

				continue
			}
			successes++
		}
	}
}

// generateRestoreContainerFunc utility function that holds logic for creating
// Root Directory or necessary functions based on path.CategoryType
func generateRestoreContainerFunc(
	gs graph.Service,
	user string,
	category path.CategoryType,
	destination string,
) func(context.Context, error, string, string, map[string]bool) (string, string, error) {
	return func(
		ctx context.Context,
		errs error,
		dirName string,
		rootFolderID string,
		pathCounter map[string]bool,
	) (string, string, error) {
		var (
			folderID string
			err      error
		)

		if rootFolderID != "" && category == path.ContactsCategory {
			return rootFolderID, rootFolderID, errs
		}

		if !pathCounter[dirName] {
			pathCounter[dirName] = true

			folderID, err = GetRestoreContainer(ctx, gs, user, category, destination)
			if err != nil {
				return "", "", support.WrapAndAppend(user+" failure during preprocessing ", err, errs)
			}

			if rootFolderID == "" {
				rootFolderID = folderID
			}
		}

		return folderID, rootFolderID, nil
	}
}
