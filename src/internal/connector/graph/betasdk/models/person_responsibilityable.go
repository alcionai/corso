package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonResponsibilityable 
type PersonResponsibilityable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCollaborationTags()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetThumbnailUrl()(*string)
    GetWebUrl()(*string)
    SetCollaborationTags(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetThumbnailUrl(value *string)()
    SetWebUrl(value *string)()
}
