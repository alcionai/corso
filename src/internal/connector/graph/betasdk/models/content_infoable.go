package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentInfoable 
type ContentInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFormat()(*ContentFormat)
    GetIdentifier()(*string)
    GetMetadata()([]KeyValuePairable)
    GetOdataType()(*string)
    GetState()(*ContentState)
    SetFormat(value *ContentFormat)()
    SetIdentifier(value *string)()
    SetMetadata(value []KeyValuePairable)()
    SetOdataType(value *string)()
    SetState(value *ContentState)()
}
