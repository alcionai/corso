package exchange

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"runtime/trace"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// RestoreExchangeObject directs restore pipeline towards restore function
// based on the path.CategoryType. All input params are necessary to perform
// the type-specific restore function.
func RestoreExchangeObject(
	ctx context.Context,
	bits []byte,
	category path.CategoryType,
	policy control.CollisionPolicy,
	service graph.Servicer,
	destination, user string,
) (*details.ExchangeInfo, error) {
	if policy != control.Copy {
		return nil, fmt.Errorf("restore policy: %s not supported for RestoreExchangeObject", policy)
	}

	switch category {
	case path.EmailCategory:
		return RestoreMailMessage(ctx, bits, service, control.Copy, destination, user)
	case path.ContactsCategory:
		return RestoreExchangeContact(ctx, bits, service, control.Copy, destination, user)
	case path.EventsCategory:
		return RestoreExchangeEvent(ctx, bits, service, control.Copy, destination, user)
	default:
		return nil, fmt.Errorf("type: %s not supported for RestoreExchangeObject", category)
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
	service graph.Servicer,
	cp control.CollisionPolicy,
	destination, user string,
) (*details.ExchangeInfo, error) {
	contact, err := support.CreateContactFromBytes(bits)
	if err != nil {
		return nil, errors.Wrap(err, "creating contact from bytes: RestoreExchangeContact")
	}

	response, err := service.Client().UsersById(user).ContactFoldersById(destination).Contacts().Post(ctx, contact, nil)
	if err != nil {
		name := *contact.GetGivenName()

		return nil, errors.Wrap(
			err,
			"uploading Contact during RestoreExchangeContact: "+name+" "+
				support.ConnectorStackErrorTrace(err),
		)
	}

	if response == nil {
		return nil, errors.New("msgraph contact post fail: REST response not received")
	}

	return ContactInfo(contact, int64(len(bits))), nil
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
	service graph.Servicer,
	cp control.CollisionPolicy,
	destination, user string,
) (*details.ExchangeInfo, error) {
	event, err := support.CreateEventFromBytes(bits)
	if err != nil {
		return nil, errors.Wrap(err, "creating event from bytes: RestoreExchangeEvent")
	}

	transformedEvent := support.ToEventSimplified(event)

	var (
		attached []models.Attachmentable
		errs     error
	)

	if *event.GetHasAttachments() {
		attached = event.GetAttachments()

		transformedEvent.SetAttachments([]models.Attachmentable{})
	}

	response, err := service.Client().UsersById(user).CalendarsById(destination).Events().Post(ctx, transformedEvent, nil)
	if err != nil {
		return nil, errors.Wrap(err,
			fmt.Sprintf(
				"uploading event during RestoreExchangeEvent: %s",
				support.ConnectorStackErrorTrace(err)),
		)
	}

	if response == nil {
		return nil, errors.New("msgraph event post fail: REST response not received")
	}

	uploader := &eventAttachmentUploader{
		calendarID: destination,
		userID:     user,
		service:    service,
		itemID:     *response.GetId(),
	}

	for _, attach := range attached {
		if err := uploadAttachment(ctx, uploader, attach); err != nil {
			errs = support.WrapAndAppend(
				fmt.Sprintf(
					"uploading attachment for message %s: %s",
					*transformedEvent.GetId(), support.ConnectorStackErrorTrace(err),
				),
				err,
				errs,
			)

			break
		}
	}

	return EventInfo(event, int64(len(bits))), errs
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
	service graph.Servicer,
	cp control.CollisionPolicy,
	destination, user string,
) (*details.ExchangeInfo, error) {
	// Creates messageable object from original bytes
	originalMessage, err := support.CreateMessageFromBytes(bits)
	if err != nil {
		return nil, errors.Wrap(err, "creating email from bytes: RestoreMailMessage")
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
		err := SendMailToBackStore(ctx, service, user, destination, clone)
		if err != nil {
			return nil, err
		}
	}

	return MessageInfo(clone, int64(len(bits))), nil
}

// attachmentBytes is a helper to retrieve the attachment content from a models.Attachmentable
// TODO: Revisit how we retrieve/persist attachment content during backup so this is not needed
func attachmentBytes(attachment models.Attachmentable) []byte {
	return reflect.Indirect(reflect.ValueOf(attachment)).FieldByName("contentBytes").Bytes()
}

// SendMailToBackStore function for transporting in-memory messageable item to M365 backstore
// @param user string represents M365 ID of user within the tenant
// @param destination represents M365 ID of a folder within the users's space
// @param message is a models.Messageable interface from "github.com/microsoftgraph/msgraph-sdk-go/models"
func SendMailToBackStore(
	ctx context.Context,
	service graph.Servicer,
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
			user+": failure sendMailAPI: Dest: "+destination+" Details: "+support.ConnectorStackErrorTrace(err),
		)
	}

	if sentMessage == nil {
		return errors.New("message not Sent: blocked by server")
	}

	id := *sentMessage.GetId()

	uploader := &mailAttachmentUploader{
		userID:   user,
		folderID: destination,
		itemID:   id,
		service:  service,
	}

	for _, attachment := range attached {
		if err := uploadAttachment(ctx, uploader, attachment); err != nil {
			errs = support.WrapAndAppend(
				fmt.Sprintf("uploading attachment for message %s: %s",
					id, support.ConnectorStackErrorTrace(err)),
				err,
				errs,
			)

			break
		}
	}

	return errs
}

// RestoreExchangeDataCollections restores M365 objects in data.Collection to MSFT
// store through GraphAPI.
// @param dest:  container destination to M365
func RestoreExchangeDataCollections(
	ctx context.Context,
	gs graph.Servicer,
	dest control.RestoreDestination,
	dcs []data.Collection,
	deets *details.Builder,
) (*support.ConnectorOperationStatus, error) {
	var (
		// map of caches... but not yet...
		directoryCaches = make(map[string]map[path.CategoryType]graph.ContainerResolver)
		metrics         support.CollectionMetrics
		errs            error
		// TODO policy to be updated from external source after completion of refactoring
		policy = control.Copy
	)

	errUpdater := func(id string, err error) {
		errs = support.WrapAndAppend(id, err, errs)
	}

	for _, dc := range dcs {
		userID := dc.FullPath().ResourceOwner()

		userCaches := directoryCaches[userID]
		if userCaches == nil {
			directoryCaches[userID] = make(map[path.CategoryType]graph.ContainerResolver)
			userCaches = directoryCaches[userID]
		}

		containerID, err := GetContainerIDFromCache(
			ctx,
			gs,
			dc.FullPath(),
			dest.ContainerName,
			userCaches)
		if err != nil {
			errs = support.WrapAndAppend(dc.FullPath().ShortRef(), err, errs)
			continue
		}

		temp, canceled := restoreCollection(ctx, gs, dc, containerID, policy, deets, errUpdater)

		metrics.Combine(temp)

		if canceled {
			break
		}
	}

	status := support.CreateStatus(ctx,
		support.Restore,
		len(dcs),
		metrics,
		errs,
		dest.ContainerName)

	return status, errs
}

// restoreCollection handles restoration of an individual collection.
func restoreCollection(
	ctx context.Context,
	gs graph.Servicer,
	dc data.Collection,
	folderID string,
	policy control.CollisionPolicy,
	deets *details.Builder,
	errUpdater func(string, error),
) (support.CollectionMetrics, bool) {
	ctx, end := D.Span(ctx, "gc:exchange:restoreCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics   support.CollectionMetrics
		items     = dc.Items()
		directory = dc.FullPath()
		service   = directory.Service()
		category  = directory.Category()
		user      = directory.ResourceOwner()
	)

	colProgress, closer := observe.CollectionProgress(user, category.String(), directory.Folder())
	defer closer()
	defer close(colProgress)

	for {
		select {
		case <-ctx.Done():
			errUpdater("context cancelled", ctx.Err())
			return metrics, true

		case itemData, ok := <-items:
			if !ok {
				return metrics, false
			}
			metrics.Objects++

			trace.Log(ctx, "gc:exchange:restoreCollection:item", itemData.UUID())

			buf := &bytes.Buffer{}

			_, err := buf.ReadFrom(itemData.ToReader())
			if err != nil {
				errUpdater(itemData.UUID()+": byteReadError during RestoreDataCollection", err)
				continue
			}

			byteArray := buf.Bytes()

			info, err := RestoreExchangeObject(ctx, byteArray, category, policy, gs, folderID, user)
			if err != nil {
				//  More information to be here
				errUpdater(
					itemData.UUID()+": failed to upload RestoreExchangeObject: "+service.String()+"-"+category.String(),
					err)

				continue
			}

			metrics.TotalBytes += int64(len(byteArray))
			metrics.Successes++

			itemPath, err := dc.FullPath().Append(itemData.UUID(), true)
			if err != nil {
				logger.Ctx(ctx).DPanicw("transforming item to full path", "error", err)
				continue
			}

			deets.Add(
				itemPath.String(),
				itemPath.ShortRef(),
				"",
				true,
				details.ItemInfo{
					Exchange: info,
				})

			colProgress <- struct{}{}
		}
	}
}

// generateRestoreContainerFunc utility function that holds logic for creating
// Root Directory or necessary functions based on path.CategoryType
// Assumption: collisionPolicy == COPY
func GetContainerIDFromCache(
	ctx context.Context,
	gs graph.Servicer,
	directory path.Path,
	destination string,
	caches map[path.CategoryType]graph.ContainerResolver,
) (string, error) {
	var (
		newCache       = false
		user           = directory.ResourceOwner()
		category       = directory.Category()
		directoryCache = caches[category]
		newPathFolders = append([]string{destination}, directory.Folders()...)
	)

	switch category {
	case path.EmailCategory:
		if directoryCache == nil {
			mfc := &mailFolderCache{
				userID: user,
				gs:     gs,
			}

			caches[category] = mfc
			newCache = true
			directoryCache = mfc
		}

		return establishMailRestoreLocation(
			ctx,
			newPathFolders,
			directoryCache,
			user,
			gs,
			newCache)
	case path.ContactsCategory:
		if directoryCache == nil {
			cfc := &contactFolderCache{
				userID: user,
				gs:     gs,
			}
			caches[category] = cfc
			newCache = true
			directoryCache = cfc
		}

		return establishContactsRestoreLocation(
			ctx,
			newPathFolders,
			directoryCache,
			user,
			gs,
			newCache)
	case path.EventsCategory:
		if directoryCache == nil {
			ecc := &eventCalendarCache{
				userID: user,
				gs:     gs,
			}
			caches[category] = ecc
			newCache = true
			directoryCache = ecc
		}

		return establishEventsRestoreLocation(
			ctx,
			newPathFolders,
			directoryCache,
			user,
			gs,
			newCache,
		)
	default:
		return "", fmt.Errorf("category: %s not support for exchange cache", category)
	}
}

// establishMailRestoreLocation creates Mail folders in sequence
// [root leaf1 leaf2] in a similar to a linked list.
// @param folders is the desired path from the root to the container
// that the items will be restored into
// @param isNewCache identifies if the cache is created and not populated
func establishMailRestoreLocation(
	ctx context.Context,
	folders []string,
	mfc graph.ContainerResolver,
	user string,
	service graph.Servicer,
	isNewCache bool,
) (string, error) {
	// Process starts with the root folder in order to recreate
	// the top-level folder with the same tactic
	folderID := rootFolderAlias
	pb := path.Builder{}

	for _, folder := range folders {
		pb = *pb.Append(folder)

		cached, ok := mfc.PathInCache(pb.String())
		if ok {
			folderID = cached
			continue
		}

		temp, err := CreateMailFolderWithParent(ctx,
			service, user, folder, folderID)
		if err != nil {
			// Should only error if cache malfunctions or incorrect parameters
			return "", errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		folderID = *temp.GetId()

		// Only populate the cache if we actually had to create it. Since we set
		// newCache to false in this we'll only try to populate it once per function
		// call even if we make a new cache.
		if isNewCache {
			if err := mfc.Populate(ctx, rootFolderAlias); err != nil {
				return "", errors.Wrap(err, "populating folder cache")
			}

			isNewCache = false
		}

		// NOOP if the folder is already in the cache.
		if err = mfc.AddToCache(ctx, temp); err != nil {
			return "", errors.Wrap(err, "adding folder to cache")
		}
	}

	return folderID, nil
}

// establishContactsRestoreLocation creates Contact Folders in sequence
// and updates the container resolver appropriately. Contact Folders are
// displayed in a flat representation. Therefore, only the root can be populated and all content
// must be restored into the root location.
// @param folders is the list of intended folders from root to leaf (e.g. [root ...])
// @param isNewCache bool representation of whether Populate function needs to be run
func establishContactsRestoreLocation(
	ctx context.Context,
	folders []string,
	cfc graph.ContainerResolver,
	user string,
	gs graph.Servicer,
	isNewCache bool,
) (string, error) {
	cached, ok := cfc.PathInCache(folders[0])
	if ok {
		return cached, nil
	}

	temp, err := CreateContactFolder(ctx, gs, user, folders[0])
	if err != nil {
		return "", errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	folderID := *temp.GetId()

	if isNewCache {
		if err := cfc.Populate(ctx, folderID, folders[0]); err != nil {
			return "", errors.Wrap(err, "populating contact cache")
		}

		if err = cfc.AddToCache(ctx, temp); err != nil {
			return "", errors.Wrap(err, "adding contact folder to cache")
		}
	}

	return folderID, nil
}

func establishEventsRestoreLocation(
	ctx context.Context,
	folders []string,
	ecc graph.ContainerResolver, // eventCalendarCache
	user string,
	gs graph.Servicer,
	isNewCache bool,
) (string, error) {
	cached, ok := ecc.PathInCache(folders[0])
	if ok {
		return cached, nil
	}

	temp, err := CreateCalendar(ctx, gs, user, folders[0])
	if err != nil {
		return "", errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	folderID := *temp.GetId()

	if isNewCache {
		if err = ecc.Populate(ctx, folderID, folders[0]); err != nil {
			return "", errors.Wrap(err, "populating event cache")
		}

		transform := CreateCalendarDisplayable(temp)
		if err = ecc.AddToCache(ctx, transform); err != nil {
			return "", errors.Wrap(err, "adding new calendar to cache")
		}
	}

	return folderID, nil
}
