package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntentDeviceSettingStateSummaryable 
type DeviceManagementIntentDeviceSettingStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantCount()(*int32)
    GetConflictCount()(*int32)
    GetErrorCount()(*int32)
    GetNonCompliantCount()(*int32)
    GetNotApplicableCount()(*int32)
    GetRemediatedCount()(*int32)
    GetSettingName()(*string)
    SetCompliantCount(value *int32)()
    SetConflictCount(value *int32)()
    SetErrorCount(value *int32)()
    SetNonCompliantCount(value *int32)()
    SetNotApplicableCount(value *int32)()
    SetRemediatedCount(value *int32)()
    SetSettingName(value *string)()
}
