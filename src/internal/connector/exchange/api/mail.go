package api

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Mail() Mail {
	return Mail{c}
}

// Mail is an interface-compliant provider of the client.
type Mail struct {
	Client
}

// ---------------------------------------------------------------------------
// methods
// ---------------------------------------------------------------------------

// CreateMailFolder makes a mail folder iff a folder of the same name does not exist
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-mailfolders?view=graph-rest-1.0&tabs=http
func (c Mail) CreateMailFolder(
	ctx context.Context,
	user, folder string,
) (models.MailFolderable, error) {
	isHidden := false
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	requestBody.SetIsHidden(&isHidden)

	return c.stable.Client().UsersById(user).MailFolders().Post(ctx, requestBody, nil)
}

func (c Mail) CreateMailFolderWithParent(
	ctx context.Context,
	user, folder, parentID string,
) (models.MailFolderable, error) {
	service, err := c.service()
	if err != nil {
		return nil, err
	}

	isHidden := false
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	requestBody.SetIsHidden(&isHidden)

	return service.
		Client().
		UsersById(user).
		MailFoldersById(parentID).
		ChildFolders().
		Post(ctx, requestBody, nil)
}

// DeleteContainer removes a mail folder with the corresponding M365 ID  from the user's M365 Exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/mailfolder-delete?view=graph-rest-1.0&tabs=http
func (c Mail) DeleteContainer(
	ctx context.Context,
	user, folderID string,
) error {
	return c.stable.Client().UsersById(user).MailFoldersById(folderID).Delete(ctx, nil)
}

func (c Mail) GetContainerByID(
	ctx context.Context,
	userID, dirID string,
) (graph.Container, error) {
	service, err := c.service()
	if err != nil {
		return nil, err
	}

	ofmf, err := optionsForMailFoldersItem([]string{"displayName", "parentFolderId"})
	if err != nil {
		return nil, errors.Wrap(err, "options for mail folder")
	}

	return service.Client().UsersById(userID).MailFoldersById(dirID).Get(ctx, ofmf)
}

// GetItem retrieves a Messageable item.  If the item contains an attachment, that
// attachment is also downloaded.
func (c Mail) GetItem(
	ctx context.Context,
	user, itemID string,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	mail, err := c.stable.Client().UsersById(user).MessagesById(itemID).Get(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	var errs *multierror.Error

	if *mail.GetHasAttachments() || HasAttachments(mail.GetBody()) {
		for count := 0; count < numberOfRetries; count++ {
			attached, err := c.largeItem.
				Client().
				UsersById(user).
				MessagesById(itemID).
				Attachments().
				Get(ctx, nil)
			if err == nil {
				mail.SetAttachments(attached.GetValue())
				break
			}

			logger.Ctx(ctx).Debugw("retrying mail attachment download", "err", err)
			errs = multierror.Append(errs, err)
		}

		if err != nil {
			logger.Ctx(ctx).Errorw("mail attachment download exceeded maximum retries", "err", errs)
			return nil, nil, support.WrapAndAppend(itemID, errors.Wrap(err, "downloading mail attachment"), nil)
		}
	}

	return mail, MailInfo(mail), nil
}

// EnumerateContainers iterates through all of the users current
// mail folders, converting each to a graph.CacheFolder, and calling
// fn(cf) on each one.  If fn(cf) errors, the error is aggregated
// into a multierror that gets returned to the caller.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Mail) EnumerateContainers(
	ctx context.Context,
	userID, baseDirID string,
	fn func(graph.CacheFolder) error,
) error {
	service, err := c.service()
	if err != nil {
		return err
	}

	var (
		resp    users.ItemMailFoldersDeltaResponseable
		errs    *multierror.Error
		builder = service.Client().
			UsersById(userID).
			MailFolders().
			Delta()
	)

	for {
		for i := 1; i <= numberOfRetries; i++ {
			resp, err = builder.Get(ctx, nil)
			if err == nil {
				break
			}

			if !graph.IsErrTimeout(err) && !graph.IsSericeUnavailable(err) {
				break
			}

			if i < numberOfRetries {
				time.Sleep(time.Duration(3*(i+1)) * time.Second)
			}
		}

		if err != nil {
			return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, v := range resp.GetValue() {
			temp := graph.NewCacheFolder(v, nil)

			if err := fn(temp); err != nil {
				errs = multierror.Append(errs, errors.Wrap(err, "iterating mail folders delta"))
				continue
			}
		}

		link := resp.GetOdataNextLink()
		if link == nil {
			break
		}

		builder = users.NewItemMailFoldersDeltaRequestBuilder(*link, service.Adapter())
	}

	return errs.ErrorOrNil()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager = &mailPager{}

type mailPager struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersItemMessagesDeltaRequestBuilder
	options *users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func (p *mailPager) getPage(ctx context.Context) (api.DeltaPageLinker, error) {
	return p.builder.Get(ctx, p.options)
}

func (p *mailPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *mailPager) valuesIn(pl api.DeltaPageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Messageable](pl)
}

func (c Mail) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	user, directoryID, oldDelta string,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	var (
		errs       *multierror.Error
		deltaURL   string
		resetDelta bool
	)

	options, err := optionsForFolderMessagesDelta([]string{"isRead"})
	if err != nil {
		return nil, nil, DeltaUpdate{}, errors.Wrap(err, "getting query options")
	}

	if len(oldDelta) > 0 {
		builder := users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(oldDelta, service.Adapter())
		pgr := &mailPager{service, builder, options}

		added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
		// note: happy path, not the error condition
		if err == nil {
			return added, removed, DeltaUpdate{deltaURL, false}, errs.ErrorOrNil()
		}
		// only return on error if it is NOT a delta issue.
		// on bad deltas we retry the call with the regular builder
		if !graph.IsErrInvalidDelta(err) {
			return nil, nil, DeltaUpdate{}, err
		}

		resetDelta = true
		errs = nil
	}

	builder := service.Client().UsersById(user).MailFoldersById(directoryID).Messages().Delta()
	pgr := &mailPager{service, builder, options}

	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return added, removed, DeltaUpdate{deltaURL, resetDelta}, errs.ErrorOrNil()
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

// Serialize transforms the mail item into a byte slice.
func (c Mail) Serialize(
	ctx context.Context,
	item serialization.Parsable,
	user, itemID string,
) ([]byte, error) {
	msg, ok := item.(models.Messageable)
	if !ok {
		return nil, fmt.Errorf("expected Messageable, got %T", item)
	}

	var (
		err    error
		writer = kioser.NewJsonSerializationWriter()
	)

	defer writer.Close()

	if err = writer.WriteObjectValue("", msg); err != nil {
		return nil, support.SetNonRecoverableError(errors.Wrap(err, itemID))
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, errors.Wrap(err, "serializing email")
	}

	return bs, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func MailInfo(msg models.Messageable) *details.ExchangeInfo {
	sender := ""
	subject := ""
	received := time.Time{}
	created := time.Time{}

	if msg.GetSender() != nil &&
		msg.GetSender().GetEmailAddress() != nil &&
		msg.GetSender().GetEmailAddress().GetAddress() != nil {
		sender = *msg.GetSender().GetEmailAddress().GetAddress()
	}

	if msg.GetSubject() != nil {
		subject = *msg.GetSubject()
	}

	if msg.GetReceivedDateTime() != nil {
		received = *msg.GetReceivedDateTime()
	}

	if msg.GetCreatedDateTime() != nil {
		created = *msg.GetCreatedDateTime()
	}

	return &details.ExchangeInfo{
		ItemType: details.ExchangeMail,
		Sender:   sender,
		Subject:  subject,
		Received: received,
		Created:  created,
		Modified: orNow(msg.GetLastModifiedDateTime()),
	}
}
