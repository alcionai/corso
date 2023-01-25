package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkSystemConfigurationable 
type TeamworkSystemConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDateTimeConfiguration()(TeamworkDateTimeConfigurationable)
    GetDefaultPassword()(*string)
    GetDeviceLockTimeout()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetIsDeviceLockEnabled()(*bool)
    GetIsLoggingEnabled()(*bool)
    GetIsPowerSavingEnabled()(*bool)
    GetIsScreenCaptureEnabled()(*bool)
    GetIsSilentModeEnabled()(*bool)
    GetLanguage()(*string)
    GetLockPin()(*string)
    GetLoggingLevel()(*string)
    GetNetworkConfiguration()(TeamworkNetworkConfigurationable)
    GetOdataType()(*string)
    SetDateTimeConfiguration(value TeamworkDateTimeConfigurationable)()
    SetDefaultPassword(value *string)()
    SetDeviceLockTimeout(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetIsDeviceLockEnabled(value *bool)()
    SetIsLoggingEnabled(value *bool)()
    SetIsPowerSavingEnabled(value *bool)()
    SetIsScreenCaptureEnabled(value *bool)()
    SetIsSilentModeEnabled(value *bool)()
    SetLanguage(value *string)()
    SetLockPin(value *string)()
    SetLoggingLevel(value *string)()
    SetNetworkConfiguration(value TeamworkNetworkConfigurationable)()
    SetOdataType(value *string)()
}
