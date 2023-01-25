package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonAnnotationable 
type PersonAnnotationable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetail()(ItemBodyable)
    GetDisplayName()(*string)
    GetThumbnailUrl()(*string)
    SetDetail(value ItemBodyable)()
    SetDisplayName(value *string)()
    SetThumbnailUrl(value *string)()
}
