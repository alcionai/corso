package m365

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/discovery"

	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
  "github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// ServiceAccess is true if a resource owner is capable of
// accessing or utilizing the specified service.
type ServiceAccess struct {
	Exchange bool
	// TODO: onedrive, sharepoint
}

// ---------------------------------------------------------------------------
// Users
// ---------------------------------------------------------------------------

// User is the minimal information required to identify and display a user.
type User struct {
	PrincipalName string
	ID            string
	Name          string
	UserPurpose   string
}

type UserInfo struct {
	ServicesEnabled ServiceAccess
}

// UsersCompat returns a list of users in the specified M365 tenant.
// TODO(ashmrtn): Remove when upstream consumers of the SDK support the fault
// package.
func UsersCompat(ctx context.Context, acct account.Account) ([]*User, error) {
	errs := fault.New(true)

	users, err := Users(ctx, acct, errs)
	if err != nil {
		return nil, err
	}

	return users, errs.Failure()
}

// Users returns a list of users in the specified M365 tenant
func Users(ctx context.Context, acct account.Account, errs *fault.Bus) ([]*User, error) {
	uapi, err := makeUserAPI(acct)
	if err != nil {
		return nil, clues.Wrap(err, "getting users").WithClues(ctx)
	}

	users, err := discovery.Users(ctx, uapi, errs)
	if err != nil {
		return nil, err
	}

	ret := make([]*User, 0, len(users))

	for _, u := range users {
		pu, err := parseUser(u)
		if err != nil {
			return nil, clues.Wrap(err, "formatting user data")
		}

		pu.UserPurpose, err = discovery.UsersDetails(ctx, acct, pu.ID, errs)
		if err != nil {
			// TODO: handle the access denied scenario
			if graph.IsErrAccessDenied(err) {
				// // clues.Wrap(err, fmt.Sprintf("access denied for %s", pu.ID))
				// errList = append(errList, fmt.Errorf("access denied for %s", pu.ID))
				logger.Ctx(ctx).Infow("access denied for %s", pu.ID)
			}
		}

		ret = append(ret, pu)
	}

	return ret, nil
}

// parseUser extracts information from `models.Userable` we care about
func parseUser(item models.Userable) (*User, error) {
	if item.GetUserPrincipalName() == nil {
		return nil, clues.New("user missing principal name").
			With("user_id", *item.GetId()) // TODO: pii
	}

	u := &User{
		PrincipalName: ptr.Val(item.GetUserPrincipalName()),
		ID:            ptr.Val(item.GetId()),
		Name:          ptr.Val(item.GetDisplayName()),
	}

	return u, nil
}

// UsersMap retrieves all users in the tenant, and returns two maps: one id-to-principalName,
// and one principalName-to-id.
func UsersMap(
	ctx context.Context,
	acct account.Account,
	errs *fault.Bus,
) (common.IDsNames, error) {
	users, err := Users(ctx, acct, errs)
	if err != nil {
		return common.IDsNames{}, err
	}

	var (
		idToName = make(map[string]string, len(users))
		nameToID = make(map[string]string, len(users))
	)

	for _, u := range users {
		id, name := strings.ToLower(u.ID), strings.ToLower(u.PrincipalName)
		idToName[id] = name
		nameToID[name] = id
	}

	ins := common.IDsNames{
		IDToName: idToName,
		NameToID: nameToID,
	}

	return ins, nil
}

// UserInfo returns the corso-specific set of user metadata.
func GetUserInfo(
	ctx context.Context,
	acct account.Account,
	userID string,
) (*UserInfo, error) {
	uapi, err := makeUserAPI(acct)
	if err != nil {
		return nil, clues.Wrap(err, "getting user info").WithClues(ctx)
	}

	ui, err := discovery.UserInfo(ctx, uapi, userID)
	if err != nil {
		return nil, err
	}

	info := UserInfo{
		ServicesEnabled: ServiceAccess{
			Exchange: ui.ServiceEnabled(path.ExchangeService),
		},
	}

	return &info, nil
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
	sites, err := discovery.Sites(ctx, acct, errs)
	if err != nil {
		return nil, clues.Wrap(err, "initializing M365 graph connection")
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
) (common.IDsNames, error) {
	sites, err := Sites(ctx, acct, errs)
	if err != nil {
		return common.IDsNames{}, err
	}

	ins := common.IDsNames{
		IDToName: make(map[string]string, len(sites)),
		NameToID: make(map[string]string, len(sites)),
	}

	for _, s := range sites {
		ins.IDToName[s.ID] = s.WebURL
		ins.NameToID[s.WebURL] = s.ID
	}

	return ins, nil
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func makeUserAPI(acct account.Account) (api.Users, error) {
	creds, err := acct.M365Config()
	if err != nil {
		return api.Users{}, clues.Wrap(err, "getting m365 account creds")
	}

	cli, err := api.NewClient(creds)
	if err != nil {
		return api.Users{}, clues.Wrap(err, "constructing api client")
	}

	return cli.Users(), nil
}
