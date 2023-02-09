package api

import (
	"context"
	"fmt"
	"time"

	"github.com/alcionai/clues"
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
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Contacts() Contacts {
	return Contacts{c}
}

// Contacts is an interface-compliant provider of the client.
type Contacts struct {
	Client
}

// ---------------------------------------------------------------------------
// methods
// ---------------------------------------------------------------------------

// CreateContactFolder makes a contact folder with the displayName of folderName.
// If successful, returns the created folder object.
func (c Contacts) CreateContactFolder(
	ctx context.Context,
	user, folderName string,
) (models.ContactFolderable, error) {
	requestBody := models.NewContactFolder()
	temp := folderName
	requestBody.SetDisplayName(&temp)

	return c.stable.Client().UsersById(user).ContactFolders().Post(ctx, requestBody, nil)
}

// DeleteContainer deletes the ContactFolder associated with the M365 ID if permissions are valid.
func (c Contacts) DeleteContainer(
	ctx context.Context,
	user, folderID string,
) error {
	return c.stable.Client().UsersById(user).ContactFoldersById(folderID).Delete(ctx, nil)
}

// GetItem retrieves a Contactable item.
func (c Contacts) GetItem(
	ctx context.Context,
	user, itemID string,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	var (
		cont models.Contactable
		err  error
	)

	err = graph.RunWithRetry(func() error {
		cont, err = c.stable.Client().UsersById(user).ContactsById(itemID).Get(ctx, nil)
		return err
	})

	if err != nil {
		return nil, nil, err
	}

	return cont, ContactInfo(cont), nil
}

func (c Contacts) GetContainerByID(
	ctx context.Context,
	userID, dirID string,
) (graph.Container, error) {
	ofcf, err := optionsForContactFolderByID([]string{"displayName", "parentFolderId"})
	if err != nil {
		return nil, errors.Wrap(err, "options for contact folder")
	}

	var resp models.ContactFolderable

	err = graph.RunWithRetry(func() error {
		resp, err = c.stable.Client().UsersById(userID).ContactFoldersById(dirID).Get(ctx, ofcf)
		return err
	})

	return resp, err
}

// EnumerateContainers iterates through all of the users current
// contacts folders, converting each to a graph.CacheFolder, and calling
// fn(cf) on each one.  If fn(cf) errors, the error is aggregated
// into a multierror that gets returned to the caller.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Contacts) EnumerateContainers(
	ctx context.Context,
	userID, baseDirID string,
	fn func(graph.CacheFolder) error,
) error {
	service, err := c.service()
	if err != nil {
		return err
	}

	var (
		errs   *multierror.Error
		resp   models.ContactFolderCollectionResponseable
		fields = []string{"displayName", "parentFolderId"}
	)

	ofcf, err := optionsForContactChildFolders(fields)
	if err != nil {
		return errors.Wrapf(err, "options for contact child folders: %v", fields)
	}

	builder := service.Client().
		UsersById(userID).
		ContactFoldersById(baseDirID).
		ChildFolders()

	for {
		err = graph.RunWithRetry(func() error {
			resp, err = builder.Get(ctx, ofcf)
			return err
		})

		if err != nil {
			return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, fold := range resp.GetValue() {
			if err := checkIDAndName(fold); err != nil {
				errs = multierror.Append(err, errs)
				continue
			}

			temp := graph.NewCacheFolder(fold, nil, nil)
			if err := fn(temp); err != nil {
				errs = multierror.Append(err, errs)
				continue
			}
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = users.NewItemContactFoldersItemChildFoldersRequestBuilder(*resp.GetOdataNextLink(), service.Adapter())
	}

	return errs.ErrorOrNil()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager = &contactPager{}

type contactPager struct {
	gs      graph.Servicer
	builder *users.ItemContactFoldersItemContactsDeltaRequestBuilder
	options *users.ItemContactFoldersItemContactsDeltaRequestBuilderGetRequestConfiguration
}

func (p *contactPager) getPage(ctx context.Context) (api.DeltaPageLinker, error) {
	var (
		resp api.DeltaPageLinker
		err  error
	)

	err = graph.RunWithRetry(func() error {
		resp, err = p.builder.Get(ctx, p.options)
		return err
	})

	return resp, err
}

func (p *contactPager) setNext(nextLink string) {
	p.builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *contactPager) valuesIn(pl api.DeltaPageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Contactable](pl)
}

func (c Contacts) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	user, directoryID, oldDelta string,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	var (
		errs       *multierror.Error
		resetDelta bool
	)

	ctx = clues.AddAll(
		ctx,
		"category", selectors.ExchangeContact,
		"folder_id", directoryID)

	options, err := optionsForContactFoldersItemDelta([]string{"parentFolderId"})
	if err != nil {
		return nil, nil, DeltaUpdate{}, errors.Wrap(err, "getting query options")
	}

	if len(oldDelta) > 0 {
		builder := users.NewItemContactFoldersItemContactsDeltaRequestBuilder(oldDelta, service.Adapter())
		pgr := &contactPager{service, builder, options}

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

	builder := service.Client().UsersById(user).ContactFoldersById(directoryID).Contacts().Delta()
	pgr := &contactPager{service, builder, options}

	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return added, removed, DeltaUpdate{deltaURL, resetDelta}, errs.ErrorOrNil()
}

// ---------------------------------------------------------------------------
// Serialization
// ---------------------------------------------------------------------------

// Serialize rserializes the item into a byte slice.
func (c Contacts) Serialize(
	ctx context.Context,
	item serialization.Parsable,
	user, itemID string,
) ([]byte, error) {
	contact, ok := item.(models.Contactable)
	if !ok {
		return nil, fmt.Errorf("expected Contactable, got %T", item)
	}

	ctx = clues.Add(ctx, "item_id", *contact.GetId())

	var (
		err    error
		writer = kioser.NewJsonSerializationWriter()
	)

	defer writer.Close()

	if err = writer.WriteObjectValue("", contact); err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, errors.Wrap(err, "serializing contact")
	}

	return bs, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func ContactInfo(contact models.Contactable) *details.ExchangeInfo {
	name := ""
	created := time.Time{}

	if contact.GetDisplayName() != nil {
		name = *contact.GetDisplayName()
	}

	if contact.GetCreatedDateTime() != nil {
		created = *contact.GetCreatedDateTime()
	}

	return &details.ExchangeInfo{
		ItemType:    details.ExchangeContact,
		ContactName: name,
		Created:     created,
		Modified:    orNow(contact.GetLastModifiedDateTime()),
	}
}
