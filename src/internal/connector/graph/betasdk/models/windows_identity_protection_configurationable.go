package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsIdentityProtectionConfigurationable 
type WindowsIdentityProtectionConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnhancedAntiSpoofingForFacialFeaturesEnabled()(*bool)
    GetPinExpirationInDays()(*int32)
    GetPinLowercaseCharactersUsage()(*ConfigurationUsage)
    GetPinMaximumLength()(*int32)
    GetPinMinimumLength()(*int32)
    GetPinPreviousBlockCount()(*int32)
    GetPinRecoveryEnabled()(*bool)
    GetPinSpecialCharactersUsage()(*ConfigurationUsage)
    GetPinUppercaseCharactersUsage()(*ConfigurationUsage)
    GetSecurityDeviceRequired()(*bool)
    GetUnlockWithBiometricsEnabled()(*bool)
    GetUseCertificatesForOnPremisesAuthEnabled()(*bool)
    GetUseSecurityKeyForSignin()(*bool)
    GetWindowsHelloForBusinessBlocked()(*bool)
    SetEnhancedAntiSpoofingForFacialFeaturesEnabled(value *bool)()
    SetPinExpirationInDays(value *int32)()
    SetPinLowercaseCharactersUsage(value *ConfigurationUsage)()
    SetPinMaximumLength(value *int32)()
    SetPinMinimumLength(value *int32)()
    SetPinPreviousBlockCount(value *int32)()
    SetPinRecoveryEnabled(value *bool)()
    SetPinSpecialCharactersUsage(value *ConfigurationUsage)()
    SetPinUppercaseCharactersUsage(value *ConfigurationUsage)()
    SetSecurityDeviceRequired(value *bool)()
    SetUnlockWithBiometricsEnabled(value *bool)()
    SetUseCertificatesForOnPremisesAuthEnabled(value *bool)()
    SetUseSecurityKeyForSignin(value *bool)()
    SetWindowsHelloForBusinessBlocked(value *bool)()
}
