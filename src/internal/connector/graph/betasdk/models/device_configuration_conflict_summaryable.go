package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationConflictSummaryable 
type DeviceConfigurationConflictSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConflictingDeviceConfigurations()([]SettingSourceable)
    GetContributingSettings()([]string)
    GetDeviceCheckinsImpacted()(*int32)
    SetConflictingDeviceConfigurations(value []SettingSourceable)()
    SetContributingSettings(value []string)()
    SetDeviceCheckinsImpacted(value *int32)()
}
