package api

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/pkg/errors"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	discover "github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/sites"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
)

// GetSitePages retrieves a collection of Pages related to the give Site.
// Returns error if error experienced during the call
func GetSitePages(
	ctx context.Context,
	serv *discover.BetaService,
	siteID string,
	pages []string,
	errs *fault.Bus,
) ([]models.SitePageable, error) {
	var (
		col         = make([]models.SitePageable, 0)
		semaphoreCh = make(chan struct{}, fetchChannelSize)
		opts        = retrieveSitePageOptions()
		el          = errs.Local()
		wg          sync.WaitGroup
		m           sync.Mutex
	)

	defer close(semaphoreCh)

	updatePages := func(page models.SitePageable) {
		m.Lock()
		defer m.Unlock()

		col = append(col, page)
	}

	for _, entry := range pages {
		if el.Failure() != nil {
			break
		}

		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(pageID string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			var (
				page models.SitePageable
				err  error
			)

			page, err = serv.Client().SitesById(siteID).PagesById(pageID).Get(ctx, opts)
			if err != nil {
				el.AddRecoverable(graph.Wrap(ctx, err, "fetching page"))
				return
			}

			updatePages(page)
		}(entry)
	}

	wg.Wait()

	return col, el.Failure()
}

// fetchPages utility function to return the tuple of item
func FetchPages(ctx context.Context, bs *discover.BetaService, siteID string) ([]NameID, error) {
	var (
		builder = bs.Client().SitesById(siteID).Pages()
		opts    = fetchPageOptions()
		pages   = make([]NameID, 0)
		resp    models.SitePageCollectionResponseable
		err     error
	)

	for {
		resp, err = builder.Get(ctx, opts)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "fetching site page")
		}

		for _, entry := range resp.GetValue() {
			var (
				pid  = *entry.GetId()
				temp = NameID{pid, pid}
			)

			name, ok := ptr.ValOK(entry.GetName())
			if ok {
				temp.Name = name
			}

			pages = append(pages, temp)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = sites.NewItemPagesRequestBuilder(*resp.GetOdataNextLink(), bs.Client().Adapter())
	}

	return pages, nil
}

// fetchPageOptions is used to return minimal information reltating to Site Pages
// Pages API: https://learn.microsoft.com/en-us/graph/api/resources/sitepage?view=graph-rest-beta
func fetchPageOptions() *sites.ItemPagesRequestBuilderGetRequestConfiguration {
	fields := []string{"id", "name"}
	options := &sites.ItemPagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemPagesRequestBuilderGetQueryParameters{
			Select: fields,
		},
	}

	return options
}

// DeleteSitePage removes the selected page from the SharePoint Site
// https://learn.microsoft.com/en-us/graph/api/sitepage-delete?view=graph-rest-beta
func DeleteSitePage(
	ctx context.Context,
	serv *discover.BetaService,
	siteID, pageID string,
) error {
	err := serv.Client().SitesById(siteID).PagesById(pageID).Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting page")
	}

	return nil
}

// retrievePageOptions returns options to expand
func retrieveSitePageOptions() *sites.ItemPagesSitePageItemRequestBuilderGetRequestConfiguration {
	fields := []string{"canvasLayout"}
	options := &sites.ItemPagesSitePageItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemPagesSitePageItemRequestBuilderGetQueryParameters{
			Expand: fields,
		},
	}

	return options
}

func RestoreSitePage(
	ctx context.Context,
	service *discover.BetaService,
	itemData data.Stream,
	siteID, destName string,
) (details.ItemInfo, error) {
	ctx, end := D.Span(ctx, "gc:sharepoint:restorePage", D.Label("item_uuid", itemData.UUID()))
	defer end()

	var (
		dii      = details.ItemInfo{}
		pageID   = itemData.UUID()
		pageName = pageID
	)

	ctx = clues.Add(ctx, "page_id", pageID)

	byteArray, err := io.ReadAll(itemData.ToReader())
	if err != nil {
		return dii, clues.Wrap(err, "reading sharepoint data").WithClues(ctx)
	}

	// Hydrate Page
	page, err := support.CreatePageFromBytes(byteArray)
	if err != nil {
		return dii, errors.Wrapf(err, "creating Page object %s", pageID)
	}

	name, ok := ptr.ValOK(page.GetName())
	if ok {
		pageName = name
	}

	newName := fmt.Sprintf("%s_%s", destName, pageName)
	page.SetName(&newName)

	// Restore is a 2-Step Process in Graph API
	// 1. Create the Page on the site
	// 2. Publish the site
	// See: https://learn.microsoft.com/en-us/graph/api/sitepage-create?view=graph-rest-beta
	restoredPage, err := service.Client().SitesById(siteID).Pages().Post(ctx, page, nil)
	if err != nil {
		return dii, graph.Wrap(ctx, err, "creating page")
	}

	pageID = ptr.Val(restoredPage.GetId())
	ctx = clues.Add(ctx, "restored_page_id", pageID)

	// Publish page to make visible
	// See https://learn.microsoft.com/en-us/graph/api/sitepage-publish?view=graph-rest-beta
	if restoredPage.GetWebUrl() == nil {
		return dii, clues.New("webURL not populated during page creation").WithClues(ctx)
	}

	err = service.Client().
		SitesById(siteID).
		PagesById(pageID).
		Publish().
		Post(ctx, nil)
	if err != nil {
		return dii, graph.Wrap(ctx, err, "publishing page")
	}

	dii.SharePoint = PageInfo(restoredPage, int64(len(byteArray)))
	// Storing new pageID in unused field.
	dii.SharePoint.ParentPath = pageID

	return dii, nil
}

// ==============================
// Helpers
// ==============================
// PageInfo extracts useful metadata into struct for book keeping
func PageInfo(page models.SitePageable, size int64) *details.SharePointInfo {
	var (
		name     = ptr.Val(page.GetTitle())
		webURL   = ptr.Val(page.GetWebUrl())
		created  = ptr.Val(page.GetCreatedDateTime())
		modified = ptr.Val(page.GetLastModifiedDateTime())
	)

	return &details.SharePointInfo{
		ItemType: details.SharePointItem,
		ItemName: name,
		Created:  created,
		Modified: modified,
		WebURL:   webURL,
		Size:     size,
	}
}
