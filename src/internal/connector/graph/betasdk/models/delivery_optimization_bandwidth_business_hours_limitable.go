package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationBandwidthBusinessHoursLimitable 
type DeliveryOptimizationBandwidthBusinessHoursLimitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBandwidthBeginBusinessHours()(*int32)
    GetBandwidthEndBusinessHours()(*int32)
    GetBandwidthPercentageDuringBusinessHours()(*int32)
    GetBandwidthPercentageOutsideBusinessHours()(*int32)
    GetOdataType()(*string)
    SetBandwidthBeginBusinessHours(value *int32)()
    SetBandwidthEndBusinessHours(value *int32)()
    SetBandwidthPercentageDuringBusinessHours(value *int32)()
    SetBandwidthPercentageOutsideBusinessHours(value *int32)()
    SetOdataType(value *string)()
}
