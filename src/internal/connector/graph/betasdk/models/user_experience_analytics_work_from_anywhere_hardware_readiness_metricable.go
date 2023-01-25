package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable 
type UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOsCheckFailedPercentage()(*float64)
    GetProcessor64BitCheckFailedPercentage()(*float64)
    GetProcessorCoreCountCheckFailedPercentage()(*float64)
    GetProcessorFamilyCheckFailedPercentage()(*float64)
    GetProcessorSpeedCheckFailedPercentage()(*float64)
    GetRamCheckFailedPercentage()(*float64)
    GetSecureBootCheckFailedPercentage()(*float64)
    GetStorageCheckFailedPercentage()(*float64)
    GetTotalDeviceCount()(*int32)
    GetTpmCheckFailedPercentage()(*float64)
    GetUpgradeEligibleDeviceCount()(*int32)
    SetOsCheckFailedPercentage(value *float64)()
    SetProcessor64BitCheckFailedPercentage(value *float64)()
    SetProcessorCoreCountCheckFailedPercentage(value *float64)()
    SetProcessorFamilyCheckFailedPercentage(value *float64)()
    SetProcessorSpeedCheckFailedPercentage(value *float64)()
    SetRamCheckFailedPercentage(value *float64)()
    SetSecureBootCheckFailedPercentage(value *float64)()
    SetStorageCheckFailedPercentage(value *float64)()
    SetTotalDeviceCount(value *int32)()
    SetTpmCheckFailedPercentage(value *float64)()
    SetUpgradeEligibleDeviceCount(value *int32)()
}
