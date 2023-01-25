package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceEncryptionStateable 
type ManagedDeviceEncryptionStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdvancedBitLockerStates()(*AdvancedBitLockerState)
    GetDeviceName()(*string)
    GetDeviceType()(*DeviceTypes)
    GetEncryptionPolicySettingState()(*ComplianceStatus)
    GetEncryptionReadinessState()(*EncryptionReadinessState)
    GetEncryptionState()(*EncryptionState)
    GetFileVaultStates()(*FileVaultState)
    GetOsVersion()(*string)
    GetPolicyDetails()([]EncryptionReportPolicyDetailsable)
    GetTpmSpecificationVersion()(*string)
    GetUserPrincipalName()(*string)
    SetAdvancedBitLockerStates(value *AdvancedBitLockerState)()
    SetDeviceName(value *string)()
    SetDeviceType(value *DeviceTypes)()
    SetEncryptionPolicySettingState(value *ComplianceStatus)()
    SetEncryptionReadinessState(value *EncryptionReadinessState)()
    SetEncryptionState(value *EncryptionState)()
    SetFileVaultStates(value *FileVaultState)()
    SetOsVersion(value *string)()
    SetPolicyDetails(value []EncryptionReportPolicyDetailsable)()
    SetTpmSpecificationVersion(value *string)()
    SetUserPrincipalName(value *string)()
}
