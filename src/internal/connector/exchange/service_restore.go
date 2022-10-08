package exchange

import (
	"bytes"
	"context"
	"fmt"
	"runtime/trace"
	"strings"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
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
) (*details.ExchangeInfo, error) {
	if policy != control.Copy {
		return nil, fmt.Errorf("restore policy: %s not supported for RestoreExchangeObject", policy)
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
	service graph.Service,
	cp control.CollisionPolicy,
	destination, user string,
) (*details.ExchangeInfo, error) {
	contact, err := support.CreateContactFromBytes(bits)
	if err != nil {
		return nil, errors.Wrap(err, "failure to create contact from bytes: RestoreExchangeContact")
	}

	response, err := service.Client().UsersById(user).ContactFoldersById(destination).Contacts().Post(ctx, contact, nil)
	if err != nil {
		name := *contact.GetGivenName()

		return nil, errors.Wrap(
			err,
			"failure to create Contact during RestoreExchangeContact: "+name+" "+
				support.ConnectorStackErrorTrace(err),
		)
	}

	if response == nil {
		return nil, errors.New("msgraph contact post fail: REST response not received")
	}

	return ContactInfo(contact), nil
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
) (*details.ExchangeInfo, error) {
	event, err := support.CreateEventFromBytes(bits)
	if err != nil {
		return nil, err
	}

	transformedEvent := support.ToEventSimplified(event)

	response, err := service.Client().UsersById(user).CalendarsById(destination).Events().Post(ctx, transformedEvent, nil)
	if err != nil {
		return nil, errors.Wrap(err,
			fmt.Sprintf(
				"failure to event creation failure during RestoreExchangeEvent: %s",
				support.ConnectorStackErrorTrace(err)),
		)
	}

	if response == nil {
		return nil, errors.New("msgraph event post fail: REST response not received")
	}

	return EventInfo(event), nil
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
) (*details.ExchangeInfo, error) {
	// Creates messageable object from original bytes
	originalMessage, err := support.CreateMessageFromBytes(bits)
	if err != nil {
		return nil, err
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
		return MessageInfo(clone), SendMailToBackStore(ctx, service, user, destination, clone)
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
	deets *details.Details,
) (*support.ConnectorOperationStatus, error) {
	var (
		pathCounter = map[string]bool{}
		// map of caches... but not yet...
		directoryCaches = make(map[path.CategoryType]graph.ContainerResolver)
		metrics         support.CollectionMetrics
		errs            error
		// TODO policy to be updated from external source after completion of refactoring
		policy = control.Copy
	)

	errUpdater := func(id string, err error) {
		errs = support.WrapAndAppend(id, err, errs)
	}

	for _, dc := range dcs {
		containerID, err := GetContainerIDFromCache(
			ctx,
			gs,
			dc.FullPath(),
			dest.ContainerName,
			directoryCaches, pathCounter)
		if err != nil {
			errs = support.WrapAndAppend(dc.FullPath().ShortRef(), err, errs)
			continue
		}

		temp, canceled := restoreCollection(ctx, gs, dc, containerID, pathCounter, policy, deets, errUpdater)

		metrics.Combine(temp)

		if canceled {
			break
		}
	}

	status := support.CreateStatus(ctx,
		support.Restore,
		len(pathCounter),
		metrics,
		errs,
		dest.ContainerName)

	return status, errs
}

// restoreCollection handles restoration of an individual collection.
func restoreCollection(
	ctx context.Context,
	gs graph.Service,
	dc data.Collection,
	folderID string,
	pathCounter map[string]bool,
	policy control.CollisionPolicy,
	deets *details.Details,
	errUpdater func(string, error),
) (support.CollectionMetrics, bool) {
	defer trace.StartRegion(ctx, "gc:exchange:restoreCollection").End()
	trace.Log(ctx, "gc:exchange:restoreCollection", dc.FullPath().String())

	var (
		metrics   support.CollectionMetrics
		items     = dc.Items()
		directory = dc.FullPath()
		service   = directory.Service()
		category  = directory.Category()
		user      = directory.ResourceOwner()
	)

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
				details.ItemInfo{
					Exchange: info,
				})
		}
	}
}

// generateRestoreContainerFunc utility function that holds logic for creating
// Root Directory or necessary functions based on path.CategoryType
// Assumption: collisionPolicy == COPY
// Constraint: Only works on exchange.Mail
func GetContainerIDFromCache(
	ctx context.Context,
	gs graph.Service,
	directory path.Path,
	destination string,
	caches map[path.CategoryType]graph.ContainerResolver,
	pathCounter map[string]bool,
) (string, error) {
	var (
		// Start with the root folder so we can create our top-level folder with
		// the same tactic.
		newCache          = make(map[path.CategoryType]bool)
		user              = directory.ResourceOwner()
		category          = directory.Category()
		directoryCache    = caches[category]
		newPathFolders    = append([]string{destination}, directory.Folders()...)
		restoreFolderPath = strings.Join(newPathFolders, "/")
	)

	if _, ok := pathCounter[restoreFolderPath]; !ok {
		pathCounter[restoreFolderPath] = true
	}

	switch category {
	case path.EmailCategory:
		if directoryCache == nil {
			mfc := &mailFolderCache{
				userID: user,
				gs:     gs,
			}

			caches[category] = mfc
			newCache[category] = true
			directoryCache = mfc
		}

		isEnabled := newCache[category]
		newCache[category] = false

		return establishMailRestoreLocation(
			ctx,
			newPathFolders,
			directoryCache,
			user,
			gs,
			isEnabled)
	case path.ContactsCategory:
		if directoryCache == nil {
			cfc := &contactFolderCache{
				userID: user,
				gs:     gs,
			}
			caches[category] = cfc
			newCache[category] = true
			directoryCache = cfc
		}

		isEnabled := newCache[category]
		newCache[category] = false

		return establishContactsRestoreLocation(
			ctx,
			newPathFolders,
			directoryCache,
			user,
			gs,
			isEnabled)
	case path.EventsCategory:
		return "", nil
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
	service graph.Service,
	isNewCache bool,
) (string, error) {
	folderID := rootFolderAlias

	for i, folder := range folders {
		lookup := pathElementStringBuilder(i+1, folders)
		cached, ok := mfc.PathInCache(lookup)

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
			if err := mfc.Populate(ctx, folderID, folder); err != nil {
				return "", errors.Wrap(err, "populating folder cache")
			}

			isNewCache = false
		}

		// NOOP if the folder is already in the cache.
		if err = mfc.AddToCache(temp); err != nil {
			return "", errors.Wrap(err, "adding folder to cache")
		}
	}

	// Lazy cache Update Cache by executing at the furthest leaf
	_, err := mfc.IDToPath(ctx, folderID)

	return folderID, err
}

// establishContactsRestoreLocation creates Contact Folders in sequence
// and updates the container resolver appropriately. Contact Folders have
// a flat representation. Therefore, only the root can be populated and all content
// must be restored into the root location.
// @param folders is the list of intended folders from root to leaf (e.g. [root ...])
// @param isNewCache bool representation of whether Populate function needs to be run
func establishContactsRestoreLocation(
	ctx context.Context,
	folders []string,
	cfc graph.ContainerResolver,
	user string,
	gs graph.Service,
	isNewCache bool,
) (string, error) {
	cached, ok := cfc.PathInCache(folders[0])
	if ok {
		fmt.Println("Cache HIT!")
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

		if err = cfc.AddToCache(temp); err != nil {
			return "", errors.Wrap(err, "adding contact folder to cache")
		}
	}

	_, err = cfc.IDToPath(ctx, folderID)

	return folderID, err
}
