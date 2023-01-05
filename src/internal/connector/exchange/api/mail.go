package api

import (
	"context"

	"github.com/hashicorp/go-multierror"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
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

// DeleteMailFolder removes a mail folder with the corresponding M365 ID  from the user's M365 Exchange account
// Reference: https://docs.microsoft.com/en-us/graph/api/mailfolder-delete?view=graph-rest-1.0&tabs=http
func (c Mail) DeleteMailFolder(
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

// RetrieveMessageDataForUser is a GraphRetrievalFunc that returns message data.
func (c Mail) RetrieveMessageDataForUser(
	ctx context.Context,
	user, m365ID string,
) (serialization.Parsable, error) {
	return c.stable.Client().UsersById(user).MessagesById(m365ID).Get(ctx, nil)
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
		errs    *multierror.Error
		builder = service.Client().
			UsersById(userID).
			MailFolders().
			Delta()
	)

	for {
		resp, err := builder.Get(ctx, nil)
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

type mailPager struct {
	gs      graph.Servicer
	builder *users.ItemMailFoldersItemMessagesDeltaRequestBuilder
	options *users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func (p *mailPager) getPage(ctx context.Context) (pageLinker, error) {
	return p.builder.Get(ctx, p.options)
}

func (p *mailPager) setNext(nextLink string) {
	p.builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
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

	errUpdater := func(err error) {
		errs = multierror.Append(errs, errors.Wrap(err, "folder "+directoryID))
	}

	options, err := optionsForFolderMessagesDelta([]string{"isRead"})
	if err != nil {
		return nil, nil, DeltaUpdate{}, errors.Wrap(err, "getting query options")
	}

	if len(oldDelta) > 0 {
		builder := users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(oldDelta, service.Adapter())
		pgr := &mailPager{service, builder, options}

		added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr, errUpdater)
		// note: happy path, not the error condition
		if err == nil {
			return added, removed, DeltaUpdate{deltaURL, false}, errs.ErrorOrNil()
		}
		// only return on error if it is NOT a delta issue.
		// on bad deltas we retry the call with the regular builder
		if graph.IsErrInvalidDelta(err) == nil {
			return nil, nil, DeltaUpdate{}, err
		}

		resetDelta = true
		errs = nil
	}

	builder := service.Client().UsersById(user).MailFoldersById(directoryID).Messages().Delta()
	pgr := &mailPager{service, builder, options}

	added, removed, deltaURL, err := getItemsAddedAndRemovedFromContainer(ctx, pgr, errUpdater)
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return added, removed, DeltaUpdate{deltaURL, resetDelta}, errs.ErrorOrNil()
}
