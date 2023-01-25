package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AirPrintDestinationable 
type AirPrintDestinationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetForceTls()(*bool)
    GetIpAddress()(*string)
    GetOdataType()(*string)
    GetPort()(*int32)
    GetResourcePath()(*string)
    SetForceTls(value *bool)()
    SetIpAddress(value *string)()
    SetOdataType(value *string)()
    SetPort(value *int32)()
    SetResourcePath(value *string)()
}
