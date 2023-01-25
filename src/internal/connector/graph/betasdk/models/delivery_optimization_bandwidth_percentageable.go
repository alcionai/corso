package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationBandwidthPercentageable 
type DeliveryOptimizationBandwidthPercentageable interface {
    DeliveryOptimizationBandwidthable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumBackgroundBandwidthPercentage()(*int32)
    GetMaximumForegroundBandwidthPercentage()(*int32)
    SetMaximumBackgroundBandwidthPercentage(value *int32)()
    SetMaximumForegroundBandwidthPercentage(value *int32)()
}
