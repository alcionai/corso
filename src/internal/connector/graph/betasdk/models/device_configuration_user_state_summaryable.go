package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationUserStateSummaryable 
type DeviceConfigurationUserStateSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompliantUserCount()(*int32)
    GetConflictUserCount()(*int32)
    GetErrorUserCount()(*int32)
    GetNonCompliantUserCount()(*int32)
    GetNotApplicableUserCount()(*int32)
    GetRemediatedUserCount()(*int32)
    GetUnknownUserCount()(*int32)
    SetCompliantUserCount(value *int32)()
    SetConflictUserCount(value *int32)()
    SetErrorUserCount(value *int32)()
    SetNonCompliantUserCount(value *int32)()
    SetNotApplicableUserCount(value *int32)()
    SetRemediatedUserCount(value *int32)()
    SetUnknownUserCount(value *int32)()
}
