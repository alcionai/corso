package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BitLockerRecoveryOptionsable 
type BitLockerRecoveryOptionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBlockDataRecoveryAgent()(*bool)
    GetEnableBitLockerAfterRecoveryInformationToStore()(*bool)
    GetEnableRecoveryInformationSaveToStore()(*bool)
    GetHideRecoveryOptions()(*bool)
    GetOdataType()(*string)
    GetRecoveryInformationToStore()(*BitLockerRecoveryInformationType)
    GetRecoveryKeyUsage()(*ConfigurationUsage)
    GetRecoveryPasswordUsage()(*ConfigurationUsage)
    SetBlockDataRecoveryAgent(value *bool)()
    SetEnableBitLockerAfterRecoveryInformationToStore(value *bool)()
    SetEnableRecoveryInformationSaveToStore(value *bool)()
    SetHideRecoveryOptions(value *bool)()
    SetOdataType(value *string)()
    SetRecoveryInformationToStore(value *BitLockerRecoveryInformationType)()
    SetRecoveryKeyUsage(value *ConfigurationUsage)()
    SetRecoveryPasswordUsage(value *ConfigurationUsage)()
}
