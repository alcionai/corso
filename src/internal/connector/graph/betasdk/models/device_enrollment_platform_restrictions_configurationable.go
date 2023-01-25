package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentPlatformRestrictionsConfigurationable 
type DeviceEnrollmentPlatformRestrictionsConfigurationable interface {
    DeviceEnrollmentConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAndroidForWorkRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetAndroidRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetIosRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetMacOSRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetMacRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetWindowsHomeSkuRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetWindowsMobileRestriction()(DeviceEnrollmentPlatformRestrictionable)
    GetWindowsRestriction()(DeviceEnrollmentPlatformRestrictionable)
    SetAndroidForWorkRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetAndroidRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetIosRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetMacOSRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetMacRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetWindowsHomeSkuRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetWindowsMobileRestriction(value DeviceEnrollmentPlatformRestrictionable)()
    SetWindowsRestriction(value DeviceEnrollmentPlatformRestrictionable)()
}
