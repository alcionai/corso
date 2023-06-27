package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// Variables
var (
	ErrMailBoxSettingsNotFound = clues.New("mailbox settings not found")
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
	service, err := c.Service()
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
		return nil, graph.Wrap(ctx, err, "getting user")
	}

	return resp, err
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

// ---------------------------------------------------------------------------
// Info
// ---------------------------------------------------------------------------

func (c Users) GetInfo(ctx context.Context, userID string) (*UserInfo, error) {
	var (
		// Assume all services are enabled
		// then filter down to only services the user has enabled
		userInfo = newUserInfo()

		mailFolderFound = true
	)

	// check whether the user is able to access their onedrive drive.
	// if they cannot, we can assume they are ineligible for onedrive backups.
	if _, err := c.GetDefaultDrive(ctx, userID); err != nil {
		if !clues.HasLabel(err, graph.LabelsMysiteNotFound) || clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			logger.CtxErr(ctx, err).Error("getting user's drive")
			return nil, graph.Wrap(ctx, err, "getting user's drive")
		}

		logger.Ctx(ctx).Info("resource owner does not have a drive")
		delete(userInfo.ServicesEnabled, path.OneDriveService)
	}

	// check whether the user is able to access their inbox.
	// if they cannot, we can assume they are ineligible for exchange backups.
	inbx, err := c.GetMailInbox(ctx, userID)
	if err != nil {
		err = graph.Stack(ctx, err)

		if graph.IsErrUserNotFound(err) {
			logger.CtxErr(ctx, err).Error("user not found")
			return nil, clues.Stack(graph.ErrResourceOwnerNotFound, err)
		}

		if !graph.IsErrExchangeMailFolderNotFound(err) {
			logger.CtxErr(ctx, err).Error("getting user's mail folder")
			return nil, err
		}

		logger.Ctx(ctx).Info("resource owner does not have a mailbox enabled")
		delete(userInfo.ServicesEnabled, path.ExchangeService)

		mailFolderFound = false
	}

	// check whether the user has accessible mailbox settings.
	// if they do, aggregate them in the MailboxInfo
	mi := MailboxInfo{
		ErrGetMailBoxSetting: []error{},
	}

	if !mailFolderFound {
		mi.ErrGetMailBoxSetting = append(mi.ErrGetMailBoxSetting, ErrMailBoxSettingsNotFound)
		userInfo.Mailbox = mi

		return userInfo, nil
	}

	mboxSettings, err := c.getMailboxSettings(ctx, userID)
	if err != nil {
		logger.CtxErr(ctx, err).Info("err getting user's mailbox settings")

		if !graph.IsErrAccessDenied(err) {
			return nil, graph.Wrap(ctx, err, "getting user's mailbox settings")
		}

		mi.ErrGetMailBoxSetting = append(mi.ErrGetMailBoxSetting, clues.New("access denied"))
	} else {
		mi = parseMailboxSettings(mboxSettings, mi)
	}

	err = c.getFirstInboxMessage(ctx, userID, ptr.Val(inbx.GetId()))
	if err != nil {
		if !graph.IsErrQuotaExceeded(err) {
			return nil, err
		}

		userInfo.Mailbox.QuotaExceeded = graph.IsErrQuotaExceeded(err)
	}

	userInfo.Mailbox = mi

	return userInfo, nil
}

func (c Users) getMailboxSettings(
	ctx context.Context,
	userID string,
) (models.Userable, error) {
	settings, err := users.
		NewUserItemRequestBuilder(
			fmt.Sprintf("https://graph.microsoft.com/v1.0/users/%s/mailboxSettings", userID),
			c.Stable.Adapter(),
		).
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
		ByMailFolderId("inbox").
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
func (c Users) getFirstInboxMessage(
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
