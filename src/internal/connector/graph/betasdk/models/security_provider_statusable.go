package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityProviderStatusable 
type SecurityProviderStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetEndpoint()(*string)
    GetOdataType()(*string)
    GetProvider()(*string)
    GetRegion()(*string)
    GetVendor()(*string)
    SetEnabled(value *bool)()
    SetEndpoint(value *string)()
    SetOdataType(value *string)()
    SetProvider(value *string)()
    SetRegion(value *string)()
    SetVendor(value *string)()
}
