package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerSystemUpdateFreezePeriodable 
type AndroidDeviceOwnerSystemUpdateFreezePeriodable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDay()(*int32)
    GetEndMonth()(*int32)
    GetOdataType()(*string)
    GetStartDay()(*int32)
    GetStartMonth()(*int32)
    SetEndDay(value *int32)()
    SetEndMonth(value *int32)()
    SetOdataType(value *string)()
    SetStartDay(value *int32)()
    SetStartMonth(value *int32)()
}
