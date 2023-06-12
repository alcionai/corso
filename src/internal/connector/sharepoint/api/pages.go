package api

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	betamodels "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	betasites "github.com/alcionai/corso/src/internal/connector/graph/betasdk/sites"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
)

type NameID struct {
	Name string
	ID   string
}

// GetSitePages retrieves a collection of Pages related to the give Site.
// Returns error if error experienced during the call
func GetSitePages(
	ctx context.Context,
	serv *BetaService,
	siteID string,
	pages []string,
	errs *fault.Bus,
) ([]betamodels.SitePageable, error) {
	var (
		col         = make([]betamodels.SitePageable, 0)
		semaphoreCh = make(chan struct{}, 5)
		opts        = retrieveSitePageOptions()
		el          = errs.Local()
		wg          sync.WaitGroup
		m           sync.Mutex
	)

	defer close(semaphoreCh)

	updatePages := func(page betamodels.SitePageable) {
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
				page betamodels.SitePageable
				err  error
			)

			page, err = serv.Client().SitesById(siteID).PagesById(pageID).Get(ctx, opts)
			if err != nil {
				el.AddRecoverable(ctx, graph.Wrap(ctx, err, "fetching page"))
				return
			}

			updatePages(page)
		}(entry)
	}

	wg.Wait()

	return col, el.Failure()
}

// fetchPages utility function to return the tuple of item
func FetchPages(ctx context.Context, bs *BetaService, siteID string) ([]NameID, error) {
	var (
		builder = bs.Client().SitesById(siteID).Pages()
		opts    = fetchPageOptions()
		pages   = make([]NameID, 0)
		resp    betamodels.SitePageCollectionResponseable
		err     error
	)

	for {
		resp, err = builder.Get(ctx, opts)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "fetching site page")
		}

		for _, entry := range resp.GetValue() {
			var (
				pid  = ptr.Val(entry.GetId())
				temp = NameID{pid, pid}
			)

			name, ok := ptr.ValOK(entry.GetName())
			if ok {
				temp.Name = name
			}

			pages = append(pages, temp)
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = betasites.NewItemPagesRequestBuilder(link, bs.Client().Adapter())
	}

	return pages, nil
}

// fetchPageOptions is used to return minimal information reltating to Site Pages
// Pages API: https://learn.microsoft.com/en-us/graph/api/resources/sitepage?view=graph-rest-beta
func fetchPageOptions() *betasites.ItemPagesRequestBuilderGetRequestConfiguration {
	fields := []string{"id", "name"}
	options := &betasites.ItemPagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &betasites.ItemPagesRequestBuilderGetQueryParameters{
			Select: fields,
		},
	}

	return options
}

// DeleteSitePage removes the selected page from the SharePoint Site
// https://learn.microsoft.com/en-us/graph/api/sitepage-delete?view=graph-rest-beta
// deletes require unique http clients
// https://github.com/alcionai/corso/issues/2707
func DeleteSitePage(
	ctx context.Context,
	serv *BetaService,
	siteID, pageID string,
) error {
	err := serv.Client().SitesById(siteID).PagesById(pageID).Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting page")
	}

	return nil
}

// retrievePageOptions returns options to expand
func retrieveSitePageOptions() *betasites.ItemPagesSitePageItemRequestBuilderGetRequestConfiguration {
	fields := []string{"canvasLayout"}
	options := &betasites.ItemPagesSitePageItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &betasites.ItemPagesSitePageItemRequestBuilderGetQueryParameters{
			Expand: fields,
		},
	}

	return options
}

func RestoreSitePage(
	ctx context.Context,
	service *BetaService,
	itemData data.Stream,
	siteID, destName string,
) (details.ItemInfo, error) {
	ctx, end := diagnostics.Span(ctx, "gc:sharepoint:restorePage", diagnostics.Label("item_uuid", itemData.UUID()))
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
	page, err := CreatePageFromBytes(byteArray)
	if err != nil {
		return dii, clues.Wrap(err, "creating Page object").WithClues(ctx)
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
func PageInfo(page betamodels.SitePageable, size int64) *details.SharePointInfo {
	var (
		name     = ptr.Val(page.GetTitle())
		webURL   = ptr.Val(page.GetWebUrl())
		created  = ptr.Val(page.GetCreatedDateTime())
		modified = ptr.Val(page.GetLastModifiedDateTime())
	)

	return &details.SharePointInfo{
		ItemType: details.SharePointPage,
		ItemName: name,
		Created:  created,
		Modified: modified,
		WebURL:   webURL,
		Size:     size,
	}
}
