package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/alcionai/clues"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
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

	var resp models.SiteCollectionResponseable

	resp, err = service.Client().Sites().Get(ctx, nil)

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
		return nil, graph.Wrap(ctx, err, "iterating all sites")
	}

	return us, el.Failure()
}

func (c Sites) GetByID(ctx context.Context, id string) (models.Siteable, error) {
	resp, err := c.stable.Client().SitesById(id).Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting site")
	}

	return resp, err
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

	id, ok := ptr.ValOK(m.GetId())
	if !ok || len(id) == 0 {
		return nil, clues.New("missing ID")
	}

	url, ok := ptr.ValOK(m.GetWebUrl())
	if !ok || len(url) == 0 {
		return nil, clues.New("missing webURL").With("site_id", id) // TODO: pii
	}

	// personal (ie: oneDrive) sites have to be filtered out server-side.
	if ok && strings.Contains(url, personalSitePath) {
		return nil, clues.Stack(errKnownSkippableCase).
			With("site_id", id, "site_url", url) // TODO: pii
	}

	if name, ok := ptr.ValOK(m.GetDisplayName()); !ok || len(name) == 0 {
		// the built-in site at "https://{tenant-domain}/search" never has a name.
		if strings.HasSuffix(url, "/search") {
			return nil, clues.Stack(errKnownSkippableCase).
				With("site_id", id, "site_url", url) // TODO: pii
		}

		return nil, clues.New("missing site display name").With("site_id", id)
	}

	return m, nil
}
