package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SitePageable 
type SitePageable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCanvasLayout()(CanvasLayoutable)
    GetContentType()(ContentTypeInfoable)
    GetPageLayout()(*PageLayoutType)
    GetPromotionKind()(*PagePromotionType)
    GetPublishingState()(PublicationFacetable)
    GetReactions()(ReactionsFacetable)
    GetShowComments()(*bool)
    GetShowRecommendedPages()(*bool)
    GetThumbnailWebUrl()(*string)
    GetTitle()(*string)
    GetTitleArea()(TitleAreaable)
    GetWebParts()([]WebPartable)
    SetCanvasLayout(value CanvasLayoutable)()
    SetContentType(value ContentTypeInfoable)()
    SetPageLayout(value *PageLayoutType)()
    SetPromotionKind(value *PagePromotionType)()
    SetPublishingState(value PublicationFacetable)()
    SetReactions(value ReactionsFacetable)()
    SetShowComments(value *bool)()
    SetShowRecommendedPages(value *bool)()
    SetThumbnailWebUrl(value *string)()
    SetTitle(value *string)()
    SetTitleArea(value TitleAreaable)()
    SetWebParts(value []WebPartable)()
}
