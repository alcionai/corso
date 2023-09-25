package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/common/tform"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type SiteOwnerType string

const (
	SiteOwnerUnknown SiteOwnerType = ""
	SiteOwnerUser    SiteOwnerType = "user"
	SiteOwnerGroup   SiteOwnerType = "group"
)

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

	OwnerType SiteOwnerType
	// OwnerID may or may not contain the site owner's ID.
	// Requires:
	// * a discoverable site owner type
	// * getByID (the drive expansion doesn't work on paginated data)
	// * lucky chance (not all responses contain an owner ID)
	OwnerID string
}

// SiteByID retrieves a specific site.
func SiteByID(
	ctx context.Context,
	acct account.Account,
	id string,
) (*Site, error) {
	ac, err := makeAC(ctx, acct, path.SharePointService)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	cc := api.CallConfig{
		Expand: []string{"drive"},
	}

	s, err := ac.Sites().GetByID(ctx, id, cc)
	if err != nil {
		return nil, clues.Stack(err)
	}

	return ParseSite(s), nil
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
		ret = append(ret, ParseSite(s))
	}

	return ret, nil
}

// ParseSite extracts the information from `models.Siteable` we care about
func ParseSite(item models.Siteable) *Site {
	s := &Site{
		ID:          ptr.Val(item.GetId()),
		WebURL:      ptr.Val(item.GetWebUrl()),
		DisplayName: ptr.Val(item.GetDisplayName()),
		OwnerType:   SiteOwnerUnknown,
	}

	if item.GetDrive() != nil &&
		item.GetDrive().GetOwner() != nil &&
		item.GetDrive().GetOwner().GetUser() != nil {
		s.OwnerType = SiteOwnerUser
		s.OwnerID = ptr.Val(item.GetDrive().GetOwner().GetUser().GetId())
	}

	if _, ok := item.GetAdditionalData()["group"]; ok {
		s.OwnerType = SiteOwnerGroup

		group, err := tform.AnyValueToT[map[string]any]("group", item.GetAdditionalData())
		if err != nil {
			return s
		}

		s.OwnerID, err = str.AnyValueToString("id", group)
		if err != nil {
			return s
		}
	}

	return s
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
