package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SitePage provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type SitePage struct {
    BaseItem
    // Indicates the layout of the content in a given SharePoint page, including horizontal sections and vertical section
    canvasLayout CanvasLayoutable
    // Inherited from baseItem.
    contentType ContentTypeInfoable
    // The name of the page layout of the page. The possible values are: microsoftReserved, article, home, unknownFutureValue.
    pageLayout *PageLayoutType
    // Indicates the promotion kind of the sitePage. The possible values are: microsoftReserved, page, newsPost, unknownFutureValue.
    promotionKind *PagePromotionType
    // The publishing status and the MM.mm version of the page.
    publishingState PublicationFacetable
    // Reactions information for the page.
    reactions ReactionsFacetable
    // Determines whether or not to show comments at the bottom of the page.
    showComments *bool
    // Determines whether or not to show recommended pages at the bottom of the page.
    showRecommendedPages *bool
    // Url of the sitePage's thumbnail image
    thumbnailWebUrl *string
    // Title of the sitePage.
    title *string
    // Title area on the SharePoint page.
    titleArea TitleAreaable
    // Collection of webparts on the SharePoint page
    webParts []WebPartable
}
// NewSitePage instantiates a new sitePage and sets the default values.
func NewSitePage()(*SitePage) {
    m := &SitePage{
        BaseItem: *NewBaseItem(),
    }
    odataTypeValue := "#microsoft.graph.sitePage";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSitePageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSitePageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSitePage(), nil
}
// GetCanvasLayout gets the canvasLayout property value. Indicates the layout of the content in a given SharePoint page, including horizontal sections and vertical section
func (m *SitePage) GetCanvasLayout()(CanvasLayoutable) {
    return m.canvasLayout
}
// GetContentType gets the contentType property value. Inherited from baseItem.
func (m *SitePage) GetContentType()(ContentTypeInfoable) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SitePage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["canvasLayout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCanvasLayoutFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanvasLayout(val.(CanvasLayoutable))
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateContentTypeInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(ContentTypeInfoable))
        }
        return nil
    }
    res["pageLayout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePageLayoutType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPageLayout(val.(*PageLayoutType))
        }
        return nil
    }
    res["promotionKind"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePagePromotionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPromotionKind(val.(*PagePromotionType))
        }
        return nil
    }
    res["publishingState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicationFacetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingState(val.(PublicationFacetable))
        }
        return nil
    }
    res["reactions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateReactionsFacetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReactions(val.(ReactionsFacetable))
        }
        return nil
    }
    res["showComments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowComments(val)
        }
        return nil
    }
    res["showRecommendedPages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowRecommendedPages(val)
        }
        return nil
    }
    res["thumbnailWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailWebUrl(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["titleArea"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTitleAreaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitleArea(val.(TitleAreaable))
        }
        return nil
    }
    res["webParts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebPartFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebPartable, len(val))
            for i, v := range val {
                res[i] = v.(WebPartable)
            }
            m.SetWebParts(res)
        }
        return nil
    }
    return res
}
// GetPageLayout gets the pageLayout property value. The name of the page layout of the page. The possible values are: microsoftReserved, article, home, unknownFutureValue.
func (m *SitePage) GetPageLayout()(*PageLayoutType) {
    return m.pageLayout
}
// GetPromotionKind gets the promotionKind property value. Indicates the promotion kind of the sitePage. The possible values are: microsoftReserved, page, newsPost, unknownFutureValue.
func (m *SitePage) GetPromotionKind()(*PagePromotionType) {
    return m.promotionKind
}
// GetPublishingState gets the publishingState property value. The publishing status and the MM.mm version of the page.
func (m *SitePage) GetPublishingState()(PublicationFacetable) {
    return m.publishingState
}
// GetReactions gets the reactions property value. Reactions information for the page.
func (m *SitePage) GetReactions()(ReactionsFacetable) {
    return m.reactions
}
// GetShowComments gets the showComments property value. Determines whether or not to show comments at the bottom of the page.
func (m *SitePage) GetShowComments()(*bool) {
    return m.showComments
}
// GetShowRecommendedPages gets the showRecommendedPages property value. Determines whether or not to show recommended pages at the bottom of the page.
func (m *SitePage) GetShowRecommendedPages()(*bool) {
    return m.showRecommendedPages
}
// GetThumbnailWebUrl gets the thumbnailWebUrl property value. Url of the sitePage's thumbnail image
func (m *SitePage) GetThumbnailWebUrl()(*string) {
    return m.thumbnailWebUrl
}
// GetTitle gets the title property value. Title of the sitePage.
func (m *SitePage) GetTitle()(*string) {
    return m.title
}
// GetTitleArea gets the titleArea property value. Title area on the SharePoint page.
func (m *SitePage) GetTitleArea()(TitleAreaable) {
    return m.titleArea
}
// GetWebParts gets the webParts property value. Collection of webparts on the SharePoint page
func (m *SitePage) GetWebParts()([]WebPartable) {
    return m.webParts
}
// Serialize serializes information the current object
func (m *SitePage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("canvasLayout", m.GetCanvasLayout())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    if m.GetPageLayout() != nil {
        cast := (*m.GetPageLayout()).String()
        err = writer.WriteStringValue("pageLayout", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPromotionKind() != nil {
        cast := (*m.GetPromotionKind()).String()
        err = writer.WriteStringValue("promotionKind", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publishingState", m.GetPublishingState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reactions", m.GetReactions())
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
    {
        err = writer.WriteObjectValue("titleArea", m.GetTitleArea())
        if err != nil {
            return err
        }
    }
    if m.GetWebParts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebParts()))
        for i, v := range m.GetWebParts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("webParts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCanvasLayout sets the canvasLayout property value. Indicates the layout of the content in a given SharePoint page, including horizontal sections and vertical section
func (m *SitePage) SetCanvasLayout(value CanvasLayoutable)() {
    m.canvasLayout = value
}
// SetContentType sets the contentType property value. Inherited from baseItem.
func (m *SitePage) SetContentType(value ContentTypeInfoable)() {
    m.contentType = value
}
// SetPageLayout sets the pageLayout property value. The name of the page layout of the page. The possible values are: microsoftReserved, article, home, unknownFutureValue.
func (m *SitePage) SetPageLayout(value *PageLayoutType)() {
    m.pageLayout = value
}
// SetPromotionKind sets the promotionKind property value. Indicates the promotion kind of the sitePage. The possible values are: microsoftReserved, page, newsPost, unknownFutureValue.
func (m *SitePage) SetPromotionKind(value *PagePromotionType)() {
    m.promotionKind = value
}
// SetPublishingState sets the publishingState property value. The publishing status and the MM.mm version of the page.
func (m *SitePage) SetPublishingState(value PublicationFacetable)() {
    m.publishingState = value
}
// SetReactions sets the reactions property value. Reactions information for the page.
func (m *SitePage) SetReactions(value ReactionsFacetable)() {
    m.reactions = value
}
// SetShowComments sets the showComments property value. Determines whether or not to show comments at the bottom of the page.
func (m *SitePage) SetShowComments(value *bool)() {
    m.showComments = value
}
// SetShowRecommendedPages sets the showRecommendedPages property value. Determines whether or not to show recommended pages at the bottom of the page.
func (m *SitePage) SetShowRecommendedPages(value *bool)() {
    m.showRecommendedPages = value
}
// SetThumbnailWebUrl sets the thumbnailWebUrl property value. Url of the sitePage's thumbnail image
func (m *SitePage) SetThumbnailWebUrl(value *string)() {
    m.thumbnailWebUrl = value
}
// SetTitle sets the title property value. Title of the sitePage.
func (m *SitePage) SetTitle(value *string)() {
    m.title = value
}
// SetTitleArea sets the titleArea property value. Title area on the SharePoint page.
func (m *SitePage) SetTitleArea(value TitleAreaable)() {
    m.titleArea = value
}
// SetWebParts sets the webParts property value. Collection of webparts on the SharePoint page
func (m *SitePage) SetWebParts(value []WebPartable)() {
    m.webParts = value
}
