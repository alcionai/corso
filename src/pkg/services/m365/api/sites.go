package api

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/alcionai/clues"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Sites() Sites {
	return Sites{c}
}

// Sites is an interface-compliant provider of the client.
type Sites struct {
	Client
}

// ---------------------------------------------------------------------------
// api calls
// ---------------------------------------------------------------------------

func (c Sites) GetRoot(
	ctx context.Context,
	cc CallConfig,
) (models.Siteable, error) {
	options := &sites.SiteItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.SiteItemRequestBuilderGetQueryParameters{},
	}

	if len(cc.Expand) > 0 {
		options.QueryParameters.Expand = cc.Expand
	}

	resp, err := c.Stable.
		Client().
		Sites().
		BySiteId("root").
		Get(ctx, options)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting root site")
	}

	return resp, nil
}

// GetAll retrieves all sites.
func (c Sites) GetAll(ctx context.Context, errs *fault.Bus) ([]models.Siteable, error) {
	resp, err := c.Stable.Client().Sites().Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting all sites")
	}

	iter, err := msgraphgocore.NewPageIterator[models.Siteable](
		resp,
		c.Stable.Adapter(),
		models.CreateSiteCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating sites iterator")
	}

	var (
		us = make([]models.Siteable, 0)
		el = errs.Local()
	)

	iterator := func(item models.Siteable) bool {
		if el.Failure() != nil {
			return false
		}

		err := ValidateSite(item)
		if errors.Is(err, ErrKnownSkippableCase) {
			// safe to no-op
			return true
		}

		if err != nil {
			el.AddRecoverable(ctx, graph.Wrap(ctx, err, "validating site"))
			return true
		}

		us = append(us, item)

		return true
	}

	if err := iter.Iterate(ctx, iterator); err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating sites")
	}

	return us, el.Failure()
}

const uuidRETmpl = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"

var uuidRE = regexp.MustCompile(uuidRETmpl)

// matches a site ID, with or without a doman name.  Ex, either one of:
// 10rqc2.sharepoint.com,deadbeef-0000-0000-0000-000000000000,beefdead-0000-0000-0000-000000000000
// deadbeef-0000-0000-0000-000000000000,beefdead-0000-0000-0000-000000000000
var siteIDRE = regexp.MustCompile(`(.+,)?` + uuidRETmpl + "," + uuidRETmpl)

const sitesWebURLGetTemplate = "https://graph.microsoft.com/v1.0/sites/%s:/%s%s"

// GetByID looks up the site matching the given identifier.  The identifier can be either a
// canonical site id or a webURL.  Assumes the webURL is complete and well formed;
// eg: https://10rqc2.sharepoint.com/sites/Example
func (c Sites) GetByID(
	ctx context.Context,
	identifier string,
	cc CallConfig,
) (models.Siteable, error) {
	var (
		resp models.Siteable
		err  error
	)

	ctx = clues.Add(ctx, "given_site_id", identifier)

	if siteIDRE.MatchString(identifier) {
		options := &sites.SiteItemRequestBuilderGetRequestConfiguration{
			QueryParameters: &sites.SiteItemRequestBuilderGetQueryParameters{},
		}

		if len(cc.Expand) > 0 {
			options.QueryParameters.Expand = cc.Expand
		}

		resp, err = c.Stable.
			Client().
			Sites().
			BySiteId(identifier).
			Get(ctx, options)
		if err != nil {
			err := graph.Wrap(ctx, err, "getting site by id")

			// a 404 when getting sites by ID returns an itemNotFound
			// error code, instead of something more sensible.
			if graph.IsErrItemNotFound(err) {
				err = clues.Stack(graph.ErrResourceOwnerNotFound, err)
			}

			return nil, err
		}

		return resp, err
	}

	// if the id is not a standard sharepoint ID, assume it's a url.
	// if it has a leading slash, assume it's only a path.  If it doesn't,
	// ensure it has a prefix https://
	if !strings.HasPrefix(identifier, "/") {
		identifier = strings.TrimPrefix(identifier, "https://")
		identifier = strings.TrimPrefix(identifier, "http://")
		identifier = "https://" + identifier
	}

	u, err := url.Parse(identifier)
	if err != nil {
		return nil, clues.Wrap(err, "site is not parseable as a url")
	}

	// don't construct a path with double leading slashes
	path := strings.TrimPrefix(u.Path, "/")

	qp := ""
	if len(cc.Expand) > 0 {
		qp = "?expand=" + strings.Join(cc.Expand, ",")
	}

	rawURL := fmt.Sprintf(sitesWebURLGetTemplate, u.Host, path, qp)

	fmt.Printf("\n-----\nrawURL %+v\n-----\n", rawURL)

	resp, err = sites.
		NewItemSitesSiteItemRequestBuilder(rawURL, c.Stable.Adapter()).
		Get(ctx, nil)
	if err != nil {
		err := graph.Wrap(ctx, err, "getting site by weburl")

		// a 404 when getting sites by ID returns an itemNotFound
		// error code, instead of something more sensible.
		if graph.IsErrItemNotFound(err) {
			err = clues.Stack(graph.ErrResourceOwnerNotFound, err)
		}

		return nil, err
	}

	return resp, err
}

// GetIDAndName looks up the site matching the given ID, and returns
// its canonical ID and the webURL as the name.  Accepts an ID or a
// WebURL as an ID.
func (c Sites) GetIDAndName(
	ctx context.Context,
	siteID string,
	cc CallConfig,
) (string, string, error) {
	s, err := c.GetByID(ctx, siteID, cc)
	if err != nil {
		return "", "", err
	}

	return ptr.Val(s.GetId()), ptr.Val(s.GetWebUrl()), nil
}

// ---------------------------------------------------------------------------
// Info
// ---------------------------------------------------------------------------

func (c Sites) GetDefaultDrive(
	ctx context.Context,
	site string,
) (models.Driveable, error) {
	d, err := c.Stable.
		Client().
		Sites().
		BySiteId(site).
		Drive().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting site's default drive")
	}

	return d, nil
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

var ErrKnownSkippableCase = clues.New("case is known and skippable")

const PersonalSitePath = "sharepoint.com/personal/"

// ValidateSite ensures the item is a Siteable, and contains the necessary
// identifiers that we handle with all users.
// returns the item as a Siteable model.
func ValidateSite(item models.Siteable) error {
	id := ptr.Val(item.GetId())
	if len(id) == 0 {
		return clues.New("missing ID")
	}

	wURL := ptr.Val(item.GetWebUrl())
	if len(wURL) == 0 {
		return clues.New("missing webURL").With("site_id", clues.Hide(id))
	}

	// personal (ie: oneDrive) sites have to be filtered out server-side.
	if strings.Contains(wURL, PersonalSitePath) {
		return clues.Stack(ErrKnownSkippableCase).
			With("site_id", clues.Hide(id), "site_web_url", clues.Hide(wURL))
	}

	name := ptr.Val(item.GetDisplayName())
	if len(name) == 0 {
		// the built-in site at "https://{tenant-domain}/search" never has a name.
		if strings.HasSuffix(wURL, "/search") {
			return clues.Stack(ErrKnownSkippableCase).
				With("site_id", clues.Hide(id), "site_web_url", clues.Hide(wURL))
		}

		return clues.New("missing site display name").With("site_id", clues.Hide(id))
	}

	return nil
}
