package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RestrictedAppsViolationable 
type RestrictedAppsViolationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceConfigurationId()(*string)
    GetDeviceConfigurationName()(*string)
    GetDeviceName()(*string)
    GetManagedDeviceId()(*string)
    GetPlatformType()(*PolicyPlatformType)
    GetRestrictedApps()([]ManagedDeviceReportedAppable)
    GetRestrictedAppsState()(*RestrictedAppsState)
    GetUserId()(*string)
    GetUserName()(*string)
    SetDeviceConfigurationId(value *string)()
    SetDeviceConfigurationName(value *string)()
    SetDeviceName(value *string)()
    SetManagedDeviceId(value *string)()
    SetPlatformType(value *PolicyPlatformType)()
    SetRestrictedApps(value []ManagedDeviceReportedAppable)()
    SetRestrictedAppsState(value *RestrictedAppsState)()
    SetUserId(value *string)()
    SetUserName(value *string)()
}
