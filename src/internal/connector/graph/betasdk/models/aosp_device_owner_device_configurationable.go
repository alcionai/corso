package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AospDeviceOwnerDeviceConfigurationable 
type AospDeviceOwnerDeviceConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppsBlockInstallFromUnknownSources()(*bool)
    GetBluetoothBlockConfiguration()(*bool)
    GetBluetoothBlocked()(*bool)
    GetCameraBlocked()(*bool)
    GetFactoryResetBlocked()(*bool)
    GetPasswordMinimumLength()(*int32)
    GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32)
    GetPasswordRequiredType()(*AndroidDeviceOwnerRequiredPasswordType)
    GetPasswordSignInFailureCountBeforeFactoryReset()(*int32)
    GetScreenCaptureBlocked()(*bool)
    GetSecurityAllowDebuggingFeatures()(*bool)
    GetStorageBlockExternalMedia()(*bool)
    GetStorageBlockUsbFileTransfer()(*bool)
    GetWifiBlockEditConfigurations()(*bool)
    SetAppsBlockInstallFromUnknownSources(value *bool)()
    SetBluetoothBlockConfiguration(value *bool)()
    SetBluetoothBlocked(value *bool)()
    SetCameraBlocked(value *bool)()
    SetFactoryResetBlocked(value *bool)()
    SetPasswordMinimumLength(value *int32)()
    SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)()
    SetPasswordRequiredType(value *AndroidDeviceOwnerRequiredPasswordType)()
    SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)()
    SetScreenCaptureBlocked(value *bool)()
    SetSecurityAllowDebuggingFeatures(value *bool)()
    SetStorageBlockExternalMedia(value *bool)()
    SetStorageBlockUsbFileTransfer(value *bool)()
    SetWifiBlockEditConfigurations(value *bool)()
}
