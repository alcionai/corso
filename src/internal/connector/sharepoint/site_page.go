package sharepoint

import (
	kioser "github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// SitePage provides operations to manage the minimal creation of a Site Page.
// Altered from original: github.com/microsoftgraph/msgraph-beta-sdk-go/models
// TODO:  remove when Issue #2086 resolved
type SitePage struct {
	models.BaseItem
	// Indicates the layout of the content in a given SharePoint page, including horizontal sections and vertical section
	// canvasLayout models.CanvasLayoutable
	// Inherited from baseItem.
	contentType models.ContentTypeInfoable
	// The name of the page layout of the page.
	// The possible values are: microsoftReserved, article, home, unknownFutureValue.
	// pageLayout *models.PageLayoutType
	// Indicates the promotion kind of the sitePage. The possible values are:
	// microsoftReserved, page, newsPost, unknownFutureValue.
	// promotionKind *models.PagePromotionType
	// The publishing status and the MM.mm version of the page.
	publishingState models.PublicationFacetable
	// Reactions information for the page.
	// reactions models.ReactionsFacetable
	// Determines whether or not to show comments at the bottom of the page.
	showComments *bool
	// Determines whether or not to show recommended pages at the bottom of the page.
	showRecommendedPages *bool
	// Url of the sitePage's thumbnail image
	//revive:disable:var-naming
	thumbnailWebUrl *string
	//revive:enable:var-naming
	// Title of the sitePage.
	title *string
}

// Title area on the SharePoint page.
// titleArea models.TitleAreaable
// Collection of webparts on the SharePoint page
// webParts []models.WebPartable

var _ SitePageable = &SitePage{}

// NewSitePage instantiates a new sitePage and sets the default values.
func NewSitePage() *SitePage {
	m := &SitePage{
		BaseItem: *models.NewBaseItem(),
	}
	odataTypeValue := "#microsoft.graph.sitePage"
	m.SetOdataType(&odataTypeValue)

	return m
}

// CreateSitePageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSitePageFromDiscriminatorValue(parseNode kioser.ParseNode) (kioser.Parsable, error) {
	return NewSitePage(), nil
}

// GetContentType gets the contentType property value. Inherited from baseItem.
func (m *SitePage) GetContentType() models.ContentTypeInfoable {
	return m.contentType
}

// GetFieldDeserializers the deserialization information for the current model
// Altered from original.
func (m *SitePage) GetFieldDeserializers() map[string]func(kioser.ParseNode) error {
	res := m.BaseItem.GetFieldDeserializers()

	return res
}

// GetPublishingState gets the publishingState property value. The publishing status and the MM.mm version of the page.
func (m *SitePage) GetPublishingState() models.PublicationFacetable {
	return m.publishingState
}

// GetShowComments gets the showComments property value.
// Determines whether or not to show comments at the bottom of the page.
func (m *SitePage) GetShowComments() *bool {
	return m.showComments
}

// GetShowRecommendedPages gets the showRecommendedPages property value.
// Determines whether or not to show recommended pages at the bottom of the page.
func (m *SitePage) GetShowRecommendedPages() *bool {
	return m.showRecommendedPages
}

// GetThumbnailWebUrl gets the thumbnailWebUrl property value. Url of the sitePage's thumbnail image
//
//revive:disable:var-naming
func (m *SitePage) GetThumbnailWebUrl() *string {
	return m.thumbnailWebUrl
}

// GetTitle gets the title property value. Title of the sitePage.
func (m *SitePage) GetTitle() *string {
	return m.title
}

// Serialize serializes information the current object
func (m *SitePage) Serialize(writer kioser.SerializationWriter) error {
	err := m.BaseItem.Serialize(writer)
	if err != nil {
		return err
	}

	if m.GetContentType() != nil {
		err = writer.WriteObjectValue("contentType", m.GetContentType())
		if err != nil {
			return err
		}
	}

	if m.GetPublishingState() != nil {
		err = writer.WriteObjectValue("publishingState", m.GetPublishingState())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteBoolValue("showComments", m.GetShowComments())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteBoolValue("showRecommendedPages", m.GetShowRecommendedPages())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteStringValue("thumbnailWebUrl", m.GetThumbnailWebUrl())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteStringValue("title", m.GetTitle())
		if err != nil {
			return err
		}
	}

	return nil
}

// SetContentType sets the contentType property value. Inherited from baseItem.
func (m *SitePage) SetContentType(value models.ContentTypeInfoable) {
	m.contentType = value
}

// SetPublishingState sets the publishingState property value. The publishing status and the MM.mm version of the page.
func (m *SitePage) SetPublishingState(value models.PublicationFacetable) {
	m.publishingState = value
}

// SetShowComments sets the showComments property value.
// Determines whether or not to show comments at the bottom of the page.
func (m *SitePage) SetShowComments(value *bool) {
	m.showComments = value
}

// SetShowRecommendedPages sets the showRecommendedPages property value.
// Determines whether or not to show recommended pages at the bottom of the page.
func (m *SitePage) SetShowRecommendedPages(value *bool) {
	m.showRecommendedPages = value
}

// SetThumbnailWebUrl sets the thumbnailWebUrl property value.
// Url of the sitePage's thumbnail image
//
//revive:disable:var-naming
func (m *SitePage) SetThumbnailWebUrl(value *string) {
	m.thumbnailWebUrl = value
}

//revive:enable:var-naming

// SetTitle sets the title property value. Title of the sitePage.
func (m *SitePage) SetTitle(value *string) {
	m.title = value
}
