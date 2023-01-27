package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebPartPositionable 
type WebPartPositionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColumnId()(*float64)
    GetHorizontalSectionId()(*float64)
    GetIsInVerticalSection()(*bool)
    GetOdataType()(*string)
    GetWebPartIndex()(*float64)
    SetColumnId(value *float64)()
    SetHorizontalSectionId(value *float64)()
    SetIsInVerticalSection(value *bool)()
    SetOdataType(value *string)()
    SetWebPartIndex(value *float64)()
}
