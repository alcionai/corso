package api

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/alcionai/clues"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/sites"
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
// methods
// ---------------------------------------------------------------------------

// GetAll retrieves all sites.
func (c Sites) GetAll(ctx context.Context, errs *fault.Bus) ([]models.Siteable, error) {
	service, err := c.service()
	if err != nil {
		return nil, err
	}

	resp, err := service.Client().Sites().Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting all sites")
	}

	iter, err := msgraphgocore.NewPageIterator(
		resp,
		service.Adapter(),
		models.CreateSiteCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating sites iterator")
	}

	var (
		us = make([]models.Siteable, 0)
		el = errs.Local()
	)

	iterator := func(item any) bool {
		if el.Failure() != nil {
			return false
		}

		s, err := validateSite(item)
		if errors.Is(err, errKnownSkippableCase) {
			// safe to no-op
			return true
		}

		if err != nil {
			el.AddRecoverable(graph.Wrap(ctx, err, "validating site"))
			return true
		}

		us = append(us, s)

		return true
	}

	if err := iter.Iterate(ctx, iterator); err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating sites")
	}

	return us, el.Failure()
}

const uuidRE = "[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}"

// matches a site ID, with or without a doman name.  Ex, either one of:
// 10rqc2.sharepoint.com,deadbeef-0000-0000-0000-000000000000,beefdead-0000-0000-0000-000000000000
// deadbeef-0000-0000-0000-000000000000,beefdead-0000-0000-0000-000000000000
var siteIDRE = regexp.MustCompile("(.+,)?" + uuidRE + "," + uuidRE)

const webURLGetTemplate = "https://graph.microsoft.com/v1.0/sites/%s:/%s"

// GetByID looks up the site matching the given ID.  The ID can be either a
// canonical site id or a webURL.  Assumes the webURL is complete and well formed;
// eg: https://10rqc2.sharepoint.com/sites/Example
func (c Sites) GetByID(ctx context.Context, id string) (models.Siteable, error) {
	var (
		resp models.Siteable
		err  error
	)

	ctx = clues.Add(ctx, "given_site_id", id)

	if siteIDRE.MatchString(id) {
		resp, err = c.stable.Client().SitesById(id).Get(ctx, nil)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "getting site by id")
		}
	} else {
		var (
			url   = strings.TrimPrefix(id, "https://")
			parts = strings.SplitN(url, "/", 1)
			host  = parts[0]
			path  string
		)

		if len(parts) > 1 {
			path = parts[1]
		}

		rawURL := fmt.Sprintf(webURLGetTemplate, host, path)
		resp, err = sites.
			NewItemSitesSiteItemRequestBuilder(rawURL, c.stable.Adapter()).
			Get(ctx, nil)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "getting site by weburl")
		}
	}

	return resp, err
}

// GetIDAndName looks up the site matching the given ID, and returns
// its canonical ID and the webURL as the name.  Accepts an ID or a
// WebURL as an ID.
func (c Sites) GetIDAndName(ctx context.Context, siteID string) (string, string, error) {
	s, err := c.GetByID(ctx, siteID)
	if err != nil {
		return "", "", err
	}

	return ptr.Val(s.GetId()), ptr.Val(s.GetWebUrl()), nil
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

var errKnownSkippableCase = clues.New("case is known and skippable")

const personalSitePath = "sharepoint.com/personal/"

// validateSite ensures the item is a Siteable, and contains the necessary
// identifiers that we handle with all users.
// returns the item as a Siteable model.
func validateSite(item any) (models.Siteable, error) {
	m, ok := item.(models.Siteable)
	if !ok {
		return nil, clues.New(fmt.Sprintf("unexpected model: %T", item))
	}

	id := ptr.Val(m.GetId())
	if len(id) == 0 {
		return nil, clues.New("missing ID")
	}

	url := ptr.Val(m.GetWebUrl())
	if len(url) == 0 {
		return nil, clues.New("missing webURL").With("site_id", id) // TODO: pii
	}

	// personal (ie: oneDrive) sites have to be filtered out server-side.
	if strings.Contains(url, personalSitePath) {
		return nil, clues.Stack(errKnownSkippableCase).
			With("site_id", id, "site_url", url) // TODO: pii
	}

	name := ptr.Val(m.GetDisplayName())
	if len(name) == 0 {
		// the built-in site at "https://{tenant-domain}/search" never has a name.
		if strings.HasSuffix(url, "/search") {
			return nil, clues.Stack(errKnownSkippableCase).
				With("site_id", id, "site_url", url) // TODO: pii
		}

		return nil, clues.New("missing site display name").With("site_id", id)
	}

	return m, nil
}
