package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/internal/common/idname"
	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/internal/common/str"
	"github.com/alcionai/canario/src/internal/common/tform"
	"github.com/alcionai/canario/src/pkg/errs/core"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/logger"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
)

type SiteOwnerType string

const (
	SiteOwnerUnknown SiteOwnerType = "unknown"
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
	// OwnerEmail may or may not contain the site owner's email.
	OwnerEmail string
}

// SiteByID retrieves a specific site.
func (c client) SiteByID(
	ctx context.Context,
	id string,
) (*Site, error) {
	cc := api.CallConfig{
		Expand: []string{"drive"},
	}

	return getSiteByID(ctx, c.AC.Sites(), id, cc)
}

func getSiteByID(
	ctx context.Context,
	ga api.GetByIDer[models.Siteable],
	id string,
	cc api.CallConfig,
) (*Site, error) {
	s, err := ga.GetByID(ctx, id, cc)
	if err != nil {
		return nil, clues.Stack(err)
	}

	return ParseSite(ctx, s), nil
}

// Sites returns a list of Sites in a specified M365 tenant
func (c client) Sites(ctx context.Context, errs *fault.Bus) ([]*Site, error) {
	return getAllSites(ctx, c.AC.Sites())
}

func getAllSites(
	ctx context.Context,
	ga getAller[models.Siteable],
) ([]*Site, error) {
	sites, err := ga.GetAll(ctx, fault.New(true))
	if err != nil {
		if clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			return nil, clues.Stack(core.ErrServiceNotEnabled, err)
		}

		return nil, clues.Wrap(err, "retrieving sites")
	}

	ret := make([]*Site, 0, len(sites))

	for _, s := range sites {
		ret = append(ret, ParseSite(ctx, s))
	}

	return ret, nil
}

// ParseSite extracts the information from `models.Siteable` we care about
func ParseSite(ctx context.Context, item models.Siteable) *Site {
	s := &Site{
		ID:          ptr.Val(item.GetId()),
		WebURL:      ptr.Val(item.GetWebUrl()),
		DisplayName: ptr.Val(item.GetDisplayName()),
		OwnerType:   SiteOwnerUnknown,
	}

	if item.GetDrive() != nil &&
		item.GetDrive().GetOwner() != nil &&
		item.GetDrive().GetOwner().GetUser() != nil &&
		// some users might come back with a nil ID
		// most likely in the case of deleted users
		item.GetDrive().GetOwner().GetUser().GetId() != nil {
		s.OwnerType = SiteOwnerUser
		s.OwnerID = ptr.Val(item.GetDrive().GetOwner().GetUser().GetId())

		addtl := item.
			GetDrive().
			GetOwner().
			GetUser().
			GetAdditionalData()

		email, err := str.AnyValueToString("email", addtl)
		if err != nil {
			return s
		}

		s.OwnerEmail = email
	} else if item.GetDrive() != nil && item.GetDrive().GetOwner() != nil {
		ownerItem := item.GetDrive().GetOwner()
		if _, ok := ownerItem.GetAdditionalData()["group"]; ok {
			s.OwnerType = SiteOwnerGroup

			group, err := tform.AnyValueToT[map[string]any]("group", ownerItem.GetAdditionalData())
			if err != nil {
				return s
			}

			// ignore the errors, these might or might not exist
			// if they don't exist, we'll just have an empty string
			s.OwnerID, err = str.AnyValueToString("id", group)
			if err != nil {
				logger.CtxErr(ctx, err).Info("could not parse owner ID")
			}

			s.OwnerEmail, err = str.AnyValueToString("email", group)
			if err != nil {
				logger.CtxErr(ctx, err).Info("could not parse owner email")
			}
		}
	}

	return s
}

// SitesMap retrieves all sites in the tenant, and returns two maps: one id-to-webURL,
// and one webURL-to-id.
func (c client) SitesMap(
	ctx context.Context,
	errs *fault.Bus,
) (idname.Cacher, error) {
	sites, err := c.Sites(ctx, errs)
	if err != nil {
		return idname.NewCache(nil), err
	}

	itn := make(map[string]string, len(sites))

	for _, s := range sites {
		itn[s.ID] = s.WebURL
	}

	return idname.NewCache(itn), nil
}
