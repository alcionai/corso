package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonWebsiteable 
type PersonWebsiteable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategories()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetThumbnailUrl()(*string)
    GetWebUrl()(*string)
    SetCategories(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetThumbnailUrl(value *string)()
    SetWebUrl(value *string)()
}
