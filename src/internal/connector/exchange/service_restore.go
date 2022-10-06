package exchange

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime/trace"

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
		directoryCaches = map[path.CategoryType]mailFolderCache{}
		rootFolder      string
		metrics         support.CollectionMetrics
		errs            error
		// TODO policy to be updated from external source after completion of refactoring
		policy = control.Copy
	)

	errUpdater := func(id string, err error) {
		errs = support.WrapAndAppend(id, err, errs)
	}

	for _, dc := range dcs {
		//PUT Caching into separate function... here
		// Inputs --> folderID, root, err := directoryCheckFunc(ctx, err, directory.String(), rootFolder, pathCounter)
		//GetContainerID()
		GetContainerIDFromCache(ctx, gs, dc.FullPath(), dest.ContainerName,
			directoryCaches, pathCounter)
		os.Exit(1)
		temp, root, canceled := restoreCollection(ctx, gs, dc, rootFolder, pathCounter, dest, policy, deets, errUpdater)

		metrics.Combine(temp)

		rootFolder = root

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
	rootFolder string,
	pathCounter map[string]bool,
	dest control.RestoreDestination,
	policy control.CollisionPolicy,
	deets *details.Details,
	errUpdater func(string, error),
) (support.CollectionMetrics, string, bool) {
	defer trace.StartRegion(ctx, "gc:exchange:restoreCollection").End()
	trace.Log(ctx, "gc:exchange:restoreCollection", dc.FullPath().String())

	var (
		metrics        support.CollectionMetrics
		folderID, root string
		err            error
		items          = dc.Items()
		directory      = dc.FullPath()
		service        = directory.Service()
		category       = directory.Category()
		user           = directory.ResourceOwner()
	)

	if err != nil { // assuming FailFast
		errUpdater(directory.String(), err)
		return metrics, rootFolder, false
	}

	for {
		select {
		case <-ctx.Done():
			errUpdater("context cancelled", ctx.Err())
			return metrics, root, true

		case itemData, ok := <-items:
			if !ok {
				return metrics, root, false
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
	caches map[path.CategoryType]mailFolderCache,
	pathCounter map[string]bool,
) (string, error) {
	var (
		folderID       string
		err            error
		user           = directory.ResourceOwner()
		category       = directory.Category()
		directoryCache = caches[category]
	)

	if directoryCache.userID == "" {
		// Creates Root Node && Enable Cache
		folderID, err = GetRestoreContainer(ctx, gs, user, category, destination)
		if err != nil {
			return "", err
		}

		mfc := mailFolderCache{
			userID: user,
			gs:     gs,
		}

		err = mfc.Populate(ctx, folderID)
		if err != nil {
			return "", err
		}

		caches[category] = mfc
	}

	// New Name would be what
	directoryCache = caches[category]
	newPath, err := path.Builder{}.Append(destination).Append(directory.Folders()...).
		ToDataLayerExchangePathForCategory(
			directory.Tenant(),
			user,
			category,
			false,
		)
	if err != nil {
		return "", err
	}

	// Check if path with the thing
	// Does the folder exist
	fmt.Println("PATH STRING: " + directory.String())
	fmt.Println("NEW PATH: " + newPath.String())
	fmt.Printf("What: %v\n", *directoryCache.cache[directoryCache.rootID].GetDisplayName())
	fmt.Printf("PotentialPath: %v\n", newPath.Folders())

	parentID := directoryCache.rootID

	folders := newPath.Folders()
	for i := 1; i < len(folders); i++ {
		temp, err := CreateMailFolderWithParent(ctx, gs, user, folders[i], parentID)
		if err != nil {
			// This means the folder already exists... We will work in this later
			return "", err
		}

		err = directoryCache.addMailFolder(temp)
		if err != nil {
			return "", err
		}

		parentID = *temp.GetId()

	}

	folderID = parentID
	_, err = directoryCache.IDToPath(ctx, folderID)
	return folderID, err
}

// Need to update cache to note that folder is in there.
/*
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
	}*/

func newMailRestorePathForCollection(collectionPath path.Path, prefix string) (path.Path, error) {

	listing := collectionPath.Folders()
	pb := collectionPath.ToBuilder()
	newBuilder, err := pb.UnescapeAndAppend(prefix)
	if err != nil {
		return nil, err
	}

	//newBuilder.Append(listing...)

	return newBuilder.Append(listing...).ToDataLayerExchangePathForCategory(
		collectionPath.Tenant(),
		collectionPath.ResourceOwner(),
		collectionPath.Category(),
		false,
	)
}
