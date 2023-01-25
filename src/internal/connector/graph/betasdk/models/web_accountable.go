package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebAccountable 
type WebAccountable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetService()(ServiceInformationable)
    GetStatusMessage()(*string)
    GetThumbnailUrl()(*string)
    GetUserId()(*string)
    GetWebUrl()(*string)
    SetDescription(value *string)()
    SetService(value ServiceInformationable)()
    SetStatusMessage(value *string)()
    SetThumbnailUrl(value *string)()
    SetUserId(value *string)()
    SetWebUrl(value *string)()
}
