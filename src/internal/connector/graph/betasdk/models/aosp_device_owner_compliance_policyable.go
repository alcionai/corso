package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AospDeviceOwnerCompliancePolicyable 
type AospDeviceOwnerCompliancePolicyable interface {
    DeviceCompliancePolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMinAndroidSecurityPatchLevel()(*string)
    GetOsMaximumVersion()(*string)
    GetOsMinimumVersion()(*string)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeLock()(*int32)
    GetPasswordRequired()(*bool)
    GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType)
    GetSecurityBlockJailbrokenDevices()(*bool)
    GetStorageRequireEncryption()(*bool)
    SetMinAndroidSecurityPatchLevel(value *string)()
    SetOsMaximumVersion(value *string)()
    SetOsMinimumVersion(value *string)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeLock(value *int32)()
    SetPasswordRequired(value *bool)()
    SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)()
    SetSecurityBlockJailbrokenDevices(value *bool)()
    SetStorageRequireEncryption(value *bool)()
}
