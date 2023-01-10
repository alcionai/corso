package sharepoint

import (
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// SitePageable adjusted from msgraph-beta-sdk-go for temporary testing
type SitePageable interface {
	models.BaseItemable
	serialization.Parsable
	GetContentType() models.ContentTypeInfoable
	GetPublishingState() models.PublicationFacetable
	GetShowComments() *bool
	GetShowRecommendedPages() *bool
	GetThumbnailWebUrl() *string
	GetTitle() *string
	SetContentType(value models.ContentTypeInfoable)
	SetPublishingState(value models.PublicationFacetable)
	SetShowComments(value *bool)
	SetShowRecommendedPages(value *bool)
	SetThumbnailWebUrl(value *string)
	SetTitle(value *string)
}
