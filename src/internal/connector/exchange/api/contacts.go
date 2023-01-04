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

// CreateContactFolder makes a contact folder with the displayName of folderName.
// If successful, returns the created folder object.
func (c Client) CreateContactFolder(
	ctx context.Context,
	user, folderName string,
) (models.ContactFolderable, error) {
	requestBody := models.NewContactFolder()
	temp := folderName
	requestBody.SetDisplayName(&temp)

	return c.stable.Client().UsersById(user).ContactFolders().Post(ctx, requestBody, nil)
}

// DeleteContactFolder deletes the ContactFolder associated with the M365 ID if permissions are valid.
// Errors returned if the function call was not successful.
func (c Client) DeleteContactFolder(
	ctx context.Context,
	user, folderID string,
) error {
	return c.stable.Client().UsersById(user).ContactFoldersById(folderID).Delete(ctx, nil)
}

// RetrieveContactDataForUser is a GraphRetrievalFun that returns all associated fields.
func (c Client) RetrieveContactDataForUser(
	ctx context.Context,
	user, m365ID string,
) (serialization.Parsable, error) {
	return c.stable.Client().UsersById(user).ContactsById(m365ID).Get(ctx, nil)
}

// GetAllContactFolderNamesForUser is a GraphQuery function for getting
// ContactFolderId and display names for contacts. All other information is omitted.
// Does not return the default Contact Folder
func (c Client) GetAllContactFolderNamesForUser(
	ctx context.Context,
	user string,
) (serialization.Parsable, error) {
	options, err := optionsForContactFolders([]string{"displayName", "parentFolderId"})
	if err != nil {
		return nil, err
	}

	return c.stable.Client().UsersById(user).ContactFolders().Get(ctx, options)
}

func (c Client) GetContactFolderByID(
	ctx context.Context,
	userID, dirID string,
	optionalFields ...string,
) (models.ContactFolderable, error) {
	fields := append([]string{"displayName", "parentFolderId"}, optionalFields...)

	ofcf, err := optionsForContactFolderByID(fields)
	if err != nil {
		return nil, errors.Wrapf(err, "options for contact folder: %v", fields)
	}

	return c.stable.Client().
		UsersById(userID).
		ContactFoldersById(dirID).
		Get(ctx, ofcf)
}

// TODO: we want this to be the full handler, not only the builder.
// but this halfway point minimizes changes for now.
func (c Client) GetContactChildFoldersBuilder(
	ctx context.Context,
	userID, baseDirID string,
	optionalFields ...string,
) (
	*users.ItemContactFoldersItemChildFoldersRequestBuilder,
	*users.ItemContactFoldersItemChildFoldersRequestBuilderGetRequestConfiguration,
	*graph.Service,
	error,
) {
	service, err := c.service()
	if err != nil {
		return nil, nil, nil, err
	}

	fields := append([]string{"displayName", "parentFolderId"}, optionalFields...)

	ofcf, err := optionsForContactChildFolders(fields)
	if err != nil {
		return nil, nil, nil, errors.Wrapf(err, "options for contact child folders: %v", fields)
	}

	builder := service.Client().
		UsersById(userID).
		ContactFoldersById(baseDirID).
		ChildFolders()

	return builder, ofcf, service, nil
}

// FetchContactIDsFromDirectory function that returns a list of  all the m365IDs of the contacts
// of the targeted directory
func (c Client) FetchContactIDsFromDirectory(
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

	options, err := optionsForContactFoldersItemDelta([]string{"parentFolderId"})
	if err != nil {
		return nil, nil, DeltaUpdate{}, errors.Wrap(err, "getting query options")
	}

	getIDs := func(builder *users.ItemContactFoldersItemContactsDeltaRequestBuilder) error {
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

			builder = users.NewItemContactFoldersItemContactsDeltaRequestBuilder(*nextLink, service.Adapter())
		}

		return nil
	}

	if len(oldDelta) > 0 {
		err := getIDs(users.NewItemContactFoldersItemContactsDeltaRequestBuilder(oldDelta, service.Adapter()))
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

	builder := service.Client().
		UsersById(user).
		ContactFoldersById(directoryID).
		Contacts().
		Delta()

	if err := getIDs(builder); err != nil {
		return nil, nil, DeltaUpdate{}, err
	}

	return ids, removedIDs, DeltaUpdate{deltaURL, resetDelta}, errs.ErrorOrNil()
}
