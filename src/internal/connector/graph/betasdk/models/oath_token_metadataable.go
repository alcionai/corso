package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OathTokenMetadataable 
type OathTokenMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetManufacturer()(*string)
    GetManufacturerProperties()([]KeyValueable)
    GetOdataType()(*string)
    GetSerialNumber()(*string)
    GetTokenType()(*string)
    SetEnabled(value *bool)()
    SetManufacturer(value *string)()
    SetManufacturerProperties(value []KeyValueable)()
    SetOdataType(value *string)()
    SetSerialNumber(value *string)()
    SetTokenType(value *string)()
}
