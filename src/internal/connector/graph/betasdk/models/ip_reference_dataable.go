package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpReferenceDataable 
type IpReferenceDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAsn()(*int64)
    GetCity()(*string)
    GetCountryOrRegionCode()(*string)
    GetOdataType()(*string)
    GetOrganization()(*string)
    GetState()(*string)
    GetVendor()(*string)
    SetAsn(value *int64)()
    SetCity(value *string)()
    SetCountryOrRegionCode(value *string)()
    SetOdataType(value *string)()
    SetOrganization(value *string)()
    SetState(value *string)()
    SetVendor(value *string)()
}
