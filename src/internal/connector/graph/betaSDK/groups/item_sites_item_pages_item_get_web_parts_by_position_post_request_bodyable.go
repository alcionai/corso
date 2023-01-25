package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBodyable 
type ItemSitesItemPagesItemGetWebPartsByPositionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    IBackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(BackingStore)
    GetColumnId()(*float64)
    GetHorizontalSectionId()(*float64)
    GetIsInVerticalSection()(*bool)
    GetWebPartIndex()(*float64)
    SetBackingStore(value BackingStore)()
    SetColumnId(value *float64)()
    SetHorizontalSectionId(value *float64)()
    SetIsInVerticalSection(value *bool)()
    SetWebPartIndex(value *float64)()
}
