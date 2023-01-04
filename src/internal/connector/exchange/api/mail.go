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

// CreateMailFolder makes a mail folder iff a folder of the same name does not exist
// Reference: https://docs.microsoft.com/en-us/graph/api/user-post-mailfolders?view=graph-rest-1.0&tabs=http
func (c Client) NewName(
	ctx context.Context,
	user, folder string,
) (models.MailFolderable, error) {
	isHidden := false
	requestBody := models.NewMailFolder()
	requestBody.SetDisplayName(&folder)
	requestBody.SetIsHidden(&isHidden)

	return c.stable.Client().UsersById(user).MailFolders().Post(ctx, requestBody, nil)
}

func (c Client) CreateMailFolderWithParent(
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
func (c Client) DeleteMailFolder(
	ctx context.Context,
	user, folderID string,
) error {
	return c.stable.Client().UsersById(user).MailFoldersById(folderID).Delete(ctx, nil)
}

// RetrieveMessageDataForUser is a GraphRetrievalFunc that returns message data.
// Attachment field is omitted due to size.
func (c Client) RetrieveMessageDataForUser(
	ctx context.Context,
	user, m365ID string,
) (serialization.Parsable, error) {
	return c.stable.Client().UsersById(user).MessagesById(m365ID).Get(ctx, nil)
}

// GetMailFoldersBuilder retrieves all of the users current mail folders.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
// TODO: we want this to be the full handler, not only the builder.
// but this halfway point minimizes changes for now.
func (c Client) GetAllMailFoldersBuilder(
	ctx context.Context,
	userID string,
) (
	*users.ItemMailFoldersDeltaRequestBuilder,
	*graph.Service,
	error,
) {
	service, err := c.service()
	if err != nil {
		return nil, nil, err
	}

	builder := service.Client().
		UsersById(userID).
		MailFolders().
		Delta()

	return builder, service, nil
}

func (c Client) GetMailFolderByID(
	ctx context.Context,
	userID, dirID string,
	optionalFields ...string,
) (models.MailFolderable, error) {
	ofmf, err := optionsForMailFoldersItem(optionalFields)
	if err != nil {
		return nil, errors.Wrapf(err, "options for mail folder: %v", optionalFields)
	}

	return c.stable.Client().UsersById(userID).MailFoldersById(dirID).Get(ctx, ofmf)
}

// FetchMessageIDsFromDirectory function that returns a list of  all the m365IDs of the exchange.Mail
// of the targeted directory
func (c Client) FetchMessageIDsFromDirectory(
	ctx context.Context,
	user, directoryID, oldDelta string,
) ([]string, []string, DeltaUpdate, error) {
	service, err := c.service()
	if err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	var (
		errs       *multierror.Error
		ids        []string
		removedIDs []string
		deltaURL   string
		resetDelta bool
	)

	options, err := optionsForFolderMessagesDelta([]string{"isRead"})
	if err != nil {
		return nil, nil, DeltaUpdate{}, errors.Wrap(err, "getting query options")
	}

	getIDs := func(builder *users.ItemMailFoldersItemMessagesDeltaRequestBuilder) error {
		for {
			resp, err := builder.Get(ctx, options)
			if err != nil {
				if err := graph.IsErrDeletedInFlight(err); err != nil {
					return err
				}

				if err := graph.IsErrInvalidDelta(err); err != nil {
					return err
				}

				return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
			}

			for _, item := range resp.GetValue() {
				if item.GetId() == nil {
					errs = multierror.Append(
						errs,
						errors.Errorf("item with nil ID in folder %s", directoryID),
					)

					// TODO(ashmrtn): Handle fail-fast.
					continue
				}

				if item.GetAdditionalData()[graph.AddtlDataRemoved] == nil {
					ids = append(ids, *item.GetId())
				} else {
					removedIDs = append(removedIDs, *item.GetId())
				}
			}

			delta := resp.GetOdataDeltaLink()
			if delta != nil && len(*delta) > 0 {
				deltaURL = *delta
			}

			nextLink := resp.GetOdataNextLink()
			if nextLink == nil || len(*nextLink) == 0 {
				break
			}

			builder = users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(*nextLink, service.Adapter())
		}

		return nil
	}

	if len(oldDelta) > 0 {
		err := getIDs(users.NewItemMailFoldersItemMessagesDeltaRequestBuilder(oldDelta, service.Adapter()))
		// note: happy path, not the error condition
		if err == nil {
			return ids, removedIDs, DeltaUpdate{deltaURL, false}, errs.ErrorOrNil()
		}
		// only return on error if it is NOT a delta issue.
		// otherwise we'll retry the call with the regular builder
		if graph.IsErrInvalidDelta(err) == nil {
			return nil, nil, DeltaUpdate{}, err
		}

		resetDelta = true
		errs = nil
	}

	builder := service.Client().UsersById(user).MailFoldersById(directoryID).Messages().Delta()

	if err := getIDs(builder); err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return ids, removedIDs, DeltaUpdate{deltaURL, resetDelta}, errs.ErrorOrNil()
}
