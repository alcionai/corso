package api

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	msmodels "github.com/microsoftgraph/msgraph-sdk-go/models"

	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	discover "github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/sites"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	D "github.com/alcionai/corso/src/internal/diagnostics"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

const (
	textWebPartType     = "#microsoft.graph.textWebPart"
	standardWebPartType = "#microsoft.graph.standardWebPart"
)

// GetSitePages retrieves a collection of Pages related to the give Site.
// Returns error if error experienced during the call
func GetSitePages(
	ctx context.Context,
	serv *discover.BetaService,
	siteID string,
	pages []string,
) ([]models.SitePageable, error) {
	var (
		col         = make([]models.SitePageable, 0)
		semaphoreCh = make(chan struct{}, fetchChannelSize)
		opts        = retrieveSitePageOptions()
		err, errs   error
		wg          sync.WaitGroup
		m           sync.Mutex
	)

	defer close(semaphoreCh)

	errUpdater := func(id string, err error) {
		m.Lock()
		errs = support.WrapAndAppend(id, err, errs)
		m.Unlock()
	}
	updatePages := func(page models.SitePageable) {
		m.Lock()
		col = append(col, page)
		m.Unlock()
	}

	for _, entry := range pages {
		semaphoreCh <- struct{}{}

		wg.Add(1)

		go func(pageID string) {
			defer wg.Done()
			defer func() { <-semaphoreCh }()

			var page models.SitePageable

			err = graph.RunWithRetry(func() error {
				page, err = serv.Client().SitesById(siteID).PagesById(pageID).Get(ctx, opts)
				return err
			})
			if err != nil {
				errUpdater(pageID, errors.Wrap(err, support.ConnectorStackErrorTrace(err)+" fetching page"))
			} else {
				updatePages(page)
			}
		}(entry)
	}

	wg.Wait()

	if errs != nil {
		return nil, errs
	}

	return col, nil
}

// fetchPages utility function to return the tuple of item
func FetchPages(ctx context.Context, bs *discover.BetaService, siteID string) ([]Tuple, error) {
	var (
		builder    = bs.Client().SitesById(siteID).Pages()
		opts       = fetchPageOptions()
		pageTuples = make([]Tuple, 0)
		resp       models.SitePageCollectionResponseable
		err        error
	)

	for {
		err = graph.RunWithRetry(func() error {
			resp, err = builder.Get(ctx, opts)
			return err
		})
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
		dii    = details.ItemInfo{}
		pageID = itemData.UUID()
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

	pageName := ptr.Val(page.GetName())
	if len(pageName) == 0 {
		pageName = pageID
	}

	newName := fmt.Sprintf("%s_%s", destName, pageName)
	page.SetName(&newName)
	pg := sanitize(page, newName)

	wtr := kioser.NewJsonSerializationWriter()
	err = wtr.WriteObjectValue("", pg)
	byteArray, err = wtr.GetSerializedContent()

	if err != nil {
		fmt.Println("What happened")
	}
	fmt.Printf("Page\n %+v\n", string(byteArray))

	fmt.Printf("Layout: %+v\n", pg.GetCanvasLayout())

	// Restore is a 2-Step Process in Graph API
	// 1. Create the Page on the site
	// 2. Publish the site
	// See: https://learn.microsoft.com/en-us/graph/api/sitepage-create?view=graph-rest-beta
	restoredPage, err := service.Client().SitesById(siteID).Pages().Post(ctx, pg, nil)
	if err != nil {
		sendErr := support.ConnectorStackErrorTraceWrap(
			err,
			"creating page: "+pageName+" API Error Details",
		)

		return dii, sendErr
	}

	pageID = *restoredPage.GetId()
	// Publish page to make visible
	// See https://learn.microsoft.com/en-us/graph/api/sitepage-publish?view=graph-rest-beta
	if restoredPage.GetWebUrl() == nil {
		return dii, fmt.Errorf("creating page %s incomplete. Field  `webURL` not populated", pageID)
	}

	err = service.Client().
		SitesById(siteID).
		PagesById(pageID).
		Publish().
		Post(ctx, nil)
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

// sanitize removes all unique M365IDs from the SitePage data type.
func sanitize(orig models.SitePageable, newName string) *models.SitePage {
	newPage := models.NewSitePage()

	layout := sanitizeCanvasLayout(orig.GetCanvasLayout())
	newPage.SetCanvasLayout(layout)

	ct := sanitizeContentType(orig.GetContentType())
	newPage.SetContentType(ct)
	newPage.SetContentType(ct)
	// Skip CreatedBy.., ..ByUser, ..DateTime
	newPage.SetDescription(orig.GetDescription())
	//  skip Etag, ID, lastModified
	//  skip ID
	// skip lastModified -> it will be the app
	newPage.SetName(orig.GetName())
	newPage.SetPageLayout(orig.GetPageLayout())
	// Parent skipped
	newPage.SetPromotionKind(nil)
	// Skip publishing state. Page will attempt to be published during restore
	newPage.SetReactions(orig.GetReactions())
	newPage.SetShowComments(nil)
	newPage.SetShowRecommendedPages(nil)
	newPage.SetThumbnailWebUrl(nil)
	newPage.SetTitle(orig.GetTitle())
	// Skip TitleArea due to Upstream Failure
	// https://github.com/microsoftgraph/msgraph-metadata/issues/258
	newPage.SetTitleArea(nil)

	wp := make([]models.WebPartable, 0)
	for _, entry := range orig.GetWebParts() {
		temp := sanitizeWebPart(entry)
		wp = append(wp, temp)
	}

	newPage.SetWebParts(wp)
	// webURL intentionally left

	return newPage
}

func sanitizeContentType(orig msmodels.ContentTypeInfoable) msmodels.ContentTypeInfoable {
	if orig == nil {
		return nil
	}

	ct := msmodels.NewContentTypeInfo()
	ct.SetName(orig.GetName())
	ct.SetOdataType(orig.GetOdataType())

	return ct
}

func sanitizeCanvasLayout(orig models.CanvasLayoutable) models.CanvasLayoutable {
	canvas := models.NewCanvasLayout()
	vert := sanitizeVertical(orig.GetVerticalSection())

	canvas.SetVerticalSection(vert)
	hzLayouts := make([]models.HorizontalSectionable, 0)
	sections := orig.GetHorizontalSections()

	for _, entry := range sections {
		temp := sanitizeHorizontal(entry)

		hzLayouts = append(hzLayouts, temp)
	}

	canvas.SetHorizontalSections(hzLayouts)
	canvas.SetHorizontalSections(nil)

	return canvas
}

func sanitizeVertical(orig models.VerticalSectionable) models.VerticalSectionable {
	if orig == nil {
		return nil
	}

	section := models.NewVerticalSection()
	wps := make([]models.WebPartable, 0)

	for _, item := range orig.GetWebparts() {
		temp := sanitizeWebPart(item)
		wps = append(wps, temp)
	}

	section.SetWebparts(wps)
	section.SetEmphasis(orig.GetEmphasis())
	section.SetOdataType(orig.GetOdataType())

	return section
}

func sanitizeHorizontal(orig models.HorizontalSectionable) models.HorizontalSectionable {
	newColumns := make([]models.HorizontalSectionColumnable, 0)
	temp := models.NewHorizontalSection()
	temp.SetEmphasis(orig.GetEmphasis())
	temp.SetLayout(orig.GetLayout())

	for _, entry := range orig.GetColumns() {
		column := sanitizeColumn(entry)
		newColumns = append(newColumns, column)
	}

	temp.SetColumns(newColumns)

	return temp
}

func sanitizeColumn(orig models.HorizontalSectionColumnable) models.HorizontalSectionColumnable {
	webparts := make([]models.WebPartable, 0)
	temp := models.NewHorizontalSectionColumn()
	temp.SetWidth(orig.GetWidth())

	parts := orig.GetWebparts()
	for _, entry := range parts {
		wp := sanitizeWebPart(entry)
		webparts = append(webparts, wp)
	}

	temp.SetWebparts(webparts)

	return temp
}

func sanitizeWebPart(orig models.WebPartable) models.WebPartable {
	fmt.Println(ptr.Val(orig.GetOdataType()))

	category := ptr.Val(orig.GetOdataType())
	switch category {
	case textWebPartType:
		temp := models.NewTextWebPart()
		cast := orig.(models.TextWebPartable)
		temp.SetInnerHtml(cast.GetInnerHtml())
		temp.SetOdataType(cast.GetOdataType())

		fmt.Println("Print Text Additional")
		printAdditional(cast.GetAdditionalData())
		fmt.Printf("WP: %+v\n", cast)

		return temp

	case standardWebPartType:
		temp := models.NewStandardWebPart()
		cast := orig.(models.StandardWebPartable)
		adtl := cast.GetAdditionalData()

		fmt.Println("Print Standard Additional")
		printAdditional(adtl)
		temp.SetData(cast.GetData())
		temp.SetOdataType(cast.GetOdataType())
		temp.SetWebPartType(cast.GetWebPartType())
		fmt.Printf("TP: %+v\n", cast)

		return temp
	default:
		return nil
	}
}

func printAdditional(mapped map[string]any) {
	if mapped == nil {
		return
	}

	fmt.Printf("Length: %d\n", len(mapped))
	for key, value := range mapped {
		switch category := value.(type) {
		case int, string, bool:
			fmt.Printf("Key: %s Value: %+v", key, value)
		default:
			fmt.Printf("Key: %s Value Type: %v\n", key, category)
		}
	}
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
