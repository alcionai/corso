package api

import (
	"context"
	"fmt"
	"os"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
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

	mdl, err := c.stable.Client().UsersById(user).ContactFolders().Post(ctx, requestBody, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating contact folder")
	}

	return mdl, nil
}

// DeleteContainer deletes the ContactFolder associated with the M365 ID if permissions are valid.
func (c Contacts) DeleteContainer(
	ctx context.Context,
	user, folderID string,
) error {
	err := c.stable.Client().UsersById(user).ContactFoldersById(folderID).Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

// GetItem retrieves a Contactable item.
func (c Contacts) GetItem(
	ctx context.Context,
	user, itemID string,
	_ *fault.Bus, // no attachments to iterate over, so this goes unused
) (serialization.Parsable, *details.ExchangeInfo, error) {
	cont, err := c.stable.Client().UsersById(user).ContactsById(itemID).Get(ctx, nil)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	return cont, ContactInfo(cont), nil
}

func (c Contacts) GetContainerByID(
	ctx context.Context,
	userID, dirID string,
) (graph.Container, error) {
	ofcf, err := optionsForContactFolderByID([]string{"displayName", "parentFolderId"})
	if err != nil {
		return nil, graph.Wrap(ctx, err, "setting contact folder options")
	}

	resp, err := c.stable.Client().UsersById(userID).ContactFoldersById(dirID).Get(ctx, ofcf)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

// EnumerateContainers iterates through all of the users current
// contacts folders, converting each to a graph.CacheFolder, and calling
// fn(cf) on each one.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Contacts) EnumerateContainers(
	ctx context.Context,
	userID, baseDirID string,
	fn func(graph.CacheFolder) error,
	errs *fault.Bus,
) error {
	service, err := c.service()
	if err != nil {
		return graph.Stack(ctx, err)
	}

	fields := []string{"displayName", "parentFolderId"}

	ofcf, err := optionsForContactChildFolders(fields)
	if err != nil {
		return graph.Wrap(ctx, err, "setting contact child folder options")
	}

	el := errs.Local()
	builder := service.Client().
		UsersById(userID).
		ContactFoldersById(baseDirID).
		ChildFolders()

	for {
		if el.Failure() != nil {
			break
		}

		resp, err := builder.Get(ctx, ofcf)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		for _, fold := range resp.GetValue() {
			if el.Failure() != nil {
				return el.Failure()
			}

			if err := checkIDAndName(fold); err != nil {
				errs.AddRecoverable(graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			fctx := clues.Add(
				ctx,
				"container_id", ptr.Val(fold.GetId()),
				"container_display_name", ptr.Val(fold.GetDisplayName()))

			temp := graph.NewCacheFolder(fold, nil, nil)
			if err := fn(temp); err != nil {
				errs.AddRecoverable(graph.Stack(fctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemContactFoldersItemChildFoldersRequestBuilder(link, service.Adapter())
	}

	return el.Failure()
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
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
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
		return nil, nil, DeltaUpdate{}, graph.Stack(ctx, err)
	}

	var resetDelta bool

	ctx = clues.Add(
		ctx,
		"category", selectors.ExchangeContact,
		"container_id", directoryID)

	options, err := optionsForContactFoldersItemDelta([]string{"parentFolderId"})
	if err != nil {
		return nil,
			nil,
			DeltaUpdate{},
			graph.Wrap(ctx, err, "setting contact folder options")
	}

	if len(oldDelta) > 0 {
		var (
			builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(oldDelta, service.Adapter())
			pgr     = &contactPager{service, builder, options}
		)

		added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
		// note: happy path, not the error condition
		if err == nil {
			return added, removed, DeltaUpdate{deltaURL, false}, err
		}

		// only return on error if it is NOT a delta issue.
		// on bad deltas we retry the call with the regular builder
		if !graph.IsErrInvalidDelta(err) {
			return nil, nil, DeltaUpdate{}, graph.Stack(ctx, err)
		}

		resetDelta = true
	}

	builder := service.Client().UsersById(user).ContactFoldersById(directoryID).Contacts().Delta()
	pgr := &contactPager{service, builder, options}

	if len(os.Getenv("CORSO_URL_LOGGING")) > 0 {
		gri, err := builder.ToGetRequestInformation(ctx, options)
		if err != nil {
			logger.Ctx(ctx).Errorw("getting builder info", "error", err)
		} else {
			logger.Ctx(ctx).
				With("user", user, "container", directoryID).
				Warnw("builder path-parameters", "path_parameters", gri.PathParameters)
		}
	}

	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return added, removed, DeltaUpdate{deltaURL, resetDelta}, nil
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
		return nil, clues.Wrap(fmt.Errorf("parseable type: %T", item), "parsable is not a Contactable")
	}

	ctx = clues.Add(ctx, "item_id", ptr.Val(contact.GetId()))

	var (
		err    error
		writer = kioser.NewJsonSerializationWriter()
	)

	defer writer.Close()

	if err = writer.WriteObjectValue("", contact); err != nil {
		return nil, graph.Stack(ctx, err)
	}

	bs, err := writer.GetSerializedContent()
	if err != nil {
		return nil, graph.Wrap(ctx, err, "serializing contact")
	}

	return bs, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func ContactInfo(contact models.Contactable) *details.ExchangeInfo {
	name := ptr.Val(contact.GetDisplayName())
	created := ptr.Val(contact.GetCreatedDateTime())

	return &details.ExchangeInfo{
		ItemType:    details.ExchangeContact,
		ContactName: name,
		Created:     created,
		Modified:    ptr.OrNow(contact.GetLastModifiedDateTime()),
	}
}
