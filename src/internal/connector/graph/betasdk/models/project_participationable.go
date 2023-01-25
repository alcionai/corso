package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProjectParticipationable 
type ProjectParticipationable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategories()([]string)
    GetClient()(CompanyDetailable)
    GetCollaborationTags()([]string)
    GetColleagues()([]RelatedPersonable)
    GetDetail()(PositionDetailable)
    GetDisplayName()(*string)
    GetSponsors()([]RelatedPersonable)
    GetThumbnailUrl()(*string)
    SetCategories(value []string)()
    SetClient(value CompanyDetailable)()
    SetCollaborationTags(value []string)()
    SetColleagues(value []RelatedPersonable)()
    SetDetail(value PositionDetailable)()
    SetDisplayName(value *string)()
    SetSponsors(value []RelatedPersonable)()
    SetThumbnailUrl(value *string)()
}
