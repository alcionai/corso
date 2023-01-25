package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompanyDetailable 
type CompanyDetailable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PhysicalAddressable)
    GetDepartment()(*string)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    GetOfficeLocation()(*string)
    GetPronunciation()(*string)
    GetWebUrl()(*string)
    SetAddress(value PhysicalAddressable)()
    SetDepartment(value *string)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
    SetOfficeLocation(value *string)()
    SetPronunciation(value *string)()
    SetWebUrl(value *string)()
}
