package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcRestorePointSettingable 
type CloudPcRestorePointSettingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFrequencyInHours()(*int32)
    GetOdataType()(*string)
    GetUserRestoreEnabled()(*bool)
    SetFrequencyInHours(value *int32)()
    SetOdataType(value *string)()
    SetUserRestoreEnabled(value *bool)()
}
