package api

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"

	discover "github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/sites"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// GetSitePages retrieves a collection of Pages related to the give Site.
// Returns error if error experienced during the call
func GetSitePage(
	ctx context.Context,
	serv *discover.BetaService,
	siteID string,
	pages []string,
) ([]models.SitePageable, error) {
	col := make([]models.SitePageable, 0)
	opts := retrieveSitePageOptions()

	for _, entry := range pages {
		page, err := serv.Client().SitesById(siteID).PagesById(entry).Get(ctx, opts)
		if err != nil {
			return nil, support.ConnectorStackErrorTraceWrap(err, "fetching page: "+entry)
		}

		col = append(col, page)
	}

	return col, nil
}

// fetchPages utility function to return the tuple of item
func FetchPages(ctx context.Context, bs *discover.BetaService, siteID string) ([]Tuple, error) {
	var (
		builder    = bs.Client().SitesById(siteID).Pages()
		opts       = fetchPageOptions()
		pageTuples = make([]Tuple, 0)
	)

	for {
		resp, err := builder.Get(ctx, opts)
		if err != nil {
			return nil, support.ConnectorStackErrorTraceWrap(err, "failed fetching site page")
		}

		for _, entry := range resp.GetValue() {
			pid := *entry.GetId()
			temp := Tuple{pid, pid}

			if entry.GetName() != nil {
				temp.Name = *entry.GetName()
			}

			pageTuples = append(pageTuples, temp)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = sites.NewItemPagesRequestBuilder(*resp.GetOdataNextLink(), bs.Client().Adapter())
	}

	return pageTuples, nil
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
		return support.ConnectorStackErrorTraceWrap(err, "deleting page: "+pageID)
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

	byteArray, err := io.ReadAll(itemData.ToReader())
	if err != nil {
		return dii, errors.Wrap(err, "reading sharepoint page bytes from stream")
	}

	// Hydrate Page
	page, err := support.CreatePageFromBytes(byteArray)
	if err != nil {
		return dii, errors.Wrapf(err, "creating Page object %s", pageID)
	}

	pageNamePtr := page.GetName()
	if pageNamePtr != nil {
		pageName = *pageNamePtr
	}

	newName := fmt.Sprintf("%s_%s", destName, pageName)
	page.SetName(&newName)

	// Restore is a 2-Step Process in Graph API
	// 1. Create the Page on the site
	// 2. Publish the site
	// See: https://learn.microsoft.com/en-us/graph/api/sitepage-create?view=graph-rest-beta
	restoredPage, err := service.Client().SitesById(siteID).Pages().Post(ctx, page, nil)
	if err != nil {
		sendErr := support.ConnectorStackErrorTraceWrap(
			err,
			"creating page from ID: %s"+pageName+" API Error Details",
		)

		return dii, sendErr
	}

	// Publish page to make visible
	// See https://learn.microsoft.com/en-us/graph/api/sitepage-publish?view=graph-rest-beta
	if restoredPage.GetWebUrl() == nil {
		return dii, fmt.Errorf("creating page %s incomplete. Field  `webURL` not populated", *restoredPage.GetId())
	}

	err = service.Client().
		SitesById(siteID).
		PagesById(*restoredPage.GetId()).Publish().Post(ctx, nil)
	if err != nil {
		return dii, support.ConnectorStackErrorTraceWrap(
			err,
			"publishing page ID: "+*restoredPage.GetId()+" API Error Details",
		)
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
		name, webURL      string
		created, modified time.Time
	)

	if page.GetTitle() != nil {
		name = *page.GetTitle()
	}

	if page.GetWebUrl() != nil {
		webURL = *page.GetWebUrl()
	}

	if page.GetCreatedDateTime() != nil {
		created = *page.GetCreatedDateTime()
	}

	if page.GetLastModifiedDateTime() != nil {
		modified = *page.GetLastModifiedDateTime()
	}

	return &details.SharePointInfo{
		ItemType: details.SharePointItem,
		ItemName: name,
		Created:  created,
		Modified: modified,
		WebURL:   webURL,
		Size:     size,
	}
}
