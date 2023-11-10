package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// Variables
var (
	ErrMailBoxNotFound             = clues.New("mailbox not found")
	ErrMailBoxSettingsAccessDenied = clues.New("mailbox settings access denied")
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Users() Users {
	return Users{c}
}

// Users is an interface-compliant provider of the client.
type Users struct {
	Client
}

// ---------------------------------------------------------------------------
// User CRUD
// ---------------------------------------------------------------------------

// Filter out both guest users, and (for on-prem installations) non-synced users.
// The latter filter makes an assumption that no on-prem users are guests; this might
// require more fine-tuned controls in the future.
// https://stackoverflow.com/questions/64044266/error-message-unsupported-or-invalid-query-filter-clause-specified-for-property
//
// ne 'Guest' ensures we don't filter out users where userType = null, which can happen
// for user accounts created prior to 2014.  In order to use the `ne` comparator, we
// MUST include $count=true and the ConsistencyLevel: eventual header.
// https://stackoverflow.com/questions/49340485/how-to-filter-users-by-usertype-null
//
//nolint:lll
var userFilterNoGuests = "onPremisesSyncEnabled eq true OR userType ne 'Guest'"

// GetAll retrieves all users.
func (c Users) GetAll(
	ctx context.Context,
	errs *fault.Bus,
) ([]models.Userable, error) {
	service, err := c.Service(c.counter)
	if err != nil {
		return nil, err
	}

	config := &users.UsersRequestBuilderGetRequestConfiguration{
		Headers: newEventualConsistencyHeaders(),
		QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
			Select: idAnd(userPrincipalName, displayName),
			Filter: &userFilterNoGuests,
			Count:  ptr.To(true),
		},
	}

	resp, err := service.Client().Users().Get(ctx, config)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting all users")
	}

	iter, err := msgraphgocore.NewPageIterator[models.Userable](
		resp,
		service.Adapter(),
		models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating users iterator")
	}

	var (
		us = make([]models.Userable, 0)
		el = errs.Local()
	)

	iterator := func(item models.Userable) bool {
		if el.Failure() != nil {
			return false
		}

		err := validateUser(item)
		if err != nil {
			el.AddRecoverable(ctx, graph.Wrap(ctx, err, "validating user"))
		} else {
			us = append(us, item)
		}

		return true
	}

	if err := iter.Iterate(ctx, iterator); err != nil {
		return nil, graph.Wrap(ctx, err, "iterating all users")
	}

	return us, el.Failure()
}

// GetByID looks up the user matching the given identifier.  The identifier can be either a
// canonical user id or a princpalName.
func (c Users) GetByID(ctx context.Context, identifier string) (models.Userable, error) {
	var (
		resp models.Userable
		err  error
	)

	resp, err = c.Stable.Client().Users().ByUserId(identifier).Get(ctx, nil)
	if err != nil {
		if graph.IsErrResourceLocked(err) {
			err = clues.Stack(graph.ErrResourceLocked, err)
		}

		return nil, graph.Stack(ctx, err)
	}

	return resp, err
}

// GetIDAndName looks up the user matching the given ID, and returns
// its canonical ID and the PrincipalName as the name.
func (c Users) GetIDAndName(
	ctx context.Context,
	userID string,
	_ CallConfig, // not currently supported
) (string, string, error) {
	u, err := c.GetByID(ctx, userID)
	if err != nil {
		return "", "", err
	}

	return ptr.Val(u.GetId()), ptr.Val(u.GetUserPrincipalName()), nil
}

// GetAllIDsAndNames retrieves all users in the tenant and returns them in an idname.Cacher
func (c Users) GetAllIDsAndNames(ctx context.Context, errs *fault.Bus) (idname.Cacher, error) {
	all, err := c.GetAll(ctx, errs)
	if err != nil {
		return nil, clues.Wrap(err, "getting all users")
	}

	idToName := make(map[string]string, len(all))

	for _, u := range all {
		id := strings.ToLower(ptr.Val(u.GetId()))
		name := strings.ToLower(ptr.Val(u.GetUserPrincipalName()))

		idToName[id] = name
	}

	return idname.NewCache(idToName), nil
}

func appendIfErr(errs []error, err error) []error {
	if err == nil {
		return errs
	}

	return append(errs, err)
}

// EvaluateMailboxError checks whether the provided error can be interpreted
// as "user does not have a mailbox", or whether it is some other error.  If
// the former (no mailbox), returns nil, otherwise returns an error.
func EvaluateMailboxError(err error) error {
	if err == nil {
		return nil
	}

	// must occur before MailFolderNotFound, due to overlapping cases.
	if graph.IsErrUserNotFound(err) {
		return clues.Stack(graph.ErrResourceOwnerNotFound, err)
	}

	if graph.IsErrResourceLocked(err) {
		return clues.Stack(graph.ErrResourceLocked, err)
	}

	if graph.IsErrExchangeMailFolderNotFound(err) || graph.IsErrAuthenticationError(err) {
		return nil
	}

	return err
}

// IsAnyErrMailboxNotFound inspects the secondary errors inside MailboxInfo and
// determines whether the resource has a mailbox.
func IsAnyErrMailboxNotFound(errs []error) bool {
	for _, err := range errs {
		if errors.Is(err, ErrMailBoxNotFound) {
			return true
		}
	}

	return false
}

func (c Users) GetMailboxSettings(
	ctx context.Context,
	userID string,
) (models.Userable, error) {
	settings, err := users.
		NewUserItemRequestBuilder(
			fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/mailboxSettings", userID),
			c.Stable.Adapter()).
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return settings, nil
}

func (c Users) GetMailInbox(
	ctx context.Context,
	userID string,
) (models.MailFolderable, error) {
	inbox, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(MailInbox).
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting MailFolders")
	}

	return inbox, nil
}

func (c Users) GetDefaultDrive(
	ctx context.Context,
	userID string,
) (models.Driveable, error) {
	d, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Drive().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting user's drive")
	}

	return d, nil
}

// TODO: This tries to determine if the user has hit their mailbox
// limit by trying to fetch an item and seeing if we get the quota
// exceeded error. Ideally(if available) we should convert this to
// pull the user's usage via an api and compare if they have used
// up their quota.
func (c Users) GetFirstInboxMessage(
	ctx context.Context,
	userID, inboxID string,
) error {
	config := &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMailFoldersItemMessagesDeltaRequestBuilderGetQueryParameters{
			Select: idAnd(),
		},
		Headers: newPreferHeaders(preferPageSize(1)),
	}

	_, err := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		MailFolders().
		ByMailFolderId(inboxID).
		Messages().
		Delta().
		Get(ctx, config)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// validateUser ensures the item is a Userable, and contains the necessary
// identifiers that we handle with all users.
func validateUser(item models.Userable) error {
	if item.GetId() == nil {
		return clues.New("missing ID")
	}

	if item.GetUserPrincipalName() == nil {
		return clues.New("missing principalName")
	}

	return nil
}
