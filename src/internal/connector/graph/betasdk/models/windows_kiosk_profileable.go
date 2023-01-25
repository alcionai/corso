package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskProfileable 
type WindowsKioskProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppConfiguration()(WindowsKioskAppConfigurationable)
    GetOdataType()(*string)
    GetProfileId()(*string)
    GetProfileName()(*string)
    GetUserAccountsConfiguration()([]WindowsKioskUserable)
    SetAppConfiguration(value WindowsKioskAppConfigurationable)()
    SetOdataType(value *string)()
    SetProfileId(value *string)()
    SetProfileName(value *string)()
    SetUserAccountsConfiguration(value []WindowsKioskUserable)()
}
