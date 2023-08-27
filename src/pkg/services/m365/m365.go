package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	commonM365 "github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// interfaces & structs
// ---------------------------------------------------------------------------

type getAller[T any] interface {
	GetAll(ctx context.Context, errs *fault.Bus) ([]T, error)
}

// ---------------------------------------------------------------------------
// Users
// ---------------------------------------------------------------------------

// User is the minimal information required to identify and display a user.
type User struct {
	PrincipalName string
	ID            string
	Name          string
	Info          api.UserInfo
}

// UserNoInfo is the minimal information required to identify and display a user.
// TODO: Remove this once `UsersCompatNoInfo` is removed
type UserNoInfo struct {
	PrincipalName string
	ID            string
	Name          string
}

// UsersCompat returns a list of users in the specified M365 tenant.
// TODO(ashmrtn): Remove when upstream consumers of the SDK support the fault
// package.
func UsersCompat(ctx context.Context, acct account.Account) ([]*User, error) {
	errs := fault.New(true)

	us, err := Users(ctx, acct, errs)
	if err != nil {
		return nil, err
	}

	return us, errs.Failure()
}

// UsersCompatNoInfo returns a list of users in the specified M365 tenant.
// TODO: Remove this once `Info` is removed from the `User` struct and callers
// have switched over
func UsersCompatNoInfo(ctx context.Context, acct account.Account) ([]*UserNoInfo, error) {
	errs := fault.New(true)

	us, err := usersNoInfo(ctx, acct, errs)
	if err != nil {
		return nil, err
	}

	return us, errs.Failure()
}

// UserHasMailbox returns true if the user has an exchange mailbox enabled
// false otherwise, and a nil pointer and an error in case of error
func UserHasMailbox(ctx context.Context, acct account.Account, userID string) (bool, error) {
	ac, err := makeAC(ctx, acct, path.ExchangeService)
	if err != nil {
		return false, clues.Stack(err).WithClues(ctx)
	}

	return commonM365.IsExchangeServiceEnabled(ctx, ac.Users(), userID)
}

func UserGetMailboxInfo(
	ctx context.Context,
	acct account.Account,
	userID string,
) (api.MailboxInfo, error) {
	ac, err := makeAC(ctx, acct, path.ExchangeService)
	if err != nil {
		return api.MailboxInfo{}, clues.Stack(err).WithClues(ctx)
	}

	mi, err := ac.Users().GetMailboxInfo(ctx, userID)
	if err != nil {
		return api.MailboxInfo{}, clues.Stack(err)
	}

	return mi, nil
}

// UserHasDrives returns true if the user has any drives
// false otherwise, and a nil pointer and an error in case of error
func UserHasDrives(ctx context.Context, acct account.Account, userID string) (bool, error) {
	ac, err := makeAC(ctx, acct, path.OneDriveService)
	if err != nil {
		return false, clues.Stack(err).WithClues(ctx)
	}

	return commonM365.IsOneDriveServiceEnabled(ctx, ac.Users(), userID)
}

// usersNoInfo returns a list of users in the specified M365 tenant - with no info
// TODO: Remove this once we remove `Info` from `Users` and instead rely on the `GetUserInfo` API
// to get user information
func usersNoInfo(ctx context.Context, acct account.Account, errs *fault.Bus) ([]*UserNoInfo, error) {
	ac, err := makeAC(ctx, acct, path.UnknownService)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	us, err := ac.Users().GetAll(ctx, errs)
	if err != nil {
		return nil, err
	}

	ret := make([]*UserNoInfo, 0, len(us))

	for _, u := range us {
		pu, err := parseUser(u)
		if err != nil {
			return nil, clues.Wrap(err, "formatting user data")
		}

		puNoInfo := &UserNoInfo{
			PrincipalName: pu.PrincipalName,
			ID:            pu.ID,
			Name:          pu.Name,
		}

		ret = append(ret, puNoInfo)
	}

	return ret, nil
}

// Users returns a list of users in the specified M365 tenant
func Users(ctx context.Context, acct account.Account, errs *fault.Bus) ([]*User, error) {
	ac, err := makeAC(ctx, acct, path.ExchangeService)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	us, err := ac.Users().GetAll(ctx, errs)
	if err != nil {
		return nil, err
	}

	ret := make([]*User, 0, len(us))

	for _, u := range us {
		pu, err := parseUser(u)
		if err != nil {
			return nil, clues.Wrap(err, "formatting user data")
		}

		userInfo, err := ac.Users().GetInfo(ctx, pu.ID)
		if err != nil {
			return nil, clues.Wrap(err, "getting user details")
		}

		pu.Info = *userInfo

		ret = append(ret, pu)
	}

	return ret, nil
}

// parseUser extracts information from `models.Userable` we care about
func parseUser(item models.Userable) (*User, error) {
	if item.GetUserPrincipalName() == nil {
		return nil, clues.New("user missing principal name").
			With("user_id", ptr.Val(item.GetId()))
	}

	u := &User{
		PrincipalName: ptr.Val(item.GetUserPrincipalName()),
		ID:            ptr.Val(item.GetId()),
		Name:          ptr.Val(item.GetDisplayName()),
	}

	return u, nil
}

// UserInfo returns the corso-specific set of user metadata.
// TODO(pandeyabs): Remove support for this API. SDK users would be using
// per service API calls - UserHasMailbox, UserGetMailboxInfo, UserHasDrive, etc.
func GetUserInfo(
	ctx context.Context,
	acct account.Account,
	userID string,
) (*api.UserInfo, error) {
	ac, err := makeAC(ctx, acct, path.ExchangeService)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	ui, err := ac.Users().GetInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	return ui, nil
}

// ---------------------------------------------------------------------------
// Sites
// ---------------------------------------------------------------------------

// Site is the minimal information required to identify and display a SharePoint site.
type Site struct {
	// WebURL is the url for the site, works as an alias for the user name.
	WebURL string

	// ID is of the format: <site collection hostname>.<site collection unique id>.<site unique id>
	// for example: contoso.sharepoint.com,abcdeab3-0ccc-4ce1-80ae-b32912c9468d,xyzud296-9f7c-44e1-af81-3c06d0d43007
	ID string

	// DisplayName is the human-readable name of the site.  Normally the plaintext name that the
	// user provided when they created the site, though it can be changed across time.
	// Ex: webUrl: https://host.com/sites/TestingSite, displayName: "Testing Site"
	DisplayName string
}

// Sites returns a list of Sites in a specified M365 tenant
func Sites(ctx context.Context, acct account.Account, errs *fault.Bus) ([]*Site, error) {
	ac, err := makeAC(ctx, acct, path.SharePointService)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return getAllSites(ctx, ac.Sites())
}

func getAllSites(
	ctx context.Context,
	ga getAller[models.Siteable],
) ([]*Site, error) {
	sites, err := ga.GetAll(ctx, fault.New(true))
	if err != nil {
		if clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			return nil, clues.Stack(graph.ErrServiceNotEnabled, err)
		}

		return nil, clues.Wrap(err, "retrieving sites")
	}

	ret := make([]*Site, 0, len(sites))

	for _, s := range sites {
		ps, err := parseSite(s)
		if err != nil {
			return nil, clues.Wrap(err, "parsing siteable")
		}

		ret = append(ret, ps)
	}

	return ret, nil
}

// parseSite extracts the information from `models.Siteable` we care about
func parseSite(item models.Siteable) (*Site, error) {
	s := &Site{
		ID:          ptr.Val(item.GetId()),
		WebURL:      ptr.Val(item.GetWebUrl()),
		DisplayName: ptr.Val(item.GetDisplayName()),
	}

	return s, nil
}

// SitesMap retrieves all sites in the tenant, and returns two maps: one id-to-webURL,
// and one webURL-to-id.
func SitesMap(
	ctx context.Context,
	acct account.Account,
	errs *fault.Bus,
) (idname.Cacher, error) {
	sites, err := Sites(ctx, acct, errs)
	if err != nil {
		return idname.NewCache(nil), err
	}

	itn := make(map[string]string, len(sites))

	for _, s := range sites {
		itn[s.ID] = s.WebURL
	}

	return idname.NewCache(itn), nil
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func makeAC(
	ctx context.Context,
	acct account.Account,
	pst path.ServiceType,
) (api.Client, error) {
	api.InitConcurrencyLimit(ctx, pst)

	creds, err := acct.M365Config()
	if err != nil {
		return api.Client{}, clues.Wrap(err, "getting m365 account creds")
	}

	cli, err := api.NewClient(creds, control.DefaultOptions())
	if err != nil {
		return api.Client{}, clues.Wrap(err, "constructing api client")
	}

	return cli, nil
}
