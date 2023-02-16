package exchange

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"runtime/trace"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
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
	errs *fault.Errors,
) (*details.ExchangeInfo, error) {
	if policy != control.Copy {
		return nil, clues.Wrap(clues.New(policy.String()), "policy not supported for Exchange restore").WithClues(ctx)
	}

	switch category {
	case path.EmailCategory:
		return RestoreMailMessage(ctx, bits, service, control.Copy, destination, user, errs)
	case path.ContactsCategory:
		return RestoreExchangeContact(ctx, bits, service, control.Copy, destination, user)
	case path.EventsCategory:
		return RestoreExchangeEvent(ctx, bits, service, control.Copy, destination, user, errs)
	default:
		return nil, clues.Wrap(clues.New(category.String()), "not supported for Exchange restore")
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
		return nil, clues.Wrap(err, "creating contact from bytes").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(contact.GetId()))

	response, err := service.Client().UsersById(user).ContactFoldersById(destination).Contacts().Post(ctx, contact, nil)
	if err != nil {
		return nil, clues.Wrap(err, "uploading Contact").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	if response == nil {
		return nil, clues.New("nil response from post").WithClues(ctx)
	}

	info := api.ContactInfo(contact)
	info.Size = int64(len(bits))

	return info, nil
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
	errs *fault.Errors,
) (*details.ExchangeInfo, error) {
	event, err := support.CreateEventFromBytes(bits)
	if err != nil {
		return nil, clues.Wrap(err, "creating event from bytes").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(event.GetId()))

	var (
		transformedEvent = support.ToEventSimplified(event)
		attached         []models.Attachmentable
	)

	if *event.GetHasAttachments() {
		attached = event.GetAttachments()

		transformedEvent.SetAttachments([]models.Attachmentable{})
	}

	response, err := service.Client().UsersById(user).CalendarsById(destination).Events().Post(ctx, transformedEvent, nil)
	if err != nil {
		return nil, clues.Wrap(err, "uploading event").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	if response == nil {
		return nil, clues.New("nil response from post").WithClues(ctx)
	}

	uploader := &eventAttachmentUploader{
		calendarID: destination,
		userID:     user,
		service:    service,
		itemID:     *response.GetId(),
	}

	for _, attach := range attached {
		if errs.Err() != nil {
			break
		}

		if err := uploadAttachment(ctx, uploader, attach); err != nil {
			errs.Add(err)
		}
	}

	info := api.EventInfo(event)
	info.Size = int64(len(bits))

	return info, errs.Err()
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
	errs *fault.Errors,
) (*details.ExchangeInfo, error) {
	// Creates messageable object from original bytes
	originalMessage, err := support.CreateMessageFromBytes(bits)
	if err != nil {
		return nil, clues.Wrap(err, "creating mail from bytes").WithClues(ctx)
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(originalMessage.GetId()))

	var (
		clone       = support.ToMessage(originalMessage)
		valueID     = MailRestorePropertyTag
		enableValue = RestoreCanonicalEnableValue
	)

	// Set Extended Properties:
	// 1st: No transmission
	// 2nd: Send Date
	// 3rd: Recv Date
	svlep := make([]models.SingleValueLegacyExtendedPropertyable, 0)
	sv1 := models.NewSingleValueLegacyExtendedProperty()
	sv1.SetId(&valueID)
	sv1.SetValue(&enableValue)
	svlep = append(svlep, sv1)

	if clone.GetSentDateTime() != nil {
		sv2 := models.NewSingleValueLegacyExtendedProperty()
		sendPropertyValue := common.FormatLegacyTime(*clone.GetSentDateTime())
		sendPropertyTag := MailSendDateTimeOverrideProperty
		sv2.SetId(&sendPropertyTag)
		sv2.SetValue(&sendPropertyValue)

		svlep = append(svlep, sv2)
	}

	if clone.GetReceivedDateTime() != nil {
		sv3 := models.NewSingleValueLegacyExtendedProperty()
		recvPropertyValue := common.FormatLegacyTime(*clone.GetReceivedDateTime())
		recvPropertyTag := MailReceiveDateTimeOverriveProperty
		sv3.SetId(&recvPropertyTag)
		sv3.SetValue(&recvPropertyValue)

		svlep = append(svlep, sv3)
	}

	clone.SetSingleValueExtendedProperties(svlep)

	if err := SendMailToBackStore(ctx, service, user, destination, clone, errs); err != nil {
		return nil, err
	}

	info := api.MailInfo(clone)
	info.Size = int64(len(bits))

	return info, nil
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
	errs *fault.Errors,
) error {
	attached := message.GetAttachments()

	// Item.Attachments --> HasAttachments doesn't always have a value populated when deserialized
	message.SetAttachments([]models.Attachmentable{})

	response, err := service.Client().UsersById(user).MailFoldersById(destination).Messages().Post(ctx, message, nil)
	if err != nil {
		return clues.Wrap(err, "restoring mail").WithClues(ctx).WithAll(graph.ErrData(err)...)
	}

	if response == nil {
		return clues.New("nil response from post").WithClues(ctx)
	}

	id := ptr.Val(response.GetId())

	uploader := &mailAttachmentUploader{
		userID:   user,
		folderID: destination,
		itemID:   id,
		service:  service,
	}

	for _, attachment := range attached {
		if errs.Err() != nil {
			break
		}

		if err := uploadAttachment(ctx, uploader, attachment); err != nil {
			if ptr.Val(attachment.GetOdataType()) == "#microsoft.graph.itemAttachment" {
				name := ptr.Val(attachment.GetName())

				logger.Ctx(ctx).
					With("err", err, "attachment_name", name).
					Infow("mail upload failed", clues.InErr(err).Slice()...)

				continue
			}

			errs.Add(errors.Wrap(err, "uploading mail attachment"))

			break
		}
	}

	return errs.Err()
}

// RestoreExchangeDataCollections restores M365 objects in data.RestoreCollection to MSFT
// store through GraphAPI.
// @param dest:  container destination to M365
func RestoreExchangeDataCollections(
	ctx context.Context,
	creds account.M365Config,
	gs graph.Servicer,
	dest control.RestoreDestination,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Errors,
) (*support.ConnectorOperationStatus, error) {
	var (
		directoryCaches = make(map[string]map[path.CategoryType]graph.ContainerResolver)
		metrics         support.CollectionMetrics
		userID          string
		// TODO policy to be updated from external source after completion of refactoring
		policy = control.Copy
	)

	if len(dcs) > 0 {
		userID = dcs[0].FullPath().ResourceOwner()
		ctx = clues.Add(ctx, "resource_owner", userID) // TODO: pii
	}

	for _, dc := range dcs {
		if errs.Err() != nil {
			return nil, errs.Err()
		}

		userCaches := directoryCaches[userID]
		if userCaches == nil {
			directoryCaches[userID] = make(map[path.CategoryType]graph.ContainerResolver)
			userCaches = directoryCaches[userID]
		}

		containerID, err := CreateContainerDestination(
			ctx,
			creds,
			dc.FullPath(),
			dest.ContainerName,
			userCaches,
			errs)
		if err != nil {
			errs.Add(clues.Wrap(err, "creating destination").WithClues(ctx))
			continue
		}

		temp, canceled := restoreCollection(ctx, gs, dc, containerID, policy, deets, errs)

		metrics.Combine(temp)

		if canceled {
			break
		}
	}

	status := support.CreateStatus(
		ctx,
		support.Restore,
		len(dcs),
		metrics,
		errs.Err(),
		dest.ContainerName)

	return status, errs.Err()
}

// restoreCollection handles restoration of an individual collection.
func restoreCollection(
	ctx context.Context,
	gs graph.Servicer,
	dc data.RestoreCollection,
	folderID string,
	policy control.CollisionPolicy,
	deets *details.Builder,
	errs *fault.Errors,
) (support.CollectionMetrics, bool) {
	ctx, end := D.Span(ctx, "gc:exchange:restoreCollection", D.Label("path", dc.FullPath()))
	defer end()

	var (
		metrics   support.CollectionMetrics
		items     = dc.Items(ctx, errs)
		directory = dc.FullPath()
		service   = directory.Service()
		category  = directory.Category()
		user      = directory.ResourceOwner()
	)

	ctx = clues.AddAll(
		ctx,
		"full_path", directory,
		"service", service,
		"category", category)

	colProgress, closer := observe.CollectionProgress(
		ctx,
		category.String(),
		observe.PII(user),
		observe.PII(directory.Folder(false)))
	defer closer()
	defer close(colProgress)

	for {
		select {
		case <-ctx.Done():
			errs.Add(clues.Wrap(ctx.Err(), "context cancelled").WithClues(ctx))
			return metrics, true

		case itemData, ok := <-items:
			if !ok || errs.Err() != nil {
				return metrics, false
			}

			ictx := clues.Add(ctx, "item_id", itemData.UUID())
			trace.Log(ictx, "gc:exchange:restoreCollection:item", itemData.UUID())
			metrics.Objects++

			buf := &bytes.Buffer{}

			_, err := buf.ReadFrom(itemData.ToReader())
			if err != nil {
				errs.Add(clues.Wrap(err, "reading item bytes").WithClues(ictx))
				continue
			}

			byteArray := buf.Bytes()

			info, err := RestoreExchangeObject(
				ictx,
				byteArray,
				category,
				policy,
				gs,
				folderID,
				user,
				errs)
			if err != nil {
				errs.Add(err)
				continue
			}

			metrics.TotalBytes += int64(len(byteArray))
			metrics.Successes++

			itemPath, err := dc.FullPath().Append(itemData.UUID(), true)
			if err != nil {
				errs.Add(clues.Wrap(err, "building full path with item").WithClues(ctx))
				continue
			}

			var locationRef string
			if category == path.ContactsCategory {
				locationRef = itemPath.Folder(false)
			}

			deets.Add(
				itemPath.String(),
				itemPath.ShortRef(),
				"",
				locationRef,
				true,
				details.ItemInfo{
					Exchange: info,
				})

			colProgress <- struct{}{}
		}
	}
}

// CreateContainerDestination builds the destination into the container
// at the provided path.  As a precondition, the destination cannot
// already exist.  If it does then an error is returned.  The provided
// containerResolver is updated with the new destination.
// @ returns the container ID of the new destination container.
func CreateContainerDestination(
	ctx context.Context,
	creds account.M365Config,
	directory path.Path,
	destination string,
	caches map[path.CategoryType]graph.ContainerResolver,
	errs *fault.Errors,
) (string, error) {
	var (
		newCache       = false
		user           = directory.ResourceOwner()
		category       = directory.Category()
		directoryCache = caches[category]
	)

	// TODO(rkeepers): pass the api client into this func, rather than generating one.
	ac, err := api.NewClient(creds)
	if err != nil {
		return "", clues.Stack(err).WithClues(ctx)
	}

	switch category {
	case path.EmailCategory:
		folders := append([]string{destination}, directory.Folders()...)

		if directoryCache == nil {
			acm := ac.Mail()
			mfc := &mailFolderCache{
				userID: user,
				enumer: acm,
				getter: acm,
			}

			caches[category] = mfc
			newCache = true
			directoryCache = mfc
		}

		return establishMailRestoreLocation(
			ctx,
			ac,
			folders,
			directoryCache,
			user,
			newCache,
			errs)

	case path.ContactsCategory:
		folders := append([]string{destination}, directory.Folders()...)

		if directoryCache == nil {
			acc := ac.Contacts()
			cfc := &contactFolderCache{
				userID: user,
				enumer: acc,
				getter: acc,
			}
			caches[category] = cfc
			newCache = true
			directoryCache = cfc
		}

		return establishContactsRestoreLocation(
			ctx,
			ac,
			folders,
			directoryCache,
			user,
			newCache,
			errs)

	case path.EventsCategory:
		dest := destination

		if directoryCache == nil {
			ace := ac.Events()
			ecc := &eventCalendarCache{
				userID: user,
				getter: ace,
				enumer: ace,
			}
			caches[category] = ecc
			newCache = true
			directoryCache = ecc
		} else if did := directoryCache.DestinationNameToID(dest); len(did) > 0 {
			// calendars are cached by ID in the resolver, not name, so once we have
			// created the destination calendar, we need to look up its id and use
			// that for resolver lookups instead of the display name.
			dest = did
		}

		folders := append([]string{dest}, directory.Folders()...)

		return establishEventsRestoreLocation(
			ctx,
			ac,
			folders,
			directoryCache,
			user,
			newCache,
			errs)

	default:
		return "", clues.Wrap(fmt.Errorf("%T", category), "not support for exchange cache").WithClues(ctx)
	}
}

// establishMailRestoreLocation creates Mail folders in sequence
// [root leaf1 leaf2] in a similar to a linked list.
// @param folders is the desired path from the root to the container
// that the items will be restored into
// @param isNewCache identifies if the cache is created and not populated
func establishMailRestoreLocation(
	ctx context.Context,
	ac api.Client,
	folders []string,
	mfc graph.ContainerResolver,
	user string,
	isNewCache bool,
	errs *fault.Errors,
) (string, error) {
	// Process starts with the root folder in order to recreate
	// the top-level folder with the same tactic
	folderID := rootFolderAlias
	pb := path.Builder{}

	ctx = clues.Add(ctx, "is_new_cache", isNewCache)

	for _, folder := range folders {
		pb = *pb.Append(folder)

		cached, ok := mfc.PathInCache(pb.String())
		if ok {
			folderID = cached
			continue
		}

		temp, err := ac.Mail().CreateMailFolderWithParent(ctx, user, folder, folderID)
		if err != nil {
			// Should only error if cache malfunctions or incorrect parameters
			return "", errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		folderID = *temp.GetId()

		// Only populate the cache if we actually had to create it. Since we set
		// newCache to false in this we'll only try to populate it once per function
		// call even if we make a new cache.
		if isNewCache {
			if err := mfc.Populate(ctx, errs, rootFolderAlias); err != nil {
				return "", errors.Wrap(err, "populating folder cache")
			}

			isNewCache = false
		}

		// NOOP if the folder is already in the cache.
		if err = mfc.AddToCache(ctx, temp, false); err != nil {
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
	ac api.Client,
	folders []string,
	cfc graph.ContainerResolver,
	user string,
	isNewCache bool,
	errs *fault.Errors,
) (string, error) {
	cached, ok := cfc.PathInCache(folders[0])
	if ok {
		return cached, nil
	}

	ctx = clues.Add(ctx, "is_new_cache", isNewCache)

	temp, err := ac.Contacts().CreateContactFolder(ctx, user, folders[0])
	if err != nil {
		return "", errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	folderID := *temp.GetId()

	if isNewCache {
		if err := cfc.Populate(ctx, errs, folderID, folders[0]); err != nil {
			return "", errors.Wrap(err, "populating contact cache")
		}

		if err = cfc.AddToCache(ctx, temp, false); err != nil {
			return "", errors.Wrap(err, "adding contact folder to cache")
		}
	}

	return folderID, nil
}

func establishEventsRestoreLocation(
	ctx context.Context,
	ac api.Client,
	folders []string,
	ecc graph.ContainerResolver, // eventCalendarCache
	user string,
	isNewCache bool,
	errs *fault.Errors,
) (string, error) {
	// Need to prefix with the "Other Calendars" folder so lookup happens properly.
	cached, ok := ecc.PathInCache(folders[0])
	if ok {
		return cached, nil
	}

	ctx = clues.Add(ctx, "is_new_cache", isNewCache)

	temp, err := ac.Events().CreateCalendar(ctx, user, folders[0])
	if err != nil {
		return "", errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	folderID := *temp.GetId()

	if isNewCache {
		if err = ecc.Populate(ctx, errs, folderID, folders[0]); err != nil {
			return "", errors.Wrap(err, "populating event cache")
		}

		displayable := api.CalendarDisplayable{Calendarable: temp}
		if err = ecc.AddToCache(ctx, displayable, true); err != nil {
			return "", errors.Wrap(err, "adding new calendar to cache")
		}
	}

	return folderID, nil
}
