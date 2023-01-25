package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BitLockerSystemDrivePolicyable 
type BitLockerSystemDrivePolicyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptionMethod()(*BitLockerEncryptionMethod)
    GetMinimumPinLength()(*int32)
    GetOdataType()(*string)
    GetPrebootRecoveryEnableMessageAndUrl()(*bool)
    GetPrebootRecoveryMessage()(*string)
    GetPrebootRecoveryUrl()(*string)
    GetRecoveryOptions()(BitLockerRecoveryOptionsable)
    GetStartupAuthenticationBlockWithoutTpmChip()(*bool)
    GetStartupAuthenticationRequired()(*bool)
    GetStartupAuthenticationTpmKeyUsage()(*ConfigurationUsage)
    GetStartupAuthenticationTpmPinAndKeyUsage()(*ConfigurationUsage)
    GetStartupAuthenticationTpmPinUsage()(*ConfigurationUsage)
    GetStartupAuthenticationTpmUsage()(*ConfigurationUsage)
    SetEncryptionMethod(value *BitLockerEncryptionMethod)()
    SetMinimumPinLength(value *int32)()
    SetOdataType(value *string)()
    SetPrebootRecoveryEnableMessageAndUrl(value *bool)()
    SetPrebootRecoveryMessage(value *string)()
    SetPrebootRecoveryUrl(value *string)()
    SetRecoveryOptions(value BitLockerRecoveryOptionsable)()
    SetStartupAuthenticationBlockWithoutTpmChip(value *bool)()
    SetStartupAuthenticationRequired(value *bool)()
    SetStartupAuthenticationTpmKeyUsage(value *ConfigurationUsage)()
    SetStartupAuthenticationTpmPinAndKeyUsage(value *ConfigurationUsage)()
    SetStartupAuthenticationTpmPinUsage(value *ConfigurationUsage)()
    SetStartupAuthenticationTpmUsage(value *ConfigurationUsage)()
}
