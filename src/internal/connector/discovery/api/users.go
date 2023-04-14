package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
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
// structs
// ---------------------------------------------------------------------------

type UserInfo struct {
	DiscoveredServices map[path.ServiceType]struct{}
}

func newUserInfo() *UserInfo {
	return &UserInfo{
		DiscoveredServices: map[path.ServiceType]struct{}{
			path.ExchangeService: {},
			path.OneDriveService: {},
		},
	}
}

// ServiceEnabled returns true if the UserInfo has an entry for the
// service.  If no entry exists, the service is assumed to not be enabled.
func (ui *UserInfo) ServiceEnabled(service path.ServiceType) bool {
	if ui == nil || len(ui.DiscoveredServices) == 0 {
		return false
	}

	_, ok := ui.DiscoveredServices[service]

	return ok
}

// ---------------------------------------------------------------------------
// methods
// ---------------------------------------------------------------------------

const (
	userSelectID            = "id"
	userSelectPrincipalName = "userPrincipalName"
	userSelectDisplayName   = "displayName"
)

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

// I can't believe I have to do this.
var t = true

func userOptions(fs *string) *users.UsersRequestBuilderGetRequestConfiguration {
	headers := abstractions.NewRequestHeaders()
	headers.Add("ConsistencyLevel", "eventual")

	return &users.UsersRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &users.UsersRequestBuilderGetQueryParameters{
			Select: []string{userSelectID, userSelectPrincipalName, userSelectDisplayName},
			Filter: fs,
			Count:  &t,
		},
	}
}

// GetAll retrieves all users.
func (c Users) GetAll(ctx context.Context, errs *fault.Bus) ([]models.Userable, error) {
	service, err := c.service()
	if err != nil {
		return nil, err
	}

	var resp models.UserCollectionResponseable

	resp, err = service.Client().Users().Get(ctx, userOptions(&userFilterNoGuests))

	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting all users")
	}

	iter, err := msgraphgocore.NewPageIterator(
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

	iterator := func(item any) bool {
		if el.Failure() != nil {
			return false
		}

		u, err := validateUser(item)
		if err != nil {
			el.AddRecoverable(graph.Wrap(ctx, err, "validating user"))
		} else {
			us = append(us, u)
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

	resp, err = c.stable.Client().UsersById(identifier).Get(ctx, nil)

	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting user")
	}

	return resp, err
}

// GetUserInfo provide information about a user like- userPurpose/type,
// if mailbox and onedrive is present for user.
func (c Users) GetUserInfo(ctx context.Context, userID string) (string, bool, bool, string, error) {
	var (
		rawURL                         = "https://graph.microsoft.com/v1.0/users/" + userID + "/mailboxSettings"
		apadtor                        = c.stable.Adapter()
		builder                        = users.NewUserItemRequestBuilder(rawURL, apadtor)
		hasMailBox, hasOneDrive, ok    bool
		userPurpose, errGetUserPurpose string
		err                            error
	)

	// verify mailbox enabled for user
	_, err = c.stable.Client().UsersById(userID).MailFolders().Get(ctx, nil)
	hasMailBox = true

	if err != nil {
		if !graph.IsErrExchangeMailFolderNotFound(err) {
			return "", false, false, "", graph.Wrap(ctx, err, "error getting mail folder")
		}

		logger.Ctx(ctx).Infof("resource owner does not have a mailbox enabled")

		hasMailBox = false
	}

	// verify onedrive enabled for user
	_, err = c.stable.Client().UsersById(userID).Drives().Get(ctx, nil)
	hasOneDrive = true

	if err != nil {
		err = graph.Stack(ctx, err)
		if !clues.HasLabel(err, graph.LabelsMysiteNotFound) {
			return "", false, false, "", graph.Wrap(ctx, err, "getting onedrive")
		}

		logger.Ctx(ctx).Infof("resource owner does not have a drive")

		hasOneDrive = false
	}

	// get userPurpose for user
	newItem, err := builder.Get(ctx, nil)
	if err == nil {
		userPurpose, ok = ptr.ValOK(newItem.GetAdditionalData()["userPurpose"].(*string))
		if !ok {
			return "", false, false, "", graph.Wrap(ctx, err, "getting purpose")
		}
	}

	if err != nil {
		if !(graph.IsErrAccessDenied(err) || graph.IsErrExchangeMailFolderNotFound(err)) {
			return "", false, false, "", graph.Wrap(ctx, err, "getting purpose")
		}

		// Not the best way to handle the errors. But since we don't want to break
		// this request and just want to convey a message that access are invalid have
		// added value here. Can also add another variable which specify the same if access
		// denied
		if graph.IsErrAccessDenied(err) {
			errGetUserPurpose = "access denied"
		}

		if graph.IsErrExchangeMailFolderNotFound(err) {
			errGetUserPurpose = "not found"
			hasMailBox = false
		}
	}

	return userPurpose, hasMailBox, hasOneDrive, errGetUserPurpose, nil
}

// GetIDAndName looks up the user matching the given ID, and returns
// its canonical ID and the PrincipalName as the name.
func (c Users) GetIDAndName(ctx context.Context, userID string) (string, string, error) {
	u, err := c.GetByID(ctx, userID)
	if err != nil {
		return "", "", err
	}

	return ptr.Val(u.GetId()), ptr.Val(u.GetUserPrincipalName()), nil
}

func (c Users) GetInfo(ctx context.Context, userID string) (*UserInfo, error) {
	// Assume all services are enabled
	// then filter down to only services the user has enabled
	var (
		err      error
		userInfo = newUserInfo()
	)

	// TODO: OneDrive
	_, err = c.stable.Client().UsersById(userID).MailFolders().Get(ctx, nil)
	if err != nil {
		if !graph.IsErrExchangeMailFolderNotFound(err) {
			return nil, graph.Wrap(ctx, err, "getting user's mail folder")
		}

		delete(userInfo.DiscoveredServices, path.ExchangeService)
	}

	return userInfo, nil
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

// validateUser ensures the item is a Userable, and contains the necessary
// identifiers that we handle with all users.
// returns the item as a Userable model.
func validateUser(item any) (models.Userable, error) {
	m, ok := item.(models.Userable)
	if !ok {
		return nil, clues.New(fmt.Sprintf("unexpected model: %T", item))
	}

	if m.GetId() == nil {
		return nil, clues.New("missing ID")
	}

	if m.GetUserPrincipalName() == nil {
		return nil, clues.New("missing principalName")
	}

	return m, nil
}
